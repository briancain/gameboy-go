package cpu

// This file contains CB-prefixed bit reset instructions:
// - RES b,r (Reset bit b of register r)
//
// These instructions reset (set to 0) a specific bit in a register or memory location.
// No flags are affected by these operations.

// RES Instructions

// 0xCB 0x80: RES 0,B - Reset bit 0 of B
func (cpu *Z80) RES_0_B() int {
	cpu.resetBit(&cpu.reg.B, 0)
	return 8
}

// 0xCB 0x81: RES 0,C - Reset bit 0 of C
func (cpu *Z80) RES_0_C() int {
	cpu.resetBit(&cpu.reg.C, 0)
	return 8
}

// 0xCB 0x82: RES 0,D - Reset bit 0 of D
func (cpu *Z80) RES_0_D() int {
	cpu.resetBit(&cpu.reg.D, 0)
	return 8
}

// 0xCB 0x83: RES 0,E - Reset bit 0 of E
func (cpu *Z80) RES_0_E() int {
	cpu.resetBit(&cpu.reg.E, 0)
	return 8
}

// 0xCB 0x84: RES 0,H - Reset bit 0 of H
func (cpu *Z80) RES_0_H() int {
	cpu.resetBit(&cpu.reg.H, 0)
	return 8
}

// 0xCB 0x85: RES 0,L - Reset bit 0 of L
func (cpu *Z80) RES_0_L() int {
	cpu.resetBit(&cpu.reg.L, 0)
	return 8
}

// 0xCB 0x86: RES 0,(HL) - Reset bit 0 of value at address HL
func (cpu *Z80) RES_0_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 0)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0x87: RES 0,A - Reset bit 0 of A
func (cpu *Z80) RES_0_A() int {
	cpu.resetBit(&cpu.reg.A, 0)
	return 8
}

// 0xCB 0x88: RES 1,B - Reset bit 1 of B
func (cpu *Z80) RES_1_B() int {
	cpu.resetBit(&cpu.reg.B, 1)
	return 8
}

// 0xCB 0x89: RES 1,C - Reset bit 1 of C
func (cpu *Z80) RES_1_C() int {
	cpu.resetBit(&cpu.reg.C, 1)
	return 8
}

// 0xCB 0x8A: RES 1,D - Reset bit 1 of D
func (cpu *Z80) RES_1_D() int {
	cpu.resetBit(&cpu.reg.D, 1)
	return 8
}

// 0xCB 0x8B: RES 1,E - Reset bit 1 of E
func (cpu *Z80) RES_1_E() int {
	cpu.resetBit(&cpu.reg.E, 1)
	return 8
}

// 0xCB 0x8C: RES 1,H - Reset bit 1 of H
func (cpu *Z80) RES_1_H() int {
	cpu.resetBit(&cpu.reg.H, 1)
	return 8
}

// 0xCB 0x8D: RES 1,L - Reset bit 1 of L
func (cpu *Z80) RES_1_L() int {
	cpu.resetBit(&cpu.reg.L, 1)
	return 8
}

// 0xCB 0x8E: RES 1,(HL) - Reset bit 1 of value at address HL
func (cpu *Z80) RES_1_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 1)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0x8F: RES 1,A - Reset bit 1 of A
func (cpu *Z80) RES_1_A() int {
	cpu.resetBit(&cpu.reg.A, 1)
	return 8
}

// 0xCB 0x90: RES 2,B - Reset bit 2 of B
func (cpu *Z80) RES_2_B() int {
	cpu.resetBit(&cpu.reg.B, 2)
	return 8
}

// 0xCB 0x91: RES 2,C - Reset bit 2 of C
func (cpu *Z80) RES_2_C() int {
	cpu.resetBit(&cpu.reg.C, 2)
	return 8
}

// 0xCB 0x92: RES 2,D - Reset bit 2 of D
func (cpu *Z80) RES_2_D() int {
	cpu.resetBit(&cpu.reg.D, 2)
	return 8
}

// 0xCB 0x93: RES 2,E - Reset bit 2 of E
func (cpu *Z80) RES_2_E() int {
	cpu.resetBit(&cpu.reg.E, 2)
	return 8
}

// 0xCB 0x94: RES 2,H - Reset bit 2 of H
func (cpu *Z80) RES_2_H() int {
	cpu.resetBit(&cpu.reg.H, 2)
	return 8
}

// 0xCB 0x95: RES 2,L - Reset bit 2 of L
func (cpu *Z80) RES_2_L() int {
	cpu.resetBit(&cpu.reg.L, 2)
	return 8
}

// 0xCB 0x96: RES 2,(HL) - Reset bit 2 of value at address HL
func (cpu *Z80) RES_2_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 2)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0x97: RES 2,A - Reset bit 2 of A
func (cpu *Z80) RES_2_A() int {
	cpu.resetBit(&cpu.reg.A, 2)
	return 8
}

// 0xCB 0x98: RES 3,B - Reset bit 3 of B
func (cpu *Z80) RES_3_B() int {
	cpu.resetBit(&cpu.reg.B, 3)
	return 8
}

// 0xCB 0x99: RES 3,C - Reset bit 3 of C
func (cpu *Z80) RES_3_C() int {
	cpu.resetBit(&cpu.reg.C, 3)
	return 8
}

// 0xCB 0x9A: RES 3,D - Reset bit 3 of D
func (cpu *Z80) RES_3_D() int {
	cpu.resetBit(&cpu.reg.D, 3)
	return 8
}

// 0xCB 0x9B: RES 3,E - Reset bit 3 of E
func (cpu *Z80) RES_3_E() int {
	cpu.resetBit(&cpu.reg.E, 3)
	return 8
}

// 0xCB 0x9C: RES 3,H - Reset bit 3 of H
func (cpu *Z80) RES_3_H() int {
	cpu.resetBit(&cpu.reg.H, 3)
	return 8
}

// 0xCB 0x9D: RES 3,L - Reset bit 3 of L
func (cpu *Z80) RES_3_L() int {
	cpu.resetBit(&cpu.reg.L, 3)
	return 8
}

// 0xCB 0x9E: RES 3,(HL) - Reset bit 3 of value at address HL
func (cpu *Z80) RES_3_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 3)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0x9F: RES 3,A - Reset bit 3 of A
func (cpu *Z80) RES_3_A() int {
	cpu.resetBit(&cpu.reg.A, 3)
	return 8
}

// 0xCB 0xA0: RES 4,B - Reset bit 4 of B
func (cpu *Z80) RES_4_B() int {
	cpu.resetBit(&cpu.reg.B, 4)
	return 8
}

// 0xCB 0xA1: RES 4,C - Reset bit 4 of C
func (cpu *Z80) RES_4_C() int {
	cpu.resetBit(&cpu.reg.C, 4)
	return 8
}

// 0xCB 0xA2: RES 4,D - Reset bit 4 of D
func (cpu *Z80) RES_4_D() int {
	cpu.resetBit(&cpu.reg.D, 4)
	return 8
}

// 0xCB 0xA3: RES 4,E - Reset bit 4 of E
func (cpu *Z80) RES_4_E() int {
	cpu.resetBit(&cpu.reg.E, 4)
	return 8
}

// 0xCB 0xA4: RES 4,H - Reset bit 4 of H
func (cpu *Z80) RES_4_H() int {
	cpu.resetBit(&cpu.reg.H, 4)
	return 8
}

// 0xCB 0xA5: RES 4,L - Reset bit 4 of L
func (cpu *Z80) RES_4_L() int {
	cpu.resetBit(&cpu.reg.L, 4)
	return 8
}

// 0xCB 0xA6: RES 4,(HL) - Reset bit 4 of value at address HL
func (cpu *Z80) RES_4_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 4)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xA7: RES 4,A - Reset bit 4 of A
func (cpu *Z80) RES_4_A() int {
	cpu.resetBit(&cpu.reg.A, 4)
	return 8
}

// 0xCB 0xA8: RES 5,B - Reset bit 5 of B
func (cpu *Z80) RES_5_B() int {
	cpu.resetBit(&cpu.reg.B, 5)
	return 8
}

// 0xCB 0xA9: RES 5,C - Reset bit 5 of C
func (cpu *Z80) RES_5_C() int {
	cpu.resetBit(&cpu.reg.C, 5)
	return 8
}

// 0xCB 0xAA: RES 5,D - Reset bit 5 of D
func (cpu *Z80) RES_5_D() int {
	cpu.resetBit(&cpu.reg.D, 5)
	return 8
}

// 0xCB 0xAB: RES 5,E - Reset bit 5 of E
func (cpu *Z80) RES_5_E() int {
	cpu.resetBit(&cpu.reg.E, 5)
	return 8
}

// 0xCB 0xAC: RES 5,H - Reset bit 5 of H
func (cpu *Z80) RES_5_H() int {
	cpu.resetBit(&cpu.reg.H, 5)
	return 8
}

// 0xCB 0xAD: RES 5,L - Reset bit 5 of L
func (cpu *Z80) RES_5_L() int {
	cpu.resetBit(&cpu.reg.L, 5)
	return 8
}

// 0xCB 0xAE: RES 5,(HL) - Reset bit 5 of value at address HL
func (cpu *Z80) RES_5_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 5)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xAF: RES 5,A - Reset bit 5 of A
func (cpu *Z80) RES_5_A() int {
	cpu.resetBit(&cpu.reg.A, 5)
	return 8
}

// 0xCB 0xB0: RES 6,B - Reset bit 6 of B
func (cpu *Z80) RES_6_B() int {
	cpu.resetBit(&cpu.reg.B, 6)
	return 8
}

// 0xCB 0xB1: RES 6,C - Reset bit 6 of C
func (cpu *Z80) RES_6_C() int {
	cpu.resetBit(&cpu.reg.C, 6)
	return 8
}

// 0xCB 0xB2: RES 6,D - Reset bit 6 of D
func (cpu *Z80) RES_6_D() int {
	cpu.resetBit(&cpu.reg.D, 6)
	return 8
}

// 0xCB 0xB3: RES 6,E - Reset bit 6 of E
func (cpu *Z80) RES_6_E() int {
	cpu.resetBit(&cpu.reg.E, 6)
	return 8
}

// 0xCB 0xB4: RES 6,H - Reset bit 6 of H
func (cpu *Z80) RES_6_H() int {
	cpu.resetBit(&cpu.reg.H, 6)
	return 8
}

// 0xCB 0xB5: RES 6,L - Reset bit 6 of L
func (cpu *Z80) RES_6_L() int {
	cpu.resetBit(&cpu.reg.L, 6)
	return 8
}

// 0xCB 0xB6: RES 6,(HL) - Reset bit 6 of value at address HL
func (cpu *Z80) RES_6_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 6)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xB7: RES 6,A - Reset bit 6 of A
func (cpu *Z80) RES_6_A() int {
	cpu.resetBit(&cpu.reg.A, 6)
	return 8
}

// 0xCB 0xB8: RES 7,B - Reset bit 7 of B
func (cpu *Z80) RES_7_B() int {
	cpu.resetBit(&cpu.reg.B, 7)
	return 8
}

// 0xCB 0xB9: RES 7,C - Reset bit 7 of C
func (cpu *Z80) RES_7_C() int {
	cpu.resetBit(&cpu.reg.C, 7)
	return 8
}

// 0xCB 0xBA: RES 7,D - Reset bit 7 of D
func (cpu *Z80) RES_7_D() int {
	cpu.resetBit(&cpu.reg.D, 7)
	return 8
}

// 0xCB 0xBB: RES 7,E - Reset bit 7 of E
func (cpu *Z80) RES_7_E() int {
	cpu.resetBit(&cpu.reg.E, 7)
	return 8
}

// 0xCB 0xBC: RES 7,H - Reset bit 7 of H
func (cpu *Z80) RES_7_H() int {
	cpu.resetBit(&cpu.reg.H, 7)
	return 8
}

// 0xCB 0xBD: RES 7,L - Reset bit 7 of L
func (cpu *Z80) RES_7_L() int {
	cpu.resetBit(&cpu.reg.L, 7)
	return 8
}

// 0xCB 0xBE: RES 7,(HL) - Reset bit 7 of value at address HL
func (cpu *Z80) RES_7_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.resetBit(&value, 7)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xBF: RES 7,A - Reset bit 7 of A
func (cpu *Z80) RES_7_A() int {
	cpu.resetBit(&cpu.reg.A, 7)
	return 8
}
