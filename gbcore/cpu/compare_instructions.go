package gbcore

// This file contains comparison instructions:
// - CP r (Compare register with A)
// - CP d8 (Compare immediate value with A)
// - SBC A,r (Subtract register and carry flag from A)
// - SBC A,d8 (Subtract immediate value and carry flag from A)

// Comparison Instructions

// 0xB8: CP B - Compare B with A
func (cpu *Z80) CP_B() int {
	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(cpu.reg.B)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.B & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 4
}

// 0xB9: CP C - Compare C with A
func (cpu *Z80) CP_C() int {
	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(cpu.reg.C)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.C & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 4
}

// 0xBA: CP D - Compare D with A
func (cpu *Z80) CP_D() int {
	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(cpu.reg.D)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.D & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 4
}

// 0xBB: CP E - Compare E with A
func (cpu *Z80) CP_E() int {
	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(cpu.reg.E)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.E & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 4
}

// 0xBC: CP H - Compare H with A
func (cpu *Z80) CP_H() int {
	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(cpu.reg.H)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.H & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 4
}

// 0xBD: CP L - Compare L with A
func (cpu *Z80) CP_L() int {
	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(cpu.reg.L)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.L & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 4
}

// 0xBE: CP (HL) - Compare value at address HL with A
func (cpu *Z80) CP_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())

	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(value)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (value & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xBF: CP A - Compare A with A
func (cpu *Z80) CP_A() int {
	// A - A is always 0

	// Set flags
	cpu.reg.F = FLAG_Z | FLAG_N // Set zero and subtract flags

	return 4
}

// 0xFE: CP d8 - Compare immediate 8-bit value with A
func (cpu *Z80) CP_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++

	// Calculate result (don't store it)
	result := int16(cpu.reg.A) - int16(value)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (value & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// Subtract with Carry Instructions

// 0x98: SBC A,B - Subtract B and carry flag from A
func (cpu *Z80) SBC_A_B() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.B) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((cpu.reg.B & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x99: SBC A,C - Subtract C and carry flag from A
func (cpu *Z80) SBC_A_C() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.C) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((cpu.reg.C & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x9A: SBC A,D - Subtract D and carry flag from A
func (cpu *Z80) SBC_A_D() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.D) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((cpu.reg.D & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x9B: SBC A,E - Subtract E and carry flag from A
func (cpu *Z80) SBC_A_E() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.E) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((cpu.reg.E & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x9C: SBC A,H - Subtract H and carry flag from A
func (cpu *Z80) SBC_A_H() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.H) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((cpu.reg.H & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x9D: SBC A,L - Subtract L and carry flag from A
func (cpu *Z80) SBC_A_L() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.L) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((cpu.reg.L & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x9E: SBC A,(HL) - Subtract value at address HL and carry flag from A
func (cpu *Z80) SBC_A_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())

	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(value) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((value & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}

// 0x9F: SBC A,A - Subtract A and carry flag from A
func (cpu *Z80) SBC_A_A() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result (A - A - carry = -carry)
	result := -int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if result == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0xDE: SBC A,d8 - Subtract immediate 8-bit value and carry flag from A
func (cpu *Z80) SBC_A_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++

	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := int16(cpu.reg.A) - int16(value) - int16(carry)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < ((value & 0x0F) + carry) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}

// 0x90: SUB B - Subtract B from A
func (cpu *Z80) SUB_B() int {
	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.B)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.B & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x91: SUB C - Subtract C from A
func (cpu *Z80) SUB_C() int {
	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.C)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.C & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x92: SUB D - Subtract D from A
func (cpu *Z80) SUB_D() int {
	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.D)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.D & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x93: SUB E - Subtract E from A
func (cpu *Z80) SUB_E() int {
	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.E)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.E & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x94: SUB H - Subtract H from A
func (cpu *Z80) SUB_H() int {
	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.H)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.H & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x95: SUB L - Subtract L from A
func (cpu *Z80) SUB_L() int {
	// Calculate result
	result := int16(cpu.reg.A) - int16(cpu.reg.L)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (cpu.reg.L & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x96: SUB (HL) - Subtract value at address HL from A
func (cpu *Z80) SUB_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())

	// Calculate result
	result := int16(cpu.reg.A) - int16(value)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (value & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}

// 0x97: SUB A - Subtract A from A
func (cpu *Z80) SUB_A() int {
	// A - A is always 0
	cpu.reg.A = 0

	// Set flags
	cpu.reg.F = FLAG_Z | FLAG_N // Set zero and subtract flags

	return 4
}

// 0xD6: SUB d8 - Subtract immediate 8-bit value from A
func (cpu *Z80) SUB_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++

	// Calculate result
	result := int16(cpu.reg.A) - int16(value)

	// Set flags
	cpu.reg.F = FLAG_N // Set subtract flag

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) < (value & 0x0F) {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if borrow
	if result < 0 {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}
