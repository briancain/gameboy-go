package cpu

// This file contains 8-bit arithmetic instructions:
// - INC r (Increment register)
// - DEC r (Decrement register)
// - ADD A,r (Add register to A)
// - ADC A,r (Add register and carry flag to A)
// - SUB r (Subtract register from A)
// - SBC A,r (Subtract register and carry flag from A)
// - ADD A,r (Add register to A)
// - ADC A,r (Add register and carry flag to A)
// - SUB r (Subtract register from A)
// - SBC A,r (Subtract register and carry flag from A)

// 8-bit Arithmetic Instructions

// 0x04: INC B - Increment B
func (cpu *Z80) INC_B() int {
	cpu.reg.B++

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (cpu.reg.B & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x0C: INC C - Increment C
func (cpu *Z80) INC_C() int {
	cpu.reg.C++

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (cpu.reg.C & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x14: INC D - Increment D
func (cpu *Z80) INC_D() int {
	cpu.reg.D++

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (cpu.reg.D & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x1C: INC E - Increment E
func (cpu *Z80) INC_E() int {
	cpu.reg.E++

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (cpu.reg.E & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x24: INC H - Increment H
func (cpu *Z80) INC_H() int {
	cpu.reg.H++

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (cpu.reg.H & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x2C: INC L - Increment L
func (cpu *Z80) INC_L() int {
	cpu.reg.L++

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (cpu.reg.L & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x3C: INC A - Increment A
func (cpu *Z80) INC_A() int {
	cpu.reg.A++

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (cpu.reg.A & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x34: INC (HL) - Increment value at address HL
func (cpu *Z80) INC_HL_() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	value++
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F &= FLAG_C // Preserve carry flag

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if (value & 0x0F) == 0 {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 12
}

// 0x05: DEC B - Decrement B
func (cpu *Z80) DEC_B() int {
	cpu.reg.B--

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.B & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x0D: DEC C - Decrement C
func (cpu *Z80) DEC_C() int {
	cpu.reg.C--

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.C & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x15: DEC D - Decrement D
func (cpu *Z80) DEC_D() int {
	cpu.reg.D--

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.D & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x1D: DEC E - Decrement E
func (cpu *Z80) DEC_E() int {
	cpu.reg.E--

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.E & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x25: DEC H - Decrement H
func (cpu *Z80) DEC_H() int {
	cpu.reg.H--

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.H & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x2D: DEC L - Decrement L
func (cpu *Z80) DEC_L() int {
	cpu.reg.L--

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.L & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x3D: DEC A - Decrement A
func (cpu *Z80) DEC_A() int {
	cpu.reg.A--

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (cpu.reg.A & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 4
}

// 0x35: DEC (HL) - Decrement value at address HL
func (cpu *Z80) DEC_HL_() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	value--
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_N // Preserve carry flag, set subtract flag

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if borrow from bit 4
	if (value & 0x0F) == 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	return 12
}

// 0x80: ADD A,B - Add B to A
func (cpu *Z80) ADD_A_B() int {
	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.B)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.B & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x81: ADD A,C - Add C to A
func (cpu *Z80) ADD_A_C() int {
	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.C)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.C & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x82: ADD A,D - Add D to A
func (cpu *Z80) ADD_A_D() int {
	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.D)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.D & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x83: ADD A,E - Add E to A
func (cpu *Z80) ADD_A_E() int {
	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.E)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.E & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x84: ADD A,H - Add H to A
func (cpu *Z80) ADD_A_H() int {
	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.H)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.H & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x85: ADD A,L - Add L to A
func (cpu *Z80) ADD_A_L() int {
	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.L)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.L & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x86: ADD A,(HL) - Add value at address HL to A
func (cpu *Z80) ADD_A_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (value & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}

// 0x87: ADD A,A - Add A to A
func (cpu *Z80) ADD_A_A() int {
	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.A)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.A & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0xC6: ADD A,d8 - Add immediate 8-bit value to A
func (cpu *Z80) ADD_A_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (value & 0x0F)) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}

// 0x88: ADC A,B - Add B and carry flag to A
func (cpu *Z80) ADC_A_B() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.B) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.B & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x89: ADC A,C - Add C and carry flag to A
func (cpu *Z80) ADC_A_C() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.C) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.C & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x8A: ADC A,D - Add D and carry flag to A
func (cpu *Z80) ADC_A_D() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.D) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.D & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x8B: ADC A,E - Add E and carry flag to A
func (cpu *Z80) ADC_A_E() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.E) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.E & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x8C: ADC A,H - Add H and carry flag to A
func (cpu *Z80) ADC_A_H() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.H) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.H & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x8D: ADC A,L - Add L and carry flag to A
func (cpu *Z80) ADC_A_L() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.L) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.L & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0x8E: ADC A,(HL) - Add value at address HL and carry flag to A
func (cpu *Z80) ADC_A_HL() int {
	// Get value from memory
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())

	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(value) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (value & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}

// 0x8F: ADC A,A - Add A and carry flag to A
func (cpu *Z80) ADC_A_A() int {
	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(cpu.reg.A) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (cpu.reg.A & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 4
}

// 0xCE: ADC A,d8 - Add immediate 8-bit value and carry flag to A
func (cpu *Z80) ADC_A_d8() int {
	// Get immediate value
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++

	// Get carry value
	carry := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		carry = 1
	}

	// Calculate result
	result := uint16(cpu.reg.A) + uint16(value) + uint16(carry)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if (result & 0xFF) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Half carry flag - set if carry from bit 3
	if ((cpu.reg.A & 0x0F) + (value & 0x0F) + carry) > 0x0F {
		cpu.reg.SetFlag(FLAG_H)
	}

	// Carry flag - set if carry from bit 7
	if result > 0xFF {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Store result
	cpu.reg.A = byte(result & 0xFF)

	return 8
}
