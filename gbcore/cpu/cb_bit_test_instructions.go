package gbcore

// This file contains CB-prefixed bit test instructions:
// - BIT b,r (Test bit b of register r)
//
// These instructions test if a specific bit in a register or memory location
// is set (1) or reset (0). The result affects the Z flag (set if bit is 0).
// The N flag is reset, and the H flag is set.

// Helper function for BIT instructions
func (cpu *Z80) testBit(value byte, bit byte) {
	// Set flags
	cpu.reg.F = (cpu.reg.F & FLAG_C) | FLAG_H // Preserve carry flag, set half carry flag

	// Zero flag - set if bit is not set
	if (value & (1 << bit)) == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	} else {
		cpu.reg.ClearFlag(FLAG_Z)
	}
}

// Helper function for RES instructions
func (cpu *Z80) resetBit(value *byte, bit byte) {
	*value &= ^(1 << bit)
}

// Helper function for SET instructions
func (cpu *Z80) setBit(value *byte, bit byte) {
	*value |= (1 << bit)
}

// BIT Instructions

// 0xCB 0x40: BIT 0,B - Test bit 0 of B
func (cpu *Z80) BIT_0_B() int {
	cpu.testBit(cpu.reg.B, 0)
	return 8
}

// 0xCB 0x41: BIT 0,C - Test bit 0 of C
func (cpu *Z80) BIT_0_C() int {
	cpu.testBit(cpu.reg.C, 0)
	return 8
}

// 0xCB 0x42: BIT 0,D - Test bit 0 of D
func (cpu *Z80) BIT_0_D() int {
	cpu.testBit(cpu.reg.D, 0)
	return 8
}

// 0xCB 0x43: BIT 0,E - Test bit 0 of E
func (cpu *Z80) BIT_0_E() int {
	cpu.testBit(cpu.reg.E, 0)
	return 8
}

// 0xCB 0x44: BIT 0,H - Test bit 0 of H
func (cpu *Z80) BIT_0_H() int {
	cpu.testBit(cpu.reg.H, 0)
	return 8
}

// 0xCB 0x45: BIT 0,L - Test bit 0 of L
func (cpu *Z80) BIT_0_L() int {
	cpu.testBit(cpu.reg.L, 0)
	return 8
}

// 0xCB 0x46: BIT 0,(HL) - Test bit 0 of value at address HL
func (cpu *Z80) BIT_0_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 0)
	return 12
}

// 0xCB 0x47: BIT 0,A - Test bit 0 of A
func (cpu *Z80) BIT_0_A() int {
	cpu.testBit(cpu.reg.A, 0)
	return 8
}

// 0xCB 0x48: BIT 1,B - Test bit 1 of B
func (cpu *Z80) BIT_1_B() int {
	cpu.testBit(cpu.reg.B, 1)
	return 8
}

// 0xCB 0x49: BIT 1,C - Test bit 1 of C
func (cpu *Z80) BIT_1_C() int {
	cpu.testBit(cpu.reg.C, 1)
	return 8
}

// 0xCB 0x4A: BIT 1,D - Test bit 1 of D
func (cpu *Z80) BIT_1_D() int {
	cpu.testBit(cpu.reg.D, 1)
	return 8
}

// 0xCB 0x4B: BIT 1,E - Test bit 1 of E
func (cpu *Z80) BIT_1_E() int {
	cpu.testBit(cpu.reg.E, 1)
	return 8
}

// 0xCB 0x4C: BIT 1,H - Test bit 1 of H
func (cpu *Z80) BIT_1_H() int {
	cpu.testBit(cpu.reg.H, 1)
	return 8
}

// 0xCB 0x4D: BIT 1,L - Test bit 1 of L
func (cpu *Z80) BIT_1_L() int {
	cpu.testBit(cpu.reg.L, 1)
	return 8
}

// 0xCB 0x4E: BIT 1,(HL) - Test bit 1 of value at address HL
func (cpu *Z80) BIT_1_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 1)
	return 12
}

// 0xCB 0x4F: BIT 1,A - Test bit 1 of A
func (cpu *Z80) BIT_1_A() int {
	cpu.testBit(cpu.reg.A, 1)
	return 8
}

// 0xCB 0x50: BIT 2,B - Test bit 2 of B
func (cpu *Z80) BIT_2_B() int {
	cpu.testBit(cpu.reg.B, 2)
	return 8
}

// 0xCB 0x51: BIT 2,C - Test bit 2 of C
func (cpu *Z80) BIT_2_C() int {
	cpu.testBit(cpu.reg.C, 2)
	return 8
}

// 0xCB 0x52: BIT 2,D - Test bit 2 of D
func (cpu *Z80) BIT_2_D() int {
	cpu.testBit(cpu.reg.D, 2)
	return 8
}

// 0xCB 0x53: BIT 2,E - Test bit 2 of E
func (cpu *Z80) BIT_2_E() int {
	cpu.testBit(cpu.reg.E, 2)
	return 8
}

// 0xCB 0x54: BIT 2,H - Test bit 2 of H
func (cpu *Z80) BIT_2_H() int {
	cpu.testBit(cpu.reg.H, 2)
	return 8
}

// 0xCB 0x55: BIT 2,L - Test bit 2 of L
func (cpu *Z80) BIT_2_L() int {
	cpu.testBit(cpu.reg.L, 2)
	return 8
}

// 0xCB 0x56: BIT 2,(HL) - Test bit 2 of value at address HL
func (cpu *Z80) BIT_2_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 2)
	return 12
}

// 0xCB 0x57: BIT 2,A - Test bit 2 of A
func (cpu *Z80) BIT_2_A() int {
	cpu.testBit(cpu.reg.A, 2)
	return 8
}

// 0xCB 0x58: BIT 3,B - Test bit 3 of B
func (cpu *Z80) BIT_3_B() int {
	cpu.testBit(cpu.reg.B, 3)
	return 8
}

// 0xCB 0x59: BIT 3,C - Test bit 3 of C
func (cpu *Z80) BIT_3_C() int {
	cpu.testBit(cpu.reg.C, 3)
	return 8
}

// 0xCB 0x5A: BIT 3,D - Test bit 3 of D
func (cpu *Z80) BIT_3_D() int {
	cpu.testBit(cpu.reg.D, 3)
	return 8
}

// 0xCB 0x5B: BIT 3,E - Test bit 3 of E
func (cpu *Z80) BIT_3_E() int {
	cpu.testBit(cpu.reg.E, 3)
	return 8
}

// 0xCB 0x5C: BIT 3,H - Test bit 3 of H
func (cpu *Z80) BIT_3_H() int {
	cpu.testBit(cpu.reg.H, 3)
	return 8
}

// 0xCB 0x5D: BIT 3,L - Test bit 3 of L
func (cpu *Z80) BIT_3_L() int {
	cpu.testBit(cpu.reg.L, 3)
	return 8
}

// 0xCB 0x5E: BIT 3,(HL) - Test bit 3 of value at address HL
func (cpu *Z80) BIT_3_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 3)
	return 12
}

// 0xCB 0x5F: BIT 3,A - Test bit 3 of A
func (cpu *Z80) BIT_3_A() int {
	cpu.testBit(cpu.reg.A, 3)
	return 8
}

// 0xCB 0x60: BIT 4,B - Test bit 4 of B
func (cpu *Z80) BIT_4_B() int {
	cpu.testBit(cpu.reg.B, 4)
	return 8
}

// 0xCB 0x61: BIT 4,C - Test bit 4 of C
func (cpu *Z80) BIT_4_C() int {
	cpu.testBit(cpu.reg.C, 4)
	return 8
}

// 0xCB 0x62: BIT 4,D - Test bit 4 of D
func (cpu *Z80) BIT_4_D() int {
	cpu.testBit(cpu.reg.D, 4)
	return 8
}

// 0xCB 0x63: BIT 4,E - Test bit 4 of E
func (cpu *Z80) BIT_4_E() int {
	cpu.testBit(cpu.reg.E, 4)
	return 8
}

// 0xCB 0x64: BIT 4,H - Test bit 4 of H
func (cpu *Z80) BIT_4_H() int {
	cpu.testBit(cpu.reg.H, 4)
	return 8
}

// 0xCB 0x65: BIT 4,L - Test bit 4 of L
func (cpu *Z80) BIT_4_L() int {
	cpu.testBit(cpu.reg.L, 4)
	return 8
}

// 0xCB 0x66: BIT 4,(HL) - Test bit 4 of value at address HL
func (cpu *Z80) BIT_4_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 4)
	return 12
}

// 0xCB 0x67: BIT 4,A - Test bit 4 of A
func (cpu *Z80) BIT_4_A() int {
	cpu.testBit(cpu.reg.A, 4)
	return 8
}

// 0xCB 0x68: BIT 5,B - Test bit 5 of B
func (cpu *Z80) BIT_5_B() int {
	cpu.testBit(cpu.reg.B, 5)
	return 8
}

// 0xCB 0x69: BIT 5,C - Test bit 5 of C
func (cpu *Z80) BIT_5_C() int {
	cpu.testBit(cpu.reg.C, 5)
	return 8
}

// 0xCB 0x6A: BIT 5,D - Test bit 5 of D
func (cpu *Z80) BIT_5_D() int {
	cpu.testBit(cpu.reg.D, 5)
	return 8
}

// 0xCB 0x6B: BIT 5,E - Test bit 5 of E
func (cpu *Z80) BIT_5_E() int {
	cpu.testBit(cpu.reg.E, 5)
	return 8
}

// 0xCB 0x6C: BIT 5,H - Test bit 5 of H
func (cpu *Z80) BIT_5_H() int {
	cpu.testBit(cpu.reg.H, 5)
	return 8
}

// 0xCB 0x6D: BIT 5,L - Test bit 5 of L
func (cpu *Z80) BIT_5_L() int {
	cpu.testBit(cpu.reg.L, 5)
	return 8
}

// 0xCB 0x6E: BIT 5,(HL) - Test bit 5 of value at address HL
func (cpu *Z80) BIT_5_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 5)
	return 12
}

// 0xCB 0x6F: BIT 5,A - Test bit 5 of A
func (cpu *Z80) BIT_5_A() int {
	cpu.testBit(cpu.reg.A, 5)
	return 8
}

// 0xCB 0x70: BIT 6,B - Test bit 6 of B
func (cpu *Z80) BIT_6_B() int {
	cpu.testBit(cpu.reg.B, 6)
	return 8
}

// 0xCB 0x71: BIT 6,C - Test bit 6 of C
func (cpu *Z80) BIT_6_C() int {
	cpu.testBit(cpu.reg.C, 6)
	return 8
}

// 0xCB 0x72: BIT 6,D - Test bit 6 of D
func (cpu *Z80) BIT_6_D() int {
	cpu.testBit(cpu.reg.D, 6)
	return 8
}

// 0xCB 0x73: BIT 6,E - Test bit 6 of E
func (cpu *Z80) BIT_6_E() int {
	cpu.testBit(cpu.reg.E, 6)
	return 8
}

// 0xCB 0x74: BIT 6,H - Test bit 6 of H
func (cpu *Z80) BIT_6_H() int {
	cpu.testBit(cpu.reg.H, 6)
	return 8
}

// 0xCB 0x75: BIT 6,L - Test bit 6 of L
func (cpu *Z80) BIT_6_L() int {
	cpu.testBit(cpu.reg.L, 6)
	return 8
}

// 0xCB 0x76: BIT 6,(HL) - Test bit 6 of value at address HL
func (cpu *Z80) BIT_6_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 6)
	return 12
}

// 0xCB 0x77: BIT 6,A - Test bit 6 of A
func (cpu *Z80) BIT_6_A() int {
	cpu.testBit(cpu.reg.A, 6)
	return 8
}

// 0xCB 0x78: BIT 7,B - Test bit 7 of B
func (cpu *Z80) BIT_7_B() int {
	cpu.testBit(cpu.reg.B, 7)
	return 8
}

// 0xCB 0x79: BIT 7,C - Test bit 7 of C
func (cpu *Z80) BIT_7_C() int {
	cpu.testBit(cpu.reg.C, 7)
	return 8
}

// 0xCB 0x7A: BIT 7,D - Test bit 7 of D
func (cpu *Z80) BIT_7_D() int {
	cpu.testBit(cpu.reg.D, 7)
	return 8
}

// 0xCB 0x7B: BIT 7,E - Test bit 7 of E
func (cpu *Z80) BIT_7_E() int {
	cpu.testBit(cpu.reg.E, 7)
	return 8
}

// 0xCB 0x7C: BIT 7,H - Test bit 7 of H
func (cpu *Z80) BIT_7_H() int {
	cpu.testBit(cpu.reg.H, 7)
	return 8
}

// 0xCB 0x7D: BIT 7,L - Test bit 7 of L
func (cpu *Z80) BIT_7_L() int {
	cpu.testBit(cpu.reg.L, 7)
	return 8
}

// 0xCB 0x7E: BIT 7,(HL) - Test bit 7 of value at address HL
func (cpu *Z80) BIT_7_HL() int {
	value := cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.testBit(value, 7)
	return 12
}

// 0xCB 0x7F: BIT 7,A - Test bit 7 of A
func (cpu *Z80) BIT_7_A() int {
	cpu.testBit(cpu.reg.A, 7)
	return 8
}
