package cpu

// This file contains basic CPU instructions:
// - NOP, HALT, STOP
// - DI, EI (Disable/Enable Interrupts)
// - 8-bit register load instructions (LD r,r', LD r,n)
// - Memory access instructions for registers (LD r,(HL), LD (HL),r)

// Basic instructions

// 0x00: NOP - No operation
func (cpu *Z80) NOP() int {
	// Do nothing
	return 4 // 1 machine cycle = 4 clock cycles
}

// 0x76: HALT - Halt the CPU until an interrupt occurs
func (cpu *Z80) HALT() int {
	// Check for HALT bug: If IME=0 and IE & IF != 0, the HALT bug occurs
	interruptFlag := cpu.mmu.ReadByte(0xFF0F) & 0x1F
	interruptEnable := cpu.mmu.ReadByte(0xFFFF) & 0x1F

	if !cpu.interruptMaster && (interruptFlag&interruptEnable) != 0 {
		// HALT bug: When interrupts are disabled (IME=0) and there are pending interrupts (IE & IF != 0),
		// the instruction following HALT is "skipped" (PC doesn't increment after fetching the opcode)
		// This causes the next instruction to be executed twice
		cpu.haltBug = true
	} else {
		cpu.halted = true
	}

	return 4
}

// 0x10: STOP - Halt CPU & LCD display until button pressed
func (cpu *Z80) STOP() int {
	cpu.stopped = true
	// Read next byte (usually 0x00)
	cpu.reg.PC++
	return 4
}

// 0xF3: DI - Disable interrupts
func (cpu *Z80) DI() int {
	// Interrupts are disabled after the next instruction
	cpu.interruptDisableScheduled = true
	return 4
}

// 0xFB: EI - Enable interrupts
func (cpu *Z80) EI() int {
	// Interrupts are enabled after the next instruction
	cpu.interruptEnableScheduled = true
	return 4
}

// 8-bit Load Instructions

// 0x06: LD B,d8 - Load immediate 8-bit value into B
func (cpu *Z80) LD_B_d8() int {
	cpu.reg.B = cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	return 8
}

// 0x0E: LD C,d8 - Load immediate 8-bit value into C
func (cpu *Z80) LD_C_d8() int {
	cpu.reg.C = cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	return 8
}

// 0x16: LD D,d8 - Load immediate 8-bit value into D
func (cpu *Z80) LD_D_d8() int {
	cpu.reg.D = cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	return 8
}

// 0x1E: LD E,d8 - Load immediate 8-bit value into E
func (cpu *Z80) LD_E_d8() int {
	cpu.reg.E = cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	return 8
}

// 0x26: LD H,d8 - Load immediate 8-bit value into H
func (cpu *Z80) LD_H_d8() int {
	cpu.reg.H = cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	return 8
}

// 0x2E: LD L,d8 - Load immediate 8-bit value into L
func (cpu *Z80) LD_L_d8() int {
	cpu.reg.L = cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	return 8
}

// 0x3E: LD A,d8 - Load immediate 8-bit value into A
func (cpu *Z80) LD_A_d8() int {
	cpu.reg.A = cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	return 8
}

// 0x40-0x7F: LD r,r' - Load register r' into register r
// 0x40: LD B,B
func (cpu *Z80) LD_B_B() int {
	// B = B (no operation needed)
	return 4
}

// 0x41: LD B,C
func (cpu *Z80) LD_B_C() int {
	cpu.reg.B = cpu.reg.C
	return 4
}

// 0x42: LD B,D
func (cpu *Z80) LD_B_D() int {
	cpu.reg.B = cpu.reg.D
	return 4
}

// 0x43: LD B,E
func (cpu *Z80) LD_B_E() int {
	cpu.reg.B = cpu.reg.E
	return 4
}

// 0x44: LD B,H
func (cpu *Z80) LD_B_H() int {
	cpu.reg.B = cpu.reg.H
	return 4
}

// 0x45: LD B,L
func (cpu *Z80) LD_B_L() int {
	cpu.reg.B = cpu.reg.L
	return 4
}

// 0x46: LD B,(HL)
func (cpu *Z80) LD_B_HL() int {
	cpu.reg.B = cpu.mmu.ReadByte(cpu.reg.GetHL())
	return 8
}

// 0x47: LD B,A
func (cpu *Z80) LD_B_A() int {
	cpu.reg.B = cpu.reg.A
	return 4
}

// 0x48: LD C,B
func (cpu *Z80) LD_C_B() int {
	cpu.reg.C = cpu.reg.B
	return 4
}

// 0x49: LD C,C
func (cpu *Z80) LD_C_C() int {
	// C = C (no operation needed)
	return 4
}

// 0x4A: LD C,D
func (cpu *Z80) LD_C_D() int {
	cpu.reg.C = cpu.reg.D
	return 4
}

// 0x4B: LD C,E
func (cpu *Z80) LD_C_E() int {
	cpu.reg.C = cpu.reg.E
	return 4
}

// 0x4C: LD C,H
func (cpu *Z80) LD_C_H() int {
	cpu.reg.C = cpu.reg.H
	return 4
}

// 0x4D: LD C,L
func (cpu *Z80) LD_C_L() int {
	cpu.reg.C = cpu.reg.L
	return 4
}

// 0x4E: LD C,(HL)
func (cpu *Z80) LD_C_HL() int {
	cpu.reg.C = cpu.mmu.ReadByte(cpu.reg.GetHL())
	return 8
}

// 0x4F: LD C,A
func (cpu *Z80) LD_C_A() int {
	cpu.reg.C = cpu.reg.A
	return 4
}

// 0x50: LD D,B
func (cpu *Z80) LD_D_B() int {
	cpu.reg.D = cpu.reg.B
	return 4
}

// 0x51: LD D,C
func (cpu *Z80) LD_D_C() int {
	cpu.reg.D = cpu.reg.C
	return 4
}

// 0x52: LD D,D
func (cpu *Z80) LD_D_D() int {
	// D = D (no operation needed)
	return 4
}

// 0x53: LD D,E
func (cpu *Z80) LD_D_E() int {
	cpu.reg.D = cpu.reg.E
	return 4
}

// 0x54: LD D,H
func (cpu *Z80) LD_D_H() int {
	cpu.reg.D = cpu.reg.H
	return 4
}

// 0x55: LD D,L
func (cpu *Z80) LD_D_L() int {
	cpu.reg.D = cpu.reg.L
	return 4
}

// 0x56: LD D,(HL)
func (cpu *Z80) LD_D_HL() int {
	cpu.reg.D = cpu.mmu.ReadByte(cpu.reg.GetHL())
	return 8
}

// 0x57: LD D,A
func (cpu *Z80) LD_D_A() int {
	cpu.reg.D = cpu.reg.A
	return 4
}

// 0x58: LD E,B
func (cpu *Z80) LD_E_B() int {
	cpu.reg.E = cpu.reg.B
	return 4
}

// 0x59: LD E,C
func (cpu *Z80) LD_E_C() int {
	cpu.reg.E = cpu.reg.C
	return 4
}

// 0x5A: LD E,D
func (cpu *Z80) LD_E_D() int {
	cpu.reg.E = cpu.reg.D
	return 4
}

// 0x5B: LD E,E
func (cpu *Z80) LD_E_E() int {
	// E = E (no operation needed)
	return 4
}

// 0x5C: LD E,H
func (cpu *Z80) LD_E_H() int {
	cpu.reg.E = cpu.reg.H
	return 4
}

// 0x5D: LD E,L
func (cpu *Z80) LD_E_L() int {
	cpu.reg.E = cpu.reg.L
	return 4
}

// 0x5E: LD E,(HL)
func (cpu *Z80) LD_E_HL() int {
	cpu.reg.E = cpu.mmu.ReadByte(cpu.reg.GetHL())
	return 8
}

// 0x5F: LD E,A
func (cpu *Z80) LD_E_A() int {
	cpu.reg.E = cpu.reg.A
	return 4
}

// 0x60: LD H,B
func (cpu *Z80) LD_H_B() int {
	cpu.reg.H = cpu.reg.B
	return 4
}

// 0x61: LD H,C
func (cpu *Z80) LD_H_C() int {
	cpu.reg.H = cpu.reg.C
	return 4
}

// 0x62: LD H,D
func (cpu *Z80) LD_H_D() int {
	cpu.reg.H = cpu.reg.D
	return 4
}

// 0x63: LD H,E
func (cpu *Z80) LD_H_E() int {
	cpu.reg.H = cpu.reg.E
	return 4
}

// 0x64: LD H,H
func (cpu *Z80) LD_H_H() int {
	// H = H (no operation needed)
	return 4
}

// 0x65: LD H,L
func (cpu *Z80) LD_H_L() int {
	cpu.reg.H = cpu.reg.L
	return 4
}

// 0x66: LD H,(HL)
func (cpu *Z80) LD_H_HL() int {
	cpu.reg.H = cpu.mmu.ReadByte(cpu.reg.GetHL())
	return 8
}

// 0x67: LD H,A
func (cpu *Z80) LD_H_A() int {
	cpu.reg.H = cpu.reg.A
	return 4
}

// 0x68: LD L,B
func (cpu *Z80) LD_L_B() int {
	cpu.reg.L = cpu.reg.B
	return 4
}

// 0x69: LD L,C
func (cpu *Z80) LD_L_C() int {
	cpu.reg.L = cpu.reg.C
	return 4
}

// 0x6A: LD L,D
func (cpu *Z80) LD_L_D() int {
	cpu.reg.L = cpu.reg.D
	return 4
}

// 0x6B: LD L,E
func (cpu *Z80) LD_L_E() int {
	cpu.reg.L = cpu.reg.E
	return 4
}

// 0x6C: LD L,H
func (cpu *Z80) LD_L_H() int {
	cpu.reg.L = cpu.reg.H
	return 4
}

// 0x6D: LD L,L
func (cpu *Z80) LD_L_L() int {
	// L = L (no operation needed)
	return 4
}

// 0x6E: LD L,(HL)
func (cpu *Z80) LD_L_HL() int {
	cpu.reg.L = cpu.mmu.ReadByte(cpu.reg.GetHL())
	return 8
}

// 0x6F: LD L,A
func (cpu *Z80) LD_L_A() int {
	cpu.reg.L = cpu.reg.A
	return 4
}

// 0x70: LD (HL),B
func (cpu *Z80) LD_HL_B() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.B)
	return 8
}

// 0x71: LD (HL),C
func (cpu *Z80) LD_HL_C() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.C)
	return 8
}

// 0x72: LD (HL),D
func (cpu *Z80) LD_HL_D() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.D)
	return 8
}

// 0x73: LD (HL),E
func (cpu *Z80) LD_HL_E() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.E)
	return 8
}

// 0x74: LD (HL),H
func (cpu *Z80) LD_HL_H() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.H)
	return 8
}

// 0x75: LD (HL),L
func (cpu *Z80) LD_HL_L() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.L)
	return 8
}

// 0x77: LD (HL),A
func (cpu *Z80) LD_HL_A() int {
	cpu.mmu.WriteByte(cpu.reg.GetHL(), cpu.reg.A)
	return 8
}

// 0x78: LD A,B
func (cpu *Z80) LD_A_B() int {
	cpu.reg.A = cpu.reg.B
	return 4
}

// 0x79: LD A,C
func (cpu *Z80) LD_A_C() int {
	cpu.reg.A = cpu.reg.C
	return 4
}

// 0x7A: LD A,D
func (cpu *Z80) LD_A_D() int {
	cpu.reg.A = cpu.reg.D
	return 4
}

// 0x7B: LD A,E
func (cpu *Z80) LD_A_E() int {
	cpu.reg.A = cpu.reg.E
	return 4
}

// 0x7C: LD A,H
func (cpu *Z80) LD_A_H() int {
	cpu.reg.A = cpu.reg.H
	return 4
}

// 0x7D: LD A,L
func (cpu *Z80) LD_A_L() int {
	cpu.reg.A = cpu.reg.L
	return 4
}

// 0x7E: LD A,(HL)
func (cpu *Z80) LD_A_HL() int {
	cpu.reg.A = cpu.mmu.ReadByte(cpu.reg.GetHL())
	return 8
}

// 0x7F: LD A,A
func (cpu *Z80) LD_A_A() int {
	// A = A (no operation needed)
	return 4
}

// 0x36: LD (HL),d8
func (cpu *Z80) LD_HL_d8() int {
	value := cpu.mmu.ReadByte(cpu.reg.PC)
	cpu.reg.PC++
	cpu.mmu.WriteByte(cpu.reg.GetHL(), value)
	return 12
}
