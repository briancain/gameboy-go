package gbcore

import (
	"log"
)

// Timer handles the Game Boy's timing system
type Timer struct {
	// Timer registers
	div  byte // Divider Register (FF04)
	tima byte // Timer Counter (FF05)
	tma  byte // Timer Modulo (FF06)
	tac  byte // Timer Control (FF07)
	
	// Internal counters
	divCounter  int // Counter for DIV register
	timaCounter int // Counter for TIMA register
	
	// Reference to MMU for memory access
	mmu MMU
}

// MMU interface for Timer to access memory
type MMU interface {
	WriteByte(addr uint16, value byte)
	ReadByte(addr uint16) byte
}

// Initialize a new Timer
func NewTimer(mmu MMU) *Timer {
	timer := &Timer{
		div:         0,
		tima:        0,
		tma:         0,
		tac:         0,
		divCounter:  0,
		timaCounter: 0,
		mmu:         mmu,
	}
	
	return timer
}

// Reset the timer
func (t *Timer) Reset() {
	t.div = 0
	t.tima = 0
	t.tma = 0
	t.tac = 0
	t.divCounter = 0
	t.timaCounter = 0
}

// Step advances the timer by the specified number of cycles
func (t *Timer) Step(cycles int) {
	// Update DIV register (increments at 16384Hz)
	// 4194304Hz / 16384Hz = 256 cycles
	t.divCounter += cycles
	if t.divCounter >= 256 {
		t.div++
		t.divCounter -= 256
	}
	
	// Check if timer is enabled
	if (t.tac & 0x04) != 0 {
		// Update TIMA register
		t.timaCounter += cycles
		
		// Get timer frequency
		var timerFreq int
		switch t.tac & 0x03 {
		case 0:
			timerFreq = 4096 // 4096Hz (1024 cycles)
		case 1:
			timerFreq = 262144 // 262144Hz (16 cycles)
		case 2:
			timerFreq = 65536 // 65536Hz (64 cycles)
		case 3:
			timerFreq = 16384 // 16384Hz (256 cycles)
		}
		
		// Calculate cycles per increment
		cyclesPerIncrement := 4194304 / timerFreq
		
		// Update TIMA
		if t.timaCounter >= cyclesPerIncrement {
			t.tima++
			t.timaCounter -= cyclesPerIncrement
			
			// Check for TIMA overflow
			if t.tima == 0 {
				// Reset TIMA to TMA
				t.tima = t.tma
				
				// Request timer interrupt
				t.requestInterrupt()
			}
		}
	}
}

// Request a timer interrupt
func (t *Timer) requestInterrupt() {
	// Set bit 2 of the IF register (0xFF0F)
	t.mmu.WriteByte(0xFF0F, t.mmu.ReadByte(0xFF0F) | 0x04)
}

// Read a timer register
func (t *Timer) ReadRegister(addr uint16) byte {
	switch addr {
	case 0xFF04:
		return t.div
	case 0xFF05:
		return t.tima
	case 0xFF06:
		return t.tma
	case 0xFF07:
		return t.tac
	default:
		return 0xFF
	}
}

// Write a timer register
func (t *Timer) WriteRegister(addr uint16, value byte) {
	switch addr {
	case 0xFF04:
		// Writing to DIV resets it to 0
		t.div = 0
		t.divCounter = 0
	case 0xFF05:
		t.tima = value
	case 0xFF06:
		t.tma = value
	case 0xFF07:
		t.tac = value
	}
}
