package gbcore

import (
	"os"
	"testing"
)

// TestMBC3ROMBanking tests the ROM banking functionality of MBC3
func TestMBC3ROMBanking(t *testing.T) {
	// Create an MBC3 cartridge with 256KB ROM
	rom := make([]byte, 256*1024)

	// Set test values in different ROM banks
	// Bank 0 (0x0000-0x3FFF)
	rom[0x1000] = 0x10
	// Bank 1 (0x4000-0x7FFF)
	rom[0x4000+0x1000] = 0x11
	// Bank 2 (0x8000-0xBFFF)
	rom[0x8000+0x1000] = 0x12
	// Bank 3 (0xC000-0xFFFF)
	rom[0xC000+0x1000] = 0x13
	// Bank 15 (0x3C000-0x3FFFF)
	rom[0x3C000+0x1000] = 0x1F

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC3
	mbc := NewMBC3(rom, 32*1024, CART_MBC3_RAM, "TESTROM", tmpDir)

	// Test ROM bank 0 (fixed)
	if mbc.ReadByte(0x1000) != 0x10 {
		t.Errorf("Expected ROM bank 0 at 0x1000 to be 0x10, got %02X", mbc.ReadByte(0x1000))
	}

	// Test ROM bank 1 (default)
	if mbc.ReadByte(0x5000) != 0x11 {
		t.Errorf("Expected ROM bank 1 at 0x5000 to be 0x11, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 2
	mbc.WriteByte(0x2000, 0x02)
	if mbc.romBank != 2 {
		t.Errorf("Expected ROM bank to be 2, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x12 {
		t.Errorf("Expected ROM bank 2 at 0x5000 to be 0x12, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 15
	mbc.WriteByte(0x2000, 0x0F)
	if mbc.romBank != 15 {
		t.Errorf("Expected ROM bank to be 15, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x1F {
		t.Errorf("Expected ROM bank 15 at 0x5000 to be 0x1F, got %02X", mbc.ReadByte(0x5000))
	}

	// Test bank 0 selection (should be treated as bank 1)
	mbc.WriteByte(0x2000, 0x00)
	if mbc.romBank != 1 {
		t.Errorf("Expected ROM bank to be 1 (bank 0 treated as 1), got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x11 {
		t.Errorf("Expected ROM bank 1 at 0x5000 to be 0x11, got %02X", mbc.ReadByte(0x5000))
	}
}

// TestMBC3RAMBanking tests the RAM banking functionality of MBC3
func TestMBC3RAMBanking(t *testing.T) {
	// Create an MBC3 cartridge with 64KB ROM and 32KB RAM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC3
	mbc := NewMBC3(rom, 32*1024, CART_MBC3_RAM, "TESTROM", tmpDir)

	// Test RAM access (disabled by default)
	mbc.WriteByte(0xA000, 0x55)
	if mbc.ReadByte(0xA000) != 0xFF {
		t.Errorf("Expected RAM read when disabled to return 0xFF, got %02X", mbc.ReadByte(0xA000))
	}

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)
	if !mbc.ramEnabled {
		t.Errorf("Expected RAM to be enabled")
	}

	// Test RAM bank 0
	mbc.WriteByte(0x4000, 0x00) // Select RAM bank 0
	mbc.WriteByte(0xA000, 0x55)
	mbc.WriteByte(0xA001, 0xAA)
	if mbc.ReadByte(0xA000) != 0x55 || mbc.ReadByte(0xA001) != 0xAA {
		t.Errorf("Expected RAM bank 0 to contain 0x55, 0xAA, got %02X, %02X",
			mbc.ReadByte(0xA000), mbc.ReadByte(0xA001))
	}

	// Test RAM bank 1
	mbc.WriteByte(0x4000, 0x01) // Select RAM bank 1
	mbc.WriteByte(0xA000, 0x33)
	mbc.WriteByte(0xA001, 0x44)
	if mbc.ReadByte(0xA000) != 0x33 || mbc.ReadByte(0xA001) != 0x44 {
		t.Errorf("Expected RAM bank 1 to contain 0x33, 0x44, got %02X, %02X",
			mbc.ReadByte(0xA000), mbc.ReadByte(0xA001))
	}

	// Switch back to RAM bank 0 and verify data
	mbc.WriteByte(0x4000, 0x00)
	if mbc.ReadByte(0xA000) != 0x55 || mbc.ReadByte(0xA001) != 0xAA {
		t.Errorf("Expected RAM bank 0 to contain 0x55, 0xAA, got %02X, %02X",
			mbc.ReadByte(0xA000), mbc.ReadByte(0xA001))
	}

	// Disable RAM
	mbc.WriteByte(0x0000, 0x00)
	if mbc.ramEnabled {
		t.Errorf("Expected RAM to be disabled")
	}

	// Test RAM access after disabling
	if mbc.ReadByte(0xA000) != 0xFF {
		t.Errorf("Expected RAM read when disabled to return 0xFF, got %02X", mbc.ReadByte(0xA000))
	}
}

// TestMBC3RTC tests the Real-Time Clock functionality of MBC3
func TestMBC3RTC(t *testing.T) {
	// Create an MBC3 cartridge with 64KB ROM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC3 with RTC
	mbc := NewMBC3(rom, 8*1024, CART_MBC3_TIMER_BAT, "TESTROM", tmpDir)

	// Enable RAM/RTC
	mbc.WriteByte(0x0000, 0x0A)
	if !mbc.ramEnabled {
		t.Errorf("Expected RAM/RTC to be enabled")
	}

	// Test halting the RTC
	mbc.WriteByte(0x4000, 0x0C) // Select RTC day counter upper register
	mbc.WriteByte(0xA000, 0x40) // Set halt bit (bit 6)

	// Verify halt bit is set
	if (mbc.rtcRegisters[RTC_DH] & RTC_DH_HALT) == 0 {
		t.Errorf("Expected RTC halt bit to be set")
	}

	// Test resuming the RTC
	mbc.WriteByte(0x4000, 0x0C) // Select RTC day counter upper register
	mbc.WriteByte(0xA000, 0x00) // Clear halt bit

	// Verify halt bit is cleared
	if (mbc.rtcRegisters[RTC_DH] & RTC_DH_HALT) != 0 {
		t.Errorf("Expected RTC halt bit to be cleared")
	}
}

// TestMBC3BatteryRAM tests the battery-backed RAM functionality of MBC3
func TestMBC3BatteryRAM(t *testing.T) {
	// Create an MBC3 cartridge with 64KB ROM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC3 with battery-backed RAM
	mbc := NewMBC3(rom, 8*1024, CART_MBC3_RAM_BAT, "TESTROM", tmpDir)

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)

	// Write some data to RAM
	mbc.WriteByte(0xA000, 0x12)
	mbc.WriteByte(0xA001, 0x34)
	mbc.WriteByte(0xA002, 0x56)

	// Save RAM to file
	mbc.saveRAM()

	// Create a new MBC3 instance that should load the saved RAM
	mbc2 := NewMBC3(rom, 8*1024, CART_MBC3_RAM_BAT, "TESTROM", tmpDir)

	// Enable RAM
	mbc2.WriteByte(0x0000, 0x0A)

	// Check that the RAM data was loaded correctly
	if mbc2.ReadByte(0xA000) != 0x12 || mbc2.ReadByte(0xA001) != 0x34 || mbc2.ReadByte(0xA002) != 0x56 {
		t.Errorf("Expected RAM to contain 0x12, 0x34, 0x56, got %02X, %02X, %02X",
			mbc2.ReadByte(0xA000), mbc2.ReadByte(0xA001), mbc2.ReadByte(0xA002))
	}

	// Test SaveBatteryRAM method
	mbc2.WriteByte(0xA000, 0x78)
	mbc2.SaveBatteryRAM()

	// Create a third MBC3 instance to verify the save
	mbc3 := NewMBC3(rom, 8*1024, CART_MBC3_RAM_BAT, "TESTROM", tmpDir)
	mbc3.WriteByte(0x0000, 0x0A)

	// Check that the updated RAM data was loaded correctly
	if mbc3.ReadByte(0xA000) != 0x78 {
		t.Errorf("Expected RAM to contain 0x78, got %02X", mbc3.ReadByte(0xA000))
	}
}

// TestMBC3RTCSave tests saving and loading RTC data
func TestMBC3RTCSave(t *testing.T) {
	// Create an MBC3 cartridge with 64KB ROM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC3 with RTC
	mbc := NewMBC3(rom, 8*1024, CART_MBC3_TIMER_BAT, "TESTROM", tmpDir)

	// Enable RAM/RTC
	mbc.WriteByte(0x0000, 0x0A)

	// Set the halt bit to stop the RTC from updating
	mbc.WriteByte(0x4000, 0x0C) // Select RTC day counter upper register
	mbc.WriteByte(0xA000, 0x40) // Set halt bit (bit 6)

	// Save RTC to file
	mbc.saveRAM()

	// Create a new MBC3 instance that should load the saved RTC
	mbc2 := NewMBC3(rom, 8*1024, CART_MBC3_TIMER_BAT, "TESTROM", tmpDir)

	// Enable RAM/RTC
	mbc2.WriteByte(0x0000, 0x0A)

	// Verify the halt bit is still set
	if (mbc2.rtcRegisters[RTC_DH] & RTC_DH_HALT) == 0 {
		t.Errorf("Expected RTC halt bit to be set after loading")
	}
}
