package main

import (
	"fmt"

	"github.com/briancain/gameboy-go/internal/ppu"
)

// Simple mock MMU for testing
type TestMMU struct {
	memory [0x10000]byte
}

func (m *TestMMU) ReadByte(addr uint16) byte {
	return m.memory[addr]
}

func (m *TestMMU) WriteByte(addr uint16, value byte) {
	m.memory[addr] = value
}

func main() {
	fmt.Println("PPU Test Program")
	fmt.Println("================")

	// Create a test MMU
	mmu := &TestMMU{}

	// Set up some basic PPU registers
	mmu.WriteByte(0xFF40, 0x91) // LCDC - LCD enabled, BG enabled
	mmu.WriteByte(0xFF47, 0xE4) // BGP - Background palette

	// Create some test tile data in VRAM
	// Let's create a simple checkerboard pattern
	tileAddr := uint16(0x8000)
	for i := 0; i < 16; i += 2 {
		if i < 8 {
			// First half of tile - alternating pattern
			mmu.WriteByte(tileAddr+uint16(i), 0xAA)   // 10101010
			mmu.WriteByte(tileAddr+uint16(i+1), 0x55) // 01010101
		} else {
			// Second half of tile - inverse pattern
			mmu.WriteByte(tileAddr+uint16(i), 0x55)   // 01010101
			mmu.WriteByte(tileAddr+uint16(i+1), 0xAA) // 10101010
		}
	}

	// Set up tile map - fill with tile 0
	for i := uint16(0x9800); i < 0x9C00; i++ {
		mmu.WriteByte(i, 0x00) // Use tile 0
	}

	// Create PPU
	ppuInstance := ppu.NewPPU(mmu)

	fmt.Printf("PPU initialized. LCD enabled: %v\n", ppuInstance.IsLCDEnabled())
	fmt.Printf("Screen dimensions: %dx%d\n", ppuInstance.GetScreenWidth(), ppuInstance.GetScreenHeight())

	// Simulate a few scanlines
	fmt.Println("\nSimulating PPU for a few lines...")
	for line := 0; line < 10; line++ {
		// Simulate one scanline worth of cycles
		// OAM (80) + VRAM (172) + HBLANK (204) = 456 cycles per line
		ppuInstance.Step(456)

		fmt.Printf("Line %d: Mode=%d, Clock=%d\n",
			ppuInstance.GetCurrentLine(),
			ppuInstance.GetCurrentMode(),
			ppuInstance.GetModeClock())
	}

	// Display a small section of the screen buffer
	fmt.Println("\nScreen buffer (top-left 20x10):")
	ppuInstance.DisplayScreenBufferSection(0, 0, 20, 10)

	// Show LCDC status
	fmt.Println("\nLCDC Status:")
	status := ppuInstance.GetLCDCStatus()
	for key, value := range status {
		fmt.Printf("  %s: %v\n", key, value)
	}

	fmt.Println("\nPPU test completed!")
}
