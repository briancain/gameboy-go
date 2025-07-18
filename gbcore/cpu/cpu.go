package gbcore

// CPU flags
const (
	FLAG_Z byte = 0x80 // Zero flag
	FLAG_N byte = 0x40 // Subtract flag
	FLAG_H byte = 0x20 // Half carry flag
	FLAG_C byte = 0x10 // Carry flag
)

// Z80 represents the Game Boy's CPU
type Z80 struct {
	// Registers
	reg Registers

	// Memory Management Unit
	mmu MMU

	// Clock
	clock Clock

	// Interrupt master enable flag
	interruptMaster bool

	// Interrupt scheduling flags
	interruptEnableScheduled  bool
	interruptDisableScheduled bool

	// Pending interrupts
	pendingInterrupts byte

	// CPU state
	halted  bool
	stopped bool
	haltBug bool
}

// Registers represents the CPU registers
type Registers struct {
	A byte // Accumulator
	F byte // Flags
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte

	PC uint16 // Program counter
	SP uint16 // Stack pointer
}

// Clock represents the CPU clock
type Clock struct {
	m int // Machine cycles
	t int // Clock cycles
}

// MMU interface for CPU to access memory
type MMU interface {
	ReadByte(addr uint16) byte
	WriteByte(addr uint16, value byte)
	ReadWord(addr uint16) uint16
	WriteWord(addr uint16, value uint16)
}

// Get 16-bit register pairs
func (r *Registers) GetAF() uint16 {
	return uint16(r.A)<<8 | uint16(r.F)
}

func (r *Registers) GetBC() uint16 {
	return uint16(r.B)<<8 | uint16(r.C)
}

func (r *Registers) GetDE() uint16 {
	return uint16(r.D)<<8 | uint16(r.E)
}

func (r *Registers) GetHL() uint16 {
	return uint16(r.H)<<8 | uint16(r.L)
}

// Set 16-bit register pairs
func (r *Registers) SetAF(value uint16) {
	r.A = byte(value >> 8)
	r.F = byte(value & 0xFF)
}

func (r *Registers) SetBC(value uint16) {
	r.B = byte(value >> 8)
	r.C = byte(value & 0xFF)
}

func (r *Registers) SetDE(value uint16) {
	r.D = byte(value >> 8)
	r.E = byte(value & 0xFF)
}

func (r *Registers) SetHL(value uint16) {
	r.H = byte(value >> 8)
	r.L = byte(value & 0xFF)
}

// Flag operations
func (r *Registers) GetFlag(flag byte) bool {
	return (r.F & flag) != 0
}

func (r *Registers) SetFlag(flag byte) {
	r.F |= flag
}

func (r *Registers) ClearFlag(flag byte) {
	r.F &= ^flag
}

// NewCPU creates a new Z80 CPU
func NewCPU(mmu MMU) (*Z80, error) {
	cpu := &Z80{mmu: mmu}
	cpu.ResetCPU()

	return cpu, nil
}

// Step executes one instruction and returns the number of cycles taken
func (cpu *Z80) Step() int {
	// Save the interrupt enable/disable scheduled flags
	interruptEnableScheduled := cpu.interruptEnableScheduled
	interruptDisableScheduled := cpu.interruptDisableScheduled

	// Clear the scheduled flags
	cpu.interruptEnableScheduled = false
	cpu.interruptDisableScheduled = false

	// Update pending interrupts
	interruptFlag := cpu.mmu.ReadByte(0xFF0F)
	interruptEnable := cpu.mmu.ReadByte(0xFFFF)
	cpu.pendingInterrupts = interruptFlag & interruptEnable & 0x1F

	// Handle interrupts
	if cpu.interruptMaster && cpu.pendingInterrupts > 0 {
		cpu.halted = false
		cpu.stopped = false
		// Process interrupts
		cpu.handleInterrupts()

		// Return cycles for interrupt handling (5 machine cycles)
		return 20
	}

	// If CPU is halted or stopped, just return cycles for one machine cycle
	if cpu.halted {
		return 4
	}
	if cpu.stopped {
		return 4
	}

	// Fetch opcode
	opcode := cpu.mmu.ReadByte(cpu.reg.PC)
	
	// Handle HALT bug
	// According to the manual, when the HALT bug occurs, the PC doesn't increment
	// after fetching the opcode, causing the next instruction to be executed twice
	if !cpu.haltBug {
		cpu.reg.PC++
	} else {
		// Clear the HALT bug flag after it's been handled
		cpu.haltBug = false
	}

	// Execute instruction
	cycles := cpu.executeInstruction(opcode)

	// Handle delayed interrupt enable/disable
	if interruptEnableScheduled {
		cpu.interruptMaster = true
	}
	if interruptDisableScheduled {
		cpu.interruptMaster = false
	}

	// Update clock
	cpu.clock.t += cycles
	cpu.clock.m += cycles / 4

	return cycles
}

// ResetCPU resets the CPU to its initial state
func (cpu *Z80) ResetCPU() {
	// Initialize registers
	cpu.reg.A = 0x01
	cpu.reg.F = 0xB0
	cpu.reg.B = 0x00
	cpu.reg.C = 0x13
	cpu.reg.D = 0x00
	cpu.reg.E = 0xD8
	cpu.reg.H = 0x01
	cpu.reg.L = 0x4D
	cpu.reg.PC = 0x0100
	cpu.reg.SP = 0xFFFE

	// Initialize clock
	cpu.clock.m = 0
	cpu.clock.t = 0

	// Initialize interrupt state
	cpu.interruptMaster = false
	cpu.interruptEnableScheduled = false
	cpu.interruptDisableScheduled = false
	cpu.pendingInterrupts = 0

	// Initialize CPU state
	cpu.halted = false
	cpu.stopped = false
	cpu.haltBug = false
}

// ResetClock resets the CPU clock
func (cpu *Z80) ResetClock() {
	cpu.clock.m = 0
	cpu.clock.t = 0
}

// Push a 16-bit value onto the stack
func (cpu *Z80) pushWord(value uint16) {
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, value)
}

func (cpu *Z80) popWord() uint16 {
	value := cpu.mmu.ReadWord(cpu.reg.SP)
	cpu.reg.SP += 2
	return value
}

// More instructions will be implemented here
