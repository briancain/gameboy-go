package cpu

// This file contains CB-prefixed shift and swap instructions:
// - SLA r (Shift Left Arithmetic)
// - SRA r (Shift Right Arithmetic)
// - SRL r (Shift Right Logical)
// - SWAP r (Swap upper and lower nibbles)

// Shift Left Arithmetic Instructions

// 0xCB 0x20: SLA B - Shift B left arithmetic (b0=0)
func (cpu *Z80) SLA_B() int {
	// Get the highest bit before shift
	carry := (cpu.reg.B & 0x80) >> 7

	// Shift left
	cpu.reg.B = cpu.reg.B << 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x21: SLA C - Shift C left arithmetic (b0=0)
func (cpu *Z80) SLA_C() int {
	// Get the highest bit before shift
	carry := (cpu.reg.C & 0x80) >> 7

	// Shift left
	cpu.reg.C = cpu.reg.C << 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x22: SLA D - Shift D left arithmetic (b0=0)
func (cpu *Z80) SLA_D() int {
	// Get the highest bit before shift
	carry := (cpu.reg.D & 0x80) >> 7

	// Shift left
	cpu.reg.D = cpu.reg.D << 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x23: SLA E - Shift E left arithmetic (b0=0)
func (cpu *Z80) SLA_E() int {
	// Get the highest bit before shift
	carry := (cpu.reg.E & 0x80) >> 7

	// Shift left
	cpu.reg.E = cpu.reg.E << 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x24: SLA H - Shift H left arithmetic (b0=0)
func (cpu *Z80) SLA_H() int {
	// Get the highest bit before shift
	carry := (cpu.reg.H & 0x80) >> 7

	// Shift left
	cpu.reg.H = cpu.reg.H << 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x25: SLA L - Shift L left arithmetic (b0=0)
func (cpu *Z80) SLA_L() int {
	// Get the highest bit before shift
	carry := (cpu.reg.L & 0x80) >> 7

	// Shift left
	cpu.reg.L = cpu.reg.L << 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x26: SLA (HL) - Shift value at address HL left arithmetic (b0=0)
func (cpu *Z80) SLA_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Get the highest bit before shift
	carry := (value & 0x80) >> 7

	// Shift left
	value = value << 1

	// Write back to memory
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 16
}

// 0xCB 0x27: SLA A - Shift A left arithmetic (b0=0)
func (cpu *Z80) SLA_A() int {
	// Get the highest bit before shift
	carry := (cpu.reg.A & 0x80) >> 7

	// Shift left
	cpu.reg.A = cpu.reg.A << 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// Shift Right Arithmetic Instructions

// 0xCB 0x28: SRA B - Shift B right arithmetic (b7=b7)
func (cpu *Z80) SRA_B() int {
	// Get the lowest bit before shift
	carry := cpu.reg.B & 0x01

	// Get the highest bit to preserve it
	highBit := cpu.reg.B & 0x80

	// Shift right
	cpu.reg.B = (cpu.reg.B >> 1) | highBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x29: SRA C - Shift C right arithmetic (b7=b7)
func (cpu *Z80) SRA_C() int {
	// Get the lowest bit before shift
	carry := cpu.reg.C & 0x01

	// Get the highest bit to preserve it
	highBit := cpu.reg.C & 0x80

	// Shift right
	cpu.reg.C = (cpu.reg.C >> 1) | highBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x2A: SRA D - Shift D right arithmetic (b7=b7)
func (cpu *Z80) SRA_D() int {
	// Get the lowest bit before shift
	carry := cpu.reg.D & 0x01

	// Get the highest bit to preserve it
	highBit := cpu.reg.D & 0x80

	// Shift right
	cpu.reg.D = (cpu.reg.D >> 1) | highBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x2B: SRA E - Shift E right arithmetic (b7=b7)
func (cpu *Z80) SRA_E() int {
	// Get the lowest bit before shift
	carry := cpu.reg.E & 0x01

	// Get the highest bit to preserve it
	highBit := cpu.reg.E & 0x80

	// Shift right
	cpu.reg.E = (cpu.reg.E >> 1) | highBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x2C: SRA H - Shift H right arithmetic (b7=b7)
func (cpu *Z80) SRA_H() int {
	// Get the lowest bit before shift
	carry := cpu.reg.H & 0x01

	// Get the highest bit to preserve it
	highBit := cpu.reg.H & 0x80

	// Shift right
	cpu.reg.H = (cpu.reg.H >> 1) | highBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x2D: SRA L - Shift L right arithmetic (b7=b7)
func (cpu *Z80) SRA_L() int {
	// Get the lowest bit before shift
	carry := cpu.reg.L & 0x01

	// Get the highest bit to preserve it
	highBit := cpu.reg.L & 0x80

	// Shift right
	cpu.reg.L = (cpu.reg.L >> 1) | highBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x2E: SRA (HL) - Shift value at address HL right arithmetic (b7=b7)
func (cpu *Z80) SRA_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Get the lowest bit before shift
	carry := value & 0x01

	// Get the highest bit to preserve it
	highBit := value & 0x80

	// Shift right
	value = (value >> 1) | highBit

	// Write back to memory
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 16
}

// 0xCB 0x2F: SRA A - Shift A right arithmetic (b7=b7)
func (cpu *Z80) SRA_A() int {
	// Get the lowest bit before shift
	carry := cpu.reg.A & 0x01

	// Get the highest bit to preserve it
	highBit := cpu.reg.A & 0x80

	// Shift right
	cpu.reg.A = (cpu.reg.A >> 1) | highBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// Swap Instructions

// 0xCB 0x30: SWAP B - Swap upper and lower nibbles of B
func (cpu *Z80) SWAP_B() int {
	// Swap nibbles
	cpu.reg.B = ((cpu.reg.B & 0x0F) << 4) | ((cpu.reg.B & 0xF0) >> 4)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 8
}

// 0xCB 0x31: SWAP C - Swap upper and lower nibbles of C
func (cpu *Z80) SWAP_C() int {
	// Swap nibbles
	cpu.reg.C = ((cpu.reg.C & 0x0F) << 4) | ((cpu.reg.C & 0xF0) >> 4)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 8
}

// 0xCB 0x32: SWAP D - Swap upper and lower nibbles of D
func (cpu *Z80) SWAP_D() int {
	// Swap nibbles
	cpu.reg.D = ((cpu.reg.D & 0x0F) << 4) | ((cpu.reg.D & 0xF0) >> 4)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 8
}

// 0xCB 0x33: SWAP E - Swap upper and lower nibbles of E
func (cpu *Z80) SWAP_E() int {
	// Swap nibbles
	cpu.reg.E = ((cpu.reg.E & 0x0F) << 4) | ((cpu.reg.E & 0xF0) >> 4)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 8
}

// 0xCB 0x34: SWAP H - Swap upper and lower nibbles of H
func (cpu *Z80) SWAP_H() int {
	// Swap nibbles
	cpu.reg.H = ((cpu.reg.H & 0x0F) << 4) | ((cpu.reg.H & 0xF0) >> 4)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 8
}

// 0xCB 0x35: SWAP L - Swap upper and lower nibbles of L
func (cpu *Z80) SWAP_L() int {
	// Swap nibbles
	cpu.reg.L = ((cpu.reg.L & 0x0F) << 4) | ((cpu.reg.L & 0xF0) >> 4)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 8
}

// 0xCB 0x36: SWAP (HL) - Swap upper and lower nibbles of value at address HL
func (cpu *Z80) SWAP_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Swap nibbles
	value = ((value & 0x0F) << 4) | ((value & 0xF0) >> 4)

	// Write back to memory
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 16
}

// 0xCB 0x37: SWAP A - Swap upper and lower nibbles of A
func (cpu *Z80) SWAP_A() int {
	// Swap nibbles
	cpu.reg.A = ((cpu.reg.A & 0x0F) << 4) | ((cpu.reg.A & 0xF0) >> 4)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	return 8
}

// Shift Right Logical Instructions

// 0xCB 0x38: SRL B - Shift B right logical (b7=0)
func (cpu *Z80) SRL_B() int {
	// Get the lowest bit before shift
	carry := cpu.reg.B & 0x01

	// Shift right
	cpu.reg.B = cpu.reg.B >> 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x39: SRL C - Shift C right logical (b7=0)
func (cpu *Z80) SRL_C() int {
	// Get the lowest bit before shift
	carry := cpu.reg.C & 0x01

	// Shift right
	cpu.reg.C = cpu.reg.C >> 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x3A: SRL D - Shift D right logical (b7=0)
func (cpu *Z80) SRL_D() int {
	// Get the lowest bit before shift
	carry := cpu.reg.D & 0x01

	// Shift right
	cpu.reg.D = cpu.reg.D >> 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x3B: SRL E - Shift E right logical (b7=0)
func (cpu *Z80) SRL_E() int {
	// Get the lowest bit before shift
	carry := cpu.reg.E & 0x01

	// Shift right
	cpu.reg.E = cpu.reg.E >> 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x3C: SRL H - Shift H right logical (b7=0)
func (cpu *Z80) SRL_H() int {
	// Get the lowest bit before shift
	carry := cpu.reg.H & 0x01

	// Shift right
	cpu.reg.H = cpu.reg.H >> 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x3D: SRL L - Shift L right logical (b7=0)
func (cpu *Z80) SRL_L() int {
	// Get the lowest bit before shift
	carry := cpu.reg.L & 0x01

	// Shift right
	cpu.reg.L = cpu.reg.L >> 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x3E: SRL (HL) - Shift value at address HL right logical (b7=0)
func (cpu *Z80) SRL_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Get the lowest bit before shift
	carry := value & 0x01

	// Shift right
	value = value >> 1

	// Write back to memory
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 16
}

// 0xCB 0x3F: SRL A - Shift A right logical (b7=0)
func (cpu *Z80) SRL_A() int {
	// Get the lowest bit before shift
	carry := cpu.reg.A & 0x01

	// Shift right
	cpu.reg.A = cpu.reg.A >> 1

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}
