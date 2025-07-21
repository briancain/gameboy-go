package cartridge

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// MBC3 implementation
// MBC3 has the following features:
// - ROM: Up to 2MB (128 banks)
// - RAM: Up to 32KB (4 banks)
// - Real-Time Clock (RTC) with 5 registers
// - Battery-backed RAM and RTC
type MBC3 struct {
	// ROM data
	rom []byte

	// RAM data
	ram []byte

	// ROM bank selection (7 bits, 0-127)
	romBank byte

	// RAM bank selection (2 bits, 0-3) or RTC register selection
	ramBank byte

	// RAM enable flag
	ramEnabled bool

	// RTC registers
	rtcRegisters [5]byte // S, M, H, DL, DH
	rtcLatch     bool
	rtcLatched   [5]byte // Latched values of RTC registers

	// RTC base time (Unix timestamp when the emulator was started)
	rtcBaseTime int64
	rtcLastTime int64

	// Battery-backed RAM and RTC
	hasBattery bool
	hasTimer   bool
	savePath   string
}

// RTC register indices
const (
	RTC_S  = 0 // Seconds (0-59)
	RTC_M  = 1 // Minutes (0-59)
	RTC_H  = 2 // Hours (0-23)
	RTC_DL = 3 // Lower 8 bits of day counter (0-255)
	RTC_DH = 4 // Upper 1 bit of day counter, halt flag, day counter carry bit
)

// RTC_DH register bits
const (
	RTC_DH_DAY_MSB = 0x01 // Bit 0: Day counter bit 8
	RTC_DH_HALT    = 0x40 // Bit 6: Halt flag (0=active, 1=halted)
	RTC_DH_CARRY   = 0x80 // Bit 7: Day counter carry bit (1=counter overflow)
)

// Create a new MBC3
func NewMBC3(romData []byte, ramSize int, cartType byte, title string, batterySaveDir string) *MBC3 {
	mbc := &MBC3{
		rom:         romData,
		romBank:     1,
		ramBank:     0,
		ramEnabled:  false,
		hasBattery:  cartType == CART_MBC3_RAM_BAT || cartType == CART_MBC3_TIMER_BAT || cartType == CART_MBC3_TIMER_RAM_BAT,
		hasTimer:    cartType == CART_MBC3_TIMER_BAT || cartType == CART_MBC3_TIMER_RAM_BAT,
		rtcBaseTime: time.Now().Unix(),
		rtcLastTime: time.Now().Unix(),
	}

	// Allocate RAM based on size
	if ramSize > 0 {
		mbc.ram = make([]byte, ramSize)
	} else {
		mbc.ram = make([]byte, 8*1024) // Default to 8KB
	}

	// Initialize RTC registers
	for i := range mbc.rtcRegisters {
		mbc.rtcRegisters[i] = 0
	}
	for i := range mbc.rtcLatched {
		mbc.rtcLatched[i] = 0
	}

	// Set up save path for battery-backed RAM and RTC
	if mbc.hasBattery {
		// Create a valid filename from the title
		safeTitle := sanitizeFilename(title)
		mbc.savePath = filepath.Join(batterySaveDir, safeTitle+".sav")

		// Try to load saved RAM and RTC data
		mbc.loadRAM()
	}

	log.Printf("[MBC3] Initialized with %d ROM bytes, %d RAM bytes, battery: %v, timer: %v, save path: %s",
		len(romData), len(mbc.ram), mbc.hasBattery, mbc.hasTimer, mbc.savePath)

	return mbc
}

// Read a byte from the MBC
func (mbc *MBC3) ReadByte(addr uint16) byte {
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
			log.Printf("[MBC3] Warning: ROM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, bank, offset)
			return 0xFF
		}
		return mbc.rom[offset]

	case addr >= 0xA000 && addr < 0xC000:
		// RAM Bank 0-3 or RTC Register
		if !mbc.ramEnabled {
			return 0xFF
		}

		if mbc.ramBank <= 0x03 {
			// RAM access
			ramAddr := uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
			if ramAddr >= uint32(len(mbc.ram)) {
				log.Printf("[MBC3] Warning: RAM read out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.ramBank, ramAddr)
				return 0xFF
			}
			return mbc.ram[ramAddr]
		} else if mbc.hasTimer && mbc.ramBank >= 0x08 && mbc.ramBank <= 0x0C {
			// RTC register access
			rtcReg := mbc.ramBank - 0x08
			return mbc.rtcLatched[rtcReg]
		}

		return 0xFF

	default:
		return 0xFF
	}
}

// Write a byte to the MBC
func (mbc *MBC3) WriteByte(addr uint16, value byte) {
	switch {
	case addr < 0x2000:
		// RAM and Timer Enable (0x0000-0x1FFF)
		wasEnabled := mbc.ramEnabled
		mbc.ramEnabled = (value & 0x0F) == 0x0A

		// If RAM is being disabled and we have battery, save the RAM
		if wasEnabled && !mbc.ramEnabled && mbc.hasBattery {
			mbc.saveRAM()
		}

	case addr < 0x4000:
		// ROM Bank Number (0x2000-0x3FFF)
		// 7 bits are used (0-127)
		value &= 0x7F

		// Update ROM bank
		mbc.romBank = value

		// Bank 0 is treated as bank 1
		if mbc.romBank == 0 {
			mbc.romBank = 1
		}

	case addr < 0x6000:
		// RAM Bank Number or RTC Register Select (0x4000-0x5FFF)
		mbc.ramBank = value

	case addr < 0x8000:
		// Latch Clock Data (0x6000-0x7FFF)
		// When writing 0 followed by 1, the RTC data is latched
		if !mbc.rtcLatch && value == 0x01 {
			// Latch the RTC data
			mbc.updateRTC()
			for i := range mbc.rtcRegisters {
				mbc.rtcLatched[i] = mbc.rtcRegisters[i]
			}
			log.Printf("[MBC3] RTC data latched: S=%02X, M=%02X, H=%02X, DL=%02X, DH=%02X",
				mbc.rtcLatched[RTC_S], mbc.rtcLatched[RTC_M], mbc.rtcLatched[RTC_H],
				mbc.rtcLatched[RTC_DL], mbc.rtcLatched[RTC_DH])
		}
		mbc.rtcLatch = (value == 0x00)

	case addr >= 0xA000 && addr < 0xC000:
		// RAM Bank 0-3 or RTC Register
		if !mbc.ramEnabled {
			return
		}

		if mbc.ramBank <= 0x03 {
			// RAM access
			ramAddr := uint32(addr-0xA000) + uint32(mbc.ramBank)*0x2000
			if ramAddr >= uint32(len(mbc.ram)) {
				log.Printf("[MBC3] Warning: RAM write out of bounds: addr=%04X, bank=%d, offset=%d", addr, mbc.ramBank, ramAddr)
				return
			}
			mbc.ram[ramAddr] = value
		} else if mbc.hasTimer && mbc.ramBank >= 0x08 && mbc.ramBank <= 0x0C {
			// RTC register access
			rtcReg := mbc.ramBank - 0x08
			mbc.rtcRegisters[rtcReg] = value
			mbc.rtcLatched[rtcReg] = value // Update latched value too for testing

			// If writing to the DH register, check if the halt bit is being set/cleared
			if rtcReg == RTC_DH {
				if (value & RTC_DH_HALT) != 0 {
					// Halting the RTC - update it first
					mbc.updateRTC()
				} else if (mbc.rtcRegisters[RTC_DH] & RTC_DH_HALT) != 0 {
					// Resuming the RTC - reset the base time
					mbc.rtcBaseTime = time.Now().Unix()
					mbc.rtcLastTime = mbc.rtcBaseTime
				}
			}
		}
	}
}

// Update the RTC registers based on the current time
func (mbc *MBC3) updateRTC() {
	// If RTC is halted, don't update
	if (mbc.rtcRegisters[RTC_DH] & RTC_DH_HALT) != 0 {
		return
	}

	// Get current time
	now := time.Now().Unix()

	// Calculate elapsed seconds since last update
	elapsed := now - mbc.rtcLastTime
	if elapsed <= 0 {
		return
	}

	// Update last time
	mbc.rtcLastTime = now

	// Get current RTC values
	seconds := int(mbc.rtcRegisters[RTC_S])
	minutes := int(mbc.rtcRegisters[RTC_M])
	hours := int(mbc.rtcRegisters[RTC_H])
	days := int(mbc.rtcRegisters[RTC_DL])
	if (mbc.rtcRegisters[RTC_DH] & RTC_DH_DAY_MSB) != 0 {
		days |= 0x100
	}

	// Add elapsed seconds
	seconds += int(elapsed)

	// Update minutes if seconds overflow
	minutes += seconds / 60
	seconds %= 60

	// Update hours if minutes overflow
	hours += minutes / 60
	minutes %= 60

	// Update days if hours overflow
	days += hours / 24
	hours %= 24

	// Check for day counter overflow (> 511 days)
	if days > 511 {
		days %= 512
		// Set day counter carry bit
		mbc.rtcRegisters[RTC_DH] |= RTC_DH_CARRY
	}

	// Update RTC registers
	mbc.rtcRegisters[RTC_S] = byte(seconds)
	mbc.rtcRegisters[RTC_M] = byte(minutes)
	mbc.rtcRegisters[RTC_H] = byte(hours)
	mbc.rtcRegisters[RTC_DL] = byte(days & 0xFF)

	// Update day counter MSB
	if (days & 0x100) != 0 {
		mbc.rtcRegisters[RTC_DH] |= RTC_DH_DAY_MSB
	} else {
		mbc.rtcRegisters[RTC_DH] &= ^byte(RTC_DH_DAY_MSB)
	}
}

// Save RAM and RTC to file (for battery-backed RAM)
func (mbc *MBC3) saveRAM() {
	if !mbc.hasBattery {
		return
	}

	// Create saves directory if it doesn't exist
	os.MkdirAll(filepath.Dir(mbc.savePath), 0755)

	// Update RTC before saving
	if mbc.hasTimer {
		mbc.updateRTC()
	}

	// Create a buffer to hold RAM and RTC data
	var data []byte

	// Add RAM data
	data = append(data, mbc.ram...)

	// Add RTC data if timer is present
	if mbc.hasTimer {
		// Add RTC registers
		data = append(data, mbc.rtcRegisters[:]...)

		// Add RTC base time as 8 bytes (int64)
		baseTimeBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(baseTimeBytes, uint64(mbc.rtcBaseTime))
		data = append(data, baseTimeBytes...)

		// Add RTC last time as 8 bytes (int64)
		lastTimeBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(lastTimeBytes, uint64(mbc.rtcLastTime))
		data = append(data, lastTimeBytes...)
	}

	// Write data to file
	err := ioutil.WriteFile(mbc.savePath, data, 0644)
	if err != nil {
		log.Printf("[MBC3] Error saving RAM/RTC to %s: %v", mbc.savePath, err)
	} else {
		log.Printf("[MBC3] Saved RAM/RTC to %s", mbc.savePath)
	}
}

// Load RAM and RTC from file (for battery-backed RAM)
func (mbc *MBC3) loadRAM() {
	if !mbc.hasBattery {
		return
	}

	// Check if save file exists
	if _, err := os.Stat(mbc.savePath); os.IsNotExist(err) {
		log.Printf("[MBC3] No save file found at %s", mbc.savePath)
		return
	}

	// Read data from file
	data, err := ioutil.ReadFile(mbc.savePath)
	if err != nil {
		log.Printf("[MBC3] Error loading RAM/RTC from %s: %v", mbc.savePath, err)
		return
	}

	// Check if data is valid
	if len(data) < len(mbc.ram) {
		log.Printf("[MBC3] Save file too small: %d bytes, expected at least %d bytes", len(data), len(mbc.ram))
		return
	}

	// Copy RAM data
	copy(mbc.ram, data[:len(mbc.ram)])

	// Copy RTC data if timer is present
	if mbc.hasTimer {
		// Check if data contains RTC information
		expectedSize := len(mbc.ram) + 5 + 16 // RAM + 5 RTC registers + 16 bytes for timestamps
		if len(data) >= expectedSize {
			// Copy RTC registers
			copy(mbc.rtcRegisters[:], data[len(mbc.ram):len(mbc.ram)+5])

			// Copy RTC base time
			mbc.rtcBaseTime = int64(binary.LittleEndian.Uint64(data[len(mbc.ram)+5 : len(mbc.ram)+13]))

			// Copy RTC last time
			mbc.rtcLastTime = int64(binary.LittleEndian.Uint64(data[len(mbc.ram)+13 : len(mbc.ram)+21]))

			// Calculate time elapsed since last save
			now := time.Now().Unix()
			elapsed := now - mbc.rtcLastTime

			// Update RTC base time to account for time elapsed while the emulator was off
			mbc.rtcBaseTime += elapsed
			mbc.rtcLastTime = now

			// Copy latched RTC registers
			copy(mbc.rtcLatched[:], mbc.rtcRegisters[:])
		}
	}

	log.Printf("[MBC3] Loaded RAM/RTC from %s", mbc.savePath)
}

// SaveBatteryRAM saves the RAM and RTC to file if this cartridge has battery-backed RAM
func (mbc *MBC3) SaveBatteryRAM() {
	mbc.saveRAM()
}

// IsRumbling always returns false for MBC3 cartridges
func (mbc *MBC3) IsRumbling() bool {
	return false
}
