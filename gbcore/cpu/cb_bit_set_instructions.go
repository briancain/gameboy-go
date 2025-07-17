package gbcore

// This file contains CB-prefixed bit set instructions:
// - SET b,r (Set bit b of register r)
// 
// These instructions set (set to 1) a specific bit in a register or memory location.
// No flags are affected by these operations.

// SET Instructions

// 0xCB 0xC0: SET 0,B - Set bit 0 of B
func (cpu *Z80) SET_0_B() int {
	cpu.setBit(&cpu.reg.B, 0)
	return 8
}

// 0xCB 0xC1: SET 0,C - Set bit 0 of C
func (cpu *Z80) SET_0_C() int {
	cpu.setBit(&cpu.reg.C, 0)
	return 8
}

// 0xCB 0xC2: SET 0,D - Set bit 0 of D
func (cpu *Z80) SET_0_D() int {
	cpu.setBit(&cpu.reg.D, 0)
	return 8
}

// 0xCB 0xC3: SET 0,E - Set bit 0 of E
func (cpu *Z80) SET_0_E() int {
	cpu.setBit(&cpu.reg.E, 0)
	return 8
}

// 0xCB 0xC4: SET 0,H - Set bit 0 of H
func (cpu *Z80) SET_0_H() int {
	cpu.setBit(&cpu.reg.H, 0)
	return 8
}

// 0xCB 0xC5: SET 0,L - Set bit 0 of L
func (cpu *Z80) SET_0_L() int {
	cpu.setBit(&cpu.reg.L, 0)
	return 8
}

// 0xCB 0xC6: SET 0,(HL) - Set bit 0 of value at address HL
func (cpu *Z80) SET_0_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 0)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xC7: SET 0,A - Set bit 0 of A
func (cpu *Z80) SET_0_A() int {
	cpu.setBit(&cpu.reg.A, 0)
	return 8
}

// 0xCB 0xC8: SET 1,B - Set bit 1 of B
func (cpu *Z80) SET_1_B() int {
	cpu.setBit(&cpu.reg.B, 1)
	return 8
}

// 0xCB 0xC9: SET 1,C - Set bit 1 of C
func (cpu *Z80) SET_1_C() int {
	cpu.setBit(&cpu.reg.C, 1)
	return 8
}

// 0xCB 0xCA: SET 1,D - Set bit 1 of D
func (cpu *Z80) SET_1_D() int {
	cpu.setBit(&cpu.reg.D, 1)
	return 8
}

// 0xCB 0xCB: SET 1,E - Set bit 1 of E
func (cpu *Z80) SET_1_E() int {
	cpu.setBit(&cpu.reg.E, 1)
	return 8
}

// 0xCB 0xCC: SET 1,H - Set bit 1 of H
func (cpu *Z80) SET_1_H() int {
	cpu.setBit(&cpu.reg.H, 1)
	return 8
}

// 0xCB 0xCD: SET 1,L - Set bit 1 of L
func (cpu *Z80) SET_1_L() int {
	cpu.setBit(&cpu.reg.L, 1)
	return 8
}

// 0xCB 0xCE: SET 1,(HL) - Set bit 1 of value at address HL
func (cpu *Z80) SET_1_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 1)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xCF: SET 1,A - Set bit 1 of A
func (cpu *Z80) SET_1_A() int {
	cpu.setBit(&cpu.reg.A, 1)
	return 8
}

// 0xCB 0xD0: SET 2,B - Set bit 2 of B
func (cpu *Z80) SET_2_B() int {
	cpu.setBit(&cpu.reg.B, 2)
	return 8
}

// 0xCB 0xD1: SET 2,C - Set bit 2 of C
func (cpu *Z80) SET_2_C() int {
	cpu.setBit(&cpu.reg.C, 2)
	return 8
}

// 0xCB 0xD2: SET 2,D - Set bit 2 of D
func (cpu *Z80) SET_2_D() int {
	cpu.setBit(&cpu.reg.D, 2)
	return 8
}

// 0xCB 0xD3: SET 2,E - Set bit 2 of E
func (cpu *Z80) SET_2_E() int {
	cpu.setBit(&cpu.reg.E, 2)
	return 8
}

// 0xCB 0xD4: SET 2,H - Set bit 2 of H
func (cpu *Z80) SET_2_H() int {
	cpu.setBit(&cpu.reg.H, 2)
	return 8
}

// 0xCB 0xD5: SET 2,L - Set bit 2 of L
func (cpu *Z80) SET_2_L() int {
	cpu.setBit(&cpu.reg.L, 2)
	return 8
}

// 0xCB 0xD6: SET 2,(HL) - Set bit 2 of value at address HL
func (cpu *Z80) SET_2_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 2)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xD7: SET 2,A - Set bit 2 of A
func (cpu *Z80) SET_2_A() int {
	cpu.setBit(&cpu.reg.A, 2)
	return 8
}

// 0xCB 0xD8: SET 3,B - Set bit 3 of B
func (cpu *Z80) SET_3_B() int {
	cpu.setBit(&cpu.reg.B, 3)
	return 8
}

// 0xCB 0xD9: SET 3,C - Set bit 3 of C
func (cpu *Z80) SET_3_C() int {
	cpu.setBit(&cpu.reg.C, 3)
	return 8
}

// 0xCB 0xDA: SET 3,D - Set bit 3 of D
func (cpu *Z80) SET_3_D() int {
	cpu.setBit(&cpu.reg.D, 3)
	return 8
}

// 0xCB 0xDB: SET 3,E - Set bit 3 of E
func (cpu *Z80) SET_3_E() int {
	cpu.setBit(&cpu.reg.E, 3)
	return 8
}

// 0xCB 0xDC: SET 3,H - Set bit 3 of H
func (cpu *Z80) SET_3_H() int {
	cpu.setBit(&cpu.reg.H, 3)
	return 8
}

// 0xCB 0xDD: SET 3,L - Set bit 3 of L
func (cpu *Z80) SET_3_L() int {
	cpu.setBit(&cpu.reg.L, 3)
	return 8
}

// 0xCB 0xDE: SET 3,(HL) - Set bit 3 of value at address HL
func (cpu *Z80) SET_3_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 3)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xDF: SET 3,A - Set bit 3 of A
func (cpu *Z80) SET_3_A() int {
	cpu.setBit(&cpu.reg.A, 3)
	return 8
}

// 0xCB 0xE0: SET 4,B - Set bit 4 of B
func (cpu *Z80) SET_4_B() int {
	cpu.setBit(&cpu.reg.B, 4)
	return 8
}

// 0xCB 0xE1: SET 4,C - Set bit 4 of C
func (cpu *Z80) SET_4_C() int {
	cpu.setBit(&cpu.reg.C, 4)
	return 8
}

// 0xCB 0xE2: SET 4,D - Set bit 4 of D
func (cpu *Z80) SET_4_D() int {
	cpu.setBit(&cpu.reg.D, 4)
	return 8
}

// 0xCB 0xE3: SET 4,E - Set bit 4 of E
func (cpu *Z80) SET_4_E() int {
	cpu.setBit(&cpu.reg.E, 4)
	return 8
}

// 0xCB 0xE4: SET 4,H - Set bit 4 of H
func (cpu *Z80) SET_4_H() int {
	cpu.setBit(&cpu.reg.H, 4)
	return 8
}

// 0xCB 0xE5: SET 4,L - Set bit 4 of L
func (cpu *Z80) SET_4_L() int {
	cpu.setBit(&cpu.reg.L, 4)
	return 8
}

// 0xCB 0xE6: SET 4,(HL) - Set bit 4 of value at address HL
func (cpu *Z80) SET_4_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 4)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xE7: SET 4,A - Set bit 4 of A
func (cpu *Z80) SET_4_A() int {
	cpu.setBit(&cpu.reg.A, 4)
	return 8
}

// 0xCB 0xE8: SET 5,B - Set bit 5 of B
func (cpu *Z80) SET_5_B() int {
	cpu.setBit(&cpu.reg.B, 5)
	return 8
}

// 0xCB 0xE9: SET 5,C - Set bit 5 of C
func (cpu *Z80) SET_5_C() int {
	cpu.setBit(&cpu.reg.C, 5)
	return 8
}

// 0xCB 0xEA: SET 5,D - Set bit 5 of D
func (cpu *Z80) SET_5_D() int {
	cpu.setBit(&cpu.reg.D, 5)
	return 8
}

// 0xCB 0xEB: SET 5,E - Set bit 5 of E
func (cpu *Z80) SET_5_E() int {
	cpu.setBit(&cpu.reg.E, 5)
	return 8
}

// 0xCB 0xEC: SET 5,H - Set bit 5 of H
func (cpu *Z80) SET_5_H() int {
	cpu.setBit(&cpu.reg.H, 5)
	return 8
}

// 0xCB 0xED: SET 5,L - Set bit 5 of L
func (cpu *Z80) SET_5_L() int {
	cpu.setBit(&cpu.reg.L, 5)
	return 8
}

// 0xCB 0xEE: SET 5,(HL) - Set bit 5 of value at address HL
func (cpu *Z80) SET_5_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 5)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xEF: SET 5,A - Set bit 5 of A
func (cpu *Z80) SET_5_A() int {
	cpu.setBit(&cpu.reg.A, 5)
	return 8
}

// 0xCB 0xF0: SET 6,B - Set bit 6 of B
func (cpu *Z80) SET_6_B() int {
	cpu.setBit(&cpu.reg.B, 6)
	return 8
}

// 0xCB 0xF1: SET 6,C - Set bit 6 of C
func (cpu *Z80) SET_6_C() int {
	cpu.setBit(&cpu.reg.C, 6)
	return 8
}

// 0xCB 0xF2: SET 6,D - Set bit 6 of D
func (cpu *Z80) SET_6_D() int {
	cpu.setBit(&cpu.reg.D, 6)
	return 8
}

// 0xCB 0xF3: SET 6,E - Set bit 6 of E
func (cpu *Z80) SET_6_E() int {
	cpu.setBit(&cpu.reg.E, 6)
	return 8
}

// 0xCB 0xF4: SET 6,H - Set bit 6 of H
func (cpu *Z80) SET_6_H() int {
	cpu.setBit(&cpu.reg.H, 6)
	return 8
}

// 0xCB 0xF5: SET 6,L - Set bit 6 of L
func (cpu *Z80) SET_6_L() int {
	cpu.setBit(&cpu.reg.L, 6)
	return 8
}

// 0xCB 0xF6: SET 6,(HL) - Set bit 6 of value at address HL
func (cpu *Z80) SET_6_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 6)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xF7: SET 6,A - Set bit 6 of A
func (cpu *Z80) SET_6_A() int {
	cpu.setBit(&cpu.reg.A, 6)
	return 8
}

// 0xCB 0xF8: SET 7,B - Set bit 7 of B
func (cpu *Z80) SET_7_B() int {
	cpu.setBit(&cpu.reg.B, 7)
	return 8
}

// 0xCB 0xF9: SET 7,C - Set bit 7 of C
func (cpu *Z80) SET_7_C() int {
	cpu.setBit(&cpu.reg.C, 7)
	return 8
}

// 0xCB 0xFA: SET 7,D - Set bit 7 of D
func (cpu *Z80) SET_7_D() int {
	cpu.setBit(&cpu.reg.D, 7)
	return 8
}

// 0xCB 0xFB: SET 7,E - Set bit 7 of E
func (cpu *Z80) SET_7_E() int {
	cpu.setBit(&cpu.reg.E, 7)
	return 8
}

// 0xCB 0xFC: SET 7,H - Set bit 7 of H
func (cpu *Z80) SET_7_H() int {
	cpu.setBit(&cpu.reg.H, 7)
	return 8
}

// 0xCB 0xFD: SET 7,L - Set bit 7 of L
func (cpu *Z80) SET_7_L() int {
	cpu.setBit(&cpu.reg.L, 7)
	return 8
}

// 0xCB 0xFE: SET 7,(HL) - Set bit 7 of value at address HL
func (cpu *Z80) SET_7_HL() int {
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)
	cpu.setBit(&value, 7)
	cpu.mmu.WriteByte(address, value)
	return 16
}

// 0xCB 0xFF: SET 7,A - Set bit 7 of A
func (cpu *Z80) SET_7_A() int {
	cpu.setBit(&cpu.reg.A, 7)
	return 8
}
