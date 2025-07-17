package gbcore

// This file contains logical operation instructions:
// - AND r (Logical AND register with A)
// - OR r (Logical OR register with A)
// - XOR r (Logical XOR register with A)
// - CP r (Compare register with A)
// - Immediate versions of the above operations

// Logical Operations

// 0xA0: AND B - Logical AND B with A
func (cpu *Z80) AND_B() int {
	// Calculate result
	cpu.reg.A &= cpu.reg.B
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xA1: AND C - Logical AND C with A
func (cpu *Z80) AND_C() int {
	// Calculate result
	cpu.reg.A &= cpu.reg.C
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xA2: AND D - Logical AND D with A
func (cpu *Z80) AND_D() int {
	// Calculate result
	cpu.reg.A &= cpu.reg.D
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xA3: AND E - Logical AND E with A
func (cpu *Z80) AND_E() int {
	// Calculate result
	cpu.reg.A &= cpu.reg.E
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xA4: AND H - Logical AND H with A
func (cpu *Z80) AND_H() int {
	// Calculate result
	cpu.reg.A &= cpu.reg.H
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xA5: AND L - Logical AND L with A
func (cpu *Z80) AND_L() int {
	// Calculate result
	cpu.reg.A &= cpu.reg.L
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xA6: AND (HL) - Logical AND value at address HL with A
func (cpu *Z80) AND_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	
	// Calculate result
	cpu.reg.A &= value
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 8
}

// 0xA7: AND A - Logical AND A with A
func (cpu *Z80) AND_A() int {
	// A & A = A, no change to A
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xE6: AND d8 - Logical AND immediate 8-bit value with A
func (cpu *Z80) AND_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	
	// Calculate result
	cpu.reg.A &= value
	
	// Set flags
	cpu.reg.F = FLAG_H // Set half carry flag
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 8
}

// 0xB0: OR B - Logical OR B with A
func (cpu *Z80) OR_B() int {
	// Calculate result
	cpu.reg.A |= cpu.reg.B
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xB1: OR C - Logical OR C with A
func (cpu *Z80) OR_C() int {
	// Calculate result
	cpu.reg.A |= cpu.reg.C
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xB2: OR D - Logical OR D with A
func (cpu *Z80) OR_D() int {
	// Calculate result
	cpu.reg.A |= cpu.reg.D
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xB3: OR E - Logical OR E with A
func (cpu *Z80) OR_E() int {
	// Calculate result
	cpu.reg.A |= cpu.reg.E
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xB4: OR H - Logical OR H with A
func (cpu *Z80) OR_H() int {
	// Calculate result
	cpu.reg.A |= cpu.reg.H
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xB5: OR L - Logical OR L with A
func (cpu *Z80) OR_L() int {
	// Calculate result
	cpu.reg.A |= cpu.reg.L
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xB6: OR (HL) - Logical OR value at address HL with A
func (cpu *Z80) OR_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	
	// Calculate result
	cpu.reg.A |= value
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 8
}

// 0xB7: OR A - Logical OR A with A
func (cpu *Z80) OR_A() int {
	// A | A = A, no change to A
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xF6: OR d8 - Logical OR immediate 8-bit value with A
func (cpu *Z80) OR_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	
	// Calculate result
	cpu.reg.A |= value
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 8
}

// 0xA8: XOR B - Logical XOR B with A
func (cpu *Z80) XOR_B() int {
	// Calculate result
	cpu.reg.A ^= cpu.reg.B
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xA9: XOR C - Logical XOR C with A
func (cpu *Z80) XOR_C() int {
	// Calculate result
	cpu.reg.A ^= cpu.reg.C
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xAA: XOR D - Logical XOR D with A
func (cpu *Z80) XOR_D() int {
	// Calculate result
	cpu.reg.A ^= cpu.reg.D
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xAB: XOR E - Logical XOR E with A
func (cpu *Z80) XOR_E() int {
	// Calculate result
	cpu.reg.A ^= cpu.reg.E
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xAC: XOR H - Logical XOR H with A
func (cpu *Z80) XOR_H() int {
	// Calculate result
	cpu.reg.A ^= cpu.reg.H
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xAD: XOR L - Logical XOR L with A
func (cpu *Z80) XOR_L() int {
	// Calculate result
	cpu.reg.A ^= cpu.reg.L
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 4
}

// 0xAE: XOR (HL) - Logical XOR value at address HL with A
func (cpu *Z80) XOR_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	
	// Calculate result
	cpu.reg.A ^= value
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 8
}

// 0xAF: XOR A - Logical XOR A with A
func (cpu *Z80) XOR_A() int {
	// A ^ A = 0
	cpu.reg.A = 0
	
	// Set flags
	cpu.reg.F = FLAG_Z // Set zero flag
	
	return 4
}

// 0xEE: XOR d8 - Logical XOR immediate 8-bit value with A
func (cpu *Z80) XOR_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	
	// Calculate result
	cpu.reg.A ^= value
	
	// Set flags
	cpu.reg.F = 0
	
	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}
	
	return 8
}
