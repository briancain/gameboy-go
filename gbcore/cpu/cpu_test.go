package gbcore

import (
	"testing"
)

// TestCPUInitialization verifies that a new CPU can be created
func TestCPUInitialization(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, err := NewCPU(mockMMU)

	// Check that there was no error
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that the CPU was initialized
	if cpu == nil {
		t.Error("Expected CPU to be initialized, got nil")
	}

	// Check initial register values
	if cpu.reg.A != 0x01 {
		t.Errorf("Expected register A to be 0x01, got %02X", cpu.reg.A)
	}

	if cpu.reg.F != 0xB0 {
		t.Errorf("Expected register F to be 0xB0, got %02X", cpu.reg.F)
	}

	if cpu.reg.PC != 0x0100 {
		t.Errorf("Expected register PC to be 0x0100, got %04X", cpu.reg.PC)
	}

	if cpu.reg.SP != 0xFFFE {
		t.Errorf("Expected register SP to be 0xFFFE, got %04X", cpu.reg.SP)
	}
}

// TestCPURegisterPairs tests the 16-bit register pair operations
func TestCPURegisterPairs(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test AF register pair
	cpu.reg.A = 0x12
	cpu.reg.F = 0x34
	if cpu.reg.GetAF() != 0x1234 {
		t.Errorf("Expected AF to be 0x1234, got %04X", cpu.reg.GetAF())
	}

	// Test BC register pair
	cpu.reg.B = 0x56
	cpu.reg.C = 0x78
	if cpu.reg.GetBC() != 0x5678 {
		t.Errorf("Expected BC to be 0x5678, got %04X", cpu.reg.GetBC())
	}

	// Test DE register pair
	cpu.reg.D = 0x9A
	cpu.reg.E = 0xBC
	if cpu.reg.GetDE() != 0x9ABC {
		t.Errorf("Expected DE to be 0x9ABC, got %04X", cpu.reg.GetDE())
	}

	// Test HL register pair
	cpu.reg.H = 0xDE
	cpu.reg.L = 0xF0
	if cpu.reg.GetHL() != 0xDEF0 {
		t.Errorf("Expected HL to be 0xDEF0, got %04X", cpu.reg.GetHL())
	}

	// Test setting register pairs
	cpu.reg.SetAF(0xABCD)
	if cpu.reg.A != 0xAB || cpu.reg.F != 0xCD {
		t.Errorf("Expected A=0xAB, F=0xCD, got A=%02X, F=%02X", cpu.reg.A, cpu.reg.F)
	}

	cpu.reg.SetBC(0x1234)
	if cpu.reg.B != 0x12 || cpu.reg.C != 0x34 {
		t.Errorf("Expected B=0x12, C=0x34, got B=%02X, C=%02X", cpu.reg.B, cpu.reg.C)
	}

	cpu.reg.SetDE(0x5678)
	if cpu.reg.D != 0x56 || cpu.reg.E != 0x78 {
		t.Errorf("Expected D=0x56, E=0x78, got D=%02X, E=%02X", cpu.reg.D, cpu.reg.E)
	}

	cpu.reg.SetHL(0x9ABC)
	if cpu.reg.H != 0x9A || cpu.reg.L != 0xBC {
		t.Errorf("Expected H=0x9A, L=0xBC, got H=%02X, L=%02X", cpu.reg.H, cpu.reg.L)
	}
}

// TestCPUFlags tests the flag operations
func TestCPUFlags(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Clear all flags
	cpu.reg.F = 0x00

	// Test setting flags
	cpu.reg.SetFlag(FLAG_Z)
	if !cpu.reg.GetFlag(FLAG_Z) {
		t.Error("Expected Z flag to be set")
	}

	cpu.reg.SetFlag(FLAG_N)
	if !cpu.reg.GetFlag(FLAG_N) {
		t.Error("Expected N flag to be set")
	}

	// Test clearing flags
	cpu.reg.ClearFlag(FLAG_Z)
	if cpu.reg.GetFlag(FLAG_Z) {
		t.Error("Expected Z flag to be cleared")
	}

	// Test multiple flags
	cpu.reg.F = 0x00
	cpu.reg.SetFlag(FLAG_Z | FLAG_C)
	if !cpu.reg.GetFlag(FLAG_Z) || !cpu.reg.GetFlag(FLAG_C) {
		t.Error("Expected Z and C flags to be set")
	}
	if cpu.reg.GetFlag(FLAG_N) || cpu.reg.GetFlag(FLAG_H) {
		t.Error("Expected N and H flags to be cleared")
	}
}

// TestCPUInstructions tests basic CPU instructions
func TestCPUInstructions(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Test NOP instruction
	cycles := cpu.NOP()
	if cycles != 4 {
		t.Errorf("Expected NOP to take 4 cycles, got %d", cycles)
	}

	// Test HALT instruction
	cpu.halted = false
	cycles = cpu.HALT()
	if cycles != 4 {
		t.Errorf("Expected HALT to take 4 cycles, got %d", cycles)
	}
	if !cpu.halted {
		t.Error("Expected CPU to be halted after HALT instruction")
	}

	// Test STOP instruction
	cpu.stopped = false
	cycles = cpu.STOP()
	if cycles != 4 {
		t.Errorf("Expected STOP to take 4 cycles, got %d", cycles)
	}
	if !cpu.stopped {
		t.Error("Expected CPU to be stopped after STOP instruction")
	}
}

// TestCPUStep tests the CPU step function
func TestCPUStep(t *testing.T) {
	// Create a mock MMU that returns NOP (0x00) for all reads
	mockMMU := &MockMMU{}

	// Create a new CPU
	cpu, _ := NewCPU(mockMMU)

	// Reset clock
	cpu.ResetClock()

	// Step the CPU (should execute a NOP)
	cycles := cpu.Step()

	// Check that the correct number of cycles were used
	if cycles != 4 {
		t.Errorf("Expected Step to take 4 cycles for NOP, got %d", cycles)
	}

	// Check that PC was incremented
	if cpu.reg.PC != 0x0101 {
		t.Errorf("Expected PC to be 0x0101 after one step, got %04X", cpu.reg.PC)
	}

	// Check that clock was updated
	if cpu.clock.t != 4 {
		t.Errorf("Expected clock.t to be 4, got %d", cpu.clock.t)
	}

	if cpu.clock.m != 1 {
		t.Errorf("Expected clock.m to be 1, got %d", cpu.clock.m)
	}
}

// MockMMU is a mock implementation of the MMU interface for testing
type MockMMU struct {
	memory [0x10000]byte // Full 64KB address space
}

func (m *MockMMU) ReadByte(addr uint16) byte {
	return m.memory[addr]
}

func (m *MockMMU) WriteByte(addr uint16, value byte) {
	m.memory[addr] = value
}

func (m *MockMMU) ReadWord(addr uint16) uint16 {
	return uint16(m.memory[addr]) | (uint16(m.memory[addr+1]) << 8)
}

func (m *MockMMU) WriteWord(addr uint16, value uint16) {
	m.memory[addr] = byte(value & 0xFF)
	m.memory[addr+1] = byte(value >> 8)
}
