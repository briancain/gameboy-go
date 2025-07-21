package gbcore

import (
	"os"
	"testing"
)

// TestMBC2ROMBanking tests the ROM banking functionality of MBC2
func TestMBC2ROMBanking(t *testing.T) {
	// Create an MBC2 cartridge with 256KB ROM
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

	// Create MBC2
	mbc := NewMBC2(rom, CART_MBC2, "TESTROM", tmpDir)

	// Test ROM bank 0 (fixed)
	if mbc.ReadByte(0x1000) != 0x10 {
		t.Errorf("Expected ROM bank 0 at 0x1000 to be 0x10, got %02X", mbc.ReadByte(0x1000))
	}

	// Test ROM bank 1 (default)
	if mbc.ReadByte(0x5000) != 0x11 {
		t.Errorf("Expected ROM bank 1 at 0x5000 to be 0x11, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 2 (using bit 8 = 1 in address)
	mbc.WriteByte(0x2100, 0x02)
	if mbc.romBank != 2 {
		t.Errorf("Expected ROM bank to be 2, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x12 {
		t.Errorf("Expected ROM bank 2 at 0x5000 to be 0x12, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 15
	mbc.WriteByte(0x2100, 0x0F)
	if mbc.romBank != 15 {
		t.Errorf("Expected ROM bank to be 15, got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x1F {
		t.Errorf("Expected ROM bank 15 at 0x5000 to be 0x1F, got %02X", mbc.ReadByte(0x5000))
	}

	// Test bank 0 selection (should be treated as bank 1)
	mbc.WriteByte(0x2100, 0x00)
	if mbc.romBank != 1 {
		t.Errorf("Expected ROM bank to be 1 (bank 0 treated as 1), got %d", mbc.romBank)
	}
	if mbc.ReadByte(0x5000) != 0x11 {
		t.Errorf("Expected ROM bank 1 at 0x5000 to be 0x11, got %02X", mbc.ReadByte(0x5000))
	}
}

// TestMBC2RAM tests the built-in RAM functionality of MBC2
func TestMBC2RAM(t *testing.T) {
	// Create an MBC2 cartridge with 64KB ROM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC2
	mbc := NewMBC2(rom, CART_MBC2, "TESTROM", tmpDir)

	// Test RAM access (disabled by default)
	mbc.WriteByte(0xA000, 0x55)
	if mbc.ReadByte(0xA000) != 0xFF {
		t.Errorf("Expected RAM read when disabled to return 0xFF, got %02X", mbc.ReadByte(0xA000))
	}

	// Enable RAM (using bit 8 = 0 in address)
	mbc.WriteByte(0x0000, 0x0A)
	if !mbc.ramEnabled {
		t.Errorf("Expected RAM to be enabled")
	}

	// Test RAM writing/reading (only 4 bits per byte)
	mbc.WriteByte(0xA000, 0x55) // 0x55 = 0101 0101, but only lower 4 bits (0101 = 5) are used
	if mbc.ReadByte(0xA000) != 0x05 {
		t.Errorf("Expected RAM at 0xA000 to be 0x05, got %02X", mbc.ReadByte(0xA000))
	}

	// Test RAM at different addresses
	mbc.WriteByte(0xA001, 0x0F)
	mbc.WriteByte(0xA002, 0xAA) // 0xAA = 1010 1010, but only lower 4 bits (1010 = A) are used
	if mbc.ReadByte(0xA001) != 0x0F || mbc.ReadByte(0xA002) != 0x0A {
		t.Errorf("Expected RAM at 0xA001, 0xA002 to be 0x0F, 0x0A, got %02X, %02X",
			mbc.ReadByte(0xA001), mbc.ReadByte(0xA002))
	}

	// Test RAM at the end of the 512×4 bits range
	mbc.WriteByte(0xA1FF, 0x0C)
	if mbc.ReadByte(0xA1FF) != 0x0C {
		t.Errorf("Expected RAM at 0xA1FF to be 0x0C, got %02X", mbc.ReadByte(0xA1FF))
	}

	// Test RAM beyond the 512×4 bits range (should return 0xFF)
	if mbc.ReadByte(0xA200) != 0xFF {
		t.Errorf("Expected RAM read beyond range to return 0xFF, got %02X", mbc.ReadByte(0xA200))
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

// TestMBC2BatteryRAM tests the battery-backed RAM functionality of MBC2
func TestMBC2BatteryRAM(t *testing.T) {
	// Create an MBC2 cartridge with 64KB ROM
	rom := make([]byte, 64*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC2 with battery-backed RAM
	mbc := NewMBC2(rom, CART_MBC2_BAT, "TESTROM", tmpDir)

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)

	// Write some data to RAM
	mbc.WriteByte(0xA000, 0x05)
	mbc.WriteByte(0xA001, 0x0A)
	mbc.WriteByte(0xA002, 0x0F)

	// Save RAM to file
	mbc.saveRAM()

	// Create a new MBC2 instance that should load the saved RAM
	mbc2 := NewMBC2(rom, CART_MBC2_BAT, "TESTROM", tmpDir)

	// Enable RAM
	mbc2.WriteByte(0x0000, 0x0A)

	// Check that the RAM data was loaded correctly
	if mbc2.ReadByte(0xA000) != 0x05 || mbc2.ReadByte(0xA001) != 0x0A || mbc2.ReadByte(0xA002) != 0x0F {
		t.Errorf("Expected RAM to contain 0x05, 0x0A, 0x0F, got %02X, %02X, %02X",
			mbc2.ReadByte(0xA000), mbc2.ReadByte(0xA001), mbc2.ReadByte(0xA002))
	}

	// Test SaveBatteryRAM method
	mbc2.WriteByte(0xA000, 0x07)
	mbc2.SaveBatteryRAM()

	// Create a third MBC2 instance to verify the save
	mbc3 := NewMBC2(rom, CART_MBC2_BAT, "TESTROM", tmpDir)
	mbc3.WriteByte(0x0000, 0x0A)

	// Check that the updated RAM data was loaded correctly
	if mbc3.ReadByte(0xA000) != 0x07 {
		t.Errorf("Expected RAM to contain 0x07, got %02X", mbc3.ReadByte(0xA000))
	}
}
