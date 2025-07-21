package cpu

// This file contains control flow instructions:
// - JP (Jump instructions)
// - JR (Jump relative instructions)
// - CALL (Call subroutine instructions)
// - RET (Return from subroutine instructions)
// - RST (Restart instructions)

// Control Flow Instructions

// 0xC3: JP a16 - Jump to address a16
func (cpu *Z80) JP_a16() int {
	// Get jump address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC = address

	return 16
}

// 0xC2: JP NZ,a16 - Jump to address a16 if Z flag is reset
func (cpu *Z80) JP_NZ_a16() int {
	// Get jump address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if !cpu.reg.GetFlag(FLAG_Z) {
		cpu.reg.PC = address
		return 16
	}

	return 12
}

// 0xCA: JP Z,a16 - Jump to address a16 if Z flag is set
func (cpu *Z80) JP_Z_a16() int {
	// Get jump address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if cpu.reg.GetFlag(FLAG_Z) {
		cpu.reg.PC = address
		return 16
	}

	return 12
}

// 0xD2: JP NC,a16 - Jump to address a16 if C flag is reset
func (cpu *Z80) JP_NC_a16() int {
	// Get jump address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if !cpu.reg.GetFlag(FLAG_C) {
		cpu.reg.PC = address
		return 16
	}

	return 12
}

// 0xDA: JP C,a16 - Jump to address a16 if C flag is set
func (cpu *Z80) JP_C_a16() int {
	// Get jump address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if cpu.reg.GetFlag(FLAG_C) {
		cpu.reg.PC = address
		return 16
	}

	return 12
}

// 0xE9: JP (HL) - Jump to address contained in HL
func (cpu *Z80) JP_HL() int {
	cpu.reg.PC = cpu.reg.GetHL()
	return 4
}

// 0x18: JR r8 - Jump relative by signed 8-bit immediate
func (cpu *Z80) JR_r8() int {
	// Get signed offset
	offset := int8(cpu.mmu.ReadByte(cpu.reg.PC))
	cpu.reg.PC++

	// Jump
	cpu.reg.PC = uint16(int32(cpu.reg.PC) + int32(offset))

	return 12
}

// 0x20: JR NZ,r8 - Jump relative by signed 8-bit immediate if Z flag is reset
func (cpu *Z80) JR_NZ_r8() int {
	// Get signed offset
	offset := int8(cpu.mmu.ReadByte(cpu.reg.PC))
	cpu.reg.PC++

	// Check condition
	if !cpu.reg.GetFlag(FLAG_Z) {
		cpu.reg.PC = uint16(int32(cpu.reg.PC) + int32(offset))
		return 12
	}

	return 8
}

// 0x28: JR Z,r8 - Jump relative by signed 8-bit immediate if Z flag is set
func (cpu *Z80) JR_Z_r8() int {
	// Get signed offset
	offset := int8(cpu.mmu.ReadByte(cpu.reg.PC))
	cpu.reg.PC++

	// Check condition
	if cpu.reg.GetFlag(FLAG_Z) {
		cpu.reg.PC = uint16(int32(cpu.reg.PC) + int32(offset))
		return 12
	}

	return 8
}

// 0x30: JR NC,r8 - Jump relative by signed 8-bit immediate if C flag is reset
func (cpu *Z80) JR_NC_r8() int {
	// Get signed offset
	offset := int8(cpu.mmu.ReadByte(cpu.reg.PC))
	cpu.reg.PC++

	// Check condition
	if !cpu.reg.GetFlag(FLAG_C) {
		cpu.reg.PC = uint16(int32(cpu.reg.PC) + int32(offset))
		return 12
	}

	return 8
}

// 0x38: JR C,r8 - Jump relative by signed 8-bit immediate if C flag is set
func (cpu *Z80) JR_C_r8() int {
	// Get signed offset
	offset := int8(cpu.mmu.ReadByte(cpu.reg.PC))
	cpu.reg.PC++

	// Check condition
	if cpu.reg.GetFlag(FLAG_C) {
		cpu.reg.PC = uint16(int32(cpu.reg.PC) + int32(offset))
		return 12
	}

	return 8
}

// 0xCD: CALL a16 - Call subroutine at address a16
func (cpu *Z80) CALL_a16() int {
	// Get call address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to call address
	cpu.reg.PC = address

	return 24
}

// 0xC4: CALL NZ,a16 - Call subroutine at address a16 if Z flag is reset
func (cpu *Z80) CALL_NZ_a16() int {
	// Get call address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if !cpu.reg.GetFlag(FLAG_Z) {
		// Push current PC onto stack
		cpu.reg.SP -= 2
		cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

		// Jump to call address
		cpu.reg.PC = address

		return 24
	}

	return 12
}

// 0xCC: CALL Z,a16 - Call subroutine at address a16 if Z flag is set
func (cpu *Z80) CALL_Z_a16() int {
	// Get call address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if cpu.reg.GetFlag(FLAG_Z) {
		// Push current PC onto stack
		cpu.reg.SP -= 2
		cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

		// Jump to call address
		cpu.reg.PC = address

		return 24
	}

	return 12
}

// 0xD4: CALL NC,a16 - Call subroutine at address a16 if C flag is reset
func (cpu *Z80) CALL_NC_a16() int {
	// Get call address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if !cpu.reg.GetFlag(FLAG_C) {
		// Push current PC onto stack
		cpu.reg.SP -= 2
		cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

		// Jump to call address
		cpu.reg.PC = address

		return 24
	}

	return 12
}

// 0xDC: CALL C,a16 - Call subroutine at address a16 if C flag is set
func (cpu *Z80) CALL_C_a16() int {
	// Get call address
	address := cpu.mmu.ReadWord(cpu.reg.PC)
	cpu.reg.PC += 2

	// Check condition
	if cpu.reg.GetFlag(FLAG_C) {
		// Push current PC onto stack
		cpu.reg.SP -= 2
		cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

		// Jump to call address
		cpu.reg.PC = address

		return 24
	}

	return 12
}

// 0xC9: RET - Return from subroutine
func (cpu *Z80) RET() int {
	// Pop address from stack
	cpu.reg.PC = cpu.mmu.ReadWord(cpu.reg.SP)
	cpu.reg.SP += 2

	return 16
}

// 0xC0: RET NZ - Return from subroutine if Z flag is reset
func (cpu *Z80) RET_NZ() int {
	// Check condition
	if !cpu.reg.GetFlag(FLAG_Z) {
		// Pop address from stack
		cpu.reg.PC = cpu.mmu.ReadWord(cpu.reg.SP)
		cpu.reg.SP += 2

		return 20
	}

	return 8
}

// 0xC8: RET Z - Return from subroutine if Z flag is set
func (cpu *Z80) RET_Z() int {
	// Check condition
	if cpu.reg.GetFlag(FLAG_Z) {
		// Pop address from stack
		cpu.reg.PC = cpu.mmu.ReadWord(cpu.reg.SP)
		cpu.reg.SP += 2

		return 20
	}

	return 8
}

// 0xD0: RET NC - Return from subroutine if C flag is reset
func (cpu *Z80) RET_NC() int {
	// Check condition
	if !cpu.reg.GetFlag(FLAG_C) {
		// Pop address from stack
		cpu.reg.PC = cpu.mmu.ReadWord(cpu.reg.SP)
		cpu.reg.SP += 2

		return 20
	}

	return 8
}

// 0xD8: RET C - Return from subroutine if C flag is set
func (cpu *Z80) RET_C() int {
	// Check condition
	if cpu.reg.GetFlag(FLAG_C) {
		// Pop address from stack
		cpu.reg.PC = cpu.mmu.ReadWord(cpu.reg.SP)
		cpu.reg.SP += 2

		return 20
	}

	return 8
}

// 0xD9: RETI - Return from interrupt
func (cpu *Z80) RETI() int {
	// Pop address from stack
	cpu.reg.PC = cpu.mmu.ReadWord(cpu.reg.SP)
	cpu.reg.SP += 2

	// Enable interrupts
	cpu.interruptMaster = true

	return 16
}

// RST instructions - Call to restart address

// 0xC7: RST 00H - Call to address 0x0000
func (cpu *Z80) RST_00H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0000

	return 16
}

// 0xCF: RST 08H - Call to address 0x0008
func (cpu *Z80) RST_08H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0008

	return 16
}

// 0xD7: RST 10H - Call to address 0x0010
func (cpu *Z80) RST_10H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0010

	return 16
}

// 0xDF: RST 18H - Call to address 0x0018
func (cpu *Z80) RST_18H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0018

	return 16
}

// 0xE7: RST 20H - Call to address 0x0020
func (cpu *Z80) RST_20H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0020

	return 16
}

// 0xEF: RST 28H - Call to address 0x0028
func (cpu *Z80) RST_28H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0028

	return 16
}

// 0xF7: RST 30H - Call to address 0x0030
func (cpu *Z80) RST_30H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0030

	return 16
}

// 0xFF: RST 38H - Call to address 0x0038
func (cpu *Z80) RST_38H() int {
	// Push current PC onto stack
	cpu.reg.SP -= 2
	cpu.mmu.WriteWord(cpu.reg.SP, cpu.reg.PC)

	// Jump to restart address
	cpu.reg.PC = 0x0038

	return 16
}
