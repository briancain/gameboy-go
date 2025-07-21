package controller

import (
	"testing"
)

func TestKeyboardInitialization(t *testing.T) {
	// Create a new keyboard controller
	keyboard := NewKeyboard()

	// Initialize the controller
	keyboard.Init()

	// Test initial state
	if keyboard.GetButtonState() != 0 {
		t.Errorf("Initial button state should be 0, got %d", keyboard.GetButtonState())
	}
}

func TestJoypadRegister(t *testing.T) {
	// Create a new keyboard controller
	keyboard := NewKeyboard()

	// Test initial joypad register value
	if keyboard.ReadJoypad() != 0xFF {
		t.Errorf("Initial joypad register should be 0xFF, got %02X", keyboard.ReadJoypad())
	}

	// Test writing to joypad register
	keyboard.WriteJoypad(0x10)

	// Test reading from joypad register
	if keyboard.ReadJoypad() == 0xFF {
		t.Errorf("Joypad register should not be 0xFF after writing")
	}
}
