package cpu

// This file contains 16-bit arithmetic instructions:
// - INC rr (Increment register pair)
// - DEC rr (Decrement register pair)
// - ADD HL,rr (Add register pair to HL)
// - ADD SP,r8 (Add signed 8-bit immediate to SP)

// 16-bit Arithmetic Instructions

// 0x03: INC BC - Increment BC
func (cpu *Z80) INC_BC() int {
	cpu.reg.SetBC(cpu.reg.GetBC() + 1)
	return 8
}

// 0x13: INC DE - Increment DE
func (cpu *Z80) INC_DE() int {
	cpu.reg.SetDE(cpu.reg.GetDE() + 1)
	return 8
}

// 0x23: INC HL - Increment HL
func (cpu *Z80) INC_HL() int {
	cpu.reg.SetHL(cpu.reg.GetHL() + 1)
	return 8
}

// 0x33: INC SP - Increment SP
func (cpu *Z80) INC_SP() int {
	cpu.reg.SP++
	return 8
}

// 0x0B: DEC BC - Decrement BC
func (cpu *Z80) DEC_BC() int {
	cpu.reg.SetBC(cpu.reg.GetBC() - 1)
	return 8
}

// 0x1B: DEC DE - Decrement DE
func (cpu *Z80) DEC_DE() int {
	cpu.reg.SetDE(cpu.reg.GetDE() - 1)
	return 8
}

// 0x2B: DEC HL - Decrement HL
func (cpu *Z80) DEC_HL() int {
	cpu.reg.SetHL(cpu.reg.GetHL() - 1)
	return 8
}

// 0x3B: DEC SP - Decrement SP
func (cpu *Z80) DEC_SP() int {
	cpu.reg.SP--
	return 8
}

// 0x09: ADD HL,BC - Add BC to HL
func (cpu *Z80) ADD_HL_BC() int {
	// Get values
	hl := cpu.reg.GetHL()
	bc := cpu.reg.GetBC()

	// Calculate result
	result := uint32(hl) + uint32(bc)

	// Set flags
	cpu.reg.F &= FLAG_Z // Preserve zero flag

	// Half carry flag - set if carry from bit 11
	if ((hl & 0x0FFF) + (bc & 0x0FFF)) > 0x0FFF {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 15
	if result > 0xFFFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.SetHL(uint16(result & 0xFFFF))

	return 8
}

// 0x19: ADD HL,DE - Add DE to HL
func (cpu *Z80) ADD_HL_DE() int {
	// Get values
	hl := cpu.reg.GetHL()
	de := cpu.reg.GetDE()

	// Calculate result
	result := uint32(hl) + uint32(de)

	// Set flags
	cpu.reg.F &= FLAG_Z // Preserve zero flag

	// Half carry flag - set if carry from bit 11
	if ((hl & 0x0FFF) + (de & 0x0FFF)) > 0x0FFF {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 15
	if result > 0xFFFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.SetHL(uint16(result & 0xFFFF))

	return 8
}

// 0x29: ADD HL,HL - Add HL to HL
func (cpu *Z80) ADD_HL_HL() int {
	// Get value
	hl := cpu.reg.GetHL()

	// Calculate result
	result := uint32(hl) + uint32(hl)

	// Set flags
	cpu.reg.F &= FLAG_Z // Preserve zero flag

	// Half carry flag - set if carry from bit 11
	if ((hl & 0x0FFF) + (hl & 0x0FFF)) > 0x0FFF {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 15
	if result > 0xFFFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.SetHL(uint16(result & 0xFFFF))

	return 8
}

// 0x39: ADD HL,SP - Add SP to HL
func (cpu *Z80) ADD_HL_SP() int {
	// Get values
	hl := cpu.reg.GetHL()
	sp := cpu.reg.SP

	// Calculate result
	result := uint32(hl) + uint32(sp)

	// Set flags
	cpu.reg.F &= FLAG_Z // Preserve zero flag

	// Half carry flag - set if carry from bit 11
	if ((hl & 0x0FFF) + (sp & 0x0FFF)) > 0x0FFF {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 15
	if result > 0xFFFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.SetHL(uint16(result & 0xFFFF))

	return 8
}

// 0xE8: ADD SP,r8 - Add signed 8-bit immediate to SP
func (cpu *Z80) ADD_SP_r8() int {
	// Get signed 8-bit immediate
	value := int8(cpu.mmu.ReadByte(cpu.reg.PC))
	cpu.reg.PC++

	// Calculate result
	result := uint32(cpu.reg.SP) + uint32(int32(value))

	// Set flags
	cpu.reg.F = 0

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.SP & 0x0F) + (uint16(value) & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if ((cpu.reg.SP & 0xFF) + (uint16(value) & 0xFF)) > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.SP = uint16(result & 0xFFFF)

	return 16
}
