package ppu

// PPU modes
const (
	MODE_HBLANK = 0
	MODE_VBLANK = 1
	MODE_OAM    = 2
	MODE_VRAM   = 3
)

// LCD Control Register (LCDC) bits
const (
	LCDC_BG_ENABLE      = 0x01 // Bit 0 - BG Display Enable
	LCDC_OBJ_ENABLE     = 0x02 // Bit 1 - OBJ (Sprite) Display Enable
	LCDC_OBJ_SIZE       = 0x04 // Bit 2 - OBJ (Sprite) Size (0=8x8, 1=8x16)
	LCDC_BG_TILEMAP     = 0x08 // Bit 3 - BG Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
	LCDC_TILE_DATA      = 0x10 // Bit 4 - BG & Window Tile Data Select (0=8800-97FF, 1=8000-8FFF)
	LCDC_WINDOW_ENABLE  = 0x20 // Bit 5 - Window Display Enable
	LCDC_WINDOW_TILEMAP = 0x40 // Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
	LCDC_DISPLAY_ENABLE = 0x80 // Bit 7 - LCD Display Enable
)

// LCD Status Register (STAT) bits
const (
	STAT_MODE       = 0x03 // Bit 0-1 - Mode Flag
	STAT_LYC_EQUAL  = 0x04 // Bit 2 - Coincidence Flag (0=LYC!=LY, 1=LYC=LY)
	STAT_HBLANK_INT = 0x08 // Bit 3 - Mode 0 H-Blank Interrupt
	STAT_VBLANK_INT = 0x10 // Bit 4 - Mode 1 V-Blank Interrupt
	STAT_OAM_INT    = 0x20 // Bit 5 - Mode 2 OAM Interrupt
	STAT_LYC_INT    = 0x40 // Bit 6 - LYC=LY Coincidence Interrupt
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
				ppu.requestVBlankInterrupt()
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
	oldMode := stat & STAT_MODE

	// Clear mode bits and set new mode
	stat &= 0xFC
	stat |= ppu.mode

	// Write the updated STAT register
	ppu.mmu.WriteByte(0xFF41, stat)

	// Check if we need to request STAT interrupt on mode change
	// Only trigger interrupt on mode transitions, not every update
	if oldMode != ppu.mode {
		ppu.checkSTATInterrupt(stat)
	}
}

// Check if a STAT interrupt should be triggered
func (ppu *PPU) checkSTATInterrupt(stat byte) {
	shouldTrigger := false

	// Check each interrupt condition
	switch ppu.mode {
	case MODE_HBLANK:
		shouldTrigger = (stat & STAT_HBLANK_INT) != 0
	case MODE_VBLANK:
		shouldTrigger = (stat & STAT_VBLANK_INT) != 0
	case MODE_OAM:
		shouldTrigger = (stat & STAT_OAM_INT) != 0
	}

	// Trigger STAT interrupt if conditions are met
	if shouldTrigger {
		ppu.requestSTATInterrupt()
	}
}

// Request a STAT interrupt
func (ppu *PPU) requestSTATInterrupt() {
	interruptFlag := ppu.mmu.ReadByte(0xFF0F)
	interruptFlag |= 0x02 // LCDC Status interrupt (bit 1)
	ppu.mmu.WriteByte(0xFF0F, interruptFlag)
}

// Request a V-Blank interrupt
func (ppu *PPU) requestVBlankInterrupt() {
	interruptFlag := ppu.mmu.ReadByte(0xFF0F)
	interruptFlag |= 0x01 // V-Blank interrupt (bit 0)
	ppu.mmu.WriteByte(0xFF0F, interruptFlag)
}

// Check LY=LYC coincidence and update STAT register
func (ppu *PPU) checkLYC() {
	stat := ppu.mmu.ReadByte(0xFF41)
	lyc := ppu.mmu.ReadByte(0xFF45)
	wasCoincident := (stat & STAT_LYC_EQUAL) != 0

	if ppu.line == lyc {
		// Set coincidence flag
		stat |= STAT_LYC_EQUAL

		// Check if we need to request STAT interrupt (only on transition to coincident)
		if !wasCoincident && (stat&STAT_LYC_INT) != 0 {
			ppu.requestSTATInterrupt()
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
	tileRow := uint16(y/8) % 32 // Wrap around at 32 tiles

	// Calculate which pixel row of the tile to use
	pixelY := y % 8

	// Render the scanline
	for x := byte(0); x < SCREEN_WIDTH; x++ {
		// Calculate which tile column to use
		pixelX := scrollX + x
		tileCol := uint16(pixelX/8) % 32 // Wrap around at 32 tiles

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
		colorValue := ((tileHigh>>colorBit)&1)<<1 | ((tileLow >> colorBit) & 1)

		// Map the color value through the palette
		colorIndex := (bgp >> (colorValue * 2)) & 0x03

		// Set the pixel in the screen buffer
		// Bounds check to prevent buffer overflow
		if ppu.line < SCREEN_HEIGHT {
			bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(x)
			ppu.screenBuffer[bufferIndex] = colorIndex
		}
	}
}

// Render the window for the current scanline
func (ppu *PPU) renderWindow() {
	lcdc := ppu.mmu.ReadByte(0xFF40)

	// Get window position
	windowY := ppu.mmu.ReadByte(0xFF4A)    // WY
	windowXRaw := ppu.mmu.ReadByte(0xFF4B) // WX

	// Check if the window is visible on this scanline
	if ppu.line < windowY {
		return
	}

	// WX values 0-6 and 167+ disable the window
	if windowXRaw == 0 || windowXRaw >= 167 {
		return
	}

	// Calculate actual window X position (WX - 7)
	// Handle the case where WX < 7 (should be treated as 0)
	var windowX byte
	if windowXRaw >= 7 {
		windowX = windowXRaw - 7
	} else {
		windowX = 0
	}

	// If window starts beyond screen width, nothing to render
	if windowX >= SCREEN_WIDTH {
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

	// Calculate which row of tiles to use (window line counter)
	windowLine := ppu.line - windowY
	tileRow := uint16(windowLine / 8)

	// Bounds check: tile map is 32x32 tiles
	if tileRow >= 32 {
		return
	}

	// Calculate which pixel row of the tile to use
	pixelY := windowLine % 8

	// Render the scanline
	for x := windowX; x < SCREEN_WIDTH; x++ {
		// Calculate which tile column to use
		pixelX := x - windowX
		tileCol := uint16(pixelX / 8)

		// Bounds check: tile map is 32x32 tiles
		if tileCol >= 32 {
			break
		}

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
		colorValue := ((tileHigh>>colorBit)&1)<<1 | ((tileLow >> colorBit) & 1)

		// Map the color value through the palette
		colorIndex := (bgp >> (colorValue * 2)) & 0x03

		// Set the pixel in the screen buffer
		// Bounds check to prevent buffer overflow
		if ppu.line < SCREEN_HEIGHT {
			bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(x)
			ppu.screenBuffer[bufferIndex] = colorIndex
		}
	}
}

// Sprite data structure for priority handling
type SpriteData struct {
	oamIndex   byte
	x          byte
	y          byte
	tileIndex  byte
	attributes byte
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

	// Collect all sprites on this scanline
	var spritesOnLine []SpriteData

	// Check all 40 sprites
	for sprite := byte(0); sprite < 40; sprite++ {
		// Get sprite attributes from OAM
		oamAddr := 0xFE00 + uint16(sprite)*4
		spriteY := ppu.mmu.ReadByte(oamAddr) - 16
		spriteX := ppu.mmu.ReadByte(oamAddr+1) - 8
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

		// Add sprite to the list
		spritesOnLine = append(spritesOnLine, SpriteData{
			oamIndex:   sprite,
			x:          spriteX,
			y:          spriteY,
			tileIndex:  tileIndex,
			attributes: attributes,
		})

		// Limit to 10 sprites per scanline
		if len(spritesOnLine) >= 10 {
			break
		}
	}

	// Sort sprites by priority:
	// 1. Lower X coordinate has higher priority (appears on top)
	// 2. If X coordinates are equal, lower OAM index has higher priority
	ppu.sortSpritesByPriority(spritesOnLine)

	// Render sprites in reverse order (lowest priority first)
	// This ensures higher priority sprites overwrite lower priority ones
	for i := len(spritesOnLine) - 1; i >= 0; i-- {
		sprite := spritesOnLine[i]
		ppu.renderSprite(sprite, spriteHeight, obp0, obp1)
	}
}

// Sort sprites by priority according to GameBoy rules
func (ppu *PPU) sortSpritesByPriority(sprites []SpriteData) {
	// Simple bubble sort - efficient for small arrays (max 10 sprites)
	n := len(sprites)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare sprites[j] and sprites[j+1]
			// Higher priority sprite should come first
			if ppu.shouldSwapSprites(sprites[j], sprites[j+1]) {
				sprites[j], sprites[j+1] = sprites[j+1], sprites[j]
			}
		}
	}
}

// Determine if two sprites should be swapped based on priority
func (ppu *PPU) shouldSwapSprites(a, b SpriteData) bool {
	// If X coordinates are different, lower X has higher priority
	if a.x != b.x {
		return a.x > b.x // Swap if a.x > b.x (b has higher priority)
	}

	// If X coordinates are the same, lower OAM index has higher priority
	return a.oamIndex > b.oamIndex // Swap if a.oamIndex > b.oamIndex (b has higher priority)
}

// Render a single sprite
func (ppu *PPU) renderSprite(sprite SpriteData, spriteHeight byte, obp0, obp1 byte) {
	// Get attributes
	yFlip := (sprite.attributes & 0x40) != 0
	xFlip := (sprite.attributes & 0x20) != 0
	priority := (sprite.attributes & 0x80) != 0
	palette := obp0
	if (sprite.attributes & 0x10) != 0 {
		palette = obp1
	}

	// Calculate which row of the sprite to use
	pixelY := ppu.line - sprite.y
	if yFlip {
		pixelY = spriteHeight - 1 - pixelY
	}

	// Get the tile data
	tileAddr := 0x8000 + uint16(sprite.tileIndex)*16 + uint16(pixelY)*2
	tileLow := ppu.mmu.ReadByte(tileAddr)
	tileHigh := ppu.mmu.ReadByte(tileAddr + 1)

	// Draw the sprite row
	for x := byte(0); x < 8; x++ {
		// Skip if sprite is off-screen
		if sprite.x+x >= SCREEN_WIDTH {
			continue
		}

		// Get the color bit
		colorBit := byte(7 - x)
		if xFlip {
			colorBit = x
		}

		// Get the color value (0-3)
		colorValue := ((tileHigh>>colorBit)&1)<<1 | ((tileLow >> colorBit) & 1)

		// Color 0 is transparent for sprites
		if colorValue == 0 {
			continue
		}

		// Check sprite priority
		if priority {
			// If priority bit is set, sprite is behind background color 1-3
			bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(sprite.x+x)
			bgColor := ppu.screenBuffer[bufferIndex]
			if bgColor != 0 {
				continue
			}
		}

		// Map the color value through the palette
		colorIndex := (palette >> (colorValue * 2)) & 0x03

		// Set the pixel in the screen buffer
		bufferIndex := uint16(ppu.line)*SCREEN_WIDTH + uint16(sprite.x+x)
		ppu.screenBuffer[bufferIndex] = colorIndex
	}
}

// Get the screen buffer
func (ppu *PPU) GetScreenBuffer() []byte {
	return ppu.screenBuffer[:]
}

// Get the screen buffer with RGB values for display
func (ppu *PPU) GetScreenBufferRGB() []byte {
	// Convert 2-bit color values to RGB (grayscale)
	// 0 = White, 1 = Light Gray, 2 = Dark Gray, 3 = Black
	colorMap := [4][3]byte{
		{255, 255, 255}, // White
		{170, 170, 170}, // Light Gray
		{85, 85, 85},    // Dark Gray
		{0, 0, 0},       // Black
	}

	rgbBuffer := make([]byte, SCREEN_WIDTH*SCREEN_HEIGHT*3)

	for i, colorIndex := range ppu.screenBuffer {
		rgb := colorMap[colorIndex&0x03] // Ensure we only use 2 bits
		rgbBuffer[i*3] = rgb[0]          // R
		rgbBuffer[i*3+1] = rgb[1]        // G
		rgbBuffer[i*3+2] = rgb[2]        // B
	}

	return rgbBuffer
}

// Get screen dimensions
func (ppu *PPU) GetScreenWidth() int {
	return SCREEN_WIDTH
}

func (ppu *PPU) GetScreenHeight() int {
	return SCREEN_HEIGHT
}

// Debug functions for PPU state inspection
func (ppu *PPU) GetCurrentMode() byte {
	return ppu.mode
}

func (ppu *PPU) GetCurrentLine() byte {
	return ppu.line
}

func (ppu *PPU) GetModeClock() int {
	return ppu.modeClock
}

// Check if LCD is enabled
func (ppu *PPU) IsLCDEnabled() bool {
	lcdc := ppu.mmu.ReadByte(0xFF40)
	return (lcdc & LCDC_DISPLAY_ENABLE) != 0
}

// Get LCDC register value with bit meanings
func (ppu *PPU) GetLCDCStatus() map[string]bool {
	lcdc := ppu.mmu.ReadByte(0xFF40)
	return map[string]bool{
		"display_enable":   (lcdc & LCDC_DISPLAY_ENABLE) != 0,
		"window_tilemap":   (lcdc & LCDC_WINDOW_TILEMAP) != 0,
		"window_enable":    (lcdc & LCDC_WINDOW_ENABLE) != 0,
		"tile_data_select": (lcdc & LCDC_TILE_DATA) != 0,
		"bg_tilemap":       (lcdc & LCDC_BG_TILEMAP) != 0,
		"obj_size":         (lcdc & LCDC_OBJ_SIZE) != 0,
		"obj_enable":       (lcdc & LCDC_OBJ_ENABLE) != 0,
		"bg_enable":        (lcdc & LCDC_BG_ENABLE) != 0,
	}
}

// Display the screen buffer in terminal (for debugging)
func (ppu *PPU) DisplayScreenBuffer() {
	// Character map for different gray levels
	chars := []rune{' ', '░', '▒', '█'} // 0=white, 1=light gray, 2=dark gray, 3=black

	for y := 0; y < SCREEN_HEIGHT; y++ {
		for x := 0; x < SCREEN_WIDTH; x++ {
			colorIndex := ppu.screenBuffer[y*SCREEN_WIDTH+x] & 0x03
			print(string(chars[colorIndex]))
		}
		println()
	}
}

// Display a small portion of the screen buffer (for debugging)
func (ppu *PPU) DisplayScreenBufferSection(startX, startY, width, height int) {
	chars := []rune{' ', '░', '▒', '█'}

	for y := startY; y < startY+height && y < SCREEN_HEIGHT; y++ {
		for x := startX; x < startX+width && x < SCREEN_WIDTH; x++ {
			colorIndex := ppu.screenBuffer[y*SCREEN_WIDTH+x] & 0x03
			print(string(chars[colorIndex]))
		}
		println()
	}
}

// WriteRegister handles writes to PPU registers with special behavior
func (ppu *PPU) WriteRegister(addr uint16, value byte) {
	switch addr {
	case 0xFF40: // LCDC - LCD Control
		ppu.handleLCDCWrite(value)
	case 0xFF41: // STAT - LCD Status
		ppu.handleSTATWrite(value)
	case 0xFF42: // SCY - Scroll Y
		// SCY can be written at any time
		// No special side effects needed, value is already stored by MMU
	case 0xFF43: // SCX - Scroll X
		// SCX can be written at any time
		// No special side effects needed, value is already stored by MMU
	case 0xFF44: // LY - LCD Y-Coordinate
		// LY is read-only, writes reset it to 0
		ppu.handleLYWrite()
	case 0xFF45: // LYC - LY Compare
		// LYC can be written at any time, check for coincidence
		// Value is already stored by MMU, just check coincidence
		ppu.checkLYC()
	case 0xFF47: // BGP - Background Palette
		// Palette can be written at any time
		// No special side effects needed, value is already stored by MMU
	case 0xFF48: // OBP0 - Object Palette 0
		// Palette can be written at any time
		// No special side effects needed, value is already stored by MMU
	case 0xFF49: // OBP1 - Object Palette 1
		// Palette can be written at any time
		// No special side effects needed, value is already stored by MMU
	case 0xFF4A: // WY - Window Y Position
		ppu.handleWYWrite(value)
	case 0xFF4B: // WX - Window X Position
		ppu.handleWXWrite(value)
	}
}

// Handle LCDC register writes
func (ppu *PPU) handleLCDCWrite(value byte) {
	// Check if LCD is being enabled or disabled
	newEnabled := (value & LCDC_DISPLAY_ENABLE) != 0

	if !newEnabled {
		// LCD is being turned off
		// This should only be done during V-Blank
		if ppu.mode != MODE_VBLANK {
			// In real hardware, turning off LCD outside V-Blank can damage the screen
			// For emulation, we'll allow it but log a warning
			// TODO: Add proper timing restriction
		}

		// When LCD is turned off, reset PPU state
		ppu.mode = MODE_HBLANK
		ppu.modeClock = 0
		ppu.line = 0

		// Don't update LY register here to avoid recursion
		// The MMU will handle storing the register value

		// Clear the screen buffer
		for i := range ppu.screenBuffer {
			ppu.screenBuffer[i] = 0
		}
	} else if newEnabled {
		// LCD is being turned on (or staying on)
		// Only reset if we were previously off (in HBLANK with line 0)
		if ppu.mode == MODE_HBLANK && ppu.line == 0 {
			ppu.mode = MODE_OAM
			ppu.modeClock = 0
			ppu.line = 0

			// Don't update LY register here to avoid recursion
			// The MMU will handle storing the register value

			// Update STAT register with new mode
			ppu.updateSTAT()
		}
	}
}

// Handle STAT register writes
func (ppu *PPU) handleSTATWrite(value byte) {
	// STAT bits 0-2 are read-only (mode and LYC=LY flag)
	// Only bits 3-6 can be written (interrupt enables)
	currentSTAT := ppu.mmu.ReadByte(0xFF41)

	// Preserve read-only bits (0-2) and update writable bits (3-6)
	newSTAT := (currentSTAT & 0x07) | (value & 0x78)

	// Update the register in memory using direct access to avoid recursion
	if mmuWithDirect, ok := ppu.mmu.(interface{ WriteIODirect(uint16, byte) }); ok {
		mmuWithDirect.WriteIODirect(0xFF41, newSTAT)
	} else {
		// Fallback: MMU doesn't support direct access
		// This shouldn't happen in normal operation, but we handle it gracefully
		// We can't update the register without causing recursion, so we log the issue
		// In a real implementation, we might want to use a logging framework
		// For now, we'll just accept that the original value remains stored
		// TODO: Consider adding proper logging or error handling
	}
}

// Handle LY register writes (resets LY to 0)
func (ppu *PPU) handleLYWrite() {
	// Writing to LY resets it to 0
	ppu.line = 0

	// Update the LY register to 0 using direct access to avoid recursion
	if mmuWithDirect, ok := ppu.mmu.(interface{ WriteIODirect(uint16, byte) }); ok {
		mmuWithDirect.WriteIODirect(0xFF44, 0)
	} else {
		// Fallback: MMU doesn't support direct access
		// This shouldn't happen in normal operation, but we handle it gracefully
		// The internal PPU state (ppu.line) is still updated correctly
		// The register value might not match, but this is better than crashing
	}

	// Check LYC coincidence with new LY value
	ppu.checkLYC()
}

// Handle WY register writes
func (ppu *PPU) handleWYWrite(value byte) {
	// WY can be written at any time
	// Changes take effect immediately for subsequent scanlines
	// No special handling needed beyond validation

	// WY values > 143 effectively disable the window for the current frame
	// This is handled in the rendering logic
}

// Handle WX register writes
func (ppu *PPU) handleWXWrite(value byte) {
	// WX can be written at any time
	// Changes take effect immediately for the current scanline if written during rendering

	// WX values 0-6 and 167+ disable the window
	// This is handled in the rendering logic

	// Note: In real hardware, changing WX during a scanline can cause glitches
	// For now, we'll allow it and handle it in the rendering logic
}
