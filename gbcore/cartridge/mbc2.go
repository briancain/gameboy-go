package gbcore

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// MBC2 implementation
// MBC2 has the following features:
// - ROM: Up to 256KB (16 banks)
// - Built-in RAM: 512×4 bits (256 bytes, but only 4 bits per byte are used)
// - Special addressing: Bit 8 of the address determines register selection
type MBC2 struct {
	// ROM data
	rom []byte

	// Built-in RAM (512×4 bits)
	ram [512]byte

	// ROM bank selection (4 bits, 0-15)
	romBank byte

	// RAM enable flag
	ramEnabled bool

	// Battery-backed RAM
	hasBattery bool
	savePath   string
}

// Create a new MBC2
func NewMBC2(romData []byte, cartType byte, title string, batterySaveDir string) *MBC2 {
	mbc := &MBC2{
		rom:        romData,
		romBank:    1,
		ramEnabled: false,
		hasBattery: cartType == CART_MBC2_BAT,
	}

	// Initialize RAM to zeros
	for i := range mbc.ram {
		mbc.ram[i] = 0
	}

	// Set up save path for battery-backed RAM
	if mbc.hasBattery {
		// Create a valid filename from the title
		safeTitle := sanitizeFilename(title)
		mbc.savePath = filepath.Join(batterySaveDir, safeTitle+".sav")

		// Try to load saved RAM data
		mbc.loadRAM()
	}

	log.Printf("[MBC2] Initialized with %d ROM bytes, 512×4 bits RAM, battery: %v, save path: %s",
		len(romData), mbc.hasBattery, mbc.savePath)

	return mbc
}

// Read a byte from the MBC
func (mbc *MBC2) ReadByte(addr uint16) byte {
	switch {
	case addr < 0x4000:
		// ROM Bank 0 (fixed)
		return mbc.rom[addr]

	case addr < 0x8000:
		// ROM Bank 1-15 (switchable)
		bank := mbc.romBank
		if bank == 0 {
			bank = 1
		}
		offset := uint32(addr-0x4000) + uint32(bank)*0x4000
		if offset >= uint32(len(mbc.rom)) {
			log.Printf("[MBC2] Warning: ROM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, bank, offset)
			return 0xFF
		}
		return mbc.rom[offset]

	case addr >= 0xA000 && addr < 0xA200:
		// Built-in RAM (512×4 bits)
		if !mbc.ramEnabled {
			return 0xFF
		}

		// MBC2 RAM is only 4 bits per byte
		ramAddr := addr - 0xA000
		return mbc.ram[ramAddr] & 0x0F

	default:
		return 0xFF
	}
}

// Write a byte to the MBC
func (mbc *MBC2) WriteByte(addr uint16, value byte) {
	switch {
	case addr < 0x4000:
		// The least significant bit of the upper address byte must be zero to enable/disable RAM
		// and must be one to select a ROM bank
		if (addr & 0x0100) == 0 {
			// RAM Enable (0x0000-0x1FFF with bit 8 = 0)
			wasEnabled := mbc.ramEnabled
			mbc.ramEnabled = (value & 0x0F) == 0x0A

			// If RAM is being disabled and we have battery, save the RAM
			if wasEnabled && !mbc.ramEnabled && mbc.hasBattery {
				mbc.saveRAM()
			}
		} else {
			// ROM Bank Number (0x2000-0x3FFF with bit 8 = 1)
			// Only the lower 4 bits are used
			mbc.romBank = value & 0x0F

			// Bank 0 is treated as bank 1
			if mbc.romBank == 0 {
				mbc.romBank = 1
			}
		}

	case addr >= 0xA000 && addr < 0xA200:
		// Built-in RAM (512×4 bits)
		if !mbc.ramEnabled {
			return
		}

		// MBC2 RAM is only 4 bits per byte
		ramAddr := addr - 0xA000
		mbc.ram[ramAddr] = value & 0x0F
	}
}

// Save RAM to file (for battery-backed RAM)
func (mbc *MBC2) saveRAM() {
	if !mbc.hasBattery {
		return
	}

	// Create saves directory if it doesn't exist
	os.MkdirAll(filepath.Dir(mbc.savePath), 0755)

	// Write RAM to file
	err := ioutil.WriteFile(mbc.savePath, mbc.ram[:], 0644)
	if err != nil {
		log.Printf("[MBC2] Error saving RAM to %s: %v", mbc.savePath, err)
	} else {
		log.Printf("[MBC2] Saved RAM to %s", mbc.savePath)
	}
}

// Load RAM from file (for battery-backed RAM)
func (mbc *MBC2) loadRAM() {
	if !mbc.hasBattery {
		return
	}

	// Check if save file exists
	if _, err := os.Stat(mbc.savePath); os.IsNotExist(err) {
		log.Printf("[MBC2] No save file found at %s", mbc.savePath)
		return
	}

	// Read RAM from file
	data, err := ioutil.ReadFile(mbc.savePath)
	if err != nil {
		log.Printf("[MBC2] Error loading RAM from %s: %v", mbc.savePath, err)
		return
	}

	// Copy data to RAM
	copy(mbc.ram[:], data)
	log.Printf("[MBC2] Loaded RAM from %s", mbc.savePath)
}

// SaveBatteryRAM saves the RAM to file if this cartridge has battery-backed RAM
func (mbc *MBC2) SaveBatteryRAM() {
	mbc.saveRAM()
}

// IsRumbling always returns false for MBC2 cartridges
func (mbc *MBC2) IsRumbling() bool {
	return false
}
