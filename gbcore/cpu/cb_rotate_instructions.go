package gbcore

// This file contains CB-prefixed rotation instructions:
// - RLC r (Rotate register left with carry)
// - RRC r (Rotate register right with carry)
// - RL r (Rotate register left through carry)
// - RR r (Rotate register right through carry)

// CB-prefixed Rotate Left Instructions

// 0xCB 0x00: RLC B - Rotate B left with carry
func (cpu *Z80) RLC_B() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.B & 0x80) >> 7

	// Rotate left
	cpu.reg.B = (cpu.reg.B << 1) | carry

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

// 0xCB 0x01: RLC C - Rotate C left with carry
func (cpu *Z80) RLC_C() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.C & 0x80) >> 7

	// Rotate left
	cpu.reg.C = (cpu.reg.C << 1) | carry

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

// 0xCB 0x02: RLC D - Rotate D left with carry
func (cpu *Z80) RLC_D() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.D & 0x80) >> 7

	// Rotate left
	cpu.reg.D = (cpu.reg.D << 1) | carry

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

// 0xCB 0x03: RLC E - Rotate E left with carry
func (cpu *Z80) RLC_E() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.E & 0x80) >> 7

	// Rotate left
	cpu.reg.E = (cpu.reg.E << 1) | carry

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

// 0xCB 0x04: RLC H - Rotate H left with carry
func (cpu *Z80) RLC_H() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.H & 0x80) >> 7

	// Rotate left
	cpu.reg.H = (cpu.reg.H << 1) | carry

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

// 0xCB 0x05: RLC L - Rotate L left with carry
func (cpu *Z80) RLC_L() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.L & 0x80) >> 7

	// Rotate left
	cpu.reg.L = (cpu.reg.L << 1) | carry

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

// 0xCB 0x06: RLC (HL) - Rotate value at address HL left with carry
func (cpu *Z80) RLC_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Get the highest bit before rotation
	carry := (value & 0x80) >> 7

	// Rotate left
	value = (value << 1) | carry

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

// 0xCB 0x07: RLC A - Rotate A left with carry
func (cpu *Z80) RLC_A() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.A & 0x80) >> 7

	// Rotate left
	cpu.reg.A = (cpu.reg.A << 1) | carry

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

// CB-prefixed Rotate Right Instructions

// 0xCB 0x08: RRC B - Rotate B right with carry
func (cpu *Z80) RRC_B() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.B & 0x01

	// Rotate right
	cpu.reg.B = (cpu.reg.B >> 1) | (carry << 7)

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

// 0xCB 0x09: RRC C - Rotate C right with carry
func (cpu *Z80) RRC_C() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.C & 0x01

	// Rotate right
	cpu.reg.C = (cpu.reg.C >> 1) | (carry << 7)

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

// 0xCB 0x0A: RRC D - Rotate D right with carry
func (cpu *Z80) RRC_D() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.D & 0x01

	// Rotate right
	cpu.reg.D = (cpu.reg.D >> 1) | (carry << 7)

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

// 0xCB 0x0B: RRC E - Rotate E right with carry
func (cpu *Z80) RRC_E() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.E & 0x01

	// Rotate right
	cpu.reg.E = (cpu.reg.E >> 1) | (carry << 7)

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

// 0xCB 0x0C: RRC H - Rotate H right with carry
func (cpu *Z80) RRC_H() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.H & 0x01

	// Rotate right
	cpu.reg.H = (cpu.reg.H >> 1) | (carry << 7)

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

// 0xCB 0x0D: RRC L - Rotate L right with carry
func (cpu *Z80) RRC_L() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.L & 0x01

	// Rotate right
	cpu.reg.L = (cpu.reg.L >> 1) | (carry << 7)

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

// 0xCB 0x0E: RRC (HL) - Rotate value at address HL right with carry
func (cpu *Z80) RRC_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Get the lowest bit before rotation
	carry := value & 0x01

	// Rotate right
	value = (value >> 1) | (carry << 7)

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

// 0xCB 0x0F: RRC A - Rotate A right with carry
func (cpu *Z80) RRC_A() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.A & 0x01

	// Rotate right
	cpu.reg.A = (cpu.reg.A >> 1) | (carry << 7)

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

// CB-prefixed Rotate Left through Carry Instructions

// 0xCB 0x10: RL B - Rotate B left through carry
func (cpu *Z80) RL_B() int {
	// Get the highest bit before rotation
	oldCarry := (cpu.reg.B & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	cpu.reg.B = (cpu.reg.B << 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x11: RL C - Rotate C left through carry
func (cpu *Z80) RL_C() int {
	// Get the highest bit before rotation
	oldCarry := (cpu.reg.C & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	cpu.reg.C = (cpu.reg.C << 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x12: RL D - Rotate D left through carry
func (cpu *Z80) RL_D() int {
	// Get the highest bit before rotation
	oldCarry := (cpu.reg.D & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	cpu.reg.D = (cpu.reg.D << 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x13: RL E - Rotate E left through carry
func (cpu *Z80) RL_E() int {
	// Get the highest bit before rotation
	oldCarry := (cpu.reg.E & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	cpu.reg.E = (cpu.reg.E << 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x14: RL H - Rotate H left through carry
func (cpu *Z80) RL_H() int {
	// Get the highest bit before rotation
	oldCarry := (cpu.reg.H & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	cpu.reg.H = (cpu.reg.H << 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x15: RL L - Rotate L left through carry
func (cpu *Z80) RL_L() int {
	// Get the highest bit before rotation
	oldCarry := (cpu.reg.L & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	cpu.reg.L = (cpu.reg.L << 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x16: RL (HL) - Rotate value at address HL left through carry
func (cpu *Z80) RL_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Get the highest bit before rotation
	oldCarry := (value & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	value = (value << 1) | newBit

	// Write back to memory
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 16
}

// 0xCB 0x17: RL A - Rotate A left through carry
func (cpu *Z80) RL_A() int {
	// Get the highest bit before rotation
	oldCarry := (cpu.reg.A & 0x80) >> 7

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 1
	}

	// Rotate left through carry
	cpu.reg.A = (cpu.reg.A << 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// CB-prefixed Rotate Right through Carry Instructions

// 0xCB 0x18: RR B - Rotate B right through carry
func (cpu *Z80) RR_B() int {
	// Get the lowest bit before rotation
	oldCarry := cpu.reg.B & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	cpu.reg.B = (cpu.reg.B >> 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.B == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x19: RR C - Rotate C right through carry
func (cpu *Z80) RR_C() int {
	// Get the lowest bit before rotation
	oldCarry := cpu.reg.C & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	cpu.reg.C = (cpu.reg.C >> 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.C == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x1A: RR D - Rotate D right through carry
func (cpu *Z80) RR_D() int {
	// Get the lowest bit before rotation
	oldCarry := cpu.reg.D & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	cpu.reg.D = (cpu.reg.D >> 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.D == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x1B: RR E - Rotate E right through carry
func (cpu *Z80) RR_E() int {
	// Get the lowest bit before rotation
	oldCarry := cpu.reg.E & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	cpu.reg.E = (cpu.reg.E >> 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.E == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x1C: RR H - Rotate H right through carry
func (cpu *Z80) RR_H() int {
	// Get the lowest bit before rotation
	oldCarry := cpu.reg.H & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	cpu.reg.H = (cpu.reg.H >> 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.H == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x1D: RR L - Rotate L right through carry
func (cpu *Z80) RR_L() int {
	// Get the lowest bit before rotation
	oldCarry := cpu.reg.L & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	cpu.reg.L = (cpu.reg.L >> 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.L == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}

// 0xCB 0x1E: RR (HL) - Rotate value at address HL right through carry
func (cpu *Z80) RR_HL() int {
	// Get value from memory
	address := cpu.reg.GetHL()
	value := cpu.mmu.ReadByte(address)

	// Get the lowest bit before rotation
	oldCarry := value & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	value = (value >> 1) | newBit

	// Write back to memory
	cpu.mmu.WriteByte(address, value)

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if value == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 16
}

// 0xCB 0x1F: RR A - Rotate A right through carry
func (cpu *Z80) RR_A() int {
	// Get the lowest bit before rotation
	oldCarry := cpu.reg.A & 0x01

	// Get current carry flag
	newBit := byte(0)
	if cpu.reg.GetFlag(FLAG_C) {
		newBit = 0x80
	}

	// Rotate right through carry
	cpu.reg.A = (cpu.reg.A >> 1) | newBit

	// Set flags
	cpu.reg.F = 0

	// Zero flag
	if cpu.reg.A == 0 {
		cpu.reg.SetFlag(FLAG_Z)
	}

	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}

	return 8
}
