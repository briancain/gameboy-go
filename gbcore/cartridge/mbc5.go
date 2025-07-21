package gbcore

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// MBC5 implementation
// MBC5 has the following features:
// - ROM: Up to 8MB (512 banks)
// - RAM: Up to 128KB (16 banks)
// - 9-bit ROM bank selection (0-511)
// - Optional rumble feature
type MBC5 struct {
	// ROM data
	rom []byte

	// RAM data
	ram []byte

	// ROM bank selection (9 bits, 0-511)
	romBank uint16

	// RAM bank selection (4 bits, 0-15)
	ramBank byte

	// RAM enable flag
	ramEnabled bool

	// Rumble feature
	hasRumble bool
	rumble    bool

	// Battery-backed RAM
	hasBattery bool
	savePath   string
}

// Create a new MBC5
func NewMBC5(romData []byte, ramSize int, cartType byte, title string, batterySaveDir string) *MBC5 {
	mbc := &MBC5{
		rom:        romData,
		romBank:    1,
		ramBank:    0,
		ramEnabled: false,
		hasRumble:  cartType == CART_MBC5_RUMBLE || cartType == CART_MBC5_RUMBLE_RAM || cartType == CART_MBC5_RUMBLE_RAM_BAT,
		hasBattery: cartType == CART_MBC5_RAM_BAT || cartType == CART_MBC5_RUMBLE_RAM_BAT,
		rumble:     false,
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
		mbc.savePath = filepath.Join(batterySaveDir, safeTitle+".sav")

		// Try to load saved RAM data
		mbc.loadRAM()
	}

	log.Printf("[MBC5] Initialized with %d ROM bytes, %d RAM bytes, battery: %v, rumble: %v, save path: %s",
		len(romData), len(mbc.ram), mbc.hasBattery, mbc.hasRumble, mbc.savePath)

	return mbc
}

// Read a byte from the MBC
func (mbc *MBC5) ReadByte(addr uint16) byte {
	switch {
	case addr < 0x4000:
		// ROM Bank 0 (fixed)
		return mbc.rom[addr]

	case addr < 0x8000:
		// ROM Bank 1-511 (switchable)
		offset := uint32(addr-0x4000) + uint32(mbc.romBank)*0x4000
		if offset >= uint32(len(mbc.rom)) {
			log.Printf("[MBC5] Warning: ROM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.romBank, offset)
			return 0xFF
		}
		return mbc.rom[offset]

	case addr >= 0xA000 && addr < 0xC000:
		// RAM Bank 0-15 (if present and enabled)
		if !mbc.ramEnabled || len(mbc.ram) == 0 {
			return 0xFF
		}

		ramAddr := uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
		if ramAddr >= uint32(len(mbc.ram)) {
			log.Printf("[MBC5] Warning: RAM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.ramBank, ramAddr)
			return 0xFF
		}

		return mbc.ram[ramAddr]

	default:
		return 0xFF
	}
}

// Write a byte to the MBC
func (mbc *MBC5) WriteByte(addr uint16, value byte) {
	switch {
	case addr < 0x2000:
		// RAM Enable (0x0000-0x1FFF)
		wasEnabled := mbc.ramEnabled
		mbc.ramEnabled = (value & 0x0F) == 0x0A

		// If RAM is being disabled and we have battery, save the RAM
		if wasEnabled && !mbc.ramEnabled && mbc.hasBattery {
			mbc.saveRAM()
		}

	case addr < 0x3000:
		// ROM Bank Number Lower 8 bits (0x2000-0x2FFF)
		mbc.romBank = (mbc.romBank & 0x100) | uint16(value)

	case addr < 0x4000:
		// ROM Bank Number Upper 1 bit (0x3000-0x3FFF)
		// Only bit 0 is used
		mbc.romBank = (mbc.romBank & 0xFF) | (uint16(value&0x01) << 8)

	case addr < 0x6000:
		// RAM Bank Number (0x4000-0x5FFF)
		if mbc.hasRumble {
			// For rumble cartridges, bit 3 controls the rumble motor
			mbc.rumble = (value & 0x08) != 0
			// Only bits 0-2 are used for RAM bank selection
			mbc.ramBank = value & 0x07
		} else {
			// For non-rumble cartridges, bits 0-3 are used for RAM bank selection
			mbc.ramBank = value & 0x0F
		}

	case addr >= 0xA000 && addr < 0xC000:
		// RAM Bank 0-15 (if present and enabled)
		if !mbc.ramEnabled || len(mbc.ram) == 0 {
			return
		}

		ramAddr := uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
		if ramAddr >= uint32(len(mbc.ram)) {
			log.Printf("[MBC5] Warning: RAM write out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.ramBank, ramAddr)
			return
		}

		mbc.ram[ramAddr] = value
	}
}

// Save RAM to file (for battery-backed RAM)
func (mbc *MBC5) saveRAM() {
	if !mbc.hasBattery || len(mbc.ram) == 0 {
		return
	}

	// Create saves directory if it doesn't exist
	os.MkdirAll(filepath.Dir(mbc.savePath), 0755)

	// Write RAM to file
	err := ioutil.WriteFile(mbc.savePath, mbc.ram, 0644)
	if err != nil {
		log.Printf("[MBC5] Error saving RAM to %s: %v", mbc.savePath, err)
	} else {
		log.Printf("[MBC5] Saved RAM to %s", mbc.savePath)
	}
}

// Load RAM from file (for battery-backed RAM)
func (mbc *MBC5) loadRAM() {
	if !mbc.hasBattery || len(mbc.ram) == 0 {
		return
	}

	// Check if save file exists
	if _, err := os.Stat(mbc.savePath); os.IsNotExist(err) {
		log.Printf("[MBC5] No save file found at %s", mbc.savePath)
		return
	}

	// Read RAM from file
	data, err := ioutil.ReadFile(mbc.savePath)
	if err != nil {
		log.Printf("[MBC5] Error loading RAM from %s: %v", mbc.savePath, err)
		return
	}

	// Copy data to RAM
	copy(mbc.ram, data)
	log.Printf("[MBC5] Loaded RAM from %s", mbc.savePath)
}

// SaveBatteryRAM saves the RAM to file if this cartridge has battery-backed RAM
func (mbc *MBC5) SaveBatteryRAM() {
	mbc.saveRAM()
}

// IsRumbling returns true if the rumble feature is currently active
func (mbc *MBC5) IsRumbling() bool {
	return mbc.hasRumble && mbc.rumble
}
