package gbcore

import (
	"testing"
)

// TestPPUInitialization verifies that a new PPU can be created
func TestPPUInitialization(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new PPU
	ppu := NewPPU(mockMMU)

	// Check that the PPU was initialized
	if ppu == nil {
		t.Error("Expected PPU to be initialized, got nil")
	}

	// Check initial state
	if ppu.mode != MODE_OAM {
		t.Errorf("Expected initial mode to be MODE_OAM, got %d", ppu.mode)
	}

	if ppu.modeClock != 0 {
		t.Errorf("Expected initial modeClock to be 0, got %d", ppu.modeClock)
	}

	if ppu.line != 0 {
		t.Errorf("Expected initial line to be 0, got %d", ppu.line)
	}

	// Check that screen buffer was initialized
	buffer := ppu.GetScreenBuffer()
	if len(buffer) != SCREEN_WIDTH*SCREEN_HEIGHT {
		t.Errorf("Expected screen buffer size to be %d, got %d",
			SCREEN_WIDTH*SCREEN_HEIGHT, len(buffer))
	}

	// Check that screen buffer was cleared
	for i, pixel := range buffer {
		if pixel != 0 {
			t.Errorf("Expected screen buffer at %d to be 0, got %d", i, pixel)
		}
	}
}

// TestPPUReset tests the PPU reset function
func TestPPUReset(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{}

	// Create a new PPU
	ppu := NewPPU(mockMMU)

	// Modify state
	ppu.mode = MODE_VBLANK
	ppu.modeClock = 100
	ppu.line = 50
	ppu.screenBuffer[0] = 3

	// Reset the PPU
	ppu.Reset()

	// Check that state was reset
	if ppu.mode != MODE_OAM {
		t.Errorf("Expected mode to be reset to MODE_OAM, got %d", ppu.mode)
	}

	if ppu.modeClock != 0 {
		t.Errorf("Expected modeClock to be reset to 0, got %d", ppu.modeClock)
	}

	if ppu.line != 0 {
		t.Errorf("Expected line to be reset to 0, got %d", ppu.line)
	}

	// Check that screen buffer was cleared
	if ppu.screenBuffer[0] != 0 {
		t.Errorf("Expected screen buffer to be cleared, got %d", ppu.screenBuffer[0])
	}
}

// TestPPUStep tests the PPU step function
func TestPPUStep(t *testing.T) {
	// Create a mock MMU
	mockMMU := &MockMMU{
		registers: map[uint16]byte{
			0xFF40: 0x91, // LCDC enabled
		},
	}

	// Create a new PPU
	ppu := NewPPU(mockMMU)

	// Step through OAM mode
	ppu.Step(80)

	// Check that mode changed to VRAM
	if ppu.mode != MODE_VRAM {
		t.Errorf("Expected mode to change to MODE_VRAM after 80 cycles, got %d", ppu.mode)
	}

	// Step through VRAM mode
	ppu.Step(172)

	// Check that mode changed to HBLANK
	if ppu.mode != MODE_HBLANK {
		t.Errorf("Expected mode to change to MODE_HBLANK after 172 cycles, got %d", ppu.mode)
	}

	// Step through HBLANK mode
	ppu.Step(204)

	// Check that mode changed back to OAM and line incremented
	if ppu.mode != MODE_OAM {
		t.Errorf("Expected mode to change to MODE_OAM after 204 cycles, got %d", ppu.mode)
	}

	if ppu.line != 1 {
		t.Errorf("Expected line to increment to 1, got %d", ppu.line)
	}

	// Step through to line 144 (end of visible screen)
	for i := 1; i < 144; i++ {
		// For each line: OAM -> VRAM -> HBLANK
		ppu.Step(80)  // OAM
		ppu.Step(172) // VRAM
		ppu.Step(204) // HBLANK
	}

	// Check that mode changed to VBLANK
	if ppu.mode != MODE_VBLANK {
		t.Errorf("Expected mode to change to MODE_VBLANK at line 144, got %d", ppu.mode)
	}

	// Check that VBlank interrupt was requested
	if mockMMU.registers[0xFF0F]&0x01 == 0 {
		t.Error("Expected VBlank interrupt to be requested")
	}
}

// TestPPUDisabled tests that the PPU does nothing when disabled
func TestPPUDisabled(t *testing.T) {
	// Create a mock MMU with LCD disabled
	mockMMU := &MockMMU{
		registers: map[uint16]byte{
			0xFF40: 0x11, // LCDC disabled (bit 7 = 0)
		},
	}

	// Create a new PPU
	ppu := NewPPU(mockMMU)

	// Set initial state
	ppu.mode = MODE_OAM
	ppu.modeClock = 0
	ppu.line = 0

	// Step the PPU
	ppu.Step(100)

	// Check that state didn't change
	if ppu.mode != MODE_OAM {
		t.Errorf("Expected mode to remain MODE_OAM when LCD is disabled, got %d", ppu.mode)
	}

	if ppu.modeClock != 0 {
		t.Errorf("Expected modeClock to remain 0 when LCD is disabled, got %d", ppu.modeClock)
	}

	if ppu.line != 0 {
		t.Errorf("Expected line to remain 0 when LCD is disabled, got %d", ppu.line)
	}
}

// MockMMU is a mock implementation of the MMU interface for testing
type MockMMU struct {
	registers map[uint16]byte
	memory    [0x10000]byte
}

func (m *MockMMU) ReadByte(addr uint16) byte {
	// Handle special registers
	if addr >= 0xFF00 && addr <= 0xFF7F {
		if val, ok := m.registers[addr]; ok {
			return val
		}
		return 0
	}
	return m.memory[addr]
}

func (m *MockMMU) WriteByte(addr uint16, value byte) {
	// Handle special registers
	if addr >= 0xFF00 && addr <= 0xFF7F {
		if m.registers == nil {
			m.registers = make(map[uint16]byte)
		}
		m.registers[addr] = value
		return
	}
	m.memory[addr] = value
}
