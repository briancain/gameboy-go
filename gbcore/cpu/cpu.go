package gbcore

import (
	"fmt"
	"log"
)

// Game Boy cpu type Z80
// Z80 Manual (http://www.zilog.com/docs/z80/um0080.pdf)
type Z80 struct {
	reg    Registers
	clock  Clock
	halted bool
	stopped bool

	interruptEnabled bool
	interruptMaster  bool // IME flag
	pendingInterrupts byte

	// Reference to MMU for memory access
	mmu MMU
}

// MMU interface for CPU to access memory
type MMU interface {
	ReadByte(addr uint16) byte
	WriteByte(addr uint16, value byte)
	ReadWord(addr uint16) uint16
	WriteWord(addr uint16, value uint16)
}

// Each struct value A-H is an 8-bit register
// Grouped together, they can form 16 bit registers
type Registers struct {
	// 8-bit registers

	// Accumulator & Flags
	A byte // Accumulator
	F byte // Flag register

	// BC
	B byte // Hi register
	C byte // Lo register

	// DE
	D byte // Hi
	E byte // Lo

	// HL
	H byte // Hi
	L byte // Lo

	// 16-bit registers
	PC uint16 // Program Counter
	SP uint16 // Stack Pointer
}

// Flag register values
const (
	FLAG_Z byte = 0x80 // Zero flag
	FLAG_N byte = 0x40 // Subtract flag
	FLAG_H byte = 0x20 // Half carry flag
	FLAG_C byte = 0x10 // Carry flag
)

type Clock struct {
	m int // Machine cycles
	t int // Clock cycles (t = m * 4)
}

// Get/Set 16-bit register pairs
func (r *Registers) GetAF() uint16 {
	return uint16(r.A)<<8 | uint16(r.F)
}

func (r *Registers) SetAF(value uint16) {
	r.A = byte(value >> 8)
	r.F = byte(value & 0xFF)
}

func (r *Registers) GetBC() uint16 {
	return uint16(r.B)<<8 | uint16(r.C)
}

func (r *Registers) SetBC(value uint16) {
	r.B = byte(value >> 8)
	r.C = byte(value & 0xFF)
}

func (r *Registers) GetDE() uint16 {
	return uint16(r.D)<<8 | uint16(r.E)
}

func (r *Registers) SetDE(value uint16) {
	r.D = byte(value >> 8)
	r.E = byte(value & 0xFF)
}

func (r *Registers) GetHL() uint16 {
	return uint16(r.H)<<8 | uint16(r.L)
}

func (r *Registers) SetHL(value uint16) {
	r.H = byte(value >> 8)
	r.L = byte(value & 0xFF)
}

// Flag operations
func (r *Registers) SetFlag(flag byte) {
	r.F |= flag
}

func (r *Registers) ClearFlag(flag byte) {
	r.F &= ^flag
}

func (r *Registers) GetFlag(flag byte) bool {
	return (r.F & flag) != 0
}

func (cpu *Z80) ResetClock() {
	cpu.clock.m = 0
	cpu.clock.t = 0
}

func (cpu *Z80) ResetCPU() {
	// Initialize registers to power-up values
	cpu.reg.A = 0x01
	cpu.reg.F = 0xB0
	cpu.reg.B = 0x00
	cpu.reg.C = 0x13
	cpu.reg.D = 0x00
	cpu.reg.E = 0xD8
	cpu.reg.H = 0x01
	cpu.reg.L = 0x4D
	cpu.reg.PC = 0x0100 // Start at 0x0100 after boot ROM
	cpu.reg.SP = 0xFFFE

	cpu.ResetClock()
	cpu.halted = false
	cpu.stopped = false
	cpu.interruptEnabled = false
	cpu.interruptMaster = false
	cpu.pendingInterrupts = 0
}

func NewCPU(mmu MMU) (*Z80, error) {
	log.Println("[Core] Initializing a new Z80 CPU ...")

	cpu := &Z80{mmu: mmu}
	cpu.ResetCPU()

	return cpu, nil
}

// Step executes one instruction and returns the number of cycles taken
func (cpu *Z80) Step() int {
	// Handle interrupts
	if cpu.interruptMaster && cpu.pendingInterrupts > 0 {
		cpu.halted = false
		cpu.stopped = false
		// Process interrupts
		cpu.handleInterrupts()
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
	cpu.reg.PC++

	// Execute instruction
	cycles := cpu.executeInstruction(opcode)

	// Update clock
	cpu.clock.t += cycles
	cpu.clock.m += cycles / 4

	return cycles
}

// Handle interrupts
func (cpu *Z80) handleInterrupts() {
	// Implementation will go here
}

// Execute instruction based on opcode
func (cpu *Z80) executeInstruction(opcode byte) int {
	switch opcode {
	case 0x00: // NOP
		return cpu.NOP()
	case 0x76: // HALT
		return cpu.HALT()
	case 0x10: // STOP
		return cpu.STOP()
	// Add more opcodes here
	default:
		log.Printf("Unknown opcode: 0x%02X at PC: 0x%04X", opcode, cpu.reg.PC-1)
		return 4
	}
}

// ************************************
// Debug Functions
// ************************************

func (cpu *Z80) DisplayCPUFrame() string {
	return fmt.Sprintf("A:  %02X\nF:  %02X\nB:  %02X\nC:  %02X\nD:  %02X\nE:  %02X\nH:  %02X\nL:  %02X\nPC: %04X\nSP: %04X\nFlags: Z:%t N:%t H:%t C:%t",
		cpu.reg.A, cpu.reg.F, cpu.reg.B, cpu.reg.C, cpu.reg.D, cpu.reg.E, cpu.reg.H, cpu.reg.L, 
		cpu.reg.PC, cpu.reg.SP,
		cpu.reg.GetFlag(FLAG_Z), cpu.reg.GetFlag(FLAG_N), cpu.reg.GetFlag(FLAG_H), cpu.reg.GetFlag(FLAG_C))
}

func (cpu *Z80) DisplayClock() string {
	return fmt.Sprintf("M: %d\nT: %d", cpu.clock.m, cpu.clock.t)
}

// ************************************
// Instruction Implementations
// ************************************

// 0x00: NOP - No operation
func (cpu *Z80) NOP() int {
	// Do nothing
	return 4 // 1 machine cycle = 4 clock cycles
}

// 0x76: HALT - Halt the CPU until an interrupt occurs
func (cpu *Z80) HALT() int {
	cpu.halted = true
	return 4
}

// 0x10: STOP - Halt CPU & LCD display until button pressed
func (cpu *Z80) STOP() int {
	cpu.stopped = true
	// Read next byte (usually 0x00)
	cpu.reg.PC++
	return 4
}

// Stack operations
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
