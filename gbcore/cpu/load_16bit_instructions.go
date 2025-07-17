package gbcore

// This file contains 16-bit load instructions:
// - LD rr,d16 (Load immediate 16-bit value into register pair)
// - LD SP,HL (Load HL into SP)
// - LD HL,SP+r8 (Add signed 8-bit immediate to SP and load result into HL)
// - LD (a16),SP (Store SP at address a16)
// - PUSH/POP instructions for register pairs

// 16-bit Load Instructions

// 0x01: LD BC,d16 - Load immediate 16-bit value into BC
func (cpu *Z80) LD_BC_d16() int {
	value := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.SetBC(value)
	return 12
}

// 0x11: LD DE,d16 - Load immediate 16-bit value into DE
func (cpu *Z80) LD_DE_d16() int {
	value := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.SetDE(value)
	return 12
}

// 0x21: LD HL,d16 - Load immediate 16-bit value into HL
func (cpu *Z80) LD_HL_d16() int {
	value := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.SetHL(value)
	return 12
}

// 0x31: LD SP,d16 - Load immediate 16-bit value into SP
func (cpu *Z80) LD_SP_d16() int {
	cpu.reg.SP = cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2
	return 12
}

// 0xF9: LD SP,HL - Load HL into SP
func (cpu *Z80) LD_SP_HL() int {
	cpu.reg.SP = cpu.reg.GetHL()
	return 8
}

// 0xF8: LD HL,SP+r8 - Add signed 8-bit immediate to SP and load result into HL
func (cpu *Z80) LD_HL_SP_r8() int {
	// Get signed 8-bit immediate
	value := int8(cpu.mmu.ReadByte(cpu.reg.PC))
	cpu.reg.PC++
	
	// Calculate result
	result := uint16(int32(cpu.reg.SP) + int32(value))
	
	// Set flags
	cpu.reg.F = 0
	
	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.SP ^ uint16(value) ^ result) & 0x10) != 0 {
		cpu.reg.SetFlag(FLAG_H)
	}
	
	// Carry flag - set if carry from bit 7
	if ((cpu.reg.SP ^ uint16(value) ^ result) & 0x100) != 0 {
		cpu.reg.SetFlag(FLAG_C)
	}
	
	// Load result into HL
	cpu.reg.SetHL(result)
	
	return 12
}

// 0x08: LD (a16),SP - Store SP at address a16
func (cpu *Z80) LD_a16_SP() int {
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.mmu.WriteWord(address, cpu.reg.SP)
	return 20
}

// 0xF5: PUSH AF - Push AF onto stack
func (cpu *Z80) PUSH_AF() int {
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.GetAF())
	return 16
}

// 0xC5: PUSH BC - Push BC onto stack
func (cpu *Z80) PUSH_BC() int {
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.GetBC())
	return 16
}

// 0xD5: PUSH DE - Push DE onto stack
func (cpu *Z80) PUSH_DE() int {
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.GetDE())
	return 16
}

// 0xE5: PUSH HL - Push HL onto stack
func (cpu *Z80) PUSH_HL() int {
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.GetHL())
	return 16
}

// 0xF1: POP AF - Pop from stack into AF
func (cpu *Z80) POP_AF() int {
	value := cpu.mmu.ReadWord(cpu.reg.SP)
	cpu.reg.SP += 2
	// Only the upper 4 bits of F are used
	cpu.reg.SetAF(value & 0xFFF0)
	return 12
}

// 0xC1: POP BC - Pop from stack into BC
func (cpu *Z80) POP_BC() int {
	value := cpu.mmu.ReadWord(cpu.reg.SP)
	cpu.reg.SP += 2
	cpu.reg.SetBC(value)
	return 12
}

// 0xD1: POP DE - Pop from stack into DE
func (cpu *Z80) POP_DE() int {
	value := cpu.mmu.ReadWord(cpu.reg.SP)
	cpu.reg.SP += 2
	cpu.reg.SetDE(value)
	return 12
}

// 0xE1: POP HL - Pop from stack into HL
func (cpu *Z80) POP_HL() int {
	value := cpu.mmu.ReadWord(cpu.reg.SP)
	cpu.reg.SP += 2
	cpu.reg.SetHL(value)
	return 12
}
