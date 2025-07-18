package gbcore

import (
	"testing"
)

// TestIncrementInstructions tests the 8-bit increment instructions (INC r)
func TestIncrementInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test INC B
	cpu.reg.B = 0x0F
	cpu.reg.F = 0x00
	cycles := cpu.INC_B()
	if cpu.reg.B != 0x10 {
		t.Errorf("INC B: Expected B to be 0x10, got %02X", cpu.reg.B)
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("INC B: Expected H flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("INC B: Expected Z flag to be clear")
	}
	if cycles != 4 {
		t.Errorf("INC B: Expected 4 cycles, got %d", cycles)
	}

	// Test INC B with overflow to zero
	cpu.reg.B = 0xFF
	cpu.reg.F = 0x00
	cpu.INC_B()
	if cpu.reg.B != 0x00 {
		t.Errorf("INC B (overflow): Expected B to be 0x00, got %02X", cpu.reg.B)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("INC B (overflow): Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("INC B (overflow): Expected H flag to be set")
	}

	// Test INC C
	cpu.reg.C = 0x42
	cpu.reg.F = 0x00
	cycles = cpu.INC_C()
	if cpu.reg.C != 0x43 {
		t.Errorf("INC C: Expected C to be 0x43, got %02X", cpu.reg.C)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("INC C: Expected Z and H flags to be clear")
	}
	if cycles != 4 {
		t.Errorf("INC C: Expected 4 cycles, got %d", cycles)
	}

	// Test INC (HL)
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0x2F)
	cycles = cpu.INC_HL_()
	if mockMMU.ReadByte(0xC000) != 0x30 {
		t.Errorf("INC (HL): Expected (HL) to be 0x30, got %02X", mockMMU.ReadByte(0xC000))
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("INC (HL): Expected H flag to be set")
	}
	if cycles != 12 {
		t.Errorf("INC (HL): Expected 12 cycles, got %d", cycles)
	}
}

// TestDecrementInstructions tests the 8-bit decrement instructions (DEC r)
func TestDecrementInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test DEC A
	cpu.reg.A = 0x01
	cpu.reg.F = 0x00
	cycles := cpu.DEC_A()
	if cpu.reg.A != 0x00 {
		t.Errorf("DEC A: Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("DEC A: Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("DEC A: Expected N flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_H) {
		t.Error("DEC A: Expected H flag to be clear")
	}
	if cycles != 4 {
		t.Errorf("DEC A: Expected 4 cycles, got %d", cycles)
	}

	// Test DEC B with half-carry
	cpu.reg.B = 0x10
	cpu.reg.F = 0x00
	cpu.DEC_B()
	if cpu.reg.B != 0x0F {
		t.Errorf("DEC B (half-carry): Expected B to be 0x0F, got %02X", cpu.reg.B)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("DEC B (half-carry): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("DEC B (half-carry): Expected N flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("DEC B (half-carry): Expected H flag to be set")
	}

	// Test DEC (HL)
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0x01)
	cycles = cpu.DEC_HL_()
	if mockMMU.ReadByte(0xC000) != 0x00 {
		t.Errorf("DEC (HL): Expected (HL) to be 0x00, got %02X", mockMMU.ReadByte(0xC000))
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("DEC (HL): Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("DEC (HL): Expected N flag to be set")
	}
	if cycles != 12 {
		t.Errorf("DEC (HL): Expected 12 cycles, got %d", cycles)
	}
}

// TestAddInstructions tests the 8-bit add instructions (ADD A,r)
func TestAddInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test ADD A,B
	cpu.reg.A = 0x3A
	cpu.reg.B = 0xC6
	cpu.reg.F = 0x00
	cycles := cpu.ADD_A_B()
	if cpu.reg.A != 0x00 {
		t.Errorf("ADD A,B: Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("ADD A,B: Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("ADD A,B: Expected H flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("ADD A,B: Expected C flag to be set")
	}
	if cycles != 4 {
		t.Errorf("ADD A,B: Expected 4 cycles, got %d", cycles)
	}

	// Test ADD A,(HL)
	cpu.reg.A = 0x3C
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0x12)
	cycles = cpu.ADD_A_HL()
	if cpu.reg.A != 0x4E {
		t.Errorf("ADD A,(HL): Expected A to be 0x4E, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("ADD A,(HL): Expected all flags to be clear")
	}
	if cycles != 8 {
		t.Errorf("ADD A,(HL): Expected 8 cycles, got %d", cycles)
	}

	// Test ADD A,d8
	cpu.reg.A = 0xFF
	cpu.reg.PC = 0xC000
	mockMMU.WriteByte(0xC000, 0x01)
	cycles = cpu.ADD_A_d8()
	if cpu.reg.A != 0x00 {
		t.Errorf("ADD A,d8: Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("ADD A,d8: Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("ADD A,d8: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("ADD A,d8: Expected 8 cycles, got %d", cycles)
	}
}

// TestAdcInstructions tests the 8-bit add with carry instructions (ADC A,r)
func TestAdcInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test ADC A,B without carry
	cpu.reg.A = 0x3A
	cpu.reg.B = 0x05
	cpu.reg.F = 0x00 // Clear carry flag
	cycles := cpu.ADC_A_B()
	if cpu.reg.A != 0x3F {
		t.Errorf("ADC A,B (no carry): Expected A to be 0x3F, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("ADC A,B (no carry): Expected all flags to be clear")
	}
	if cycles != 4 {
		t.Errorf("ADC A,B (no carry): Expected 4 cycles, got %d", cycles)
	}

	// Test ADC A,B with carry
	cpu.reg.A = 0x3A
	cpu.reg.B = 0x05
	cpu.reg.F = FLAG_C // Set carry flag
	cpu.ADC_A_B()
	if cpu.reg.A != 0x40 {
		t.Errorf("ADC A,B (with carry): Expected A to be 0x40, got %02X", cpu.reg.A)
	}
	// The half-carry flag is set because (0x3A & 0x0F) + (0x05 & 0x0F) + 1 > 0x0F
	// (0x0A + 0x05 + 0x01 = 0x10)
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || !cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("ADC A,B (with carry): Expected only H flag to be set")
	}

	// Test ADC A,A with carry (should set half-carry and carry flags)
	cpu.reg.A = 0x80
	cpu.reg.F = FLAG_C // Set carry flag
	cpu.ADC_A_A()
	if cpu.reg.A != 0x01 {
		t.Errorf("ADC A,A (with carry): Expected A to be 0x01, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("ADC A,A (with carry): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("ADC A,A (with carry): Expected C flag to be set")
	}
}

// TestSubInstructions tests the 8-bit subtraction instructions (SUB r)
func TestSubInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test SUB B
	cpu.reg.A = 0x3E
	cpu.reg.B = 0x3E
	cpu.reg.F = 0x00
	cycles := cpu.SUB_B()
	if cpu.reg.A != 0x00 {
		t.Errorf("SUB B: Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("SUB B: Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("SUB B: Expected N flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("SUB B: Expected H and C flags to be clear")
	}
	if cycles != 4 {
		t.Errorf("SUB B: Expected 4 cycles, got %d", cycles)
	}

	// Test SUB B with half-carry and carry
	cpu.reg.A = 0x10
	cpu.reg.B = 0x21
	cpu.reg.F = 0x00
	cpu.SUB_B()
	if cpu.reg.A != 0xEF {
		t.Errorf("SUB B (with carry): Expected A to be 0xEF, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("SUB B (with carry): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("SUB B (with carry): Expected N flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("SUB B (with carry): Expected H flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("SUB B (with carry): Expected C flag to be set")
	}

	// Test SUB (HL)
	cpu.reg.A = 0x42
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0x12)
	cycles = cpu.SUB_HL()
	if cpu.reg.A != 0x30 {
		t.Errorf("SUB (HL): Expected A to be 0x30, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("SUB (HL): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("SUB (HL): Expected N flag to be set")
	}
	if cycles != 8 {
		t.Errorf("SUB (HL): Expected 8 cycles, got %d", cycles)
	}
}

// TestSbcInstructions tests the 8-bit subtract with carry instructions (SBC A,r)
func TestSbcInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test SBC A,B without carry
	cpu.reg.A = 0x3E
	cpu.reg.B = 0x0F
	cpu.reg.F = 0x00 // Clear carry flag
	cycles := cpu.SBC_A_B()
	if cpu.reg.A != 0x2F {
		t.Errorf("SBC A,B (no carry): Expected A to be 0x2F, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("SBC A,B (no carry): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("SBC A,B (no carry): Expected N flag to be set")
	}
	if cycles != 4 {
		t.Errorf("SBC A,B (no carry): Expected 4 cycles, got %d", cycles)
	}

	// Test SBC A,B with carry
	cpu.reg.A = 0x3E
	cpu.reg.B = 0x0F
	cpu.reg.F = FLAG_C // Set carry flag
	cpu.SBC_A_B()
	if cpu.reg.A != 0x2E {
		t.Errorf("SBC A,B (with carry): Expected A to be 0x2E, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("SBC A,B (with carry): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("SBC A,B (with carry): Expected N flag to be set")
	}

	// Test SBC A,A with carry
	cpu.reg.A = 0x01
	cpu.reg.F = FLAG_C // Set carry flag
	cpu.SBC_A_A()
	if cpu.reg.A != 0xFF {
		t.Errorf("SBC A,A (with carry): Expected A to be 0xFF, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("SBC A,A (with carry): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("SBC A,A (with carry): Expected N flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("SBC A,A (with carry): Expected H flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("SBC A,A (with carry): Expected C flag to be set")
	}
}

// TestCompareInstructions tests the compare instructions (CP r)
func TestCompareInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test CP B (equal)
	cpu.reg.A = 0x42
	cpu.reg.B = 0x42
	cpu.reg.F = 0x00
	cycles := cpu.CP_B()
	if cpu.reg.A != 0x42 {
		t.Errorf("CP B (equal): Expected A to remain 0x42, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("CP B (equal): Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("CP B (equal): Expected N flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("CP B (equal): Expected H and C flags to be clear")
	}
	if cycles != 4 {
		t.Errorf("CP B (equal): Expected 4 cycles, got %d", cycles)
	}

	// Test CP B (A > B)
	cpu.reg.A = 0x42
	cpu.reg.B = 0x21
	cpu.reg.F = 0x00
	cpu.CP_B()
	if cpu.reg.A != 0x42 {
		t.Errorf("CP B (A > B): Expected A to remain 0x42, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("CP B (A > B): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("CP B (A > B): Expected N flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_C) {
		t.Error("CP B (A > B): Expected C flag to be clear")
	}

	// Test CP B (A < B)
	cpu.reg.A = 0x21
	cpu.reg.B = 0x42
	cpu.reg.F = 0x00
	cpu.CP_B()
	if cpu.reg.A != 0x21 {
		t.Errorf("CP B (A < B): Expected A to remain 0x21, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("CP B (A < B): Expected Z flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("CP B (A < B): Expected N flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("CP B (A < B): Expected C flag to be set")
	}

	// Test CP (HL)
	cpu.reg.A = 0x42
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0x42)
	cycles = cpu.CP_HL()
	if cpu.reg.A != 0x42 {
		t.Errorf("CP (HL): Expected A to remain 0x42, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("CP (HL): Expected Z flag to be set")
	}
	if cycles != 8 {
		t.Errorf("CP (HL): Expected 8 cycles, got %d", cycles)
	}
}

// TestLogicalInstructions tests the logical instructions (AND, OR, XOR)
func TestLogicalInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test AND B
	cpu.reg.A = 0xAA // 10101010
	cpu.reg.B = 0x55 // 01010101
	cpu.reg.F = 0x00
	cycles := cpu.AND_B()
	if cpu.reg.A != 0x00 {
		t.Errorf("AND B: Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("AND B: Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("AND B: Expected H flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("AND B: Expected N and C flags to be clear")
	}
	if cycles != 4 {
		t.Errorf("AND B: Expected 4 cycles, got %d", cycles)
	}

	// Test AND B (non-zero result)
	cpu.reg.A = 0xAA // 10101010
	cpu.reg.B = 0xAA // 10101010
	cpu.reg.F = 0x00
	cpu.AND_B()
	if cpu.reg.A != 0xAA {
		t.Errorf("AND B (non-zero): Expected A to be 0xAA, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("AND B (non-zero): Expected Z flag to be clear")
	}

	// Test OR C
	cpu.reg.A = 0xAA // 10101010
	cpu.reg.C = 0x55 // 01010101
	cpu.reg.F = 0x00
	cycles = cpu.OR_C()
	if cpu.reg.A != 0xFF {
		t.Errorf("OR C: Expected A to be 0xFF, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("OR C: Expected all flags to be clear")
	}
	if cycles != 4 {
		t.Errorf("OR C: Expected 4 cycles, got %d", cycles)
	}

	// Test OR C (zero result)
	cpu.reg.A = 0x00
	cpu.reg.C = 0x00
	cpu.reg.F = 0x00
	cpu.OR_C()
	if cpu.reg.A != 0x00 {
		t.Errorf("OR C (zero): Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("OR C (zero): Expected Z flag to be set")
	}

	// Test XOR D
	cpu.reg.A = 0xAA // 10101010
	cpu.reg.D = 0x55 // 01010101
	cpu.reg.F = 0x00
	cycles = cpu.XOR_D()
	if cpu.reg.A != 0xFF {
		t.Errorf("XOR D: Expected A to be 0xFF, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("XOR D: Expected all flags to be clear")
	}
	if cycles != 4 {
		t.Errorf("XOR D: Expected 4 cycles, got %d", cycles)
	}

	// Test XOR A (always results in zero)
	cpu.reg.A = 0xAA
	cpu.reg.F = 0x00
	cpu.XOR_A()
	if cpu.reg.A != 0x00 {
		t.Errorf("XOR A: Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("XOR A: Expected Z flag to be set")
	}
}

// TestBitOperationInstructions tests the bit operation instructions (BIT, SET, RES)
func TestBitOperationInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test BIT 0,B (bit is 0)
	cpu.reg.B = 0xFE // 11111110 (bit 0 is 0)
	cpu.reg.F = 0x00
	cycles := cpu.BIT_0_B()
	if cpu.reg.B != 0xFE {
		t.Errorf("BIT 0,B: Expected B to remain 0xFE, got %02X", cpu.reg.B)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("BIT 0,B (bit is 0): Expected Z flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_N) {
		t.Error("BIT 0,B (bit is 0): Expected N flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("BIT 0,B (bit is 0): Expected H flag to be set")
	}
	if cycles != 8 {
		t.Errorf("BIT 0,B: Expected 8 cycles, got %d", cycles)
	}

	// Test BIT 0,B (bit is 1)
	cpu.reg.B = 0xFF // 11111111 (bit 0 is 1)
	cpu.reg.F = 0x00
	cpu.BIT_0_B()
	if cpu.reg.B != 0xFF {
		t.Errorf("BIT 0,B: Expected B to remain 0xFF, got %02X", cpu.reg.B)
	}
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("BIT 0,B (bit is 1): Expected Z flag to be clear")
	}

	// Test SET 0,C
	cpu.reg.C = 0xFE // 11111110 (bit 0 is 0)
	cycles = cpu.SET_0_C()
	if cpu.reg.C != 0xFF {
		t.Errorf("SET 0,C: Expected C to be 0xFF, got %02X", cpu.reg.C)
	}
	if cycles != 8 {
		t.Errorf("SET 0,C: Expected 8 cycles, got %d", cycles)
	}

	// Test RES 7,D
	cpu.reg.D = 0xFF // 11111111 (bit 7 is 1)
	cycles = cpu.RES_7_D()
	if cpu.reg.D != 0x7F {
		t.Errorf("RES 7,D: Expected D to be 0x7F, got %02X", cpu.reg.D)
	}
	if cycles != 8 {
		t.Errorf("RES 7,D: Expected 8 cycles, got %d", cycles)
	}

	// Test BIT 4,(HL)
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0xEF) // 11101111 (bit 4 is 0)
	cycles = cpu.BIT_4_HL()
	if mockMMU.ReadByte(0xC000) != 0xEF {
		t.Errorf("BIT 4,(HL): Expected (HL) to remain 0xEF, got %02X", mockMMU.ReadByte(0xC000))
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("BIT 4,(HL): Expected Z flag to be set")
	}
	if cycles != 12 {
		t.Errorf("BIT 4,(HL): Expected 12 cycles, got %d", cycles)
	}
}

// Test16BitArithmeticInstructions tests the 16-bit arithmetic instructions
func Test16BitArithmeticInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test INC BC
	cpu.reg.SetBC(0x1234)
	cycles := cpu.INC_BC()
	if cpu.reg.GetBC() != 0x1235 {
		t.Errorf("INC BC: Expected BC to be 0x1235, got %04X", cpu.reg.GetBC())
	}
	if cycles != 8 {
		t.Errorf("INC BC: Expected 8 cycles, got %d", cycles)
	}

	// Test DEC DE
	cpu.reg.SetDE(0x1234)
	cycles = cpu.DEC_DE()
	if cpu.reg.GetDE() != 0x1233 {
		t.Errorf("DEC DE: Expected DE to be 0x1233, got %04X", cpu.reg.GetDE())
	}
	if cycles != 8 {
		t.Errorf("DEC DE: Expected 8 cycles, got %d", cycles)
	}

	// Test ADD HL,BC
	cpu.reg.SetHL(0x1234)
	cpu.reg.SetBC(0x0DEF)
	cpu.reg.F = 0x00
	cycles = cpu.ADD_HL_BC()
	if cpu.reg.GetHL() != 0x2023 {
		t.Errorf("ADD HL,BC: Expected HL to be 0x2023, got %04X", cpu.reg.GetHL())
	}
	if cpu.reg.GetFlag(FLAG_N) {
		t.Error("ADD HL,BC: Expected N flag to be clear")
	}
	if cycles != 8 {
		t.Errorf("ADD HL,BC: Expected 8 cycles, got %d", cycles)
	}

	// Test ADD HL,BC with half-carry and carry
	cpu.reg.SetHL(0xF234)
	cpu.reg.SetBC(0x1DEF)
	cpu.reg.F = 0x00
	cpu.ADD_HL_BC()
	if cpu.reg.GetHL() != 0x1023 {
		t.Errorf("ADD HL,BC (with carry): Expected HL to be 0x1023, got %04X", cpu.reg.GetHL())
	}
	if cpu.reg.GetFlag(FLAG_N) {
		t.Error("ADD HL,BC (with carry): Expected N flag to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("ADD HL,BC (with carry): Expected H flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("ADD HL,BC (with carry): Expected C flag to be set")
	}
}

// TestRotationInstructions tests the rotation instructions (RLC, RRC, RL, RR)
func TestRotationInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test RLCA (rotate A left)
	cpu.reg.A = 0x85 // 10000101
	cpu.reg.F = 0x00
	cycles := cpu.RLCA()
	if cpu.reg.A != 0x0B { // 00001011
		t.Errorf("RLCA: Expected A to be 0x0B, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RLCA: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RLCA: Expected C flag to be set")
	}
	if cycles != 4 {
		t.Errorf("RLCA: Expected 4 cycles, got %d", cycles)
	}

	// Test RRCA (rotate A right)
	cpu.reg.A = 0x01 // 00000001
	cpu.reg.F = 0x00
	cycles = cpu.RRCA()
	if cpu.reg.A != 0x80 { // 10000000
		t.Errorf("RRCA: Expected A to be 0x80, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RRCA: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RRCA: Expected C flag to be set")
	}
	if cycles != 4 {
		t.Errorf("RRCA: Expected 4 cycles, got %d", cycles)
	}

	// Test RLA (rotate A left through carry)
	cpu.reg.A = 0x95 // 10010101
	cpu.reg.F = 0x00 // Carry flag is clear
	cycles = cpu.RLA()
	if cpu.reg.A != 0x2A { // 00101010
		t.Errorf("RLA: Expected A to be 0x2A, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RLA: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RLA: Expected C flag to be set")
	}
	if cycles != 4 {
		t.Errorf("RLA: Expected 4 cycles, got %d", cycles)
	}

	// Test RRA (rotate A right through carry)
	cpu.reg.A = 0x81   // 10000001
	cpu.reg.F = FLAG_C // Carry flag is set
	cycles = cpu.RRA()
	if cpu.reg.A != 0xC0 { // 11000000
		t.Errorf("RRA: Expected A to be 0xC0, got %02X", cpu.reg.A)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RRA: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RRA: Expected C flag to be set")
	}
	if cycles != 4 {
		t.Errorf("RRA: Expected 4 cycles, got %d", cycles)
	}
}

// TestCBRotationInstructions tests the CB-prefixed rotation instructions
func TestCBRotationInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test RLC B (rotate B left)
	cpu.reg.B = 0x85 // 10000101
	cpu.reg.F = 0x00
	cycles := cpu.RLC_B()
	if cpu.reg.B != 0x0B { // 00001011
		t.Errorf("RLC B: Expected B to be 0x0B, got %02X", cpu.reg.B)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RLC B: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RLC B: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("RLC B: Expected 8 cycles, got %d", cycles)
	}

	// Test RRC C (rotate C right)
	cpu.reg.C = 0x01 // 00000001
	cpu.reg.F = 0x00
	cycles = cpu.RRC_C()
	if cpu.reg.C != 0x80 { // 10000000
		t.Errorf("RRC C: Expected C to be 0x80, got %02X", cpu.reg.C)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RRC C: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RRC C: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("RRC C: Expected 8 cycles, got %d", cycles)
	}

	// Test RL D (rotate D left through carry)
	cpu.reg.D = 0x80   // 10000000
	cpu.reg.F = FLAG_C // Carry flag is set
	cycles = cpu.RL_D()
	if cpu.reg.D != 0x01 { // 00000001
		t.Errorf("RL D: Expected D to be 0x01, got %02X", cpu.reg.D)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RL D: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RL D: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("RL D: Expected 8 cycles, got %d", cycles)
	}

	// Test RR E (rotate E right through carry)
	cpu.reg.E = 0x01 // 00000001
	cpu.reg.F = 0x00 // Carry flag is clear
	cycles = cpu.RR_E()
	if cpu.reg.E != 0x00 { // 00000000
		t.Errorf("RR E: Expected E to be 0x00, got %02X", cpu.reg.E)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("RR E: Expected Z flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("RR E: Expected N and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("RR E: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("RR E: Expected 8 cycles, got %d", cycles)
	}
}

// TestShiftInstructions tests the CB-prefixed shift instructions
func TestShiftInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test SLA B (shift B left arithmetic)
	cpu.reg.B = 0x85 // 10000101
	cpu.reg.F = 0x00
	cycles := cpu.SLA_B()
	if cpu.reg.B != 0x0A { // 00001010
		t.Errorf("SLA B: Expected B to be 0x0A, got %02X", cpu.reg.B)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("SLA B: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("SLA B: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("SLA B: Expected 8 cycles, got %d", cycles)
	}

	// Test SRA C (shift C right arithmetic)
	cpu.reg.C = 0x85 // 10000101
	cpu.reg.F = 0x00
	cycles = cpu.SRA_C()
	if cpu.reg.C != 0xC2 { // 11000010
		t.Errorf("SRA C: Expected C to be 0xC2, got %02X", cpu.reg.C)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("SRA C: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("SRA C: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("SRA C: Expected 8 cycles, got %d", cycles)
	}

	// Test SRL D (shift D right logical)
	cpu.reg.D = 0x85 // 10000101
	cpu.reg.F = 0x00
	cycles = cpu.SRL_D()
	if cpu.reg.D != 0x42 { // 01000010
		t.Errorf("SRL D: Expected D to be 0x42, got %02X", cpu.reg.D)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("SRL D: Expected Z, N, and H flags to be clear")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("SRL D: Expected C flag to be set")
	}
	if cycles != 8 {
		t.Errorf("SRL D: Expected 8 cycles, got %d", cycles)
	}

	// Test SWAP E (swap nibbles in E)
	cpu.reg.E = 0x1F // 00011111
	cpu.reg.F = 0x00
	cycles = cpu.SWAP_E()
	if cpu.reg.E != 0xF1 { // 11110001
		t.Errorf("SWAP E: Expected E to be 0xF1, got %02X", cpu.reg.E)
	}
	if cpu.reg.GetFlag(FLAG_Z) || cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) || cpu.reg.GetFlag(FLAG_C) {
		t.Error("SWAP E: Expected all flags to be clear")
	}
	if cycles != 8 {
		t.Errorf("SWAP E: Expected 8 cycles, got %d", cycles)
	}
}

// TestLoadInstructions tests the load instructions
func TestLoadInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test LD B,d8 (load immediate value into B)
	cpu.reg.PC = 0xC000
	mockMMU.WriteByte(0xC000, 0x42)
	cycles := cpu.LD_B_d8()
	if cpu.reg.B != 0x42 {
		t.Errorf("LD B,d8: Expected B to be 0x42, got %02X", cpu.reg.B)
	}
	if cpu.reg.PC != 0xC001 {
		t.Errorf("LD B,d8: Expected PC to be 0xC001, got %04X", cpu.reg.PC)
	}
	if cycles != 8 {
		t.Errorf("LD B,d8: Expected 8 cycles, got %d", cycles)
	}

	// Test LD C,B (load B into C)
	cpu.reg.B = 0x42
	cpu.reg.C = 0x00
	cycles = cpu.LD_C_B()
	if cpu.reg.C != 0x42 {
		t.Errorf("LD C,B: Expected C to be 0x42, got %02X", cpu.reg.C)
	}
	if cycles != 4 {
		t.Errorf("LD C,B: Expected 4 cycles, got %d", cycles)
	}

	// Test LD (HL),A (load A into memory at HL)
	cpu.reg.A = 0x42
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0x00)
	cycles = cpu.LD_HL_A()
	if mockMMU.ReadByte(0xC000) != 0x42 {
		t.Errorf("LD (HL),A: Expected (HL) to be 0x42, got %02X", mockMMU.ReadByte(0xC000))
	}
	if cycles != 8 {
		t.Errorf("LD (HL),A: Expected 8 cycles, got %d", cycles)
	}

	// Test LD A,(HL) (load memory at HL into A)
	cpu.reg.A = 0x00
	cpu.reg.SetHL(0xC000)
	mockMMU.WriteByte(0xC000, 0x42)
	cycles = cpu.LD_A_HL()
	if cpu.reg.A != 0x42 {
		t.Errorf("LD A,(HL): Expected A to be 0x42, got %02X", cpu.reg.A)
	}
	if cycles != 8 {
		t.Errorf("LD A,(HL): Expected 8 cycles, got %d", cycles)
	}

	// Test LD (BC),A (load A into memory at BC)
	cpu.reg.A = 0x42
	cpu.reg.SetBC(0xC000)
	mockMMU.WriteByte(0xC000, 0x00)
	cycles = cpu.LD_BC_A()
	if mockMMU.ReadByte(0xC000) != 0x42 {
		t.Errorf("LD (BC),A: Expected (BC) to be 0x42, got %02X", mockMMU.ReadByte(0xC000))
	}
	if cycles != 8 {
		t.Errorf("LD (BC),A: Expected 8 cycles, got %d", cycles)
	}

	// Test LD A,(BC) (load memory at BC into A)
	cpu.reg.A = 0x00
	cpu.reg.SetBC(0xC000)
	mockMMU.WriteByte(0xC000, 0x42)
	cycles = cpu.LD_A_BC()
	if cpu.reg.A != 0x42 {
		t.Errorf("LD A,(BC): Expected A to be 0x42, got %02X", cpu.reg.A)
	}
	if cycles != 8 {
		t.Errorf("LD A,(BC): Expected 8 cycles, got %d", cycles)
	}
}

// TestControlFlowInstructions tests the control flow instructions
func TestControlFlowInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test JP a16 (jump to immediate address)
	cpu.reg.PC = 0xC000
	mockMMU.WriteByte(0xC000, 0x34)
	mockMMU.WriteByte(0xC001, 0x12)
	cycles := cpu.JP_a16()
	if cpu.reg.PC != 0x1234 {
		t.Errorf("JP a16: Expected PC to be 0x1234, got %04X", cpu.reg.PC)
	}
	if cycles != 16 {
		t.Errorf("JP a16: Expected 16 cycles, got %d", cycles)
	}

	// Test JP HL (jump to address in HL)
	cpu.reg.SetHL(0x5678)
	cycles = cpu.JP_HL()
	if cpu.reg.PC != 0x5678 {
		t.Errorf("JP HL: Expected PC to be 0x5678, got %04X", cpu.reg.PC)
	}
	if cycles != 4 {
		t.Errorf("JP HL: Expected 4 cycles, got %d", cycles)
	}

	// Test JP Z,a16 (conditional jump if Z flag is set)
	cpu.reg.PC = 0xC000
	mockMMU.WriteByte(0xC000, 0x34)
	mockMMU.WriteByte(0xC001, 0x12)
	cpu.reg.F = FLAG_Z // Set Z flag
	cycles = cpu.JP_Z_a16()
	if cpu.reg.PC != 0x1234 {
		t.Errorf("JP Z,a16 (Z set): Expected PC to be 0x1234, got %04X", cpu.reg.PC)
	}
	if cycles != 16 {
		t.Errorf("JP Z,a16 (Z set): Expected 16 cycles, got %d", cycles)
	}

	// Test JP Z,a16 (conditional jump if Z flag is clear)
	cpu.reg.PC = 0xC000
	mockMMU.WriteByte(0xC000, 0x34)
	mockMMU.WriteByte(0xC001, 0x12)
	cpu.reg.F = 0x00 // Clear Z flag
	cycles = cpu.JP_Z_a16()
	if cpu.reg.PC != 0xC002 {
		t.Errorf("JP Z,a16 (Z clear): Expected PC to be 0xC002, got %04X", cpu.reg.PC)
	}
	if cycles != 12 {
		t.Errorf("JP Z,a16 (Z clear): Expected 12 cycles, got %d", cycles)
	}

	// Test JR r8 (relative jump)
	cpu.reg.PC = 0xC000
	mockMMU.WriteByte(0xC000, 0x05) // Jump forward 5 bytes
	cycles = cpu.JR_r8()
	if cpu.reg.PC != 0xC006 {
		t.Errorf("JR r8 (forward): Expected PC to be 0xC006, got %04X", cpu.reg.PC)
	}
	if cycles != 12 {
		t.Errorf("JR r8 (forward): Expected 12 cycles, got %d", cycles)
	}

	// Test JR r8 (relative jump backward)
	cpu.reg.PC = 0xC010
	mockMMU.WriteByte(0xC010, 0xFA) // Jump backward 6 bytes (0xFA = -6 in two's complement)
	cycles = cpu.JR_r8()
	if cpu.reg.PC != 0xC00B {
		t.Errorf("JR r8 (backward): Expected PC to be 0xC00B, got %04X", cpu.reg.PC)
	}
	if cycles != 12 {
		t.Errorf("JR r8 (backward): Expected 12 cycles, got %d", cycles)
	}
}

// TestStackInstructions tests the stack operation instructions
func TestStackInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test PUSH BC
	cpu.reg.SetBC(0x1234)
	cpu.reg.SP = 0xC000
	cycles := cpu.PUSH_BC()
	if cpu.reg.SP != 0xBFFE {
		t.Errorf("PUSH BC: Expected SP to be 0xBFFE, got %04X", cpu.reg.SP)
	}
	if mockMMU.ReadWord(0xBFFE) != 0x1234 {
		t.Errorf("PUSH BC: Expected word at 0xBFFE to be 0x1234, got %04X", mockMMU.ReadWord(0xBFFE))
	}
	if cycles != 16 {
		t.Errorf("PUSH BC: Expected 16 cycles, got %d", cycles)
	}

	// Test POP DE
	cpu.reg.SP = 0xBFFE
	mockMMU.WriteWord(0xBFFE, 0x5678)
	cycles = cpu.POP_DE()
	if cpu.reg.SP != 0xC000 {
		t.Errorf("POP DE: Expected SP to be 0xC000, got %04X", cpu.reg.SP)
	}
	if cpu.reg.GetDE() != 0x5678 {
		t.Errorf("POP DE: Expected DE to be 0x5678, got %04X", cpu.reg.GetDE())
	}
	if cycles != 12 {
		t.Errorf("POP DE: Expected 12 cycles, got %d", cycles)
	}

	// Test CALL a16
	cpu.reg.PC = 0xC000
	cpu.reg.SP = 0xC100
	mockMMU.WriteByte(0xC000, 0x34)
	mockMMU.WriteByte(0xC001, 0x12)
	cycles = cpu.CALL_a16()
	if cpu.reg.PC != 0x1234 {
		t.Errorf("CALL a16: Expected PC to be 0x1234, got %04X", cpu.reg.PC)
	}
	if cpu.reg.SP != 0xC0FE {
		t.Errorf("CALL a16: Expected SP to be 0xC0FE, got %04X", cpu.reg.SP)
	}
	if mockMMU.ReadWord(0xC0FE) != 0xC002 {
		t.Errorf("CALL a16: Expected word at 0xC0FE to be 0xC002, got %04X", mockMMU.ReadWord(0xC0FE))
	}
	if cycles != 24 {
		t.Errorf("CALL a16: Expected 24 cycles, got %d", cycles)
	}

	// Test RET
	cpu.reg.SP = 0xC0FE
	mockMMU.WriteWord(0xC0FE, 0xABCD)
	cycles = cpu.RET()
	if cpu.reg.PC != 0xABCD {
		t.Errorf("RET: Expected PC to be 0xABCD, got %04X", cpu.reg.PC)
	}
	if cpu.reg.SP != 0xC100 {
		t.Errorf("RET: Expected SP to be 0xC100, got %04X", cpu.reg.SP)
	}
	if cycles != 16 {
		t.Errorf("RET: Expected 16 cycles, got %d", cycles)
	}

	// Test RST 00H
	cpu.reg.PC = 0xC000
	cpu.reg.SP = 0xC100
	cycles = cpu.RST_00H()
	if cpu.reg.PC != 0x0000 {
		t.Errorf("RST 00H: Expected PC to be 0x0000, got %04X", cpu.reg.PC)
	}
	if cpu.reg.SP != 0xC0FE {
		t.Errorf("RST 00H: Expected SP to be 0xC0FE, got %04X", cpu.reg.SP)
	}
	if mockMMU.ReadWord(0xC0FE) != 0xC000 {
		t.Errorf("RST 00H: Expected word at 0xC0FE to be 0xC000, got %04X", mockMMU.ReadWord(0xC0FE))
	}
	if cycles != 16 {
		t.Errorf("RST 00H: Expected 16 cycles, got %d", cycles)
	}
}

// TestSpecialInstructions tests special instructions like DAA, CPL, etc.
func TestSpecialInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test DAA (decimal adjust A) after addition
	cpu.reg.A = 0x9A // Invalid BCD result after addition
	cpu.reg.F = 0x00 // No flags set
	cycles := cpu.DAA()
	if cpu.reg.A != 0x00 {
		t.Errorf("DAA (after addition): Expected A to be 0x00, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("DAA (after addition): Expected Z flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("DAA (after addition): Expected C flag to be set")
	}
	if cycles != 4 {
		t.Errorf("DAA: Expected 4 cycles, got %d", cycles)
	}

	// Test DAA (decimal adjust A) after subtraction
	cpu.reg.A = 0x35
	cpu.reg.F = FLAG_N // Subtraction flag set
	cpu.DAA()
	if cpu.reg.A != 0x35 {
		t.Errorf("DAA (after subtraction): Expected A to be 0x35, got %02X", cpu.reg.A)
	}

	// Test CPL (complement A)
	cpu.reg.A = 0x35
	cpu.reg.F = 0x00
	cycles = cpu.CPL()
	if cpu.reg.A != 0xCA {
		t.Errorf("CPL: Expected A to be 0xCA, got %02X", cpu.reg.A)
	}
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("CPL: Expected N flag to be set")
	}
	if !cpu.reg.GetFlag(FLAG_H) {
		t.Error("CPL: Expected H flag to be set")
	}
	if cycles != 4 {
		t.Errorf("CPL: Expected 4 cycles, got %d", cycles)
	}

	// Test CCF (complement carry flag)
	cpu.reg.F = FLAG_C | FLAG_N | FLAG_H // Set all flags
	cycles = cpu.CCF()
	if cpu.reg.GetFlag(FLAG_C) {
		t.Error("CCF: Expected C flag to be clear")
	}
	if cpu.reg.GetFlag(FLAG_N) {
		t.Error("CCF: Expected N flag to be clear")
	}
	if cpu.reg.GetFlag(FLAG_H) {
		t.Error("CCF: Expected H flag to be clear")
	}
	if cycles != 4 {
		t.Errorf("CCF: Expected 4 cycles, got %d", cycles)
	}

	// Test SCF (set carry flag)
	cpu.reg.F = FLAG_N | FLAG_H // Set N and H flags
	cycles = cpu.SCF()
	if !cpu.reg.GetFlag(FLAG_C) {
		t.Error("SCF: Expected C flag to be set")
	}
	if cpu.reg.GetFlag(FLAG_N) {
		t.Error("SCF: Expected N flag to be clear")
	}
	if cpu.reg.GetFlag(FLAG_H) {
		t.Error("SCF: Expected H flag to be clear")
	}
	if cycles != 4 {
		t.Errorf("SCF: Expected 4 cycles, got %d", cycles)
	}

	// Test DI (disable interrupts)
	cpu.interruptMaster = true
	cycles = cpu.DI()
	if cpu.interruptMaster {
		t.Error("DI: Expected interrupts to be disabled")
	}
	if cycles != 4 {
		t.Errorf("DI: Expected 4 cycles, got %d", cycles)
	}

	// Test EI (enable interrupts)
	cpu.interruptMaster = false
	cycles = cpu.EI()
	if !cpu.interruptMaster {
		t.Error("EI: Expected interrupts to be enabled")
	}
	if cycles != 4 {
		t.Errorf("EI: Expected 4 cycles, got %d", cycles)
	}
}
