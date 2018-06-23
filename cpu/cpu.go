package cpu

import (
	"fmt"
	"log"
)

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

// Game Boy cpu type Z80
type Z80 struct {
	reg   Registers
	clock Clock
}

func ResetClock(clock *Clock) {
	log.Print("[DEBUG] Resetting clock")
	clock.m = 0
	clock.t = 0
}

func ResetCPU(cpu *Z80) {
	cpu.reg.A = 0
	cpu.reg.B = 0
	cpu.reg.C = 0
	cpu.reg.D = 0
	cpu.reg.E = 0
	cpu.reg.F = 0
	cpu.reg.PC = 0
	cpu.reg.SP = 0
	ResetClock(&cpu.reg.LastInstructionClock)
}

func DisplayCPUFrame(cpu Z80) string {
	return fmt.Sprintf("A: %X\nB: %X\nC: %X\nD: %X\nE: %X\nF: %X",
		cpu.reg.A, cpu.reg.B, cpu.reg.C, cpu.reg.D, cpu.reg.E, cpu.reg.F)
}

func NewCPU() *Z80 {
	cpu := new(Z80)
	ResetCPU(cpu)
	return cpu
}
