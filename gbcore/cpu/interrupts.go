package gbcore

// Interrupt flags
const (
	INT_VBLANK = 0x01 // V-Blank interrupt (bit 0)
	INT_LCDC   = 0x02 // LCDC Status interrupt (bit 1)
	INT_TIMER  = 0x04 // Timer Overflow interrupt (bit 2)
	INT_SERIAL = 0x08 // Serial Transfer Completion interrupt (bit 3)
	INT_JOYPAD = 0x10 // Joypad interrupt (bit 4)
)

// Interrupt vectors
const (
	VBLANK_VECTOR = 0x0040 // V-Blank interrupt vector
	LCDC_VECTOR   = 0x0048 // LCDC Status interrupt vector
	TIMER_VECTOR  = 0x0050 // Timer Overflow interrupt vector
	SERIAL_VECTOR = 0x0058 // Serial Transfer Completion interrupt vector
	JOYPAD_VECTOR = 0x0060 // Joypad interrupt vector
)

// Handle interrupts
func (cpu *Z80) handleInterrupts() {
	// Get interrupt flags (IF) and interrupt enable (IE)
	interruptFlag := cpu.mmu.ReadByte(0xFF0F)
	interruptEnable := cpu.mmu.ReadByte(0xFFFF)

	// Calculate pending interrupts
	pendingInterrupts := interruptFlag & interruptEnable & 0x1F

	// If there are no pending interrupts, return
	if pendingInterrupts == 0 {
		return
	}

	// Disable interrupts
	cpu.interruptMaster = false

	// Handle interrupts in priority order
	if pendingInterrupts&INT_VBLANK != 0 {
		// Clear the interrupt flag
		cpu.mmu.WriteByte(0xFF0F, interruptFlag&^INT_VBLANK)

		// Call the interrupt handler
		cpu.callInterrupt(VBLANK_VECTOR)

	} else if pendingInterrupts&INT_LCDC != 0 {
		// Clear the interrupt flag
		cpu.mmu.WriteByte(0xFF0F, interruptFlag&^INT_LCDC)

		// Call the interrupt handler
		cpu.callInterrupt(LCDC_VECTOR)

	} else if pendingInterrupts&INT_TIMER != 0 {
		// Clear the interrupt flag
		cpu.mmu.WriteByte(0xFF0F, interruptFlag&^INT_TIMER)

		// Call the interrupt handler
		cpu.callInterrupt(TIMER_VECTOR)

	} else if pendingInterrupts&INT_SERIAL != 0 {
		// Clear the interrupt flag
		cpu.mmu.WriteByte(0xFF0F, interruptFlag&^INT_SERIAL)

		// Call the interrupt handler
		cpu.callInterrupt(SERIAL_VECTOR)

	} else if pendingInterrupts&INT_JOYPAD != 0 {
		// Clear the interrupt flag
		cpu.mmu.WriteByte(0xFF0F, interruptFlag&^INT_JOYPAD)

		// Call the interrupt handler
		cpu.callInterrupt(JOYPAD_VECTOR)
	}
}

// Call an interrupt handler
func (cpu *Z80) callInterrupt(vector uint16) {
	// Push PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to interrupt handler
	cpu.reg.PC = vector
}
