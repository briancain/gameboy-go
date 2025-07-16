package gbcore

import (
	"log"
)

// MBC1 implementation
type MBC1 struct {
	// ROM data
	rom []byte
	
	// RAM data
	ram []byte
	
	// ROM bank selection (5 bits, 0-31)
	romBank byte
	
	// RAM bank selection (2 bits, 0-3)
	ramBank byte
	
	// RAM enable flag
	ramEnabled bool
	
	// Banking mode (0 = ROM banking, 1 = RAM banking)
	bankingMode byte
}

// Create a new MBC1
func NewMBC1(romData []byte, ramSize int) *MBC1 {
	mbc := &MBC1{
		rom:         romData,
		romBank:     1,
		ramBank:     0,
		ramEnabled:  false,
		bankingMode: 0,
	}
	
	// Allocate RAM based on size
	if ramSize > 0 {
		mbc.ram = make([]byte, ramSize)
	} else {
		mbc.ram = make([]byte, 8*1024) // Default to 8KB
	}
	
	log.Printf("[MBC1] Initialized with %d ROM bytes, %d RAM bytes", len(romData), len(mbc.ram))
	
	return mbc
}

// Read a byte from the MBC
func (mbc *MBC1) ReadByte(addr uint16) byte {
	switch {
	case addr < 0x4000:
		// ROM Bank 0 (fixed)
		return mbc.rom[addr]
	
	case addr < 0x8000:
		// ROM Bank 1-127 (switchable)
		bank := mbc.romBank
		if bank == 0 {
			bank = 1
		}
		offset := uint32(addr-0x4000) + uint32(bank)*0x4000
		if offset >= uint32(len(mbc.rom)) {
			log.Printf("[MBC1] Warning: ROM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, bank, offset)
			return 0xFF
		}
		return mbc.rom[offset]
	
	case addr >= 0xA000 && addr < 0xC000:
		// RAM Bank 0-3 (if present and enabled)
		if !mbc.ramEnabled || len(mbc.ram) == 0 {
			return 0xFF
		}
		
		var ramAddr uint32
		if mbc.bankingMode == 1 {
			// RAM banking mode
			ramAddr = uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
		} else {
			// ROM banking mode
			ramAddr = uint32(addr - 0xA000)
		}
		
		if ramAddr >= uint32(len(mbc.ram)) {
			log.Printf("[MBC1] Warning: RAM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.ramBank, ramAddr)
			return 0xFF
		}
		
		return mbc.ram[ramAddr]
	
	default:
		return 0xFF
	}
}

// Write a byte to the MBC
func (mbc *MBC1) WriteByte(addr uint16, value byte) {
	switch {
	case addr < 0x2000:
		// RAM Enable (0x0000-0x1FFF)
		mbc.ramEnabled = (value & 0x0F) == 0x0A
	
	case addr < 0x4000:
		// ROM Bank Number (0x2000-0x3FFF)
		// Lower 5 bits of the ROM bank number
		mbc.romBank = (mbc.romBank & 0x60) | (value & 0x1F)
		if mbc.romBank == 0 {
			mbc.romBank = 1
		}
	
	case addr < 0x6000:
		// RAM Bank Number or Upper ROM Bank Number (0x4000-0x5FFF)
		if mbc.bankingMode == 0 {
			// ROM banking mode - set upper bits of ROM bank
			mbc.romBank = (mbc.romBank & 0x1F) | ((value & 0x03) << 5)
			if mbc.romBank == 0 {
				mbc.romBank = 1
			}
		} else {
			// RAM banking mode - set RAM bank
			mbc.ramBank = value & 0x03
		}
	
	case addr < 0x8000:
		// Banking Mode Select (0x6000-0x7FFF)
		mbc.bankingMode = value & 0x01
	
	case addr >= 0xA000 && addr < 0xC000:
		// RAM Bank 0-3 (if present and enabled)
		if !mbc.ramEnabled || len(mbc.ram) == 0 {
			return
		}
		
		var ramAddr uint32
		if mbc.bankingMode == 1 {
			// RAM banking mode
			ramAddr = uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
		} else {
			// ROM banking mode
			ramAddr = uint32(addr - 0xA000)
		}
		
		if ramAddr >= uint32(len(mbc.ram)) {
			log.Printf("[MBC1] Warning: RAM write out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.ramBank, ramAddr)
			return
		}
		
		mbc.ram[ramAddr] = value
	}
}
