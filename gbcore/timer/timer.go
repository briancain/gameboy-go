package gbcore

// Timer handles the Game Boy's timing system
// According to the Game Boy CPU manual section 2.10:
// - DIV register increments at 16384Hz
// - TIMA register increments at a frequency selected by TAC
// - When TIMA overflows, it's reset to TMA and a timer interrupt is requested
type Timer struct {
	// Timer registers
	div  byte // Divider Register (FF04) - Increments at 16384Hz
	tima byte // Timer Counter (FF05) - Increments at frequency selected by TAC
	tma  byte // Timer Modulo (FF06) - Loaded into TIMA when it overflows
	tac  byte // Timer Control (FF07) - Controls timer enable and frequency

	// Internal counters
	divCounter  int  // Counter for DIV register
	timaCounter int  // Counter for TIMA register
	prevTimerOn bool // Previous state of timer enable bit

	// Reference to MMU for memory access
	mmu MMU
}

// TAC register bits
const (
	TAC_ENABLE    = 0x04 // Bit 2 - Timer Enable
	TAC_FREQ_MASK = 0x03 // Bits 0-1 - Input Clock Select
)

// Timer frequencies in Hz
const (
	FREQ_4096   = 4096   // 00: 4096Hz (1024 cycles per increment)
	FREQ_262144 = 262144 // 01: 262144Hz (16 cycles per increment)
	FREQ_65536  = 65536  // 10: 65536Hz (64 cycles per increment)
	FREQ_16384  = 16384  // 11: 16384Hz (256 cycles per increment)
)

// CPU clock speed in Hz
const CPU_CLOCK = 4194304

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
		prevTimerOn: false,
		mmu:         mmu,
	}

	return timer
}

// Reset the timer to its initial state
func (t *Timer) Reset() {
	t.div = 0
	t.tima = 0
	t.tma = 0
	t.tac = 0
	t.divCounter = 0
	t.timaCounter = 0
	t.prevTimerOn = false
}

// Step advances the timer by the specified number of cycles
func (t *Timer) Step(cycles int) {
	// Update DIV register (increments at 16384Hz)
	// 4194304Hz / 16384Hz = 256 cycles
	t.divCounter += cycles
	for t.divCounter >= 256 {
		t.div++
		t.divCounter -= 256
	}

	// Check if timer is enabled
	timerOn := (t.tac & TAC_ENABLE) != 0

	// Get timer frequency
	var cyclesPerIncrement int
	switch t.tac & TAC_FREQ_MASK {
	case 0:
		cyclesPerIncrement = CPU_CLOCK / FREQ_4096 // 1024 cycles
	case 1:
		cyclesPerIncrement = CPU_CLOCK / FREQ_262144 // 16 cycles
	case 2:
		cyclesPerIncrement = CPU_CLOCK / FREQ_65536 // 64 cycles
	case 3:
		cyclesPerIncrement = CPU_CLOCK / FREQ_16384 // 256 cycles
	}

	// Update TIMA if timer is enabled
	if timerOn {
		t.timaCounter += cycles

		// Increment TIMA when the counter reaches the threshold
		for t.timaCounter >= cyclesPerIncrement {
			// Increment TIMA
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

	// Store current timer state for edge detection
	t.prevTimerOn = timerOn
}

// Request a timer interrupt
func (t *Timer) requestInterrupt() {
	// Set bit 2 of the IF register (0xFF0F)
	interruptFlag := t.mmu.ReadByte(0xFF0F)
	interruptFlag |= 0x04 // Timer interrupt (bit 2)
	t.mmu.WriteByte(0xFF0F, interruptFlag)
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
		// Check for edge-triggered behavior when changing timer enable
		oldTac := t.tac
		t.tac = value & 0x07 // Only bits 0-2 are used

		// If timer was just enabled, reset the TIMA counter
		if ((oldTac & TAC_ENABLE) == 0) && ((t.tac & TAC_ENABLE) != 0) {
			t.timaCounter = 0
		}
	}
}
