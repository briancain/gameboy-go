package display

import (
	"testing"
)

// Mock emulator for testing
type MockEmulator struct {
	running      bool
	screenBuffer []byte
}

func (m *MockEmulator) Step() error {
	return nil
}

func (m *MockEmulator) StepInstruction() (int, error) {
	return 4, nil // Return a typical instruction cycle count
}

func (m *MockEmulator) GetPPUDebugInfo() map[string]interface{} {
	return map[string]interface{}{
		"lcdc_enabled": true,
		"bg_enabled":   true,
		"lcdc_value":   0x91,
	}
}

func (m *MockEmulator) GetScreenBuffer() []byte {
	if m.screenBuffer == nil {
		// Create a simple test pattern
		m.screenBuffer = make([]byte, SCREEN_WIDTH*SCREEN_HEIGHT)
		for i := range m.screenBuffer {
			m.screenBuffer[i] = byte(i % 4) // Create a pattern with all 4 colors
		}
	}
	return m.screenBuffer
}

func (m *MockEmulator) IsRunning() bool {
	return m.running
}

func (m *MockEmulator) Exit() {
	m.running = false
}

// Mock input handler for testing
type MockInputHandler struct {
	buttonStates map[string]bool
}

func (m *MockInputHandler) SetButtonState(button string, pressed bool) {
	if m.buttonStates == nil {
		m.buttonStates = make(map[string]bool)
	}
	m.buttonStates[button] = pressed
}

func TestEbitenDisplayCreation(t *testing.T) {
	mockEmulator := &MockEmulator{running: true}
	mockInputHandler := &MockInputHandler{}

	display := NewEbitenDisplay(mockEmulator, mockInputHandler, 2, false)

	if display == nil {
		t.Error("Display should not be nil")
	}

	if display.scale != 2 {
		t.Errorf("Expected scale 2, got %d", display.scale)
	}
}

func TestColorConversion(t *testing.T) {
	mockEmulator := &MockEmulator{running: true}
	mockInputHandler := &MockInputHandler{}

	display := NewEbitenDisplay(mockEmulator, mockInputHandler, 1, false)

	// Test with a full screen buffer (this is how it's used in practice)
	testBuffer := make([]byte, SCREEN_WIDTH*SCREEN_HEIGHT)
	// Set first few pixels to different colors
	testBuffer[0] = 0 // White
	testBuffer[1] = 1 // Light gray
	testBuffer[2] = 2 // Dark gray
	testBuffer[3] = 3 // Black

	rgbData := display.convertToRGB(testBuffer)

	// Should be full screen * 4 bytes per pixel (RGBA)
	expectedLength := SCREEN_WIDTH * SCREEN_HEIGHT * 4
	if len(rgbData) != expectedLength {
		t.Errorf("Expected RGB data length %d, got %d", expectedLength, len(rgbData))
	}

	// Check first pixel (color 0 = white)
	if rgbData[0] != 255 || rgbData[1] != 255 || rgbData[2] != 255 || rgbData[3] != 255 {
		t.Error("First pixel should be white (255,255,255,255)")
	}

	// Check second pixel (color 1 = light gray)
	if rgbData[4] != 170 || rgbData[5] != 170 || rgbData[6] != 170 || rgbData[7] != 255 {
		t.Error("Second pixel should be light gray (170,170,170,255)")
	}

	// Check fourth pixel (color 3 = black)
	if rgbData[12] != 0 || rgbData[13] != 0 || rgbData[14] != 0 || rgbData[15] != 255 {
		t.Error("Fourth pixel should be black (0,0,0,255)")
	}
}

func TestInputHandling(t *testing.T) {
	mockEmulator := &MockEmulator{running: true}
	mockInputHandler := &MockInputHandler{}

	display := NewEbitenDisplay(mockEmulator, mockInputHandler, 1, false)

	// Test that input handler is set correctly
	if display.inputHandler != mockInputHandler {
		t.Error("Input handler should be set correctly")
	}
}
