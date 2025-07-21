package cartridge

import (
	"os"
	"testing"
)

// TestMBC5ROMBanking tests the ROM banking functionality of MBC5
func TestMBC5ROMBanking(t *testing.T) {
	// Create an MBC5 cartridge with 1MB ROM (64 banks)
	rom := make([]byte, 1024*1024)

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
	// Bank 16 (0x40000-0x43FFF)
	rom[0x40000+0x1000] = 0x20
	// Bank 63 (0xFC000-0xFFFFF)
	rom[0xFC000+0x1000] = 0x3F

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC5
	mbc := NewMBC5(rom, 32*1024, CART_MBC5, "TESTROM", tmpDir)

	// Test ROM bank 0 (fixed)
	if mbc.ReadByte(0x1000) != 0x10 {
		t.Errorf("Expected ROM bank 0 at 0x1000 to be 0x10, got %02X", mbc.ReadByte(0x1000))
	}

	// Test ROM bank 1 (default)
	if mbc.ReadByte(0x5000) != 0x11 {
		t.Errorf("Expected ROM bank 1 at 0x5000 to be 0x11, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 2 (using lower 8 bits)
	mbc.WriteByte(0x2000, 0x02)
	if mbc.romBank != 2 {
		t.Errorf("Expected ROM bank to be 2, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x12 {
		t.Errorf("Expected ROM bank 2 at 0x5000 to be 0x12, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 15 (using lower 8 bits)
	mbc.WriteByte(0x2000, 0x0F)
	if mbc.romBank != 15 {
		t.Errorf("Expected ROM bank to be 15, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x1F {
		t.Errorf("Expected ROM bank 15 at 0x5000 to be 0x1F, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 16 (using lower 8 bits)
	mbc.WriteByte(0x2000, 0x10)
	if mbc.romBank != 16 {
		t.Errorf("Expected ROM bank to be 16, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x20 {
		t.Errorf("Expected ROM bank 16 at 0x5000 to be 0x20, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 63 (using lower 8 bits)
	mbc.WriteByte(0x2000, 0x3F)
	if mbc.romBank != 63 {
		t.Errorf("Expected ROM bank to be 63, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x3F {
		t.Errorf("Expected ROM bank 63 at 0x5000 to be 0x3F, got %02X", mbc.ReadByte(0x5000))
	}

	// Test 9-bit ROM bank selection (using upper 1 bit)
	// We can't test this fully without a larger ROM, but we can test the bit setting
	mbc.WriteByte(0x2000, 0x00) // Lower 8 bits = 0
	mbc.WriteByte(0x3000, 0x01) // Upper 1 bit = 1
	if mbc.romBank != 256 {
		t.Errorf("Expected ROM bank to be 256, got %d", mbc.romBank)
	}

	// Test bank 0 selection (valid in MBC5, unlike MBC1-3)
	mbc.WriteByte(0x2000, 0x00) // Lower 8 bits = 0
	mbc.WriteByte(0x3000, 0x00) // Upper 1 bit = 0
	if mbc.romBank != 0 {
		t.Errorf("Expected ROM bank to be 0, got %d", mbc.romBank)
	}
	// In MBC5, bank 0 is actually accessible (unlike MBC1-3)
	// But we can't test this properly because we're reading from the same ROM data
	// Just verify that the bank number is set correctly
}

// TestMBC5RAMBanking tests the RAM banking functionality of MBC5
func TestMBC5RAMBanking(t *testing.T) {
	// Create an MBC5 cartridge with 64KB ROM and 128KB RAM (16 banks)
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC5
	mbc := NewMBC5(rom, 128*1024, CART_MBC5_RAM, "TESTROM", tmpDir)

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

	// Test RAM bank 15 (max for MBC5)
	mbc.WriteByte(0x4000, 0x0F) // Select RAM bank 15
	mbc.WriteByte(0xA000, 0x77)
	mbc.WriteByte(0xA001, 0x88)
	if mbc.ReadByte(0xA000) != 0x77 || mbc.ReadByte(0xA001) != 0x88 {
		t.Errorf("Expected RAM bank 15 to contain 0x77, 0x88, got %02X, %02X",
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

// TestMBC5Rumble tests the rumble functionality of MBC5
func TestMBC5Rumble(t *testing.T) {
	// Create an MBC5 cartridge with 64KB ROM and 32KB RAM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC5 with rumble
	mbc := NewMBC5(rom, 32*1024, CART_MBC5_RUMBLE_RAM, "TESTROM", tmpDir)

	// Check that rumble is initially off
	if mbc.IsRumbling() {
		t.Errorf("Expected rumble to be off initially")
	}

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)

	// Turn on rumble (bit 3)
	mbc.WriteByte(0x4000, 0x08)
	if !mbc.IsRumbling() {
		t.Errorf("Expected rumble to be on after setting bit 3")
	}

	// Turn off rumble
	mbc.WriteByte(0x4000, 0x00)
	if mbc.IsRumbling() {
		t.Errorf("Expected rumble to be off after clearing bit 3")
	}

	// Test that RAM bank selection still works with rumble
	// RAM bank 0 with rumble off
	mbc.WriteByte(0x4000, 0x00)
	mbc.WriteByte(0xA000, 0x55)

	// RAM bank 1 with rumble on
	mbc.WriteByte(0x4000, 0x09) // Bank 1 (0x01) with rumble on (0x08)
	mbc.WriteByte(0xA000, 0xAA)

	// Check that rumble is on
	if !mbc.IsRumbling() {
		t.Errorf("Expected rumble to be on with bank 1")
	}

	// Check RAM bank 1 data
	if mbc.ReadByte(0xA000) != 0xAA {
		t.Errorf("Expected RAM bank 1 to contain 0xAA, got %02X", mbc.ReadByte(0xA000))
	}

	// Switch back to RAM bank 0 with rumble off
	mbc.WriteByte(0x4000, 0x00)

	// Check that rumble is off
	if mbc.IsRumbling() {
		t.Errorf("Expected rumble to be off with bank 0")
	}

	// Check RAM bank 0 data
	if mbc.ReadByte(0xA000) != 0x55 {
		t.Errorf("Expected RAM bank 0 to contain 0x55, got %02X", mbc.ReadByte(0xA000))
	}
}

// TestMBC5BatteryRAM tests the battery-backed RAM functionality of MBC5
func TestMBC5BatteryRAM(t *testing.T) {
	// Create an MBC5 cartridge with 64KB ROM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC5 with battery-backed RAM
	mbc := NewMBC5(rom, 8*1024, CART_MBC5_RAM_BAT, "TESTROM", tmpDir)

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)

	// Write some data to RAM
	mbc.WriteByte(0xA000, 0x12)
	mbc.WriteByte(0xA001, 0x34)
	mbc.WriteByte(0xA002, 0x56)

	// Save RAM to file
	mbc.saveRAM()

	// Create a new MBC5 instance that should load the saved RAM
	mbc2 := NewMBC5(rom, 8*1024, CART_MBC5_RAM_BAT, "TESTROM", tmpDir)

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

	// Create a third MBC5 instance to verify the save
	mbc3 := NewMBC5(rom, 8*1024, CART_MBC5_RAM_BAT, "TESTROM", tmpDir)
	mbc3.WriteByte(0x0000, 0x0A)

	// Check that the updated RAM data was loaded correctly
	if mbc3.ReadByte(0xA000) != 0x78 {
		t.Errorf("Expected RAM to contain 0x78, got %02X", mbc3.ReadByte(0xA000))
	}
}
