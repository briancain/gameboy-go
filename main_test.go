package main

import (
	"flag"
	"os"
	"testing"
)

// TestFlagParsing tests that command line flags are parsed correctly
func TestFlagParsing(t *testing.T) {
	// Save original flag values
	origCartridgePath := CartridgePath
	origHelp := Help
	origDebugOutput := DebugOutput
	origScale := Scale
	origHeadless := Headless

	// Save original os.Args
	origArgs := os.Args

	// Restore original values when test completes
	defer func() {
		CartridgePath = origCartridgePath
		Help = origHelp
		DebugOutput = origDebugOutput
		Scale = origScale
		Headless = origHeadless
		os.Args = origArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}()

	// Set up test flags
	os.Args = []string{"gameboy-go", "-rom-file=test.gb", "-debug", "-scale=3", "-headless"}

	// Reset flags for testing
	CartridgePath = ""
	Help = false
	DebugOutput = false
	Scale = 2
	Headless = false

	// Re-initialize flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.BoolVar(&Help, "help", false, "Displays help")
	flag.StringVar(&CartridgePath, "rom-file", "", "A path to a cartridge ROM file")
	flag.BoolVar(&DebugOutput, "debug", false, "Displays debug output")
	flag.IntVar(&Scale, "scale", 2, "Screen scale factor (1-4)")
	flag.BoolVar(&Headless, "headless", false, "Run without display (for testing)")

	// Parse flags
	flag.Parse()

	// Check that flags were parsed correctly
	if CartridgePath != "test.gb" {
		t.Errorf("Expected CartridgePath to be 'test.gb', got '%s'", CartridgePath)
	}

	if !DebugOutput {
		t.Errorf("Expected DebugOutput to be true")
	}

	if Scale != 3 {
		t.Errorf("Expected Scale to be 3, got %d", Scale)
	}

	if !Headless {
		t.Errorf("Expected Headless to be true")
	}
}
