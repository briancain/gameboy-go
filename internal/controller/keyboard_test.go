package controller

import (
	"testing"
)

// TestKeyboardInitialization verifies that a new keyboard controller can be created
func TestKeyboardInitialization(t *testing.T) {
	// Create a new keyboard controller
	keyboard := NewKeyboard()

	// Check that the keyboard was initialized
	if keyboard == nil {
		t.Error("Expected keyboard to be initialized, got nil")
	}

	// Initialize the keyboard
	keyboard.Init()

	// Test button state
	if keyboard.GetButtonState() != 0 {
		t.Errorf("Expected initial button state to be 0, got %d", keyboard.GetButtonState())
	}
}

// TestJoypadRegister tests the joypad register read/write operations
func TestJoypadRegister(t *testing.T) {
	// Create a new keyboard controller
	keyboard := NewKeyboard()

	// Test initial joypad register value
	initialValue := keyboard.ReadJoypad()
	if initialValue != 0xFF {
		t.Errorf("Expected initial joypad register to be 0xFF, got %02X", initialValue)
	}

	// Test writing to joypad register
	keyboard.WriteJoypad(0x30) // Select both button types

	// Read back the value
	readValue := keyboard.ReadJoypad()

	// Only bits 4-5 should be affected
	expectedValue := byte(0xCF | 0x30)
	if readValue != expectedValue {
		t.Errorf("Expected joypad register to be %02X, got %02X", expectedValue, readValue)
	}
}
