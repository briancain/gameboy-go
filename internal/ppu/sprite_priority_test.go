package ppu

import (
	"testing"
)

// Test sprite priority sorting
func TestSpritePrioritySorting(t *testing.T) {
	ppu := &PPU{}

	// Create test sprites with different X coordinates and OAM indices
	sprites := []SpriteData{
		{oamIndex: 2, x: 50, y: 50, tileIndex: 0, attributes: 0},
		{oamIndex: 0, x: 40, y: 50, tileIndex: 1, attributes: 0},
		{oamIndex: 1, x: 40, y: 50, tileIndex: 2, attributes: 0}, // Same X as above
		{oamIndex: 3, x: 60, y: 50, tileIndex: 3, attributes: 0},
	}

	ppu.sortSpritesByPriority(sprites)

	// Expected order after sorting (highest priority first):
	// 1. oamIndex: 0, x: 40 (lowest X, lowest OAM index)
	// 2. oamIndex: 1, x: 40 (lowest X, higher OAM index)
	// 3. oamIndex: 2, x: 50 (middle X)
	// 4. oamIndex: 3, x: 60 (highest X)

	expectedOrder := []byte{0, 1, 2, 3}
	for i, expected := range expectedOrder {
		if sprites[i].oamIndex != expected {
			t.Errorf("Sprite at position %d should have OAM index %d, got %d", i, expected, sprites[i].oamIndex)
		}
	}
}

// Test sprite priority comparison
func TestShouldSwapSprites(t *testing.T) {
	ppu := &PPU{}

	// Test case 1: Different X coordinates
	spriteA := SpriteData{oamIndex: 0, x: 50, y: 50, tileIndex: 0, attributes: 0}
	spriteB := SpriteData{oamIndex: 1, x: 40, y: 50, tileIndex: 1, attributes: 0}

	// B should have higher priority (lower X), so we should swap
	if !ppu.shouldSwapSprites(spriteA, spriteB) {
		t.Error("Should swap sprites when B has lower X coordinate")
	}

	// Test case 2: Same X coordinates, different OAM indices
	spriteC := SpriteData{oamIndex: 2, x: 40, y: 50, tileIndex: 0, attributes: 0}
	spriteD := SpriteData{oamIndex: 1, x: 40, y: 50, tileIndex: 1, attributes: 0}

	// D should have higher priority (lower OAM index), so we should swap
	if !ppu.shouldSwapSprites(spriteC, spriteD) {
		t.Error("Should swap sprites when B has lower OAM index with same X coordinate")
	}

	// Test case 3: No swap needed
	spriteE := SpriteData{oamIndex: 1, x: 40, y: 50, tileIndex: 0, attributes: 0}
	spriteF := SpriteData{oamIndex: 2, x: 50, y: 50, tileIndex: 1, attributes: 0}

	// E already has higher priority, no swap needed
	if ppu.shouldSwapSprites(spriteE, spriteF) {
		t.Error("Should not swap sprites when A already has higher priority")
	}
}
