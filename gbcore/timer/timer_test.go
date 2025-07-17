package gbcore

import (
	"testing"
)

// TestTimerInitialization verifies that a new Timer can be created
func TestTimerInitialization(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Check that the Timer was initialized
	if timer == nil {
		t.Error("Expected Timer to be initialized, got nil")
	}

	// Check initial register values
	if timer.div != 0 {
		t.Errorf("Expected DIV to be 0, got %02X", timer.div)
	}

	if timer.tima != 0 {
		t.Errorf("Expected TIMA to be 0, got %02X", timer.tima)
	}

	if timer.tma != 0 {
		t.Errorf("Expected TMA to be 0, got %02X", timer.tma)
	}

	if timer.tac != 0 {
		t.Errorf("Expected TAC to be 0, got %02X", timer.tac)
	}
}

// TestTimerReset tests the Timer reset function
func TestTimerReset(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Modify register values
	timer.div = 0x42
	timer.tima = 0x42
	timer.tma = 0x42
	timer.tac = 0x42
	timer.divCounter = 100
	timer.timaCounter = 100

	// Reset the Timer
	timer.Reset()

	// Check that registers were reset
	if timer.div != 0 {
		t.Errorf("Expected DIV to be reset to 0, got %02X", timer.div)
	}

	if timer.tima != 0 {
		t.Errorf("Expected TIMA to be reset to 0, got %02X", timer.tima)
	}

	if timer.tma != 0 {
		t.Errorf("Expected TMA to be reset to 0, got %02X", timer.tma)
	}

	if timer.tac != 0 {
		t.Errorf("Expected TAC to be reset to 0, got %02X", timer.tac)
	}

	if timer.divCounter != 0 {
		t.Errorf("Expected divCounter to be reset to 0, got %d", timer.divCounter)
	}

	if timer.timaCounter != 0 {
		t.Errorf("Expected timaCounter to be reset to 0, got %d", timer.timaCounter)
	}
}

// TestTimerDIV tests the DIV register
func TestTimerDIV(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Step the timer by 256 cycles (should increment DIV once)
	timer.Step(256)

	// Check that DIV was incremented
	if timer.div != 1 {
		t.Errorf("Expected DIV to be incremented to 1, got %02X", timer.div)
	}

	// Step the timer by 255 cycles (should not increment DIV again)
	timer.Step(255)

	// Check that DIV was not incremented
	if timer.div != 1 {
		t.Errorf("Expected DIV to remain 1, got %02X", timer.div)
	}

	// Step the timer by 1 more cycle (should increment DIV again)
	timer.Step(1)

	// Check that DIV was incremented
	if timer.div != 2 {
		t.Errorf("Expected DIV to be incremented to 2, got %02X", timer.div)
	}

	// Test writing to DIV (should reset it)
	timer.WriteRegister(0xFF04, 0x42)

	// Check that DIV was reset
	if timer.div != 0 {
		t.Errorf("Expected DIV to be reset to 0, got %02X", timer.div)
	}

	// Check that divCounter was reset
	if timer.divCounter != 0 {
		t.Errorf("Expected divCounter to be reset to 0, got %d", timer.divCounter)
	}
}

// TestTimerTIMA tests the TIMA register
func TestTimerTIMA(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Enable timer with 4096Hz frequency (1024 cycles per increment)
	timer.tac = 0x04 // Timer enabled, 4096Hz

	// Step the timer by 1024 cycles (should increment TIMA once)
	timer.Step(1024)

	// Check that TIMA was incremented
	if timer.tima != 1 {
		t.Errorf("Expected TIMA to be incremented to 1, got %02X", timer.tima)
	}

	// Test TIMA overflow
	timer.tima = 0xFF
	timer.tma = 0x42 // Modulo value

	// Step the timer by 1024 cycles (should overflow TIMA and set it to TMA)
	timer.Step(1024)

	// Check that TIMA was set to TMA
	if timer.tima != 0x42 {
		t.Errorf("Expected TIMA to be set to TMA (0x42), got %02X", timer.tima)
	}

	// Check that interrupt was requested
	if mockMMU.interruptFlag&0x04 == 0 {
		t.Error("Expected timer interrupt to be requested")
	}
}

// TestTimerFrequency tests the different timer frequencies
func TestTimerFrequency(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Test each frequency
	testCases := []struct {
		tac      byte
		cycles   int
		expected byte
	}{
		{0x04, 1024, 1}, // 4096Hz (1024 cycles per increment)
		{0x05, 16, 1},   // 262144Hz (16 cycles per increment)
		{0x06, 64, 1},   // 65536Hz (64 cycles per increment)
		{0x07, 256, 1},  // 16384Hz (256 cycles per increment)
	}

	for _, tc := range testCases {
		// Reset timer
		timer.Reset()

		// Set TAC
		timer.tac = tc.tac

		// Step the timer
		timer.Step(tc.cycles)

		// Check that TIMA was incremented correctly
		if timer.tima != tc.expected {
			t.Errorf("With TAC=%02X, expected TIMA to be %d after %d cycles, got %02X",
				tc.tac, tc.expected, tc.cycles, timer.tima)
		}
	}
}

// TestTimerDisabled tests that the timer does nothing when disabled
func TestTimerDisabled(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Disable timer
	timer.tac = 0x00

	// Step the timer by a large number of cycles
	timer.Step(10000)

	// Check that TIMA was not incremented
	if timer.tima != 0 {
		t.Errorf("Expected TIMA to remain 0 when timer is disabled, got %02X", timer.tima)
	}

	// DIV should still increment
	if timer.div == 0 {
		t.Error("Expected DIV to increment even when timer is disabled")
	}
}

// MockMMU is a mock implementation of the MMU interface for testing
type MockMMU struct {
	interruptFlag byte
}

func (m *MockMMU) WriteByte(addr uint16, value byte) {
	if addr == 0xFF0F {
		m.interruptFlag = value
	}
}

func (m *MockMMU) ReadByte(addr uint16) byte {
	if addr == 0xFF0F {
		return m.interruptFlag
	}
	return 0
}
