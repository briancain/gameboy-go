package gbcore

import (
	"fmt"
)

// DisplayCPUFrame displays the current CPU state
func (cpu *Z80) DisplayCPUFrame() {
	fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: %04X\n",
		cpu.reg.A, cpu.reg.F, cpu.reg.B, cpu.reg.C, cpu.reg.D, cpu.reg.E, cpu.reg.H, cpu.reg.L, cpu.reg.SP, cpu.reg.PC)

	fmt.Printf("Flags: Z:%d N:%d H:%d C:%d\n",
		boolToInt(cpu.reg.GetFlag(FLAG_Z)), boolToInt(cpu.reg.GetFlag(FLAG_N)),
		boolToInt(cpu.reg.GetFlag(FLAG_H)), boolToInt(cpu.reg.GetFlag(FLAG_C)))
}

// DisplayClock displays the current CPU clock state
func (cpu *Z80) DisplayClock() {
	fmt.Printf("Clock - M: %d T: %d\n", cpu.clock.m, cpu.clock.t)
}

// Helper function to convert bool to int
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
