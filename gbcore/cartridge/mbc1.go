package gbcore

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// MBC1 implementation
// MBC1 has two modes:
// - 16Mbit ROM/8KByte RAM mode (default): ROM banks 0-127, RAM bank 0
// - 4Mbit ROM/32KByte RAM mode: ROM banks 0-31, RAM banks 0-3
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

	// Battery-backed RAM
	hasBattery bool
	savePath   string
}

// Create a new MBC1
func NewMBC1(romData []byte, ramSize int, cartType byte, title string, saveDir string) *MBC1 {
	mbc := &MBC1{
		rom:         romData,
		romBank:     1,
		ramBank:     0,
		ramEnabled:  false,
		bankingMode: 0,
		hasBattery:  cartType == CART_MBC1_RAM_BAT,
	}

	// Allocate RAM based on size
	if ramSize > 0 {
		mbc.ram = make([]byte, ramSize)
	} else {
		mbc.ram = make([]byte, 8*1024) // Default to 8KB
	}

	// Set up save path for battery-backed RAM
	if mbc.hasBattery {
		// Create a valid filename from the title
		safeTitle := sanitizeFilename(title)
		mbc.savePath = filepath.Join(saveDir, safeTitle+".sav")

		// Try to load saved RAM data
		mbc.loadRAM()
	}

	log.Printf("[MBC1] Initialized with %d ROM bytes, %d RAM bytes, battery: %v, save path: %s",
		len(romData), len(mbc.ram), mbc.hasBattery, mbc.savePath)

	return mbc
}

// Sanitize a string to be used as a filename
func sanitizeFilename(name string) string {
	// Replace invalid characters with underscores
	result := ""
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			result += string(r)
		} else {
			result += "_"
		}
	}
	return result
}

// Read a byte from the MBC
func (mbc *MBC1) ReadByte(addr uint16) byte {
	switch {
	case addr < 0x4000:
		// ROM Bank 0 (fixed)
		// In RAM banking mode, this can be affected by the upper bits
		if mbc.bankingMode == 1 {
			// In RAM banking mode, the upper 2 bits of the ROM bank number
			// can affect the ROM bank 0 area
			bank := (mbc.ramBank << 5) & 0x60
			if bank == 0 {
				return mbc.rom[addr]
			}
			offset := uint32(addr) + uint32(bank)*0x4000
			if offset >= uint32(len(mbc.rom)) {
				log.Printf("[MBC1] Warning: ROM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, bank, offset)
				return 0xFF
			}
			return mbc.rom[offset]
		}
		// In ROM banking mode, always use bank 0
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
		if mbc.bankingMode == 1 && len(mbc.ram) > 0x2000 {
			// RAM banking mode - use ramBank to select the bank
			ramAddr = uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
		} else {
			// ROM banking mode or small RAM - always use bank 0
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
		wasEnabled := mbc.ramEnabled
		mbc.ramEnabled = (value & 0x0F) == 0x0A

		// If RAM is being disabled and we have battery, save the RAM
		if wasEnabled && !mbc.ramEnabled && mbc.hasBattery {
			mbc.saveRAM()
		}

	case addr < 0x4000:
		// ROM Bank Number (0x2000-0x3FFF)
		// Lower 5 bits of the ROM bank number
		value &= 0x1F // Only use the lower 5 bits

		// Update the lower 5 bits of the ROM bank number
		mbc.romBank = (mbc.romBank & 0x60) | value

		// Bank 0 is treated as bank 1
		if (mbc.romBank & 0x1F) == 0 {
			mbc.romBank = (mbc.romBank & 0x60) | 0x01
		}

	case addr < 0x6000:
		// RAM Bank Number or Upper ROM Bank Number (0x4000-0x5FFF)
		value &= 0x03 // Only use the lower 2 bits

		if mbc.bankingMode == 0 {
			// ROM banking mode - set upper bits of ROM bank
			mbc.romBank = (mbc.romBank & 0x1F) | (value << 5)
		} else {
			// RAM banking mode - set RAM bank
			mbc.ramBank = value
		}

	case addr < 0x8000:
		// Banking Mode Select (0x6000-0x7FFF)
		oldMode := mbc.bankingMode
		mbc.bankingMode = value & 0x01

		// If switching modes, we might need to adjust the bank numbers
		if oldMode != mbc.bankingMode {
			if mbc.bankingMode == 0 {
				// Switching to ROM banking mode
				// Move the upper bits from ramBank to romBank
				mbc.romBank = (mbc.romBank & 0x1F) | (mbc.ramBank << 5)
				// Bank 0 is treated as bank 1
				if (mbc.romBank & 0x1F) == 0 {
					mbc.romBank = (mbc.romBank & 0x60) | 0x01
				}
			} else {
				// Switching to RAM banking mode
				// Move the upper bits from romBank to ramBank
				mbc.ramBank = (mbc.romBank >> 5) & 0x03
			}
		}

	case addr >= 0xA000 && addr < 0xC000:
		// RAM Bank 0-3 (if present and enabled)
		if !mbc.ramEnabled || len(mbc.ram) == 0 {
			return
		}

		var ramAddr uint32
		if mbc.bankingMode == 1 && len(mbc.ram) > 0x2000 {
			// RAM banking mode - use ramBank to select the bank
			ramAddr = uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
		} else {
			// ROM banking mode or small RAM - always use bank 0
			ramAddr = uint32(addr - 0xA000)
		}

		if ramAddr >= uint32(len(mbc.ram)) {
			log.Printf("[MBC1] Warning: RAM write out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.ramBank, ramAddr)
			return
		}

		mbc.ram[ramAddr] = value
	}
}

// Save RAM to file (for battery-backed RAM)
func (mbc *MBC1) saveRAM() {
	if !mbc.hasBattery || len(mbc.ram) == 0 {
		return
	}

	// Create saves directory if it doesn't exist
	os.MkdirAll(filepath.Dir(mbc.savePath), 0755)

	// Write RAM to file
	err := ioutil.WriteFile(mbc.savePath, mbc.ram, 0644)
	if err != nil {
		log.Printf("[MBC1] Error saving RAM to %s: %v", mbc.savePath, err)
	} else {
		log.Printf("[MBC1] Saved RAM to %s", mbc.savePath)
	}
}

// Load RAM from file (for battery-backed RAM)
func (mbc *MBC1) loadRAM() {
	if !mbc.hasBattery || len(mbc.ram) == 0 {
		return
	}

	// Check if save file exists
	if _, err := os.Stat(mbc.savePath); os.IsNotExist(err) {
		log.Printf("[MBC1] No save file found at %s", mbc.savePath)
		return
	}

	// Read RAM from file
	data, err := ioutil.ReadFile(mbc.savePath)
	if err != nil {
		log.Printf("[MBC1] Error loading RAM from %s: %v", mbc.savePath, err)
		return
	}

	// Copy data to RAM
	copy(mbc.ram, data)
	log.Printf("[MBC1] Loaded RAM from %s", mbc.savePath)
}

// SaveBatteryRAM saves the RAM to file if this cartridge has battery-backed RAM
func (mbc *MBC1) SaveBatteryRAM() {
	mbc.saveRAM()
}
