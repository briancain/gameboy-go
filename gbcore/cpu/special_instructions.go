package gbcore

// This file contains special instructions:
// - DAA (Decimal Adjust Accumulator)
// - CPL (Complement A)
// - CCF (Complement Carry Flag)
// - SCF (Set Carry Flag)
// - Memory access instructions (LD (addr),A, LD A,(addr))

// Special Instructions

// 0x27: DAA - Decimal Adjust Accumulator
func (cpu *Z80) DAA() int {
	var adjust byte = 0

	if cpu.reg.GetFlag(FLAG_H) || (!cpu.reg.GetFlag(FLAG_N) && (cpu.reg.A&0x0F) > 9) {
		adjust |= 0x06
	}

	if cpu.reg.GetFlag(FLAG_C) || (!cpu.reg.GetFlag(FLAG_N) && cpu.reg.A > 0x99) {
		adjust |= 0x60
		cpu.reg.SetFlag(FLAG_C)
	}

	if cpu.reg.GetFlag(FLAG_N) {
		cpu.reg.A -= adjust
	} else {
		cpu.reg.A += adjust
	}

	// Reset half carry flag
	cpu.reg.ClearFlag(FLAG_H)

	// Set zero flag if result is zero
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	} else {
		cpu.reg.ClearFlag(FLAG_Z)
	}

	return 4
}

// 0x2F: CPL - Complement A (flip all bits)
func (cpu *Z80) CPL() int {
	cpu.reg.A = ^cpu.reg.A

	// Set flags
	cpu.reg.SetFlag(FLAG_N | FLAG_H)

	return 4
}

// 0x3F: CCF - Complement Carry Flag
func (cpu *Z80) CCF() int {
	// Flip carry flag
	if cpu.reg.GetFlag(FLAG_C) {
		cpu.reg.ClearFlag(FLAG_C)
	} else {
		cpu.reg.SetFlag(FLAG_C)
	}

	// Reset flags
	cpu.reg.ClearFlag(FLAG_N | FLAG_H)

	return 4
}

// 0x37: SCF - Set Carry Flag
func (cpu *Z80) SCF() int {
	// Set carry flag
	cpu.reg.SetFlag(FLAG_C)

	// Reset flags
	cpu.reg.ClearFlag(FLAG_N | FLAG_H)

	return 4
}

// Memory Access Instructions

// 0x02: LD (BC),A - Store A in address pointed to by BC
func (cpu *Z80) LD_BC_A() int {
	cpu.mmu.WriteByte(cpu.reg.GetBC(), cpu.reg.A)
	return 8
}

// 0x12: LD (DE),A - Store A in address pointed to by DE
func (cpu *Z80) LD_DE_A() int {
	cpu.mmu.WriteByte(cpu.reg.GetDE(), cpu.reg.A)
	return 8
}

// 0x0A: LD A,(BC) - Load A from address pointed to by BC
func (cpu *Z80) LD_A_BC() int {
	cpu.reg.A = cpu.mmu.ReadByte(cpu.reg.GetBC())
	return 8
}

// 0x1A: LD A,(DE) - Load A from address pointed to by DE
func (cpu *Z80) LD_A_DE() int {
	cpu.reg.A = cpu.mmu.ReadByte(cpu.reg.GetDE())
	return 8
}

// 0xEA: LD (a16),A - Store A in address a16
func (cpu *Z80) LD_a16_A() int {
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.mmu.WriteByte(address, cpu.reg.A)
	return 16
}

// 0xFA: LD A,(a16) - Load A from address a16
func (cpu *Z80) LD_A_a16() int {
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.A = cpu.mmu.ReadByte(address)
	return 16
}

// 0xE0: LDH (a8),A - Store A in address 0xFF00+a8
func (cpu *Z80) LDH_a8_A() int {
	offset := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	cpu.mmu.WriteByte(0xFF00+uint16(offset), cpu.reg.A)
	return 12
}

// 0xF0: LDH A,(a8) - Load A from address 0xFF00+a8
func (cpu *Z80) LDH_A_a8() int {
	offset := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	cpu.reg.A = cpu.mmu.ReadByte(0xFF00 + uint16(offset))
	return 12
}

// 0xE2: LD (C),A - Store A in address 0xFF00+C
func (cpu *Z80) LD_C_mem_A() int {
	cpu.mmu.WriteByte(0xFF00+uint16(cpu.reg.C), cpu.reg.A)
	return 8
}

// 0xF2: LD A,(C) - Load A from address 0xFF00+C
func (cpu *Z80) LD_A_C_mem() int {
	cpu.reg.A = cpu.mmu.ReadByte(0xFF00 + uint16(cpu.reg.C))
	return 8
}

// 0x22: LD (HL+),A - Store A in address HL and increment HL
func (cpu *Z80) LD_HLI_A() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.A)
	cpu.reg.SetHL(cpu.reg.GetHL() + 1)
	return 8
}

// 0x2A: LD A,(HL+) - Load A from address HL and increment HL
func (cpu *Z80) LD_A_HLI() int {
	cpu.reg.A = cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.reg.SetHL(cpu.reg.GetHL() + 1)
	return 8
}

// 0x32: LD (HL-),A - Store A in address HL and decrement HL
func (cpu *Z80) LD_HLD_A() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.A)
	cpu.reg.SetHL(cpu.reg.GetHL() - 1)
	return 8
}

// 0x3A: LD A,(HL-) - Load A from address HL and decrement HL
func (cpu *Z80) LD_A_HLD() int {
	cpu.reg.A = cpu.mmu.ReadByte(cpu.reg.GetHL())
	cpu.reg.SetHL(cpu.reg.GetHL() - 1)
	return 8
}
