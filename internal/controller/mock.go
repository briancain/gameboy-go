package controller

import "fmt"

// A mock controller for testing
type MockController struct {
	// Button state for direction buttons (P14 low)
	dpadState byte

	// Button state for action buttons (P15 low)
	buttonState byte

	// Joypad register state
	joypadRegister byte
}

// Create a new mock controller
func NewMockController() *MockController {
	return &MockController{
		dpadState:      0,
		buttonState:    0,
		joypadRegister: 0xFF,
	}
}

func (m *MockController) Init() {
	// No initialization needed for mock
}

func (m *MockController) Update() bool {
	// No update needed for mock
	return false
}

// Get the current button state
func (m *MockController) GetButtonState() byte {
	// This is a combined state for testing
	return (m.buttonState << 4) | m.dpadState
}

// Set the button state (for testing)
func (m *MockController) SetButtonState(state byte) {
	// Split the state into direction and action buttons
	m.dpadState = state & 0x0F
	m.buttonState = (state >> 4) & 0x0F
}

// Process a joypad register write
func (m *MockController) WriteJoypad(value byte) {
	// Only bits 4-5 are writable
	m.joypadRegister = value
	fmt.Printf("WriteJoypad: value=0x%02X, joypadRegister=0x%02X\n", value, m.joypadRegister)
}

// Read the joypad register
func (m *MockController) ReadJoypad() byte {
	// Start with all bits set except 4-5 which come from the register
	result := byte(0xCF) | (m.joypadRegister & 0x30)

	// If bit 4 is clear, we're selecting direction buttons (P14 low)
	if (m.joypadRegister & 0x10) == 0 {
		// Check if any direction buttons are pressed
		if (m.dpadState & 0x01) != 0 {
			result &= ^byte(0x01) // Clear bit 0
		}
		if (m.dpadState & 0x02) != 0 {
			result &= ^byte(0x02) // Clear bit 1
		}
		if (m.dpadState & 0x04) != 0 {
			result &= ^byte(0x04) // Clear bit 2
		}
		if (m.dpadState & 0x08) != 0 {
			result &= ^byte(0x08) // Clear bit 3
		}
	}

	// If bit 5 is clear, we're selecting action buttons (P15 low)
	if (m.joypadRegister & 0x20) == 0 {
		// Check if any action buttons are pressed
		if (m.buttonState & 0x01) != 0 {
			result &= ^byte(0x01) // Clear bit 0
		}
		if (m.buttonState & 0x02) != 0 {
			result &= ^byte(0x02) // Clear bit 1
		}
		if (m.buttonState & 0x04) != 0 {
			result &= ^byte(0x04) // Clear bit 2
		}
		if (m.buttonState & 0x08) != 0 {
			result &= ^byte(0x08) // Clear bit 3
		}
	}

	fmt.Printf("ReadJoypad: joypadRegister=0x%02X, dpadState=0x%02X, buttonState=0x%02X, result=0x%02X\n",
		m.joypadRegister, m.dpadState, m.buttonState, result)

	return result
}

// Check if a joypad interrupt should be triggered
func (m *MockController) CheckInterrupt() bool {
	// For testing, always return false
	return false
}

// Clean up resources
func (m *MockController) Cleanup() {
	// No cleanup needed for mock
}
