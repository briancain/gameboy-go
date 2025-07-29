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
	fmt.Println("Window Rendering Test")
	fmt.Println("====================")

	// Create a test MMU
	mmu := &TestMMU{}

	// Set up PPU registers
	mmu.WriteByte(0xFF40, 0xA1) // LCDC - LCD on, BG on, Window on
	mmu.WriteByte(0xFF47, 0xE4) // BGP - Background palette
	mmu.WriteByte(0xFF4A, 0)    // WY = 0 (window starts at line 0)
	mmu.WriteByte(0xFF4B, 3)    // WX = 3 (should be treated as 0)

	// Set up tile data
	mmu.WriteByte(0x8000, 0xFF) // Tile 0, line 0, low byte (all pixels set)
	mmu.WriteByte(0x8001, 0x00) // Tile 0, line 0, high byte (color 1)
	mmu.WriteByte(0x9800, 0x00) // Tile map entry 0 -> use tile 0

	// Create PPU
	ppuInstance := ppu.NewPPU(mmu)

	fmt.Printf("LCDC: 0x%02X\n", mmu.ReadByte(0xFF40))
	fmt.Printf("WY: %d, WX: %d\n", mmu.ReadByte(0xFF4A), mmu.ReadByte(0xFF4B))
	fmt.Printf("BGP: 0x%02X\n", mmu.ReadByte(0xFF47))

	// Clear screen buffer
	buffer := ppuInstance.GetScreenBuffer()
	for i := range buffer {
		buffer[i] = 0
	}

	fmt.Println("\nTesting window rendering...")

	// Reset PPU and manually test window rendering on line 0
	ppuInstance.Reset()
	
	// The issue is that we need to test the window rendering when PPU is actually on line 0
	// Let's create a simple test by directly calling the render functions
	
	// First, let's see what happens with background rendering
	fmt.Println("Testing background rendering on line 0...")
	
	// Set up background tile data too
	mmu.WriteByte(0x8010, 0xAA) // Tile 1, line 0, low byte (alternating pattern)
	mmu.WriteByte(0x8011, 0x55) // Tile 1, line 0, high byte
	mmu.WriteByte(0x9800, 0x01) // Background tile map entry 0 -> use tile 1
	
	// Clear buffer
	for i := range buffer {
		buffer[i] = 0
	}
	
	// We need to simulate being on line 0 in VRAM mode
	// Let's step the PPU carefully
	stepCount := 0
	for stepCount < 1000 { // Safety limit
		ppuInstance.Step(1)
		stepCount++
		
		line := ppuInstance.GetCurrentLine()
		mode := ppuInstance.GetCurrentMode()
		
		if line == 0 && mode == 3 { // VRAM mode on line 0
			fmt.Printf("Found line 0, VRAM mode after %d steps\n", stepCount)
			break
		}
		
		if line > 0 {
			fmt.Printf("Moved to line %d, resetting...\n", line)
			ppuInstance.Reset()
			stepCount = 0
		}
	}

	fmt.Printf("Final PPU state - Line: %d, Mode: %d\n", ppuInstance.GetCurrentLine(), ppuInstance.GetCurrentMode())

	// Display first few pixels
	fmt.Println("\nFirst 10 pixels of screen buffer:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Pixel %d: %d\n", i, buffer[i])
	}

	// Test different WX values
	fmt.Println("\nTesting different WX values:")
	
	testCases := []struct {
		wx   byte
		desc string
	}{
		{0, "WX=0 (disabled)"},
		{3, "WX=3 (should be treated as 0)"},
		{7, "WX=7 (window at position 0)"},
		{10, "WX=10 (window at position 3)"},
		{167, "WX=167 (disabled)"},
	}

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		mmu.WriteByte(0xFF4B, tc.wx)
		
		// Clear buffer
		for i := range buffer {
			buffer[i] = 0
		}
		
		// Reset PPU to line 0
		ppuInstance.Reset()
		
		// Step to render one scanline
		for i := 0; i < 456; i++ {
			ppuInstance.Step(1)
		}
		
		// Show first 5 pixels
		for i := 0; i < 5; i++ {
			fmt.Printf("  Pixel %d: %d\n", i, buffer[i])
		}
	}

	fmt.Println("\nWindow test completed!")
}
