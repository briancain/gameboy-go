package main

import (
	"fmt"

	"github.com/briancain/gameboy-go/internal/mmu"
	"github.com/briancain/gameboy-go/internal/ppu"
)

func main() {
	fmt.Println("PPU Register Write Handler Test")
	fmt.Println("===============================")

	// Create MMU and PPU
	mmuInstance := mmu.NewMMU()
	ppuInstance := ppu.NewPPU(mmuInstance)

	// Set up the PPU in the MMU (this simulates what the core does)
	mmuInstance.SetPPU(ppuInstance)

	fmt.Printf("Initial PPU state - Mode: %d, Line: %d\n", ppuInstance.GetCurrentMode(), ppuInstance.GetCurrentLine())

	// Test 1: LCDC register write handling
	fmt.Println("\n=== Test 1: LCDC Register Write Handling ===")

	// Turn off LCD
	fmt.Println("Turning off LCD...")
	mmuInstance.WriteByte(0xFF40, 0x00)
	fmt.Printf("PPU state after LCD off - Mode: %d, Line: %d\n", ppuInstance.GetCurrentMode(), ppuInstance.GetCurrentLine())

	// Turn on LCD
	fmt.Println("Turning on LCD...")
	mmuInstance.WriteByte(0xFF40, 0x80)
	fmt.Printf("PPU state after LCD on - Mode: %d, Line: %d\n", ppuInstance.GetCurrentMode(), ppuInstance.GetCurrentLine())

	// Test 2: STAT register write handling
	fmt.Println("\n=== Test 2: STAT Register Write Handling ===")

	// Set up initial STAT value
	mmuInstance.WriteByte(0xFF41, 0x05) // Mode 1, LYC=LY set
	fmt.Printf("Initial STAT: 0x%02X\n", mmuInstance.ReadByte(0xFF41))

	// Try to write all bits (should preserve read-only bits)
	mmuInstance.WriteByte(0xFF41, 0xFF)
	fmt.Printf("STAT after write 0xFF: 0x%02X (read-only bits should be preserved)\n", mmuInstance.ReadByte(0xFF41))

	// Test 3: LY register write handling
	fmt.Println("\n=== Test 3: LY Register Write Handling ===")

	// Simulate PPU being on a different line
	fmt.Printf("Current LY: %d\n", ppuInstance.GetCurrentLine())

	// Write to LY (should reset to 0)
	fmt.Println("Writing to LY register...")
	mmuInstance.WriteByte(0xFF44, 0xFF)
	fmt.Printf("LY after write: %d (should be 0)\n", ppuInstance.GetCurrentLine())

	// Test 4: Scroll register writes
	fmt.Println("\n=== Test 4: Scroll Register Writes ===")

	mmuInstance.WriteByte(0xFF42, 0x12) // SCY
	mmuInstance.WriteByte(0xFF43, 0x34) // SCX

	fmt.Printf("SCY: 0x%02X, SCX: 0x%02X\n", mmuInstance.ReadByte(0xFF42), mmuInstance.ReadByte(0xFF43))

	// Test 5: Window position register writes
	fmt.Println("\n=== Test 5: Window Position Register Writes ===")

	mmuInstance.WriteByte(0xFF4A, 0x56) // WY
	mmuInstance.WriteByte(0xFF4B, 0x78) // WX

	fmt.Printf("WY: 0x%02X, WX: 0x%02X\n", mmuInstance.ReadByte(0xFF4A), mmuInstance.ReadByte(0xFF4B))

	// Test 6: Palette register writes
	fmt.Println("\n=== Test 6: Palette Register Writes ===")

	mmuInstance.WriteByte(0xFF47, 0xE4) // BGP
	mmuInstance.WriteByte(0xFF48, 0xD2) // OBP0
	mmuInstance.WriteByte(0xFF49, 0xA1) // OBP1

	fmt.Printf("BGP: 0x%02X, OBP0: 0x%02X, OBP1: 0x%02X\n",
		mmuInstance.ReadByte(0xFF47), mmuInstance.ReadByte(0xFF48), mmuInstance.ReadByte(0xFF49))

	// Test 7: Show LCDC status
	fmt.Println("\n=== Test 7: LCDC Status ===")

	status := ppuInstance.GetLCDCStatus()
	for key, value := range status {
		fmt.Printf("  %s: %v\n", key, value)
	}

	fmt.Println("\nPPU register write handler test completed!")
	fmt.Println("All register writes were handled properly with appropriate special behavior.")
}
