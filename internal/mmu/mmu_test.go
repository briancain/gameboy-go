package mmu

import (
	"testing"
)

// TestMMUInitialization verifies that a new MMU can be created
func TestMMUInitialization(t *testing.T) {
	// Create a new MMU
	mmu := NewMMU()

	// Check that the MMU was initialized
	if mmu == nil {
		t.Error("Expected MMU to be initialized, got nil")
	}

	// Test basic read/write operations
	addr := uint16(0xFF80) // High RAM
	value := byte(0x42)

	// Write a value
	mmu.WriteByte(addr, value)

	// Read it back
	readValue := mmu.ReadByte(addr)

	// Check that the value was stored correctly
	if readValue != value {
		t.Errorf("Expected %02X, got %02X", value, readValue)
	}
}

// TestMMUReset tests the MMU reset function
func TestMMUReset(t *testing.T) {
	// Create a new MMU
	mmu := NewMMU()

	// Write some values to memory
	mmu.WriteByte(0xFF80, 0x42) // HRAM
	mmu.WriteByte(0xFF00, 0x55) // I/O
	mmu.WriteByte(0xC000, 0xAA) // WRAM

	// Reset the MMU
	mmu.Reset()

	// Check that memory was reset
	if mmu.ReadByte(0xFF80) != 0 {
		t.Errorf("Expected HRAM to be reset to 0, got %02X", mmu.ReadByte(0xFF80))
	}

	// Check that I/O registers were initialized to their default values
	if mmu.ReadByte(0xFF40) != 0x91 { // LCDC
		t.Errorf("Expected LCDC to be initialized to 0x91, got %02X", mmu.ReadByte(0xFF40))
	}
}

// TestMMUMemoryMap tests the MMU memory mapping
func TestMMUMemoryMap(t *testing.T) {
	// Create a new MMU
	mmu := NewMMU()

	// Create a mock cartridge
	mockCart := &MockCartridge{}
	mmu.SetCartridge(mockCart)

	// Test ROM access (should go to cartridge)
	mmu.ReadByte(0x0000)
	if mockCart.lastReadAddr != 0x0000 {
		t.Errorf("Expected ROM read at 0x0000 to access cartridge, got address %04X", mockCart.lastReadAddr)
	}

	mmu.ReadByte(0x4000)
	if mockCart.lastReadAddr != 0x4000 {
		t.Errorf("Expected ROM read at 0x4000 to access cartridge, got address %04X", mockCart.lastReadAddr)
	}

	// Test VRAM access
	value := byte(0x42)
	mmu.WriteByte(0x8000, value)
	if mmu.ReadByte(0x8000) != value {
		t.Errorf("Expected VRAM read at 0x8000 to return %02X, got %02X", value, mmu.ReadByte(0x8000))
	}

	// Test external RAM access (should go to cartridge)
	mmu.WriteByte(0xA000, value)
	if mockCart.lastWriteAddr != 0xA000 || mockCart.lastWriteValue != value {
		t.Errorf("Expected external RAM write at 0xA000 to access cartridge")
	}

	// Test WRAM access
	mmu.WriteByte(0xC000, value)
	if mmu.ReadByte(0xC000) != value {
		t.Errorf("Expected WRAM read at 0xC000 to return %02X, got %02X", value, mmu.ReadByte(0xC000))
	}

	// Test echo RAM
	mmu.WriteByte(0xC100, value)
	if mmu.ReadByte(0xE100) != value {
		t.Errorf("Expected echo RAM read at 0xE100 to return %02X, got %02X", value, mmu.ReadByte(0xE100))
	}

	// Test OAM access
	mmu.WriteByte(0xFE00, value)
	if mmu.ReadByte(0xFE00) != value {
		t.Errorf("Expected OAM read at 0xFE00 to return %02X, got %02X", value, mmu.ReadByte(0xFE00))
	}

	// Test I/O register access
	mmu.WriteByte(0xFF00, value)
	// Note: The joypad register (0xFF00) has special handling, so let's use a different register
	mmu.WriteByte(0xFF01, value)
	if mmu.ReadByte(0xFF01) != value {
		t.Errorf("Expected I/O read at 0xFF01 to return %02X, got %02X", value, mmu.ReadByte(0xFF01))
	}

	// Test HRAM access
	mmu.WriteByte(0xFF80, value)
	if mmu.ReadByte(0xFF80) != value {
		t.Errorf("Expected HRAM read at 0xFF80 to return %02X, got %02X", value, mmu.ReadByte(0xFF80))
	}

	// Test IE register
	mmu.WriteByte(0xFFFF, value)
	if mmu.ReadByte(0xFFFF) != value {
		t.Errorf("Expected IE read at 0xFFFF to return %02X, got %02X", value, mmu.ReadByte(0xFFFF))
	}
}

// TestMMUWordOperations tests the 16-bit read/write operations
func TestMMUWordOperations(t *testing.T) {
	// Create a new MMU
	mmu := NewMMU()

	// Test word write/read
	addr := uint16(0xC000)
	value := uint16(0xABCD)

	mmu.WriteWord(addr, value)

	// Check that the bytes were written correctly (little-endian)
	if mmu.ReadByte(addr) != 0xCD || mmu.ReadByte(addr+1) != 0xAB {
		t.Errorf("Expected bytes at %04X to be CD AB, got %02X %02X",
			addr, mmu.ReadByte(addr), mmu.ReadByte(addr+1))
	}

	// Check that the word is read back correctly
	readValue := mmu.ReadWord(addr)
	if readValue != value {
		t.Errorf("Expected ReadWord to return %04X, got %04X", value, readValue)
	}
}

// TestMMUDMATransfer tests the OAM DMA transfer
func TestMMUDMATransfer(t *testing.T) {
	// Create a new MMU
	mmu := NewMMU()

	// Fill source area with test data
	sourceBase := uint16(0xC000)
	for i := uint16(0); i < 160; i++ {
		mmu.WriteByte(sourceBase+i, byte(i))
	}

	// Trigger DMA transfer from 0xC000
	mmu.WriteByte(0xFF46, 0xC0)

	// Check that OAM was filled with the correct data
	for i := uint16(0); i < 160; i++ {
		oamAddr := uint16(0xFE00) + i
		expected := byte(i)
		if mmu.ReadByte(oamAddr) != expected {
			t.Errorf("Expected OAM at %04X to be %02X, got %02X",
				oamAddr, expected, mmu.ReadByte(oamAddr))
		}
	}
}

// MockCartridge is a mock implementation of the Cartridge interface for testing
type MockCartridge struct {
	lastReadAddr   uint16
	lastWriteAddr  uint16
	lastWriteValue byte
}

func (m *MockCartridge) ReadByte(addr uint16) byte {
	m.lastReadAddr = addr
	return 0 // Return 0 for all reads
}

func (m *MockCartridge) WriteByte(addr uint16, value byte) {
	m.lastWriteAddr = addr
	m.lastWriteValue = value
}
