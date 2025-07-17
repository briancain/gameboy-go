package gbcore

import (
	"testing"
)

// TestGameBoyCoreInitialization verifies that a new GameBoyCore can be created
func TestGameBoyCoreInitialization(t *testing.T) {
	// Create a new GameBoyCore
	gb, err := NewGameBoyCore(true) // With debug enabled

	// Check that there was no error
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that the GameBoyCore was initialized
	if gb == nil {
		t.Error("Expected GameBoyCore to be initialized, got nil")
	}

	// Check debug flag
	if !gb.debug {
		t.Error("Expected debug flag to be true")
	}

	// Check FPS
	if gb.FPS != 60 {
		t.Errorf("Expected FPS to be 60, got %d", gb.FPS)
	}
}

// TestGameBoyCoreExit tests the Exit method
func TestGameBoyCoreExit(t *testing.T) {
	// Create a new GameBoyCore
	gb, _ := NewGameBoyCore(false)

	// Check initial exit flag
	if gb.exit {
		t.Error("Expected exit flag to be false initially")
	}

	// Call Exit
	gb.Exit()

	// Check that exit flag was set
	if !gb.exit {
		t.Error("Expected exit flag to be true after calling Exit")
	}
}
