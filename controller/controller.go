package controller

// Game Boy button bits
const (
	BUTTON_RIGHT  = 0x01
	BUTTON_LEFT   = 0x02
	BUTTON_UP     = 0x04
	BUTTON_DOWN   = 0x08
	BUTTON_A      = 0x10
	BUTTON_B      = 0x20
	BUTTON_SELECT = 0x40
	BUTTON_START  = 0x80
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
}
