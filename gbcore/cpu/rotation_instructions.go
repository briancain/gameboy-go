package gbcore

// This file contains standard rotation instructions:
// - RLCA (Rotate A left)
// - RRCA (Rotate A right)
// - RLA (Rotate A left through carry)
// - RRA (Rotate A right through carry)

// Standard Rotation Instructions

// 0x07: RLCA - Rotate A left (simple)
func (cpu *Z80) RLCA() int {
	// Get the highest bit before rotation
	carry := (cpu.reg.A & 0x80) >> 7
	
	// Rotate left
	cpu.reg.A = (cpu.reg.A << 1) | carry
	
	// Set flags
	cpu.reg.F = 0
	
	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}
	
	return 4
}

// 0x0F: RRCA - Rotate A right (simple)
func (cpu *Z80) RRCA() int {
	// Get the lowest bit before rotation
	carry := cpu.reg.A & 0x01
	
	// Rotate right
	cpu.reg.A = (cpu.reg.A >> 1) | (carry << 7)
	
	// Set flags
	cpu.reg.F = 0
	
	// Carry flag - set to the bit that was shifted out
	if carry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}
	
	return 4
}

// 0x17: RLA - Rotate A left through carry (simple)
func (cpu *Z80) RLA() int {
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
	
	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}
	
	return 4
}

// 0x1F: RRA - Rotate A right through carry (simple)
func (cpu *Z80) RRA() int {
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
	
	// Carry flag - set to the bit that was shifted out
	if oldCarry == 1 {
		cpu.reg.SetFlag(FLAG_C)
	}
	
	return 4
}
