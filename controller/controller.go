package controller

// Generic interface for handling user input for controlling the game
type Controller interface {
	// Initializes the controller
	Init()

	// Gets input from user and handles it. Returns true if input detected, and
	// false if no input detected
	Update() bool
}
