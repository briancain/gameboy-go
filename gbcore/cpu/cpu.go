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

type Registers struct {
	// 8-bit registers
	A byte
	B byte
	C byte
	D byte
	E byte
	F byte // Flag register

	// 16-bit registers
	PC                   byte // Program Counter
	SP                   byte // Stack Pointer
	LastInstructionClock Clock
}

type Clock struct {
	m byte
	t byte
}

// Flags
const (
	zero      = 0x80
	operation = 0x40
	halfcarry = 0x20
	carry     = 0x10
)

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

// TODO: Implement cpu ops codes next with functions
// https://gbdev.io/gb-opcodes/optables/
// TODO: move this into its own component?

func (cpu *Z80) NOP() {
	cpu.clock.m = 1
	cpu.clock.t = 4
}

func (cpu *Z80) HALT() {
	cpu.clock.m = 1
	cpu.clock.t = 4
}

func (cpu *Z80) STOP() {
	cpu.clock.m = 1
	cpu.clock.t = 4
}
