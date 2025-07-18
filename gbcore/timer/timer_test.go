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
	timer.prevTimerOn = true

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

	if timer.prevTimerOn != false {
		t.Errorf("Expected prevTimerOn to be reset to false, got %v", timer.prevTimerOn)
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
	timer.tac = TAC_ENABLE // Timer enabled, 4096Hz

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
		{TAC_ENABLE, 1024, 1},       // 4096Hz (1024 cycles per increment)
		{TAC_ENABLE | 0x01, 16, 1},  // 262144Hz (16 cycles per increment)
		{TAC_ENABLE | 0x02, 64, 1},  // 65536Hz (64 cycles per increment)
		{TAC_ENABLE | 0x03, 256, 1}, // 16384Hz (256 cycles per increment)
		{TAC_ENABLE, 2048, 2},       // 4096Hz (2048 cycles = 2 increments)
		{TAC_ENABLE | 0x01, 32, 2},  // 262144Hz (32 cycles = 2 increments)
		{TAC_ENABLE | 0x02, 128, 2}, // 65536Hz (128 cycles = 2 increments)
		{TAC_ENABLE | 0x03, 512, 2}, // 16384Hz (512 cycles = 2 increments)
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

// TestTimerEnableEdge tests the edge-triggered behavior when enabling the timer
func TestTimerEnableEdge(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Set initial state (timer disabled)
	timer.tac = 0x00
	timer.timaCounter = 500 // Set a non-zero counter

	// Enable timer
	timer.WriteRegister(0xFF07, TAC_ENABLE)

	// Check that timaCounter was reset
	if timer.timaCounter != 0 {
		t.Errorf("Expected timaCounter to be reset when timer is enabled, got %d", timer.timaCounter)
	}

	// Check that timer is now enabled
	if (timer.tac & TAC_ENABLE) == 0 {
		t.Error("Expected timer to be enabled after writing TAC_ENABLE")
	}
}

// TestMultipleOverflows tests multiple TIMA overflows in a single step
func TestMultipleOverflows(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new Timer
	timer := NewTimer(mockMMU)

	// Enable timer with 262144Hz frequency (16 cycles per increment)
	timer.tac = TAC_ENABLE | 0x01
	timer.tima = 0xFE // Almost at overflow
	timer.tma = 0x42  // Modulo value

	// Step the timer by 32 cycles (should cause 2 increments, 1 overflow)
	timer.Step(32)

	// Check that TIMA has the correct value
	// 0xFE + 2 = 0x100, which overflows once and becomes 0x42
	if timer.tima != 0x42 {
		t.Errorf("Expected TIMA to be 0x42 after overflow, got %02X", timer.tima)
	}

	// Check that interrupt was requested
	if mockMMU.interruptFlag&0x04 == 0 {
		t.Error("Expected timer interrupt to be requested")
	}

	// Reset interrupt flag
	mockMMU.interruptFlag = 0

	// Step again to test another overflow
	timer.Step(16)

	// Check that TIMA has the correct value
	// 0x42 + 1 = 0x43
	if timer.tima != 0x43 {
		t.Errorf("Expected TIMA to be 0x43 after second step, got %02X", timer.tima)
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
