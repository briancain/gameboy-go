package gbcore

import "log"

// Execute CB-prefixed instruction
func (cpu *Z80) executeCBInstruction(opcode byte) int {
	switch opcode {
	case 0x00: // RLC B
		return cpu.RLC_B()
	case 0x01: // RLC C
		return cpu.RLC_C()
	case 0x02: // RLC D
		return cpu.RLC_D()
	case 0x03: // RLC E
		return cpu.RLC_E()
	case 0x04: // RLC H
		return cpu.RLC_H()
	case 0x05: // RLC L
		return cpu.RLC_L()
	case 0x06: // RLC (HL)
		return cpu.RLC_HL()
	case 0x07: // RLC A
		return cpu.RLC_A()
	case 0x08: // RRC B
		return cpu.RRC_B()
	case 0x09: // RRC C
		return cpu.RRC_C()
	case 0x0A: // RRC D
		return cpu.RRC_D()
	case 0x0B: // RRC E
		return cpu.RRC_E()
	case 0x0C: // RRC H
		return cpu.RRC_H()
	case 0x0D: // RRC L
		return cpu.RRC_L()
	case 0x0E: // RRC (HL)
		return cpu.RRC_HL()
	case 0x0F: // RRC A
		return cpu.RRC_A()
	case 0x10: // RL B
		return cpu.RL_B()
	case 0x11: // RL C
		return cpu.RL_C()
	case 0x12: // RL D
		return cpu.RL_D()
	case 0x13: // RL E
		return cpu.RL_E()
	case 0x14: // RL H
		return cpu.RL_H()
	case 0x15: // RL L
		return cpu.RL_L()
	case 0x16: // RL (HL)
		return cpu.RL_HL()
	case 0x17: // RL A
		return cpu.RL_A()
	case 0x18: // RR B
		return cpu.RR_B()
	case 0x19: // RR C
		return cpu.RR_C()
	case 0x1A: // RR D
		return cpu.RR_D()
	case 0x1B: // RR E
		return cpu.RR_E()
	case 0x1C: // RR H
		return cpu.RR_H()
	case 0x1D: // RR L
		return cpu.RR_L()
	case 0x1E: // RR (HL)
		return cpu.RR_HL()
	case 0x1F: // RR A
		return cpu.RR_A()
	case 0x20: // SLA B
		return cpu.SLA_B()
	case 0x21: // SLA C
		return cpu.SLA_C()
	case 0x22: // SLA D
		return cpu.SLA_D()
	case 0x23: // SLA E
		return cpu.SLA_E()
	case 0x24: // SLA H
		return cpu.SLA_H()
	case 0x25: // SLA L
		return cpu.SLA_L()
	case 0x26: // SLA (HL)
		return cpu.SLA_HL()
	case 0x27: // SLA A
		return cpu.SLA_A()
	case 0x28: // SRA B
		return cpu.SRA_B()
	case 0x29: // SRA C
		return cpu.SRA_C()
	case 0x2A: // SRA D
		return cpu.SRA_D()
	case 0x2B: // SRA E
		return cpu.SRA_E()
	case 0x2C: // SRA H
		return cpu.SRA_H()
	case 0x2D: // SRA L
		return cpu.SRA_L()
	case 0x2E: // SRA (HL)
		return cpu.SRA_HL()
	case 0x2F: // SRA A
		return cpu.SRA_A()
	case 0x30: // SWAP B
		return cpu.SWAP_B()
	case 0x31: // SWAP C
		return cpu.SWAP_C()
	case 0x32: // SWAP D
		return cpu.SWAP_D()
	case 0x33: // SWAP E
		return cpu.SWAP_E()
	case 0x34: // SWAP H
		return cpu.SWAP_H()
	case 0x35: // SWAP L
		return cpu.SWAP_L()
	case 0x36: // SWAP (HL)
		return cpu.SWAP_HL()
	case 0x37: // SWAP A
		return cpu.SWAP_A()
	case 0x38: // SRL B
		return cpu.SRL_B()
	case 0x39: // SRL C
		return cpu.SRL_C()
	case 0x3A: // SRL D
		return cpu.SRL_D()
	case 0x3B: // SRL E
		return cpu.SRL_E()
	case 0x3C: // SRL H
		return cpu.SRL_H()
	case 0x3D: // SRL L
		return cpu.SRL_L()
	case 0x3E: // SRL (HL)
		return cpu.SRL_HL()
	case 0x3F: // SRL A
		return cpu.SRL_A()
	case 0x40: // BIT 0,B
		return cpu.BIT_0_B()
	case 0x41: // BIT 0,C
		return cpu.BIT_0_C()
	case 0x42: // BIT 0,D
		return cpu.BIT_0_D()
	case 0x43: // BIT 0,E
		return cpu.BIT_0_E()
	case 0x44: // BIT 0,H
		return cpu.BIT_0_H()
	case 0x45: // BIT 0,L
		return cpu.BIT_0_L()
	case 0x46: // BIT 0,(HL)
		return cpu.BIT_0_HL()
	case 0x47: // BIT 0,A
		return cpu.BIT_0_A()
	case 0x48: // BIT 1,B
		return cpu.BIT_1_B()
	case 0x49: // BIT 1,C
		return cpu.BIT_1_C()
	case 0x4A: // BIT 1,D
		return cpu.BIT_1_D()
	case 0x4B: // BIT 1,E
		return cpu.BIT_1_E()
	case 0x4C: // BIT 1,H
		return cpu.BIT_1_H()
	case 0x4D: // BIT 1,L
		return cpu.BIT_1_L()
	case 0x4E: // BIT 1,(HL)
		return cpu.BIT_1_HL()
	case 0x4F: // BIT 1,A
		return cpu.BIT_1_A()
	case 0x50: // BIT 2,B
		return cpu.BIT_2_B()
	case 0x51: // BIT 2,C
		return cpu.BIT_2_C()
	case 0x52: // BIT 2,D
		return cpu.BIT_2_D()
	case 0x53: // BIT 2,E
		return cpu.BIT_2_E()
	case 0x54: // BIT 2,H
		return cpu.BIT_2_H()
	case 0x55: // BIT 2,L
		return cpu.BIT_2_L()
	case 0x56: // BIT 2,(HL)
		return cpu.BIT_2_HL()
	case 0x57: // BIT 2,A
		return cpu.BIT_2_A()
	case 0x58: // BIT 3,B
		return cpu.BIT_3_B()
	case 0x59: // BIT 3,C
		return cpu.BIT_3_C()
	case 0x5A: // BIT 3,D
		return cpu.BIT_3_D()
	case 0x5B: // BIT 3,E
		return cpu.BIT_3_E()
	case 0x5C: // BIT 3,H
		return cpu.BIT_3_H()
	case 0x5D: // BIT 3,L
		return cpu.BIT_3_L()
	case 0x5E: // BIT 3,(HL)
		return cpu.BIT_3_HL()
	case 0x5F: // BIT 3,A
		return cpu.BIT_3_A()
	case 0x60: // BIT 4,B
		return cpu.BIT_4_B()
	case 0x61: // BIT 4,C
		return cpu.BIT_4_C()
	case 0x62: // BIT 4,D
		return cpu.BIT_4_D()
	case 0x63: // BIT 4,E
		return cpu.BIT_4_E()
	case 0x64: // BIT 4,H
		return cpu.BIT_4_H()
	case 0x65: // BIT 4,L
		return cpu.BIT_4_L()
	case 0x66: // BIT 4,(HL)
		return cpu.BIT_4_HL()
	case 0x67: // BIT 4,A
		return cpu.BIT_4_A()
	case 0x68: // BIT 5,B
		return cpu.BIT_5_B()
	case 0x69: // BIT 5,C
		return cpu.BIT_5_C()
	case 0x6A: // BIT 5,D
		return cpu.BIT_5_D()
	case 0x6B: // BIT 5,E
		return cpu.BIT_5_E()
	case 0x6C: // BIT 5,H
		return cpu.BIT_5_H()
	case 0x6D: // BIT 5,L
		return cpu.BIT_5_L()
	case 0x6E: // BIT 5,(HL)
		return cpu.BIT_5_HL()
	case 0x6F: // BIT 5,A
		return cpu.BIT_5_A()
	case 0x70: // BIT 6,B
		return cpu.BIT_6_B()
	case 0x71: // BIT 6,C
		return cpu.BIT_6_C()
	case 0x72: // BIT 6,D
		return cpu.BIT_6_D()
	case 0x73: // BIT 6,E
		return cpu.BIT_6_E()
	case 0x74: // BIT 6,H
		return cpu.BIT_6_H()
	case 0x75: // BIT 6,L
		return cpu.BIT_6_L()
	case 0x76: // BIT 6,(HL)
		return cpu.BIT_6_HL()
	case 0x77: // BIT 6,A
		return cpu.BIT_6_A()
	case 0x78: // BIT 7,B
		return cpu.BIT_7_B()
	case 0x79: // BIT 7,C
		return cpu.BIT_7_C()
	case 0x7A: // BIT 7,D
		return cpu.BIT_7_D()
	case 0x7B: // BIT 7,E
		return cpu.BIT_7_E()
	case 0x7C: // BIT 7,H
		return cpu.BIT_7_H()
	case 0x7D: // BIT 7,L
		return cpu.BIT_7_L()
	case 0x7E: // BIT 7,(HL)
		return cpu.BIT_7_HL()
	case 0x7F: // BIT 7,A
		return cpu.BIT_7_A()
	case 0x80: // RES 0,B
		return cpu.RES_0_B()
	case 0x81: // RES 0,C
		return cpu.RES_0_C()
	case 0x82: // RES 0,D
		return cpu.RES_0_D()
	case 0x83: // RES 0,E
		return cpu.RES_0_E()
	case 0x84: // RES 0,H
		return cpu.RES_0_H()
	case 0x85: // RES 0,L
		return cpu.RES_0_L()
	case 0x86: // RES 0,(HL)
		return cpu.RES_0_HL()
	case 0x87: // RES 0,A
		return cpu.RES_0_A()
	case 0x88: // RES 1,B
		return cpu.RES_1_B()
	case 0x89: // RES 1,C
		return cpu.RES_1_C()
	case 0x8A: // RES 1,D
		return cpu.RES_1_D()
	case 0x8B: // RES 1,E
		return cpu.RES_1_E()
	case 0x8C: // RES 1,H
		return cpu.RES_1_H()
	case 0x8D: // RES 1,L
		return cpu.RES_1_L()
	case 0x8E: // RES 1,(HL)
		return cpu.RES_1_HL()
	case 0x8F: // RES 1,A
		return cpu.RES_1_A()
	case 0x90: // RES 2,B
		return cpu.RES_2_B()
	case 0x91: // RES 2,C
		return cpu.RES_2_C()
	case 0x92: // RES 2,D
		return cpu.RES_2_D()
	case 0x93: // RES 2,E
		return cpu.RES_2_E()
	case 0x94: // RES 2,H
		return cpu.RES_2_H()
	case 0x95: // RES 2,L
		return cpu.RES_2_L()
	case 0x96: // RES 2,(HL)
		return cpu.RES_2_HL()
	case 0x97: // RES 2,A
		return cpu.RES_2_A()
	case 0x98: // RES 3,B
		return cpu.RES_3_B()
	case 0x99: // RES 3,C
		return cpu.RES_3_C()
	case 0x9A: // RES 3,D
		return cpu.RES_3_D()
	case 0x9B: // RES 3,E
		return cpu.RES_3_E()
	case 0x9C: // RES 3,H
		return cpu.RES_3_H()
	case 0x9D: // RES 3,L
		return cpu.RES_3_L()
	case 0x9E: // RES 3,(HL)
		return cpu.RES_3_HL()
	case 0x9F: // RES 3,A
		return cpu.RES_3_A()
	case 0xA0: // RES 4,B
		return cpu.RES_4_B()
	case 0xA1: // RES 4,C
		return cpu.RES_4_C()
	case 0xA2: // RES 4,D
		return cpu.RES_4_D()
	case 0xA3: // RES 4,E
		return cpu.RES_4_E()
	case 0xA4: // RES 4,H
		return cpu.RES_4_H()
	case 0xA5: // RES 4,L
		return cpu.RES_4_L()
	case 0xA6: // RES 4,(HL)
		return cpu.RES_4_HL()
	case 0xA7: // RES 4,A
		return cpu.RES_4_A()
	case 0xA8: // RES 5,B
		return cpu.RES_5_B()
	case 0xA9: // RES 5,C
		return cpu.RES_5_C()
	case 0xAA: // RES 5,D
		return cpu.RES_5_D()
	case 0xAB: // RES 5,E
		return cpu.RES_5_E()
	case 0xAC: // RES 5,H
		return cpu.RES_5_H()
	case 0xAD: // RES 5,L
		return cpu.RES_5_L()
	case 0xAE: // RES 5,(HL)
		return cpu.RES_5_HL()
	case 0xAF: // RES 5,A
		return cpu.RES_5_A()
	case 0xB0: // RES 6,B
		return cpu.RES_6_B()
	case 0xB1: // RES 6,C
		return cpu.RES_6_C()
	case 0xB2: // RES 6,D
		return cpu.RES_6_D()
	case 0xB3: // RES 6,E
		return cpu.RES_6_E()
	case 0xB4: // RES 6,H
		return cpu.RES_6_H()
	case 0xB5: // RES 6,L
		return cpu.RES_6_L()
	case 0xB6: // RES 6,(HL)
		return cpu.RES_6_HL()
	case 0xB7: // RES 6,A
		return cpu.RES_6_A()
	case 0xB8: // RES 7,B
		return cpu.RES_7_B()
	case 0xB9: // RES 7,C
		return cpu.RES_7_C()
	case 0xBA: // RES 7,D
		return cpu.RES_7_D()
	case 0xBB: // RES 7,E
		return cpu.RES_7_E()
	case 0xBC: // RES 7,H
		return cpu.RES_7_H()
	case 0xBD: // RES 7,L
		return cpu.RES_7_L()
	case 0xBE: // RES 7,(HL)
		return cpu.RES_7_HL()
	case 0xBF: // RES 7,A
		return cpu.RES_7_A()
	case 0xC0: // SET 0,B
		return cpu.SET_0_B()
	case 0xC1: // SET 0,C
		return cpu.SET_0_C()
	case 0xC2: // SET 0,D
		return cpu.SET_0_D()
	case 0xC3: // SET 0,E
		return cpu.SET_0_E()
	case 0xC4: // SET 0,H
		return cpu.SET_0_H()
	case 0xC5: // SET 0,L
		return cpu.SET_0_L()
	case 0xC6: // SET 0,(HL)
		return cpu.SET_0_HL()
	case 0xC7: // SET 0,A
		return cpu.SET_0_A()
	case 0xC8: // SET 1,B
		return cpu.SET_1_B()
	case 0xC9: // SET 1,C
		return cpu.SET_1_C()
	case 0xCA: // SET 1,D
		return cpu.SET_1_D()
	case 0xCB: // SET 1,E
		return cpu.SET_1_E()
	case 0xCC: // SET 1,H
		return cpu.SET_1_H()
	case 0xCD: // SET 1,L
		return cpu.SET_1_L()
	case 0xCE: // SET 1,(HL)
		return cpu.SET_1_HL()
	case 0xCF: // SET 1,A
		return cpu.SET_1_A()
	case 0xD0: // SET 2,B
		return cpu.SET_2_B()
	case 0xD1: // SET 2,C
		return cpu.SET_2_C()
	case 0xD2: // SET 2,D
		return cpu.SET_2_D()
	case 0xD3: // SET 2,E
		return cpu.SET_2_E()
	case 0xD4: // SET 2,H
		return cpu.SET_2_H()
	case 0xD5: // SET 2,L
		return cpu.SET_2_L()
	case 0xD6: // SET 2,(HL)
		return cpu.SET_2_HL()
	case 0xD7: // SET 2,A
		return cpu.SET_2_A()
	case 0xD8: // SET 3,B
		return cpu.SET_3_B()
	case 0xD9: // SET 3,C
		return cpu.SET_3_C()
	case 0xDA: // SET 3,D
		return cpu.SET_3_D()
	case 0xDB: // SET 3,E
		return cpu.SET_3_E()
	case 0xDC: // SET 3,H
		return cpu.SET_3_H()
	case 0xDD: // SET 3,L
		return cpu.SET_3_L()
	case 0xDE: // SET 3,(HL)
		return cpu.SET_3_HL()
	case 0xDF: // SET 3,A
		return cpu.SET_3_A()
	case 0xE0: // SET 4,B
		return cpu.SET_4_B()
	case 0xE1: // SET 4,C
		return cpu.SET_4_C()
	case 0xE2: // SET 4,D
		return cpu.SET_4_D()
	case 0xE3: // SET 4,E
		return cpu.SET_4_E()
	case 0xE4: // SET 4,H
		return cpu.SET_4_H()
	case 0xE5: // SET 4,L
		return cpu.SET_4_L()
	case 0xE6: // SET 4,(HL)
		return cpu.SET_4_HL()
	case 0xE7: // SET 4,A
		return cpu.SET_4_A()
	case 0xE8: // SET 5,B
		return cpu.SET_5_B()
	case 0xE9: // SET 5,C
		return cpu.SET_5_C()
	case 0xEA: // SET 5,D
		return cpu.SET_5_D()
	case 0xEB: // SET 5,E
		return cpu.SET_5_E()
	case 0xEC: // SET 5,H
		return cpu.SET_5_H()
	case 0xED: // SET 5,L
		return cpu.SET_5_L()
	case 0xEE: // SET 5,(HL)
		return cpu.SET_5_HL()
	case 0xEF: // SET 5,A
		return cpu.SET_5_A()
	case 0xF0: // SET 6,B
		return cpu.SET_6_B()
	case 0xF1: // SET 6,C
		return cpu.SET_6_C()
	case 0xF2: // SET 6,D
		return cpu.SET_6_D()
	case 0xF3: // SET 6,E
		return cpu.SET_6_E()
	case 0xF4: // SET 6,H
		return cpu.SET_6_H()
	case 0xF5: // SET 6,L
		return cpu.SET_6_L()
	case 0xF6: // SET 6,(HL)
		return cpu.SET_6_HL()
	case 0xF7: // SET 6,A
		return cpu.SET_6_A()
	case 0xF8: // SET 7,B
		return cpu.SET_7_B()
	case 0xF9: // SET 7,C
		return cpu.SET_7_C()
	case 0xFA: // SET 7,D
		return cpu.SET_7_D()
	case 0xFB: // SET 7,E
		return cpu.SET_7_E()
	case 0xFC: // SET 7,H
		return cpu.SET_7_H()
	case 0xFD: // SET 7,L
		return cpu.SET_7_L()
	case 0xFE: // SET 7,(HL)
		return cpu.SET_7_HL()
	case 0xFF: // SET 7,A
		return cpu.SET_7_A()
	default:
		log.Printf("[CPU] Unknown CB-prefixed opcode: 0x%02X at PC: 0x%04X", opcode, cpu.reg.PC-2)
		return 8
	}
}
