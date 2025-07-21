package controller

// Game Boy button bits
const (
	BUTTON_RIGHT  = 0x01
	BUTTON_LEFT   = 0x02
	BUTTON_UP     = 0x04
	BUTTON_DOWN   = 0x08
	BUTTON_A      = 0x01
	BUTTON_B      = 0x02
	BUTTON_SELECT = 0x04
	BUTTON_START  = 0x08
)

// Direction buttons (P14 low)
const (
	DPAD_RIGHT = 0x01
	DPAD_LEFT  = 0x02
	DPAD_UP    = 0x04
	DPAD_DOWN  = 0x08
)

// Action buttons (P15 low)
const (
	BTN_A      = 0x01
	BTN_B      = 0x02
	BTN_SELECT = 0x04
	BTN_START  = 0x08
)

// Generic interface for handling user input for controlling the game
type Controller interface {
	// Initializes the controller
	Init()

	// Gets input from user and handles it. Returns true if input detected, and
	// false if no input detected
	Update() bool

	// Get the current button state
	GetButtonState() byte

	// Process a joypad register write
	WriteJoypad(value byte)

	// Read the joypad register
	ReadJoypad() byte

	// Check if a joypad interrupt should be triggered
	CheckInterrupt() bool

	// Clean up resources
	Cleanup()
}
