package cpu

import (
	"fmt"
	"github.com/briancain/gameboy-go/mmu"
	"log"
)

// Game Boy cpu type Z80
type Z80 struct {
	reg    Registers
	clock  Clock
	m      mmu.MemoryManagedUnit
	halted bool
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

func ResetClock(clock *Clock) {
	log.Print("[DEBUG] Resetting clock")
	clock.m = 0
	clock.t = 0
}

func ResetCPU(cpu *Z80) {
	log.Print("[DEBUG] Resetting CPU")
	cpu.reg.A = 0
	cpu.reg.B = 0
	cpu.reg.C = 0
	cpu.reg.D = 0
	cpu.reg.E = 0
	cpu.reg.F = 0
	cpu.reg.PC = 0
	cpu.reg.SP = 0
	ResetClock(&cpu.reg.LastInstructionClock)
	cpu.halted = false
}

func NewCPU() *Z80 {
	cpu := new(Z80)
	ResetCPU(cpu)
	return cpu
}

// ************************************
// Debug Functions
// ************************************

func DisplayCPUFrame(cpu Z80) string {
	return fmt.Sprintf("A:  %X\nB:  %X\nC:  %X\nD:  %X\nE:  %X\nF:  %X\nPC: %X\nSP: %X",
		cpu.reg.A, cpu.reg.B, cpu.reg.C, cpu.reg.D, cpu.reg.E, cpu.reg.F, cpu.reg.PC, cpu.reg.SP)
}

func DisplayClock(cpu Z80) string {
	return fmt.Sprintf("M: %X\nT: %X", cpu.clock.m, cpu.clock.t)
}

// ************************************
// Opscode Functions
// ************************************

// TODO: Implement cpu ops codes next with functions

func NOP() {
}
