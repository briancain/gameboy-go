package gbcore

// PPU modes
const (
	MODE_HBLANK = 0
	MODE_VBLANK = 1
	MODE_OAM    = 2
	MODE_VRAM   = 3
)

// LCD Control Register (LCDC) bits
const (
	LCDC_BG_ENABLE        = 0x01 // Bit 0 - BG Display Enable
	LCDC_OBJ_ENABLE       = 0x02 // Bit 1 - OBJ (Sprite) Display Enable
	LCDC_OBJ_SIZE         = 0x04 // Bit 2 - OBJ (Sprite) Size (0=8x8, 1=8x16)
	LCDC_BG_TILEMAP       = 0x08 // Bit 3 - BG Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
	LCDC_TILE_DATA        = 0x10 // Bit 4 - BG & Window Tile Data Select (0=8800-97FF, 1=8000-8FFF)
	LCDC_WINDOW_ENABLE    = 0x20 // Bit 5 - Window Display Enable
	LCDC_WINDOW_TILEMAP   = 0x40 // Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
	LCDC_DISPLAY_ENABLE   = 0x80 // Bit 7 - LCD Display Enable
)

// LCD Status Register (STAT) bits
const (
	STAT_MODE          = 0x03 // Bit 0-1 - Mode Flag
	STAT_LYC_EQUAL     = 0x04 // Bit 2 - Coincidence Flag (0=LYC!=LY, 1=LYC=LY)
	STAT_HBLANK_INT    = 0x08 // Bit 3 - Mode 0 H-Blank Interrupt
	STAT_VBLANK_INT    = 0x10 // Bit 4 - Mode 1 V-Blank Interrupt
	STAT_OAM_INT       = 0x20 // Bit 5 - Mode 2 OAM Interrupt
	STAT_LYC_INT       = 0x40 // Bit 6 - LYC=LY Coincidence Interrupt
)

// Game Boy screen dimensions
const (
	SCREEN_WIDTH  = 160
	SCREEN_HEIGHT = 144
)

// PPU (Picture Processing Unit) handles the Game Boy's graphics
type PPU struct {
	// Screen buffer (160x144 pixels, 4 colors per pixel)
	screenBuffer [SCREEN_WIDTH * SCREEN_HEIGHT]byte

	// Current PPU state
	mode      byte
	modeClock int
	line      byte // LY register

	// Reference to MMU for memory access
	mmu MMU
}

// MMU interface for PPU to access memory
type MMU interface {
	ReadByte(addr uint16) byte
	WriteByte(addr uint16, value byte)
}

// Initialize a new PPU
func NewPPU(mmu MMU) *PPU {
	ppu := &PPU{
		mmu:       mmu,
		mode:      MODE_OAM,
		modeClock: 0,
		line:      0,
	}
	
	// Clear screen buffer
	for i := range ppu.screenBuffer {
		ppu.screenBuffer[i] = 0
	}
	
	return ppu
}

// Reset the PPU to initial state
func (ppu *PPU) Reset() {
	ppu.mode = MODE_OAM
	ppu.modeClock = 0
	ppu.line = 0
	
	// Clear screen buffer
	for i := range ppu.screenBuffer {
		ppu.screenBuffer[i] = 0
	}
}

// Step advances the PPU by the specified number of cycles
func (ppu *PPU) Step(cycles int) {
	// Check if LCD is enabled
	lcdc := ppu.mmu.ReadByte(0xFF40)
	if (lcdc & LCDC_DISPLAY_ENABLE) == 0 {
		// LCD is disabled
		return
	}
	
	// Update mode clock
	ppu.modeClock += cycles
	
	// Process based on current mode
	switch ppu.mode {
	case MODE_OAM:
		// OAM Search - 80 cycles
		if ppu.modeClock >= 80 {
			ppu.modeClock -= 80
			ppu.mode = MODE_VRAM
			ppu.updateSTAT()
		}
	
	case MODE_VRAM:
		// Pixel Transfer - 172 cycles
		if ppu.modeClock >= 172 {
			ppu.modeClock -= 172
			ppu.mode = MODE_HBLANK
			ppu.updateSTAT()
			
			// Render scanline
			ppu.renderScanline()
		}
	
	case MODE_HBLANK:
		// H-Blank - 204 cycles
		if ppu.modeClock >= 204 {
			ppu.modeClock -= 204
			ppu.line++
			
			// Check if we've reached the bottom of the screen
			if ppu.line == 144 {
				ppu.mode = MODE_VBLANK
				ppu.updateSTAT()
				
				// Request V-Blank interrupt
				interruptFlag := ppu.mmu.ReadByte(0xFF0F)
				ppu.mmu.WriteByte(0xFF0F, interruptFlag | 0x01)
			} else {
				ppu.mode = MODE_OAM
				ppu.updateSTAT()
			}
			
			// Update LY register
			ppu.mmu.WriteByte(0xFF44, ppu.line)
			
			// Check LY=LYC coincidence
			ppu.checkLYC()
		}
	
	case MODE_VBLANK:
		// V-Blank - 4560 cycles (10 lines, each 456 cycles)
		if ppu.modeClock >= 456 {
			ppu.modeClock -= 456
			ppu.line++
			
			// Update LY register
			ppu.mmu.WriteByte(0xFF44, ppu.line)
			
			// Check LY=LYC coincidence
			ppu.checkLYC()
			
			// End of V-Blank
			if ppu.line > 153 {
				ppu.mode = MODE_OAM
				ppu.line = 0
				ppu.mmu.WriteByte(0xFF44, 0)
				ppu.updateSTAT()
				ppu.checkLYC()
			}
		}
	}
}

// Update the STAT register based on current mode
func (ppu *PPU) updateSTAT() {
	stat := ppu.mmu.ReadByte(0xFF41)
	
	// Clear mode bits and set new mode
	stat &= 0xFC
	stat |= ppu.mode
	
	// Check if we need to request STAT interrupt
	if ((stat & STAT_HBLANK_INT) != 0 && ppu.mode == MODE_HBLANK) ||
	   ((stat & STAT_VBLANK_INT) != 0 && ppu.mode == MODE_VBLANK) ||
	   ((stat & STAT_OAM_INT) != 0 && ppu.mode == MODE_OAM) {
		// Request STAT interrupt
		interruptFlag := ppu.mmu.ReadByte(0xFF0F)
		ppu.mmu.WriteByte(0xFF0F, interruptFlag | 0x02)
	}
	
	ppu.mmu.WriteByte(0xFF41, stat)
}

// Check LY=LYC coincidence and update STAT register
func (ppu *PPU) checkLYC() {
	stat := ppu.mmu.ReadByte(0xFF41)
	lyc := ppu.mmu.ReadByte(0xFF45)
	
	if ppu.line == lyc {
		// Set coincidence flag
		stat |= STAT_LYC_EQUAL
		
		// Check if we need to request STAT interrupt
		if (stat & STAT_LYC_INT) != 0 {
			interruptFlag := ppu.mmu.ReadByte(0xFF0F)
			ppu.mmu.WriteByte(0xFF0F, interruptFlag | 0x02)
		}
	} else {
		// Clear coincidence flag
		stat &= ^byte(STAT_LYC_EQUAL)
	}
	
	ppu.mmu.WriteByte(0xFF41, stat)
}

// Render the current scanline
func (ppu *PPU) renderScanline() {
	lcdc := ppu.mmu.ReadByte(0xFF40)
	
	// Render background if enabled
	if (lcdc & LCDC_BG_ENABLE) != 0 {
		ppu.renderBackground()
	}
	
	// Render window if enabled
	if (lcdc & LCDC_WINDOW_ENABLE) != 0 {
		ppu.renderWindow()
	}
	
	// Render sprites if enabled
	if (lcdc & LCDC_OBJ_ENABLE) != 0 {
		ppu.renderSprites()
	}
}

// Render the background for the current scanline
func (ppu *PPU) renderBackground() {
	lcdc := ppu.mmu.ReadByte(0xFF40)
	
	// Get background tile map address
	tileMapAddr := uint16(0x9800)
	if (lcdc & LCDC_BG_TILEMAP) != 0 {
		tileMapAddr = 0x9C00
	}
	
	// Get tile data address
	tileDataAddr := uint16(0x8800)
	tileDataSigned := true
	if (lcdc & LCDC_TILE_DATA) != 0 {
		tileDataAddr = 0x8000
		tileDataSigned = false
	}
	
	// Get scroll positions
	scrollY := ppu.mmu.ReadByte(0xFF42)
	scrollX := ppu.mmu.ReadByte(0xFF43)
	
	// Get palette
	bgp := ppu.mmu.ReadByte(0xFF47)
	
	// Calculate which row of tiles to use
	y := scrollY + ppu.line
	tileRow := uint16(y / 8)
	
	// Calculate which pixel row of the tile to use
	pixelY := y % 8
	
	// Render the scanline
	for x := byte(0); x < SCREEN_WIDTH; x++ {
		// Calculate which tile column to use
		pixelX := scrollX + x
		tileCol := uint16(pixelX / 8)
		
		// Get the tile index
		tileMapOffset := tileRow*32 + tileCol
		tileIndex := ppu.mmu.ReadByte(tileMapAddr + tileMapOffset)
		
		// Get the tile data address
		var tileDataOffset uint16
		if tileDataSigned {
			// 8800 method - tile index is signed
			tileDataOffset = uint16(int16(0x1000) + int16(int8(tileIndex))*16)
		} else {
			// 8000 method - tile index is unsigned
			tileDataOffset = uint16(tileIndex) * 16
		}
		
		// Get the pixel data from the tile
		tileAddr := tileDataAddr + tileDataOffset + uint16(pixelY*2)
		tileLow := ppu.mmu.ReadByte(tileAddr)
		tileHigh := ppu.mmu.ReadByte(tileAddr + 1)
		
		// Get the color value (0-3) for this pixel
		colorBit := 7 - (pixelX % 8)
		colorValue := ((tileHigh >> colorBit) & 1) << 1 | ((tileLow >> colorBit) & 1)
		
		// Map the color value through the palette
		colorIndex := (bgp >> (colorValue * 2)) & 0x03
		
		// Set the pixel in the screen buffer
		bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(x)
		ppu.screenBuffer[bufferIndex] = colorIndex
	}
}

// Render the window for the current scanline
func (ppu *PPU) renderWindow() {
	lcdc := ppu.mmu.ReadByte(0xFF40)
	
	// Get window position
	windowY := ppu.mmu.ReadByte(0xFF4A)
	windowX := ppu.mmu.ReadByte(0xFF4B) - 7
	
	// Check if the window is visible on this scanline
	if ppu.line < windowY {
		return
	}
	
	// Get window tile map address
	tileMapAddr := uint16(0x9800)
	if (lcdc & LCDC_WINDOW_TILEMAP) != 0 {
		tileMapAddr = 0x9C00
	}
	
	// Get tile data address
	tileDataAddr := uint16(0x8800)
	tileDataSigned := true
	if (lcdc & LCDC_TILE_DATA) != 0 {
		tileDataAddr = 0x8000
		tileDataSigned = false
	}
	
	// Get palette
	bgp := ppu.mmu.ReadByte(0xFF47)
	
	// Calculate which row of tiles to use
	y := ppu.line - windowY
	tileRow := uint16(y / 8)
	
	// Calculate which pixel row of the tile to use
	pixelY := y % 8
	
	// Render the scanline
	for x := byte(0); x < SCREEN_WIDTH; x++ {
		// Skip if this pixel is not in the window
		if x < windowX {
			continue
		}
		
		// Calculate which tile column to use
		pixelX := x - windowX
		tileCol := uint16(pixelX / 8)
		
		// Get the tile index
		tileMapOffset := tileRow*32 + tileCol
		tileIndex := ppu.mmu.ReadByte(tileMapAddr + tileMapOffset)
		
		// Get the tile data address
		var tileDataOffset uint16
		if tileDataSigned {
			// 8800 method - tile index is signed
			tileDataOffset = uint16(int16(0x1000) + int16(int8(tileIndex))*16)
		} else {
			// 8000 method - tile index is unsigned
			tileDataOffset = uint16(tileIndex) * 16
		}
		
		// Get the pixel data from the tile
		tileAddr := tileDataAddr + tileDataOffset + uint16(pixelY*2)
		tileLow := ppu.mmu.ReadByte(tileAddr)
		tileHigh := ppu.mmu.ReadByte(tileAddr + 1)
		
		// Get the color value (0-3) for this pixel
		colorBit := 7 - (pixelX % 8)
		colorValue := ((tileHigh >> colorBit) & 1) << 1 | ((tileLow >> colorBit) & 1)
		
		// Map the color value through the palette
		colorIndex := (bgp >> (colorValue * 2)) & 0x03
		
		// Set the pixel in the screen buffer
		bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(x)
		ppu.screenBuffer[bufferIndex] = colorIndex
	}
}

// Render sprites for the current scanline
func (ppu *PPU) renderSprites() {
	lcdc := ppu.mmu.ReadByte(0xFF40)
	
	// Determine sprite size (8x8 or 8x16)
	spriteHeight := byte(8)
	if (lcdc & LCDC_OBJ_SIZE) != 0 {
		spriteHeight = 16
	}
	
	// Get sprite palettes
	obp0 := ppu.mmu.ReadByte(0xFF48)
	obp1 := ppu.mmu.ReadByte(0xFF49)
	
	// Maximum of 10 sprites per scanline
	spritesOnLine := 0
	
	// Check all 40 sprites
	for sprite := byte(0); sprite < 40; sprite++ {
		// Get sprite attributes from OAM
		oamAddr := 0xFE00 + uint16(sprite)*4
		spriteY := ppu.mmu.ReadByte(oamAddr) - 16
		spriteX := ppu.mmu.ReadByte(oamAddr + 1) - 8
		tileIndex := ppu.mmu.ReadByte(oamAddr + 2)
		attributes := ppu.mmu.ReadByte(oamAddr + 3)
		
		// If using 8x16 sprites, the lower bit of the tile index is ignored
		if spriteHeight == 16 {
			tileIndex &= 0xFE
		}
		
		// Check if sprite is on this scanline
		if ppu.line < spriteY || ppu.line >= spriteY+spriteHeight {
			continue
		}
		
		// Limit to 10 sprites per scanline
		spritesOnLine++
		if spritesOnLine > 10 {
			break
		}
		
		// Get attributes
		yFlip := (attributes & 0x40) != 0
		xFlip := (attributes & 0x20) != 0
		priority := (attributes & 0x80) != 0
		palette := obp0
		if (attributes & 0x10) != 0 {
			palette = obp1
		}
		
		// Calculate which row of the sprite to use
		pixelY := ppu.line - spriteY
		if yFlip {
			pixelY = spriteHeight - 1 - pixelY
		}
		
		// Get the tile data
		tileAddr := 0x8000 + uint16(tileIndex)*16 + uint16(pixelY)*2
		tileLow := ppu.mmu.ReadByte(tileAddr)
		tileHigh := ppu.mmu.ReadByte(tileAddr + 1)
		
		// Draw the sprite row
		for x := byte(0); x < 8; x++ {
			// Skip if sprite is off-screen
			if spriteX+x >= SCREEN_WIDTH {
				continue
			}
			
			// Get the color bit
			colorBit := byte(7 - x)
			if xFlip {
				colorBit = x
			}
			
			// Get the color value (0-3)
			colorValue := ((tileHigh >> colorBit) & 1) << 1 | ((tileLow >> colorBit) & 1)
			
			// Color 0 is transparent for sprites
			if colorValue == 0 {
				continue
			}
			
			// Check sprite priority
			if priority {
				// If priority bit is set, sprite is behind background color 1-3
				bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(spriteX+x)
				bgColor := ppu.screenBuffer[bufferIndex]
				if bgColor != 0 {
					continue
				}
			}
			
			// Map the color value through the palette
			colorIndex := (palette >> (colorValue * 2)) & 0x03
			
			// Set the pixel in the screen buffer
			bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(spriteX+x)
			ppu.screenBuffer[bufferIndex] = colorIndex
		}
	}
}

// Get the screen buffer
func (ppu *PPU) GetScreenBuffer() []byte {
	return ppu.screenBuffer[:]
}
