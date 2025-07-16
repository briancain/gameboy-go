package gbcore

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

// MockCartridge is a mock implementation of the Cartridge interface for testing
type MockCartridge struct{}

func (m *MockCartridge) ReadByte(addr uint16) byte {
	return 0
}

func (m *MockCartridge) WriteByte(addr uint16, value byte) {
}
