package gbcore

import (
	"testing"
)

// TestCPUInitialization verifies that a new CPU can be created
func TestCPUInitialization(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, err := NewCPU(mockMMU)

	// Check that there was no error
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that the CPU was initialized
	if cpu == nil {
		t.Error("Expected CPU to be initialized, got nil")
	}
}

// MockMMU is a mock implementation of the MMU interface for testing
type MockMMU struct{}

func (m *MockMMU) ReadByte(addr uint16) byte {
	return 0
}

func (m *MockMMU) WriteByte(addr uint16, value byte) {
}

func (m *MockMMU) ReadWord(addr uint16) uint16 {
	return 0
}

func (m *MockMMU) WriteWord(addr uint16, value uint16) {
}
