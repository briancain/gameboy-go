//go:build !sdl
// +build !sdl

package controller

import (
	"log"
)

// A stub keyboard controller for non-SDL builds
type Keyboard struct {
	// Current button state (1 = pressed)
	buttonState byte

	// Joypad register state
	joypadRegister byte

	// Previous button state for detecting changes
	prevButtonState byte
}

// Create a new keyboard controller
func NewKeyboard() *Keyboard {
	return &Keyboard{
		buttonState:     0,
		joypadRegister:  0xFF,
		prevButtonState: 0,
	}
}

func (k *Keyboard) Init() {
	log.Println("[DEBUG] Initializing stub Keyboard controller...")
}

func (k *Keyboard) Update() bool {
	// No update in stub implementation
	return false
}

// Get the current button state
func (k *Keyboard) GetButtonState() byte {
	return k.buttonState
}

// Process a joypad register write
func (k *Keyboard) WriteJoypad(value byte) {
	// Only bits 4-5 are writable
	k.joypadRegister = (k.joypadRegister & 0xCF) | (value & 0x30)
}

// Read the joypad register
func (k *Keyboard) ReadJoypad() byte {
	result := k.joypadRegister | 0xCF // Set bits 0-3 and 6-7 to 1

	// Check which button type is selected
	if (k.joypadRegister & 0x10) == 0 {
		// Direction buttons
		if (k.buttonState & BUTTON_RIGHT) != 0 {
			result &= ^byte(0x01)
		}
		if (k.buttonState & BUTTON_LEFT) != 0 {
			result &= ^byte(0x02)
		}
		if (k.buttonState & BUTTON_UP) != 0 {
			result &= ^byte(0x04)
		}
		if (k.buttonState & BUTTON_DOWN) != 0 {
			result &= ^byte(0x08)
		}
	}

	if (k.joypadRegister & 0x20) == 0 {
		// Action buttons
		if (k.buttonState & BUTTON_A) != 0 {
			result &= ^byte(0x01)
		}
		if (k.buttonState & BUTTON_B) != 0 {
			result &= ^byte(0x02)
		}
		if (k.buttonState & BUTTON_SELECT) != 0 {
			result &= ^byte(0x04)
		}
		if (k.buttonState & BUTTON_START) != 0 {
			result &= ^byte(0x08)
		}
	}

	return result
}

// CheckInterrupt checks if a joypad interrupt should be triggered
// Returns true if an interrupt should be triggered
func (k *Keyboard) CheckInterrupt() bool {
	// No interrupt in stub implementation
	return false
}

// Clean up resources
func (k *Keyboard) Cleanup() {
	// No cleanup needed for stub
}
