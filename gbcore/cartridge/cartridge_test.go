package gbcore

import (
	"os"
	"testing"
)

// TestROMOnlyCartridge tests the ROM-only cartridge implementation
func TestROMOnlyCartridge(t *testing.T) {
	// Create a ROM-only cartridge
	rom := make([]byte, 32*1024) // 32KB ROM
	ram := make([]byte, 8*1024)  // 8KB RAM

	// Set some test values in ROM
	rom[0x100] = 0x42
	rom[0x4000] = 0x84

	cart := &ROMOnly{
		rom: rom,
		ram: ram,
	}

	// Test ROM reading
	if cart.ReadByte(0x100) != 0x42 {
		t.Errorf("Expected ROM[0x100] to be 0x42, got %02X", cart.ReadByte(0x100))
	}

	if cart.ReadByte(0x4000) != 0x84 {
		t.Errorf("Expected ROM[0x4000] to be 0x84, got %02X", cart.ReadByte(0x4000))
	}

	// Test RAM reading/writing
	cart.WriteByte(0xA000, 0x55)
	if cart.ReadByte(0xA000) != 0x55 {
		t.Errorf("Expected RAM[0x0000] to be 0x55, got %02X", cart.ReadByte(0xA000))
	}

	// Test out-of-bounds access
	if cart.ReadByte(0xFFFF) != 0xFF {
		t.Errorf("Expected out-of-bounds read to return 0xFF, got %02X", cart.ReadByte(0xFFFF))
	}
}

// TestMBC1 tests the MBC1 implementation
func TestMBC1(t *testing.T) {
	// Create an MBC1 cartridge with 128KB ROM and 8KB RAM
	rom := make([]byte, 128*1024)

	// Set some test values in different ROM banks
	rom[0x100] = 0x42  // Bank 0
	rom[0x4100] = 0x43 // Bank 1
	rom[0x8100] = 0x44 // Bank 2
	rom[0xC100] = 0x45 // Bank 3

	mbc := NewMBC1(rom, 8*1024)

	// Test ROM bank 0 (fixed)
	if mbc.ReadByte(0x100) != 0x42 {
		t.Errorf("Expected ROM bank 0 at 0x100 to be 0x42, got %02X", mbc.ReadByte(0x100))
	}

	// Test ROM bank 1 (default)
	if mbc.ReadByte(0x4100) != 0x43 {
		t.Errorf("Expected ROM bank 1 at 0x4100 to be 0x43, got %02X", mbc.ReadByte(0x4100))
	}

	// Switch to ROM bank 2
	mbc.WriteByte(0x2000, 0x02)

	// Test ROM bank 2
	if mbc.ReadByte(0x4100) != 0x44 {
		t.Errorf("Expected ROM bank 2 at 0x4100 to be 0x44, got %02X", mbc.ReadByte(0x4100))
	}

	// Test RAM access (disabled by default)
	mbc.WriteByte(0xA000, 0x55)
	if mbc.ReadByte(0xA000) != 0xFF {
		t.Errorf("Expected RAM read when disabled to return 0xFF, got %02X", mbc.ReadByte(0xA000))
	}

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)

	// Test RAM writing/reading
	mbc.WriteByte(0xA000, 0x55)
	if mbc.ReadByte(0xA000) != 0x55 {
		t.Errorf("Expected RAM at 0xA000 to be 0x55, got %02X", mbc.ReadByte(0xA000))
	}
}

// TestCartridgeHelpers tests the helper functions in the cartridge package
func TestCartridgeHelpers(t *testing.T) {
	// Test ROM size calculation
	if getROMSize(0x00) != 32*1024 {
		t.Errorf("Expected ROM size for 0x00 to be 32KB, got %d", getROMSize(0x00))
	}

	if getROMSize(0x01) != 64*1024 {
		t.Errorf("Expected ROM size for 0x01 to be 64KB, got %d", getROMSize(0x01))
	}

	// Test RAM size calculation
	if getRAMSize(RAM_NONE) != 0 {
		t.Errorf("Expected RAM size for RAM_NONE to be 0, got %d", getRAMSize(RAM_NONE))
	}

	if getRAMSize(RAM_8KB) != 8*1024 {
		t.Errorf("Expected RAM size for RAM_8KB to be 8KB, got %d", getRAMSize(RAM_8KB))
	}
}

// TestNewCartridge tests the NewCartridge function
func TestNewCartridge(t *testing.T) {
	// Create a temporary test ROM file
	tmpFile, err := os.CreateTemp("", "test_rom_*.gb")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Test with valid file
	cart, err := NewCartridge(tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error for valid file, got %v", err)
	}

	if cart == nil {
		t.Error("Expected cartridge to be initialized, got nil")
	}

	// Test with invalid file
	_, err = NewCartridge("/nonexistent/file.gb")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}
