package gbcore

import (
	"testing"
)

// TestSoundInitialization verifies that a new Sound system can be created
func TestSoundInitialization(t *testing.T) {
	// Create a new Sound system
	sound := NewSound()

	// Check that the Sound system was initialized
	if sound == nil {
		t.Error("Expected Sound to be initialized, got nil")
	}

	// Check initial state
	if !sound.enabled {
		t.Error("Expected sound to be enabled initially")
	}

	// Check that channels are initialized
	if sound.channel1.enabled {
		t.Error("Expected channel 1 to be disabled initially")
	}

	if sound.channel2.enabled {
		t.Error("Expected channel 2 to be disabled initially")
	}

	if sound.channel3.enabled {
		t.Error("Expected channel 3 to be disabled initially")
	}

	if sound.channel4.enabled {
		t.Error("Expected channel 4 to be disabled initially")
	}
}

// TestSoundReset tests the Sound reset function
func TestSoundReset(t *testing.T) {
	// Create a new Sound system
	sound := NewSound()

	// Modify state
	sound.enabled = false
	sound.channel1.enabled = true
	sound.channel2.enabled = true
	sound.channel3.enabled = true
	sound.channel4.enabled = true

	// Reset the Sound system
	sound.Reset()

	// Check that state was reset
	if !sound.enabled {
		t.Error("Expected sound to be enabled after reset")
	}

	if sound.channel1.enabled {
		t.Error("Expected channel 1 to be disabled after reset")
	}

	if sound.channel2.enabled {
		t.Error("Expected channel 2 to be disabled after reset")
	}

	if sound.channel3.enabled {
		t.Error("Expected channel 3 to be disabled after reset")
	}

	if sound.channel4.enabled {
		t.Error("Expected channel 4 to be disabled after reset")
	}
}

// TestSoundRegisterRead tests reading from sound registers
func TestSoundRegisterRead(t *testing.T) {
	// Create a new Sound system
	sound := NewSound()

	// Test reading NR52 (sound on/off)
	value := sound.ReadRegister(0xFF26)
	if value&0x80 == 0 {
		t.Error("Expected bit 7 of NR52 to be set when sound is enabled")
	}

	// Disable sound
	sound.enabled = false

	// Test reading NR52 again
	value = sound.ReadRegister(0xFF26)
	if value&0x80 != 0 {
		t.Error("Expected bit 7 of NR52 to be clear when sound is disabled")
	}

	// Test reading NR51 (sound panning)
	value = sound.ReadRegister(0xFF25)
	if value != 0xF3 {
		t.Errorf("Expected NR51 default value to be 0xF3, got %02X", value)
	}

	// Test reading NR50 (master volume)
	value = sound.ReadRegister(0xFF24)
	if value != 0x77 {
		t.Errorf("Expected NR50 default value to be 0x77, got %02X", value)
	}
}

// TestChannelTrigger tests triggering sound channels
func TestChannelTrigger(t *testing.T) {
	// Create a new Sound system
	sound := NewSound()

	// Trigger channel 1
	sound.channel1.registers[4] = 0x80 // Set trigger bit
	sound.updateChannel1()

	// Check that channel 1 was enabled
	if !sound.channel1.enabled {
		t.Error("Expected channel 1 to be enabled after trigger")
	}

	// Trigger channel 2
	sound.channel2.registers[4] = 0x80 // Set trigger bit
	sound.updateChannel2()

	// Check that channel 2 was enabled
	if !sound.channel2.enabled {
		t.Error("Expected channel 2 to be enabled after trigger")
	}

	// Trigger channel 3
	sound.channel3.registers[0] = 0x80 // Enable channel 3
	sound.channel3.registers[4] = 0x80 // Set trigger bit
	sound.updateChannel3()

	// Check that channel 3 was enabled
	if !sound.channel3.enabled {
		t.Error("Expected channel 3 to be enabled after trigger")
	}

	// Trigger channel 4
	sound.channel4.registers[4] = 0x80 // Set trigger bit
	sound.updateChannel4()

	// Check that channel 4 was enabled
	if !sound.channel4.enabled {
		t.Error("Expected channel 4 to be enabled after trigger")
	}
}
