package gbcore

import (
	"testing"
)

func TestInterruptHandling(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Set up a simple interrupt handler at the V-Blank vector
	mockMMU.WriteByte(VBLANK_VECTOR, 0xC9) // RET instruction

	// Enable interrupts
	cpu.interruptMaster = true

	// Set the interrupt enable register to enable V-Blank interrupts
	mockMMU.WriteByte(0xFFFF, INT_VBLANK)

	// Set the interrupt flag register to trigger a V-Blank interrupt
	mockMMU.WriteByte(0xFF0F, INT_VBLANK)

	// Update the pendingInterrupts flag
	cpu.pendingInterrupts = INT_VBLANK

	// Save the current PC
	oldPC := cpu.reg.PC

	// Step the CPU
	cpu.Step()

	// Check that the PC was changed to the V-Blank vector
	if cpu.reg.PC != VBLANK_VECTOR {
		t.Errorf("Expected PC to be %04X (V-Blank vector), got %04X", VBLANK_VECTOR, cpu.reg.PC)
	}

	// Check that interrupts were disabled
	if cpu.interruptMaster {
		t.Error("Expected interrupts to be disabled after handling an interrupt")
	}

	// Check that the interrupt flag was cleared
	if mockMMU.ReadByte(0xFF0F)&INT_VBLANK != 0 {
		t.Error("Expected V-Blank interrupt flag to be cleared")
	}

	// Check that the return address was pushed onto the stack
	if mockMMU.ReadWord(cpu.reg.SP) != oldPC {
		t.Errorf("Expected return address on stack to be %04X, got %04X", oldPC, mockMMU.ReadWord(cpu.reg.SP))
	}
}

func TestInterruptPriority(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Set up simple interrupt handlers
	mockMMU.WriteByte(VBLANK_VECTOR, 0xC9) // RET instruction
	mockMMU.WriteByte(LCDC_VECTOR, 0xC9)   // RET instruction
	mockMMU.WriteByte(TIMER_VECTOR, 0xC9)  // RET instruction

	// Enable interrupts
	cpu.interruptMaster = true

	// Set the interrupt enable register to enable all interrupts
	mockMMU.WriteByte(0xFFFF, INT_VBLANK|INT_LCDC|INT_TIMER)

	// Test 1: V-Blank has highest priority
	mockMMU.WriteByte(0xFF0F, INT_VBLANK|INT_LCDC|INT_TIMER)
	cpu.pendingInterrupts = INT_VBLANK | INT_LCDC | INT_TIMER
	cpu.Step()
	if cpu.reg.PC != VBLANK_VECTOR {
		t.Errorf("Expected PC to be %04X (V-Blank vector), got %04X", VBLANK_VECTOR, cpu.reg.PC)
	}
	if mockMMU.ReadByte(0xFF0F)&INT_VBLANK != 0 {
		t.Error("Expected V-Blank interrupt flag to be cleared")
	}
	if mockMMU.ReadByte(0xFF0F)&(INT_LCDC|INT_TIMER) == 0 {
		t.Error("Expected other interrupt flags to remain set")
	}

	// Reset CPU
	cpu.ResetCPU()
	cpu.interruptMaster = true

	// Test 2: LCDC has second highest priority
	mockMMU.WriteByte(0xFF0F, INT_LCDC|INT_TIMER)
	cpu.pendingInterrupts = INT_LCDC | INT_TIMER
	cpu.Step()
	if cpu.reg.PC != LCDC_VECTOR {
		t.Errorf("Expected PC to be %04X (LCDC vector), got %04X", LCDC_VECTOR, cpu.reg.PC)
	}
	if mockMMU.ReadByte(0xFF0F)&INT_LCDC != 0 {
		t.Error("Expected LCDC interrupt flag to be cleared")
	}
	if mockMMU.ReadByte(0xFF0F)&INT_TIMER == 0 {
		t.Error("Expected Timer interrupt flag to remain set")
	}

	// Reset CPU
	cpu.ResetCPU()
	cpu.interruptMaster = true

	// Test 3: Timer has third highest priority
	mockMMU.WriteByte(0xFF0F, INT_TIMER)
	cpu.pendingInterrupts = INT_TIMER
	cpu.Step()
	if cpu.reg.PC != TIMER_VECTOR {
		t.Errorf("Expected PC to be %04X (Timer vector), got %04X", TIMER_VECTOR, cpu.reg.PC)
	}
	if mockMMU.ReadByte(0xFF0F)&INT_TIMER != 0 {
		t.Error("Expected Timer interrupt flag to be cleared")
	}
}

func TestEIAndDIInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test EI instruction
	cpu.interruptMaster = false
	cpu.EI()
	// EI should not enable interrupts immediately
	if cpu.interruptMaster {
		t.Error("EI should not enable interrupts immediately")
	}
	if !cpu.interruptEnableScheduled {
		t.Error("EI should schedule interrupt enable")
	}

	// Execute a Step to trigger the delayed enable
	cpu.Step()
	if !cpu.interruptMaster {
		t.Error("Interrupts should be enabled after one instruction")
	}

	// Test DI instruction
	cpu.interruptMaster = true
	cpu.DI()
	// DI should not disable interrupts immediately
	if !cpu.interruptMaster {
		t.Error("DI should not disable interrupts immediately")
	}
	if !cpu.interruptDisableScheduled {
		t.Error("DI should schedule interrupt disable")
	}

	// Execute a Step to trigger the delayed disable
	cpu.Step()
	if cpu.interruptMaster {
		t.Error("Interrupts should be disabled after one instruction")
	}
}

func TestHALTBug(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Set up a scenario where the HALT bug would occur:
	// - Interrupts are disabled (IME=0)
	// - An interrupt is pending (IE & IF != 0)
	cpu.interruptMaster = false
	mockMMU.WriteByte(0xFFFF, INT_VBLANK) // Enable V-Blank interrupt
	mockMMU.WriteByte(0xFF0F, INT_VBLANK) // Set V-Blank interrupt flag

	// Execute HALT instruction directly
	cpu.HALT()

	// Check that the HALT bug flag is set
	if !cpu.haltBug {
		t.Error("HALT bug flag should be set")
	}

	// Check that the CPU is not halted
	if cpu.halted {
		t.Error("CPU should not be halted when HALT bug occurs")
	}
}
