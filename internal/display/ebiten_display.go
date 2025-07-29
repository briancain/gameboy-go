package display

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	// GameBoy screen dimensions
	SCREEN_WIDTH  = 160
	SCREEN_HEIGHT = 144
)

// EbitenDisplay handles the visual output using ebiten
type EbitenDisplay struct {
	// Screen buffer image
	screenImage *ebiten.Image

	// Scale factor for display
	scale int

	// Reference to emulator core for stepping and getting screen data
	emulator Emulator

	// Input handling
	inputHandler InputHandler

	// Debug mode
	debug bool
}

// Emulator interface for the display to interact with the core
type Emulator interface {
	Step() error
	StepInstruction() (int, error)
	GetScreenBuffer() []byte
	IsRunning() bool
	Exit()
}

// InputHandler interface for handling input
type InputHandler interface {
	SetButtonState(button string, pressed bool)
}

// NewEbitenDisplay creates a new ebiten-based display
func NewEbitenDisplay(emulator Emulator, inputHandler InputHandler, scale int, debug bool) *EbitenDisplay {
	if scale < 1 || scale > 4 {
		scale = 2 // Default scale
	}

	return &EbitenDisplay{
		screenImage:  ebiten.NewImage(SCREEN_WIDTH, SCREEN_HEIGHT),
		scale:        scale,
		emulator:     emulator,
		inputHandler: inputHandler,
		debug:        debug,
	}
}

// Update is called every frame by ebiten
func (d *EbitenDisplay) Update() error {
	// Handle input
	d.handleInput()

	// Step the emulator for the right number of cycles per frame
	// GameBoy runs at ~70,224 cycles per frame at 60 FPS
	if d.emulator.IsRunning() {
		// Run approximately 1/60th of GameBoy cycles per update
		cyclesThisUpdate := 0
		targetCycles := 1170 // 70224 / 60

		for cyclesThisUpdate < targetCycles {
			cycles, err := d.emulator.StepInstruction()
			if err != nil {
				log.Printf("Emulator step error: %v", err)
				return err
			}
			cyclesThisUpdate += cycles
		}
	}

	return nil
}

// Draw is called every frame by ebiten to render the screen
func (d *EbitenDisplay) Draw(screen *ebiten.Image) {
	// Get the screen buffer from the emulator
	screenBuffer := d.emulator.GetScreenBuffer()

	// Convert GameBoy 2-bit colors to RGB
	rgbData := d.convertToRGB(screenBuffer)

	// Update the screen image
	d.screenImage.WritePixels(rgbData)

	// Draw the screen image scaled up
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(float64(d.scale), float64(d.scale))
	screen.DrawImage(d.screenImage, opts)

	// Draw debug info if enabled
	if d.debug {
		d.drawDebugInfo(screen)
	}
}

// Layout returns the screen size
func (d *EbitenDisplay) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH * d.scale, SCREEN_HEIGHT * d.scale
}

// convertToRGB converts GameBoy 2-bit color values to RGB
func (d *EbitenDisplay) convertToRGB(screenBuffer []byte) []byte {
	// GameBoy color palette (grayscale)
	// 0 = White, 1 = Light Gray, 2 = Dark Gray, 3 = Black
	colorMap := [4][4]byte{
		{255, 255, 255, 255}, // White (RGBA)
		{170, 170, 170, 255}, // Light Gray
		{85, 85, 85, 255},    // Dark Gray
		{0, 0, 0, 255},       // Black
	}

	rgbData := make([]byte, SCREEN_WIDTH*SCREEN_HEIGHT*4) // RGBA

	for i, colorIndex := range screenBuffer {
		if i >= SCREEN_WIDTH*SCREEN_HEIGHT {
			break // Safety check
		}

		color := colorMap[colorIndex&0x03] // Ensure we only use 2 bits
		rgbData[i*4] = color[0]            // R
		rgbData[i*4+1] = color[1]          // G
		rgbData[i*4+2] = color[2]          // B
		rgbData[i*4+3] = color[3]          // A
	}

	return rgbData
}

// handleInput processes keyboard input and maps it to GameBoy buttons
func (d *EbitenDisplay) handleInput() {
	if d.inputHandler == nil {
		return
	}

	// GameBoy button mappings
	buttonMappings := map[ebiten.Key]string{
		ebiten.KeyArrowUp:    "up",
		ebiten.KeyArrowDown:  "down",
		ebiten.KeyArrowLeft:  "left",
		ebiten.KeyArrowRight: "right",
		ebiten.KeyZ:          "a",      // Z = A button
		ebiten.KeyX:          "b",      // X = B button
		ebiten.KeyEnter:      "start",  // Enter = Start
		ebiten.KeySpace:      "select", // Space = Select
	}

	// Check each button
	for key, button := range buttonMappings {
		if inpututil.IsKeyJustPressed(key) {
			d.inputHandler.SetButtonState(button, true)
		} else if inpututil.IsKeyJustReleased(key) {
			d.inputHandler.SetButtonState(button, false)
		}
	}

	// Handle quit (ESC key)
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		d.emulator.Exit()
	}
}

// drawDebugInfo draws debug information on screen
func (d *EbitenDisplay) drawDebugInfo(screen *ebiten.Image) {
	// Draw FPS
	fps := ebiten.ActualFPS()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f", fps))
}

// Run starts the ebiten game loop
func (d *EbitenDisplay) Run() error {
	// Set window properties
	ebiten.SetWindowSize(SCREEN_WIDTH*d.scale, SCREEN_HEIGHT*d.scale)
	ebiten.SetWindowTitle("GameBoy Go Emulator")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Set TPS to 60 to match GameBoy frame rate
	ebiten.SetTPS(60)

	// Run the game
	return ebiten.RunGame(d)
}

// SetCustomPalette allows setting a custom color palette
func (d *EbitenDisplay) SetCustomPalette(colors [4][3]byte) {
	// This could be implemented to allow custom color schemes
	// For now, we'll stick with the classic GameBoy green palette
}
