package gbcore

import (
	"os"
	"testing"
)

// TestMBC1BankingModes tests the different banking modes of MBC1
func TestMBC1BankingModes(t *testing.T) {
	// Create an MBC1 cartridge with 256KB ROM and 32KB RAM
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
	// Bank 15 (0x3C000-0x3FFFF) - within our 256KB ROM size
	rom[0x3C000+0x1000] = 0x20

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC1 with 32KB RAM
	mbc := NewMBC1(rom, 32*1024, CART_MBC1_RAM, "TESTROM", tmpDir)

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)

	// Test ROM banking mode (default)
	if mbc.bankingMode != 0 {
		t.Errorf("Expected default banking mode to be 0, got %d", mbc.bankingMode)
	}

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
	if mbc.ReadByte(0x5000) != 0x12 {
		t.Errorf("Expected ROM bank 2 at 0x5000 to be 0x12, got %02X", mbc.ReadByte(0x5000))
	}

	// Switch to ROM bank 15 (using upper bits)
	mbc.WriteByte(0x2000, 0x0F) // Lower 5 bits = 15
	// Should be bank 15 (0x0F = 15)
	if mbc.ReadByte(0x5000) != 0x20 {
		t.Errorf("Expected ROM bank 15 at 0x5000 to be 0x20, got %02X", mbc.ReadByte(0x5000))
	}

	// Test RAM banking in ROM mode
	// In ROM banking mode, only RAM bank 0 is accessible
	mbc.WriteByte(0xA000, 0xAA)
	mbc.WriteByte(0xA001, 0xBB)

	// Switch to RAM banking mode
	mbc.WriteByte(0x6000, 0x01)
	if mbc.bankingMode != 1 {
		t.Errorf("Expected banking mode to be 1, got %d", mbc.bankingMode)
	}

	// Write to RAM bank 0 first
	mbc.WriteByte(0xA000, 0xEE)
	mbc.WriteByte(0xA001, 0xFF)

	// Set RAM bank to 1
	mbc.WriteByte(0x4000, 0x01)
	if mbc.ramBank != 1 {
		t.Errorf("Expected RAM bank to be 1, got %d", mbc.ramBank)
	}

	// Write to RAM bank 1
	mbc.WriteByte(0xA000, 0xCC)
	mbc.WriteByte(0xA001, 0xDD)

	// Switch to RAM bank 0
	mbc.WriteByte(0x4000, 0x00)
	// Check that RAM bank 0 has our updated values
	if mbc.ReadByte(0xA000) != 0xEE || mbc.ReadByte(0xA001) != 0xFF {
		t.Errorf("Expected RAM bank 0 to contain 0xEE, 0xFF, got %02X, %02X",
			mbc.ReadByte(0xA000), mbc.ReadByte(0xA001))
	}

	// Switch back to RAM bank 1
	mbc.WriteByte(0x4000, 0x01)
	// Check that RAM bank 1 has our values
	if mbc.ReadByte(0xA000) != 0xCC || mbc.ReadByte(0xA001) != 0xDD {
		t.Errorf("Expected RAM bank 1 to contain 0xCC, 0xDD, got %02X, %02X",
			mbc.ReadByte(0xA000), mbc.ReadByte(0xA001))
	}

	// Switch back to ROM banking mode
	mbc.WriteByte(0x6000, 0x00)
	if mbc.bankingMode != 0 {
		t.Errorf("Expected banking mode to be 0, got %d", mbc.bankingMode)
	}

	// The upper bits should now affect ROM bank selection again
	// ROM bank should now be 1 << 5 | 15 = 47
	expectedBank := byte(0x2F) // 0x20 | 0x0F
	if mbc.romBank != expectedBank {
		t.Errorf("Expected ROM bank to be 0x%02X, got 0x%02X", expectedBank, mbc.romBank)
	}
}

// TestMBC1BatteryRAM tests the battery-backed RAM functionality of MBC1
func TestMBC1BatteryRAM(t *testing.T) {
	// Create an MBC1 cartridge with 128KB ROM and 8KB RAM
	rom := make([]byte, 128*1024)

	// Create a temporary directory for save files
	tmpDir, err := os.MkdirTemp("", "gameboy_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create MBC1 with battery-backed RAM
	mbc := NewMBC1(rom, 8*1024, CART_MBC1_RAM_BAT, "TESTROM", tmpDir)

	// Enable RAM
	mbc.WriteByte(0x0000, 0x0A)

	// Write some data to RAM
	mbc.WriteByte(0xA000, 0x12)
	mbc.WriteByte(0xA001, 0x34)
	mbc.WriteByte(0xA002, 0x56)

	// Save RAM to file
	mbc.saveRAM()

	// Create a new MBC1 instance that should load the saved RAM
	mbc2 := NewMBC1(rom, 8*1024, CART_MBC1_RAM_BAT, "TESTROM", tmpDir)

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

	// Create a third MBC1 instance to verify the save
	mbc3 := NewMBC1(rom, 8*1024, CART_MBC1_RAM_BAT, "TESTROM", tmpDir)
	mbc3.WriteByte(0x0000, 0x0A)

	// Check that the updated RAM data was loaded correctly
	if mbc3.ReadByte(0xA000) != 0x78 {
		t.Errorf("Expected RAM to contain 0x78, got %02X", mbc3.ReadByte(0xA000))
	}
}

// TestSanitizeFilename tests the sanitizeFilename function
func TestSanitizeFilename(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"TESTROM", "TESTROM"},
		{"TEST ROM", "TEST_ROM"},
		{"TEST-ROM", "TEST-ROM"},
		{"TEST/ROM", "TEST_ROM"},
		{"TEST\\ROM", "TEST_ROM"},
		{"TEST:ROM", "TEST_ROM"},
		{"TEST*ROM", "TEST_ROM"},
		{"TEST?ROM", "TEST_ROM"},
		{"TEST\"ROM", "TEST_ROM"},
		{"TEST<ROM", "TEST_ROM"},
		{"TEST>ROM", "TEST_ROM"},
		{"TEST|ROM", "TEST_ROM"},
	}

	for _, test := range tests {
		result := sanitizeFilename(test.input)
		if result != test.expected {
			t.Errorf("sanitizeFilename(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
