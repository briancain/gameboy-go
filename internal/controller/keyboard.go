//go:build sdl
// +build sdl

package controller

import (
	"log"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

// A keyboard controller
type Keyboard struct {
	// Current button state (1 = pressed)
	buttonState byte

	// Joypad register state
	joypadRegister byte

	// Previous button state for detecting changes
	prevButtonState byte

	// Mutex for thread safety
	mutex sync.Mutex

	// Key mapping
	keyMap map[sdl.Keycode]byte

	// Flag to indicate if SDL is initialized
	sdlInitialized bool
}

// Create a new keyboard controller
func NewKeyboard() *Keyboard {
	return &Keyboard{
		buttonState:     0,
		joypadRegister:  0xFF,
		prevButtonState: 0,
		keyMap:          make(map[sdl.Keycode]byte),
		sdlInitialized:  false,
	}
}

func (k *Keyboard) Init() {
	log.Println("[DEBUG] Initializing Keyboard controller...")

	// Initialize SDL for keyboard input
	if err := sdl.Init(sdl.INIT_EVENTS); err != nil {
		log.Printf("[ERROR] Failed to initialize SDL: %v", err)
		return
	}

	k.sdlInitialized = true

	// Set up default key mapping
	k.keyMap[sdl.K_RIGHT] = BUTTON_RIGHT
	k.keyMap[sdl.K_LEFT] = BUTTON_LEFT
	k.keyMap[sdl.K_UP] = BUTTON_UP
	k.keyMap[sdl.K_DOWN] = BUTTON_DOWN
	k.keyMap[sdl.K_z] = BUTTON_A
	k.keyMap[sdl.K_x] = BUTTON_B
	k.keyMap[sdl.K_RETURN] = BUTTON_START
	k.keyMap[sdl.K_SPACE] = BUTTON_SELECT

	log.Println("[DEBUG] Keyboard controller initialized")
}

func (k *Keyboard) Update() bool {
	if !k.sdlInitialized {
		return false
	}

	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Store previous button state
	k.prevButtonState = k.buttonState

	// Process SDL events
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.KeyboardEvent:
			if button, ok := k.keyMap[e.Keysym.Sym]; ok {
				if e.Type == sdl.KEYDOWN {
					k.buttonState |= button
				} else if e.Type == sdl.KEYUP {
					k.buttonState &= ^button
				}
			}
		case *sdl.QuitEvent:
			// Handle quit event if needed
		}
	}

	// Return true if button state changed
	return k.buttonState != k.prevButtonState
}

// Get the current button state
func (k *Keyboard) GetButtonState() byte {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	return k.buttonState
}

// SetButtonState sets the state of a specific button (for external input handling)
func (k *Keyboard) SetButtonState(button string, pressed bool) {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Store previous state for interrupt detection
	k.prevButtonState = k.buttonState

	// Map button names to button bits
	var buttonBit byte
	switch button {
	case "up":
		buttonBit = BUTTON_UP
	case "down":
		buttonBit = BUTTON_DOWN
	case "left":
		buttonBit = BUTTON_LEFT
	case "right":
		buttonBit = BUTTON_RIGHT
	case "a":
		buttonBit = BUTTON_A
	case "b":
		buttonBit = BUTTON_B
	case "select":
		buttonBit = BUTTON_SELECT
	case "start":
		buttonBit = BUTTON_START
	default:
		return // Unknown button
	}

	// Set or clear the button bit
	if pressed {
		k.buttonState |= buttonBit
	} else {
		k.buttonState &= ^buttonBit
	}
}

// Process a joypad register write
func (k *Keyboard) WriteJoypad(value byte) {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Only bits 4-5 are writable
	k.joypadRegister = (k.joypadRegister & 0xCF) | (value & 0x30)
}

// Read the joypad register
func (k *Keyboard) ReadJoypad() byte {
	k.mutex.Lock()
	defer k.mutex.Unlock()

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
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Check if any button has transitioned from not pressed to pressed
	buttonTransition := (k.buttonState & ^k.prevButtonState) != 0

	// Check if the joypad register is selecting the buttons that changed
	directionSelected := (k.joypadRegister & 0x10) == 0
	actionSelected := (k.joypadRegister & 0x20) == 0

	// Check if any direction buttons changed and are selected
	directionChange := (k.buttonState & ^k.prevButtonState & 0x0F) != 0 && directionSelected

	// Check if any action buttons changed and are selected
	actionChange := (k.buttonState & ^k.prevButtonState & 0xF0) != 0 && actionSelected

	// Return true if any selected buttons changed from not pressed to pressed
	return buttonTransition && (directionChange || actionChange)
}

// Clean up resources
func (k *Keyboard) Cleanup() {
	if k.sdlInitialized {
		sdl.Quit()
		k.sdlInitialized = false
	}
}
