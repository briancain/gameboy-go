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

	interruptEnabled bool
}

// Each struct value A-H is an 8-bit register
// Grouped together, they can form 16 bit registers
type Registers struct {
	// 8-bit registers

	// Accumulator & Flags
	A byte // Hi register
	F byte // Flag register

	// BC
	B byte // Hi register
	C byte // Lo register

	D byte // Hi
	E byte // Lo

	L byte // Hi
	H byte // Lo

	// 16-bit registers
	PC                   byte // Program Counter
	SP                   byte // Stack Pointer
	LastInstructionClock Clock
}

// Flag register values
const (
	zero      = 0x80
	operation = 0x40
	halfcarry = 0x20
	carry     = 0x10
)

type Clock struct {
	m byte
	t byte
}

func (cpu *Z80) ResetClock() error {
	log.Print("[DEBUG] Resetting clock")
	cpu.clock.m = 0
	cpu.clock.t = 0

	return nil
}

func (cpu *Z80) ResetCPU() error {
	log.Print("[DEBUG] Resetting CPU")
	cpu.reg.A = 0
	cpu.reg.B = 0
	cpu.reg.C = 0
	cpu.reg.D = 0
	cpu.reg.E = 0
	cpu.reg.F = 0
	cpu.reg.L = 0
	cpu.reg.H = 0

	cpu.reg.PC = 0
	cpu.reg.SP = 0

	cpu.ResetClock()
	cpu.halted = false

	return nil
}

func NewCPU() (*Z80, error) {
	log.Println("[Core] Initializing a new Z80 CPU ...")

	cpu := &Z80{}
	if err := cpu.ResetCPU(); err != nil {
		return nil, err
	}

	return cpu, nil
}

// ************************************
// Debug Functions
// ************************************

func (cpu *Z80) DisplayCPUFrame() string {
	return fmt.Sprintf("A:  %X\nB:  %X\nC:  %X\nD:  %X\nE:  %X\nF:  %X\nPC: %X\nSP: %X",
		cpu.reg.A, cpu.reg.B, cpu.reg.C, cpu.reg.D, cpu.reg.E, cpu.reg.F, cpu.reg.PC, cpu.reg.SP)
}

func (cpu *Z80) DisplayClock() string {
	return fmt.Sprintf("M: %X\nT: %X", cpu.clock.m, cpu.clock.t)
}

// ************************************
// Opscode Functions
// ************************************

// http://gameboy.mongenel.com/dmg/opcodes.html

func (cpu *Z80) ExecuteOpCode(code byte) error {
	return nil
}

// TODO: Implement cpu ops codes next with functions
// https://gbdev.io/gb-opcodes/optables/
// TODO: move this into its own component?

// 0x00
func (cpu *Z80) NOP() {
	// do nothing
}

func (cpu *Z80) HALT() {
}

func (cpu *Z80) STOP() {
}
