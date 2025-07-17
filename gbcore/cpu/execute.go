package gbcore

import "log"

// Execute instruction based on opcode
func (cpu *Z80) executeInstruction(opcode byte) int {
	// Check for CB prefix
	if opcode == 0xCB {
		// Get the next byte for the CB-prefixed opcode
		cbOpcode := cpu.mmu.ReadByte(cpu.reg.PC)
		cpu.reg.PC++
		return cpu.executeCBInstruction(cbOpcode)
	}

	// Execute unprefixed instruction
	switch opcode {
	case 0x00: // NOP
		return cpu.NOP()
	case 0x01: // LD BC,d16
		return cpu.LD_BC_d16()
	case 0x02: // LD (BC),A
		return cpu.LD_BC_A()
	case 0x03: // INC BC
		return cpu.INC_BC()
	case 0x04: // INC B
		return cpu.INC_B()
	case 0x05: // DEC B
		return cpu.DEC_B()
	case 0x06: // LD B,d8
		return cpu.LD_B_d8()
	case 0x07: // RLCA
		return cpu.RLCA()
	case 0x08: // LD (a16),SP
		return cpu.LD_a16_SP()
	case 0x09: // ADD HL,BC
		return cpu.ADD_HL_BC()
	case 0x0A: // LD A,(BC)
		return cpu.LD_A_BC()
	case 0x0B: // DEC BC
		return cpu.DEC_BC()
	case 0x0C: // INC C
		return cpu.INC_C()
	case 0x0D: // DEC C
		return cpu.DEC_C()
	case 0x0E: // LD C,d8
		return cpu.LD_C_d8()
	case 0x0F: // RRCA
		return cpu.RRCA()
	case 0x10: // STOP
		return cpu.STOP()
	case 0x11: // LD DE,d16
		return cpu.LD_DE_d16()
	case 0x12: // LD (DE),A
		return cpu.LD_DE_A()
	case 0x13: // INC DE
		return cpu.INC_DE()
	case 0x14: // INC D
		return cpu.INC_D()
	case 0x15: // DEC D
		return cpu.DEC_D()
	case 0x16: // LD D,d8
		return cpu.LD_D_d8()
	case 0x17: // RLA
		return cpu.RLA()
	case 0x18: // JR r8
		return cpu.JR_r8()
	case 0x19: // ADD HL,DE
		return cpu.ADD_HL_DE()
	case 0x1A: // LD A,(DE)
		return cpu.LD_A_DE()
	case 0x1B: // DEC DE
		return cpu.DEC_DE()
	case 0x1C: // INC E
		return cpu.INC_E()
	case 0x1D: // DEC E
		return cpu.DEC_E()
	case 0x1E: // LD E,d8
		return cpu.LD_E_d8()
	case 0x1F: // RRA
		return cpu.RRA()
	case 0x20: // JR NZ,r8
		return cpu.JR_NZ_r8()
	case 0x21: // LD HL,d16
		return cpu.LD_HL_d16()
	case 0x22: // LD (HL+),A
		return cpu.LD_HLI_A()
	case 0x23: // INC HL
		return cpu.INC_HL()
	case 0x24: // INC H
		return cpu.INC_H()
	case 0x25: // DEC H
		return cpu.DEC_H()
	case 0x26: // LD H,d8
		return cpu.LD_H_d8()
	case 0x27: // DAA
		return cpu.DAA()
	case 0x28: // JR Z,r8
		return cpu.JR_Z_r8()
	case 0x29: // ADD HL,HL
		return cpu.ADD_HL_HL()
	case 0x2A: // LD A,(HL+)
		return cpu.LD_A_HLI()
	case 0x2B: // DEC HL
		return cpu.DEC_HL()
	case 0x2C: // INC L
		return cpu.INC_L()
	case 0x2D: // DEC L
		return cpu.DEC_L()
	case 0x2E: // LD L,d8
		return cpu.LD_L_d8()
	case 0x2F: // CPL
		return cpu.CPL()
	case 0x30: // JR NC,r8
		return cpu.JR_NC_r8()
	case 0x31: // LD SP,d16
		return cpu.LD_SP_d16()
	case 0x32: // LD (HL-),A
		return cpu.LD_HLD_A()
	case 0x33: // INC SP
		return cpu.INC_SP()
	case 0x34: // INC (HL)
		return cpu.INC_HL_()
	case 0x35: // DEC (HL)
		return cpu.DEC_HL_()
	case 0x36: // LD (HL),d8
		return cpu.LD_HL_d8()
	case 0x37: // SCF
		return cpu.SCF()
	case 0x38: // JR C,r8
		return cpu.JR_C_r8()
	case 0x39: // ADD HL,SP
		return cpu.ADD_HL_SP()
	case 0x3A: // LD A,(HL-)
		return cpu.LD_A_HLD()
	case 0x3B: // DEC SP
		return cpu.DEC_SP()
	case 0x3C: // INC A
		return cpu.INC_A()
	case 0x3D: // DEC A
		return cpu.DEC_A()
	case 0x3E: // LD A,d8
		return cpu.LD_A_d8()
	case 0x3F: // CCF
		return cpu.CCF()
	case 0x40: // LD B,B
		return cpu.LD_B_B()
	case 0x41: // LD B,C
		return cpu.LD_B_C()
	case 0x42: // LD B,D
		return cpu.LD_B_D()
	case 0x43: // LD B,E
		return cpu.LD_B_E()
	case 0x44: // LD B,H
		return cpu.LD_B_H()
	case 0x45: // LD B,L
		return cpu.LD_B_L()
	case 0x46: // LD B,(HL)
		return cpu.LD_B_HL()
	case 0x47: // LD B,A
		return cpu.LD_B_A()
	case 0x48: // LD C,B
		return cpu.LD_C_B()
	case 0x49: // LD C,C
		return cpu.LD_C_C()
	case 0x4A: // LD C,D
		return cpu.LD_C_D()
	case 0x4B: // LD C,E
		return cpu.LD_C_E()
	case 0x4C: // LD C,H
		return cpu.LD_C_H()
	case 0x4D: // LD C,L
		return cpu.LD_C_L()
	case 0x4E: // LD C,(HL)
		return cpu.LD_C_HL()
	case 0x4F: // LD C,A
		return cpu.LD_C_A()
	case 0x50: // LD D,B
		return cpu.LD_D_B()
	case 0x51: // LD D,C
		return cpu.LD_D_C()
	case 0x52: // LD D,D
		return cpu.LD_D_D()
	case 0x53: // LD D,E
		return cpu.LD_D_E()
	case 0x54: // LD D,H
		return cpu.LD_D_H()
	case 0x55: // LD D,L
		return cpu.LD_D_L()
	case 0x56: // LD D,(HL)
		return cpu.LD_D_HL()
	case 0x57: // LD D,A
		return cpu.LD_D_A()
	case 0x58: // LD E,B
		return cpu.LD_E_B()
	case 0x59: // LD E,C
		return cpu.LD_E_C()
	case 0x5A: // LD E,D
		return cpu.LD_E_D()
	case 0x5B: // LD E,E
		return cpu.LD_E_E()
	case 0x5C: // LD E,H
		return cpu.LD_E_H()
	case 0x5D: // LD E,L
		return cpu.LD_E_L()
	case 0x5E: // LD E,(HL)
		return cpu.LD_E_HL()
	case 0x5F: // LD E,A
		return cpu.LD_E_A()
	case 0x60: // LD H,B
		return cpu.LD_H_B()
	case 0x61: // LD H,C
		return cpu.LD_H_C()
	case 0x62: // LD H,D
		return cpu.LD_H_D()
	case 0x63: // LD H,E
		return cpu.LD_H_E()
	case 0x64: // LD H,H
		return cpu.LD_H_H()
	case 0x65: // LD H,L
		return cpu.LD_H_L()
	case 0x66: // LD H,(HL)
		return cpu.LD_H_HL()
	case 0x67: // LD H,A
		return cpu.LD_H_A()
	case 0x68: // LD L,B
		return cpu.LD_L_B()
	case 0x69: // LD L,C
		return cpu.LD_L_C()
	case 0x6A: // LD L,D
		return cpu.LD_L_D()
	case 0x6B: // LD L,E
		return cpu.LD_L_E()
	case 0x6C: // LD L,H
		return cpu.LD_L_H()
	case 0x6D: // LD L,L
		return cpu.LD_L_L()
	case 0x6E: // LD L,(HL)
		return cpu.LD_L_HL()
	case 0x6F: // LD L,A
		return cpu.LD_L_A()
	case 0x70: // LD (HL),B
		return cpu.LD_HL_B()
	case 0x71: // LD (HL),C
		return cpu.LD_HL_C()
	case 0x72: // LD (HL),D
		return cpu.LD_HL_D()
	case 0x73: // LD (HL),E
		return cpu.LD_HL_E()
	case 0x74: // LD (HL),H
		return cpu.LD_HL_H()
	case 0x75: // LD (HL),L
		return cpu.LD_HL_L()
	case 0x76: // HALT
		return cpu.HALT()
	case 0x77: // LD (HL),A
		return cpu.LD_HL_A()
	case 0x78: // LD A,B
		return cpu.LD_A_B()
	case 0x79: // LD A,C
		return cpu.LD_A_C()
	case 0x7A: // LD A,D
		return cpu.LD_A_D()
	case 0x7B: // LD A,E
		return cpu.LD_A_E()
	case 0x7C: // LD A,H
		return cpu.LD_A_H()
	case 0x7D: // LD A,L
		return cpu.LD_A_L()
	case 0x7E: // LD A,(HL)
		return cpu.LD_A_HL()
	case 0x7F: // LD A,A
		return cpu.LD_A_A()
	case 0x80: // ADD A,B
		return cpu.ADD_A_B()
	case 0x81: // ADD A,C
		return cpu.ADD_A_C()
	case 0x82: // ADD A,D
		return cpu.ADD_A_D()
	case 0x83: // ADD A,E
		return cpu.ADD_A_E()
	case 0x84: // ADD A,H
		return cpu.ADD_A_H()
	case 0x85: // ADD A,L
		return cpu.ADD_A_L()
	case 0x86: // ADD A,(HL)
		return cpu.ADD_A_HL()
	case 0x87: // ADD A,A
		return cpu.ADD_A_A()
	case 0x88: // ADC A,B
		return cpu.ADC_A_B()
	case 0x89: // ADC A,C
		return cpu.ADC_A_C()
	case 0x8A: // ADC A,D
		return cpu.ADC_A_D()
	case 0x8B: // ADC A,E
		return cpu.ADC_A_E()
	case 0x8C: // ADC A,H
		return cpu.ADC_A_H()
	case 0x8D: // ADC A,L
		return cpu.ADC_A_L()
	case 0x8E: // ADC A,(HL)
		return cpu.ADC_A_HL()
	case 0x8F: // ADC A,A
		return cpu.ADC_A_A()
	case 0x90: // SUB B
		return cpu.SUB_B()
	case 0x91: // SUB C
		return cpu.SUB_C()
	case 0x92: // SUB D
		return cpu.SUB_D()
	case 0x93: // SUB E
		return cpu.SUB_E()
	case 0x94: // SUB H
		return cpu.SUB_H()
	case 0x95: // SUB L
		return cpu.SUB_L()
	case 0x96: // SUB (HL)
		return cpu.SUB_HL()
	case 0x97: // SUB A
		return cpu.SUB_A()
	case 0x98: // SBC A,B
		return cpu.SBC_A_B()
	case 0x99: // SBC A,C
		return cpu.SBC_A_C()
	case 0x9A: // SBC A,D
		return cpu.SBC_A_D()
	case 0x9B: // SBC A,E
		return cpu.SBC_A_E()
	case 0x9C: // SBC A,H
		return cpu.SBC_A_H()
	case 0x9D: // SBC A,L
		return cpu.SBC_A_L()
	case 0x9E: // SBC A,(HL)
		return cpu.SBC_A_HL()
	case 0x9F: // SBC A,A
		return cpu.SBC_A_A()
	case 0xA0: // AND B
		return cpu.AND_B()
	case 0xA1: // AND C
		return cpu.AND_C()
	case 0xA2: // AND D
		return cpu.AND_D()
	case 0xA3: // AND E
		return cpu.AND_E()
	case 0xA4: // AND H
		return cpu.AND_H()
	case 0xA5: // AND L
		return cpu.AND_L()
	case 0xA6: // AND (HL)
		return cpu.AND_HL()
	case 0xA7: // AND A
		return cpu.AND_A()
	case 0xA8: // XOR B
		return cpu.XOR_B()
	case 0xA9: // XOR C
		return cpu.XOR_C()
	case 0xAA: // XOR D
		return cpu.XOR_D()
	case 0xAB: // XOR E
		return cpu.XOR_E()
	case 0xAC: // XOR H
		return cpu.XOR_H()
	case 0xAD: // XOR L
		return cpu.XOR_L()
	case 0xAE: // XOR (HL)
		return cpu.XOR_HL()
	case 0xAF: // XOR A
		return cpu.XOR_A()
	case 0xB0: // OR B
		return cpu.OR_B()
	case 0xB1: // OR C
		return cpu.OR_C()
	case 0xB2: // OR D
		return cpu.OR_D()
	case 0xB3: // OR E
		return cpu.OR_E()
	case 0xB4: // OR H
		return cpu.OR_H()
	case 0xB5: // OR L
		return cpu.OR_L()
	case 0xB6: // OR (HL)
		return cpu.OR_HL()
	case 0xB7: // OR A
		return cpu.OR_A()
	case 0xB8: // CP B
		return cpu.CP_B()
	case 0xB9: // CP C
		return cpu.CP_C()
	case 0xBA: // CP D
		return cpu.CP_D()
	case 0xBB: // CP E
		return cpu.CP_E()
	case 0xBC: // CP H
		return cpu.CP_H()
	case 0xBD: // CP L
		return cpu.CP_L()
	case 0xBE: // CP (HL)
		return cpu.CP_HL()
	case 0xBF: // CP A
		return cpu.CP_A()
	case 0xC0: // RET NZ
		return cpu.RET_NZ()
	case 0xC1: // POP BC
		return cpu.POP_BC()
	case 0xC2: // JP NZ,a16
		return cpu.JP_NZ_a16()
	case 0xC3: // JP a16
		return cpu.JP_a16()
	case 0xC4: // CALL NZ,a16
		return cpu.CALL_NZ_a16()
	case 0xC5: // PUSH BC
		return cpu.PUSH_BC()
	case 0xC6: // ADD A,d8
		return cpu.ADD_A_d8()
	case 0xC7: // RST 00H
		return cpu.RST_00H()
	case 0xC8: // RET Z
		return cpu.RET_Z()
	case 0xC9: // RET
		return cpu.RET()
	case 0xCA: // JP Z,a16
		return cpu.JP_Z_a16()
	case 0xCB: // CB prefix - handled above
		return 4 // Should never reach here
	case 0xCC: // CALL Z,a16
		return cpu.CALL_Z_a16()
	case 0xCD: // CALL a16
		return cpu.CALL_a16()
	case 0xCE: // ADC A,d8
		return cpu.ADC_A_d8()
	case 0xCF: // RST 08H
		return cpu.RST_08H()
	case 0xD0: // RET NC
		return cpu.RET_NC()
	case 0xD1: // POP DE
		return cpu.POP_DE()
	case 0xD2: // JP NC,a16
		return cpu.JP_NC_a16()
	case 0xD3: // Invalid opcode
		return 4
	case 0xD4: // CALL NC,a16
		return cpu.CALL_NC_a16()
	case 0xD5: // PUSH DE
		return cpu.PUSH_DE()
	case 0xD6: // SUB d8
		return cpu.SUB_d8()
	case 0xD7: // RST 10H
		return cpu.RST_10H()
	case 0xD8: // RET C
		return cpu.RET_C()
	case 0xD9: // RETI
		return cpu.RETI()
	case 0xDA: // JP C,a16
		return cpu.JP_C_a16()
	case 0xDB: // Invalid opcode
		return 4
	case 0xDC: // CALL C,a16
		return cpu.CALL_C_a16()
	case 0xDD: // Invalid opcode
		return 4
	case 0xDE: // SBC A,d8
		return cpu.SBC_A_d8()
	case 0xDF: // RST 18H
		return cpu.RST_18H()
	case 0xE0: // LDH (a8),A
		return cpu.LDH_a8_A()
	case 0xE1: // POP HL
		return cpu.POP_HL()
	case 0xE2: // LD (C),A
		return cpu.LD_C_mem_A()
	case 0xE3: // Invalid opcode
		return 4
	case 0xE4: // Invalid opcode
		return 4
	case 0xE5: // PUSH HL
		return cpu.PUSH_HL()
	case 0xE6: // AND d8
		return cpu.AND_d8()
	case 0xE7: // RST 20H
		return cpu.RST_20H()
	case 0xE8: // ADD SP,r8
		return cpu.ADD_SP_r8()
	case 0xE9: // JP (HL)
		return cpu.JP_HL()
	case 0xEA: // LD (a16),A
		return cpu.LD_a16_A()
	case 0xEB: // Invalid opcode
		return 4
	case 0xEC: // Invalid opcode
		return 4
	case 0xED: // Invalid opcode
		return 4
	case 0xEE: // XOR d8
		return cpu.XOR_d8()
	case 0xEF: // RST 28H
		return cpu.RST_28H()
	case 0xF0: // LDH A,(a8)
		return cpu.LDH_A_a8()
	case 0xF1: // POP AF
		return cpu.POP_AF()
	case 0xF2: // LD A,(C)
		return cpu.LD_A_C_mem()
	case 0xF3: // DI
		return cpu.DI()
	case 0xF4: // Invalid opcode
		return 4
	case 0xF5: // PUSH AF
		return cpu.PUSH_AF()
	case 0xF6: // OR d8
		return cpu.OR_d8()
	case 0xF7: // RST 30H
		return cpu.RST_30H()
	case 0xF8: // LD HL,SP+r8
		return cpu.LD_HL_SP_r8()
	case 0xF9: // LD SP,HL
		return cpu.LD_SP_HL()
	case 0xFA: // LD A,(a16)
		return cpu.LD_A_a16()
	case 0xFB: // EI
		return cpu.EI()
	case 0xFC: // Invalid opcode
		return 4
	case 0xFD: // Invalid opcode
		return 4
	case 0xFE: // CP d8
		return cpu.CP_d8()
	case 0xFF: // RST 38H
		return cpu.RST_38H()
	default:
		log.Printf("[CPU] Unknown opcode: 0x%02X at PC: 0x%04X", opcode, cpu.reg.PC-1)
		return 4
	}
}
