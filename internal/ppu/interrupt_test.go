package ppu

import (
	"testing"
)

// Test V-Blank interrupt generation
func TestVBlankInterruptGeneration(t *testing.T) {
	mockMMU := &MockMMU{
		registers: make(map[uint16]byte),
	}
	ppu := NewPPU(mockMMU)

	// Clear interrupt flag initially
	mockMMU.WriteByte(0xFF0F, 0x00)

	// Request V-Blank interrupt
	ppu.requestVBlankInterrupt()

	// Check that the V-Blank interrupt flag is set
	interruptFlag := mockMMU.ReadByte(0xFF0F)
	if (interruptFlag & 0x01) == 0 {
		t.Error("V-Blank interrupt flag should be set")
	}
}

// Test STAT interrupt generation
func TestSTATInterruptGeneration(t *testing.T) {
	mockMMU := &MockMMU{
		registers: make(map[uint16]byte),
	}
	ppu := NewPPU(mockMMU)

	// Clear interrupt flag initially
	mockMMU.WriteByte(0xFF0F, 0x00)

	// Request STAT interrupt
	ppu.requestSTATInterrupt()

	// Check that the STAT interrupt flag is set
	interruptFlag := mockMMU.ReadByte(0xFF0F)
	if (interruptFlag & 0x02) == 0 {
		t.Error("STAT interrupt flag should be set")
	}
}

// Test LYC coincidence interrupt
func TestLYCCoincidenceInterrupt(t *testing.T) {
	mockMMU := &MockMMU{
		registers: make(map[uint16]byte),
	}
	ppu := NewPPU(mockMMU)

	// Set up STAT register to enable LYC interrupt
	mockMMU.WriteByte(0xFF41, STAT_LYC_INT)

	// Set LYC to match current line
	mockMMU.WriteByte(0xFF45, 0) // LYC = 0
	ppu.line = 0                 // LY = 0

	// Clear interrupt flag initially
	mockMMU.WriteByte(0xFF0F, 0x00)

	// Check LYC coincidence
	ppu.checkLYC()

	// Verify coincidence flag is set
	stat := mockMMU.ReadByte(0xFF41)
	if (stat & STAT_LYC_EQUAL) == 0 {
		t.Error("LYC coincidence flag should be set")
	}

	// Verify STAT interrupt was requested
	interruptFlag := mockMMU.ReadByte(0xFF0F)
	if (interruptFlag & 0x02) == 0 {
		t.Error("STAT interrupt should be requested on LYC coincidence")
	}
}

// Test that LYC interrupt is only triggered on transition
func TestLYCInterruptOnlyOnTransition(t *testing.T) {
	mockMMU := &MockMMU{
		registers: make(map[uint16]byte),
	}
	ppu := NewPPU(mockMMU)

	// Set up STAT register to enable LYC interrupt and set coincidence flag
	mockMMU.WriteByte(0xFF41, STAT_LYC_INT|STAT_LYC_EQUAL)

	// Set LYC to match current line
	mockMMU.WriteByte(0xFF45, 0) // LYC = 0
	ppu.line = 0                 // LY = 0

	// Clear interrupt flag initially
	mockMMU.WriteByte(0xFF0F, 0x00)

	// Check LYC coincidence (should not trigger interrupt since already coincident)
	ppu.checkLYC()

	// Verify no interrupt was requested (since we were already coincident)
	interruptFlag := mockMMU.ReadByte(0xFF0F)
	if (interruptFlag & 0x02) != 0 {
		t.Error("STAT interrupt should not be requested when already coincident")
	}
}
