package sound

import (
	"log"
)

// Sound handles the Game Boy's audio system
type Sound struct {
	// Sound channels
	channel1 Channel // Tone & Sweep
	channel2 Channel // Tone
	channel3 Channel // Wave Output
	channel4 Channel // Noise

	enabled bool
}

// Channel represents a sound channel
type Channel struct {
	enabled bool

	// Channel-specific registers
	registers [5]byte

	// Sound generation state
	frequency    uint16
	lengthTimer  byte
	volumeTimer  byte
	sweepTimer   byte
	envelopeStep byte

	// Output state
	outputVolume byte
	outputValue  byte
}

// Initialize a new Sound system
func NewSound() *Sound {
	sound := &Sound{
		enabled: true,
	}

	// Initialize channels
	sound.channel1 = Channel{enabled: false}
	sound.channel2 = Channel{enabled: false}
	sound.channel3 = Channel{enabled: false}
	sound.channel4 = Channel{enabled: false}

	return sound
}

// Reset the sound system
func (s *Sound) Reset() {
	s.enabled = true

	// Reset channels
	s.channel1 = Channel{enabled: false}
	s.channel2 = Channel{enabled: false}
	s.channel3 = Channel{enabled: false}
	s.channel4 = Channel{enabled: false}
}

// Step advances the sound system by the specified number of cycles
func (s *Sound) Step(cycles int) {
	// TODO: Implement sound generation
	// This is a placeholder for now
}

// Read a sound register
func (s *Sound) ReadRegister(addr uint16) byte {
	// Sound registers are from 0xFF10 to 0xFF3F
	regAddr := addr - 0xFF10

	switch {
	case regAddr < 0x05:
		// Channel 1 - Tone & Sweep
		return s.channel1.registers[regAddr]
	case regAddr < 0x0A:
		// Channel 2 - Tone
		return s.channel2.registers[regAddr-0x05]
	case regAddr < 0x0F:
		// Channel 3 - Wave Output
		return s.channel3.registers[regAddr-0x0A]
	case regAddr < 0x14:
		// Channel 4 - Noise
		return s.channel4.registers[regAddr-0x0F]
	case regAddr == 0x14:
		// NR50 - Channel control / ON-OFF / Volume (R/W)
		return 0x77 // Default value
	case regAddr == 0x15:
		// NR51 - Selection of Sound output terminal (R/W)
		return 0xF3 // Default value
	case regAddr == 0x16:
		// NR52 - Sound on/off
		value := byte(0x70) // Bits 4-6 are unused and always return 1
		if s.enabled {
			value |= 0x80
		}
		if s.channel1.enabled {
			value |= 0x01
		}
		if s.channel2.enabled {
			value |= 0x02
		}
		if s.channel3.enabled {
			value |= 0x04
		}
		if s.channel4.enabled {
			value |= 0x08
		}
		return value
	case regAddr >= 0x30 && regAddr < 0x40:
		// Wave Pattern RAM
		return 0xFF // Default value
	default:
		return 0xFF
	}
}

// Write a sound register
func (s *Sound) WriteRegister(addr uint16, value byte) {
	// Sound registers are from 0xFF10 to 0xFF3F
	regAddr := addr - 0xFF10

	// If sound is disabled, only NR52 can be written
	if !s.enabled && regAddr != 0x16 {
		return
	}

	switch {
	case regAddr < 0x05:
		// Channel 1 - Tone & Sweep
		s.channel1.registers[regAddr] = value
		s.updateChannel1()
	case regAddr < 0x0A:
		// Channel 2 - Tone
		s.channel2.registers[regAddr-0x05] = value
		s.updateChannel2()
	case regAddr < 0x0F:
		// Channel 3 - Wave Output
		s.channel3.registers[regAddr-0x0A] = value
		s.updateChannel3()
	case regAddr < 0x14:
		// Channel 4 - Noise
		s.channel4.registers[regAddr-0x0F] = value
		s.updateChannel4()
	case regAddr == 0x14:
		// NR50 - Channel control / ON-OFF / Volume (R/W)
		// TODO: Implement volume control
	case regAddr == 0x15:
		// NR51 - Selection of Sound output terminal (R/W)
		// TODO: Implement panning
	case regAddr == 0x16:
		// NR52 - Sound on/off
		s.enabled = (value & 0x80) != 0
		if !s.enabled {
			// Disable all sound channels when sound is turned off
			s.Reset()
		}
	case regAddr >= 0x30 && regAddr < 0x40:
		// Wave Pattern RAM
		// TODO: Implement wave pattern RAM
	}
}

// Update Channel 1 (Tone & Sweep)
func (s *Sound) updateChannel1() {
	// Check if channel was triggered
	if (s.channel1.registers[4] & 0x80) != 0 {
		s.channel1.enabled = true

		// Reset length timer if needed
		if s.channel1.lengthTimer == 0 {
			s.channel1.lengthTimer = 64
		}

		// Set frequency
		s.channel1.frequency = uint16(s.channel1.registers[3]&0x07)<<8 | uint16(s.channel1.registers[2])

		// Set volume
		s.channel1.outputVolume = (s.channel1.registers[2] >> 4) & 0x0F

		// Reset envelope
		s.channel1.envelopeStep = 0

		// Reset sweep
		s.channel1.sweepTimer = (s.channel1.registers[0] >> 4) & 0x07
		if s.channel1.sweepTimer == 0 {
			s.channel1.sweepTimer = 8
		}

		log.Printf("[Sound] Channel 1 triggered, freq=%d, vol=%d", s.channel1.frequency, s.channel1.outputVolume)
	}
}

// Update Channel 2 (Tone)
func (s *Sound) updateChannel2() {
	// Check if channel was triggered
	if (s.channel2.registers[4] & 0x80) != 0 {
		s.channel2.enabled = true

		// Reset length timer if needed
		if s.channel2.lengthTimer == 0 {
			s.channel2.lengthTimer = 64
		}

		// Set frequency
		s.channel2.frequency = uint16(s.channel2.registers[3]&0x07)<<8 | uint16(s.channel2.registers[2])

		// Set volume
		s.channel2.outputVolume = (s.channel2.registers[2] >> 4) & 0x0F

		// Reset envelope
		s.channel2.envelopeStep = 0

		log.Printf("[Sound] Channel 2 triggered, freq=%d, vol=%d", s.channel2.frequency, s.channel2.outputVolume)
	}
}

// Update Channel 3 (Wave Output)
func (s *Sound) updateChannel3() {
	// Check if channel was triggered
	if (s.channel3.registers[4] & 0x80) != 0 {
		s.channel3.enabled = true

		// Check if channel is enabled
		if (s.channel3.registers[0] & 0x80) == 0 {
			s.channel3.enabled = false
		}

		// Reset length timer if needed
		if s.channel3.lengthTimer == 0 {
			s.channel3.lengthTimer = 255 // Max value for byte
		}

		// Set frequency
		s.channel3.frequency = uint16(s.channel3.registers[3]&0x07)<<8 | uint16(s.channel3.registers[2])

		// Set volume
		volumeCode := (s.channel3.registers[2] >> 5) & 0x03
		switch volumeCode {
		case 0:
			s.channel3.outputVolume = 0
		case 1:
			s.channel3.outputVolume = 15
		case 2:
			s.channel3.outputVolume = 7
		case 3:
			s.channel3.outputVolume = 3
		}

		log.Printf("[Sound] Channel 3 triggered, freq=%d, vol=%d", s.channel3.frequency, s.channel3.outputVolume)
	}
}

// Update Channel 4 (Noise)
func (s *Sound) updateChannel4() {
	// Check if channel was triggered
	if (s.channel4.registers[4] & 0x80) != 0 {
		s.channel4.enabled = true

		// Reset length timer if needed
		if s.channel4.lengthTimer == 0 {
			s.channel4.lengthTimer = 64
		}

		// Set volume
		s.channel4.outputVolume = (s.channel4.registers[2] >> 4) & 0x0F

		// Reset envelope
		s.channel4.envelopeStep = 0

		log.Printf("[Sound] Channel 4 triggered, vol=%d", s.channel4.outputVolume)
	}
}
