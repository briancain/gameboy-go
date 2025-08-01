package core

import (
	"log"
	"time"

	"github.com/briancain/gameboy-go/internal/cartridge"
	"github.com/briancain/gameboy-go/internal/controller"
	"github.com/briancain/gameboy-go/internal/cpu"
	"github.com/briancain/gameboy-go/internal/mmu"
	"github.com/briancain/gameboy-go/internal/ppu"
	"github.com/briancain/gameboy-go/internal/snapshot"
	"github.com/briancain/gameboy-go/internal/sound"
	"github.com/briancain/gameboy-go/internal/timer"
)

// Core emulator implementation
type GameBoyCore struct {
	// Core gameboy components
	Cpu       *cpu.Z80
	Mmu       *mmu.MemoryManagedUnit
	Ppu       *ppu.PPU
	Sound     *sound.Sound
	Timer     *timer.Timer
	Cartridge *cartridge.Cartridge

	// Speed options
	FPS int

	Controller controller.Controller

	Snapshots []snapshot.Snapshot

	// Private vars
	exit           bool
	debug          bool
	batterySaveDir string

	// Timing
	cyclesPerFrame int
	lastFrameTime  time.Time
}

func NewGameBoyCore(debug bool) (*GameBoyCore, error) {
	return &GameBoyCore{
		debug:          debug,
		FPS:            60,
		cyclesPerFrame: 70224, // 4194304 Hz / 60 FPS = ~70224 cycles per frame
		lastFrameTime:  time.Now(),
	}, nil
}

func (gb *GameBoyCore) Init(cartPath string) error {
	// Initialize core components
	gb.Mmu = mmu.NewMMU()

	// Initialize and read cartridge file
	crt, err := cartridge.NewCartridge(cartPath)
	if err != nil {
		return err
	}
	gb.Cartridge = crt

	// Set the save directory for the cartridge
	if gb.batterySaveDir != "" {
		crt.SetSaveDirectory(gb.batterySaveDir)
	}

	// Load the cartridge rom file from disk
	if err := crt.LoadCartridge(); err != nil {
		return err
	}

	// Set up the cartridge in the MMU
	gb.Mmu.SetCartridge(crt)

	// Initialize CPU with reference to MMU
	gb.Cpu, err = cpu.NewCPU(gb.Mmu)
	if err != nil {
		return err
	}

	// Initialize PPU with reference to MMU
	gb.Ppu = ppu.NewPPU(gb.Mmu)

	// Set the PPU in the MMU for register write handling
	gb.Mmu.SetPPU(gb.Ppu)

	// Initialize Timer with reference to MMU
	gb.Timer = timer.NewTimer(gb.Mmu)

	// Set the timer in the MMU
	gb.Mmu.SetTimer(gb.Timer)

	// Initialize Sound
	gb.Sound = sound.NewSound()

	// Initialize hardware controller
	gb.Controller = controller.NewKeyboard()
	gb.Controller.Init()

	// Set the controller in the MMU
	gb.Mmu.SetController(gb.Controller)

	// Initialize to post-boot state (simulate boot ROM completion)
	gb.Initialize()

	return nil
}

// Run runs the main emulator loop by progressing the CPU tick
func (gb *GameBoyCore) Run() error {
	log.Println("[Core] Starting emulator loop...")

	for {
		// Process one frame
		if err := gb.runFrame(); err != nil {
			return err
		}

		// Process controller input
		if gb.Controller.Update() {
			// Check if a joypad interrupt should be triggered
			if gb.Controller.CheckInterrupt() {
				// Set the joypad interrupt flag (bit 4)
				interruptFlags := gb.Mmu.ReadByte(0xFF0F)
				gb.Mmu.WriteByte(0xFF0F, interruptFlags|0x10)
			}
		}

		// Throttle to target FPS
		gb.throttleFPS()

		if gb.exit {
			log.Println("[Core] Exiting emulator...")
			return nil
		}
	}
}

// runFrame executes one frame of emulation
func (gb *GameBoyCore) runFrame() error {
	cyclesThisFrame := 0

	// Run until we've executed enough cycles for one frame
	for cyclesThisFrame < gb.cyclesPerFrame {
		// Execute one CPU instruction
		cycles := gb.Cpu.Step()

		// Update PPU
		gb.Ppu.Step(cycles)

		// Update Sound
		gb.Sound.Step(cycles)

		// Update Timer
		gb.Timer.Step(cycles)

		cyclesThisFrame += cycles

		// Debug output
		if gb.debug {
			log.Printf("[DEBUG] Executed instruction, cycles: %d", cycles)
		}
	}

	// Debug output for frame
	if gb.debug {
		log.Print("[DEBUG] CPU Frame:")
		gb.Cpu.DisplayCPUFrame()
		log.Print("[DEBUG] CPU Clock:")
		gb.Cpu.DisplayClock()
	}

	return nil
}

// throttleFPS limits the emulation speed to the target FPS
func (gb *GameBoyCore) throttleFPS() {
	// Calculate target frame time
	targetFrameTime := time.Second / time.Duration(gb.FPS)

	// Calculate elapsed time since last frame
	elapsed := time.Since(gb.lastFrameTime)

	// Sleep if we're running too fast
	if elapsed < targetFrameTime {
		time.Sleep(targetFrameTime - elapsed)
	}

	// Update last frame time
	gb.lastFrameTime = time.Now()
}

// Update is the core game loop function
func (gb *GameBoyCore) Update() error {
	// This is now handled by runFrame
	return nil
}

// Exit sets the exit flag to stop the emulator
func (gb *GameBoyCore) Exit() {
	// Save battery RAM if available
	if gb.Cartridge != nil && gb.Cartridge.GetMBC() != nil {
		log.Println("[Core] Saving battery RAM...")
		gb.Cartridge.GetMBC().SaveBatteryRAM()
	}

	// Clean up controller resources
	if gb.Controller != nil {
		gb.Controller.Cleanup()
	}

	gb.exit = true
}

// Step executes one frame of emulation (for display integration)
func (gb *GameBoyCore) Step() error {
	return gb.runFrame()
}

// StepInstruction executes a single CPU instruction (for more granular control)
func (gb *GameBoyCore) StepInstruction() (int, error) {
	// Execute one CPU instruction
	cycles := gb.Cpu.Step()

	// Update PPU
	gb.Ppu.Step(cycles)

	// Update Sound
	gb.Sound.Step(cycles)

	// Update Timer
	gb.Timer.Step(cycles)

	return cycles, nil
}

// Initialize sets up the GameBoy to the post-boot state
func (gb *GameBoyCore) Initialize() {
	// Most importantly, enable the LCD (what boot ROM would do)
	gb.Mmu.WriteByte(0xFF40, 0x91) // LCDC - LCD enabled, BG enabled
	gb.Mmu.WriteByte(0xFF42, 0x00) // SCY - Scroll Y
	gb.Mmu.WriteByte(0xFF43, 0x00) // SCX - Scroll X
	gb.Mmu.WriteByte(0xFF45, 0x00) // LYC - LY Compare
	gb.Mmu.WriteByte(0xFF47, 0xFC) // BGP - BG Palette
	gb.Mmu.WriteByte(0xFF48, 0xFF) // OBP0 - Object Palette 0
	gb.Mmu.WriteByte(0xFF49, 0xFF) // OBP1 - Object Palette 1
	gb.Mmu.WriteByte(0xFF4A, 0x00) // WY - Window Y
	gb.Mmu.WriteByte(0xFF4B, 0x00) // WX - Window X

	// Initialize sound registers
	gb.Mmu.WriteByte(0xFF26, 0xF1) // NR52 - Sound on/off

	// Initialize timer
	gb.Mmu.WriteByte(0xFF07, 0x00) // TAC - Timer control
}
func (gb *GameBoyCore) GetPPUDebugInfo() map[string]interface{} {
	lcdc := gb.Mmu.ReadByte(0xFF40) // LCDC register
	stat := gb.Mmu.ReadByte(0xFF41) // STAT register
	ly := gb.Mmu.ReadByte(0xFF44)   // LY register

	return map[string]interface{}{
		"lcdc_enabled":    (lcdc & 0x80) != 0,
		"bg_enabled":      (lcdc & 0x01) != 0,
		"sprites_enabled": (lcdc & 0x02) != 0,
		"window_enabled":  (lcdc & 0x20) != 0,
		"lcdc_value":      lcdc,
		"stat_value":      stat,
		"ly_value":        ly,
		"ppu_mode":        stat & 0x03,
	}
}

// GetScreenBuffer returns the current screen buffer from the PPU
func (gb *GameBoyCore) GetScreenBuffer() []byte {
	return gb.Ppu.GetScreenBuffer()
}

// IsRunning returns whether the emulator is still running
func (gb *GameBoyCore) IsRunning() bool {
	return !gb.exit
}

// SetButtonState sets the state of a GameBoy button (for input handling)
func (gb *GameBoyCore) SetButtonState(button string, pressed bool) {
	if gb.Controller != nil {
		gb.Controller.SetButtonState(button, pressed)
	}
}

// Takes a snapshot of the current state of the system that can be loaded later
func (gb *GameBoyCore) SaveState() error {
	// TODO: Implement save state functionality
	return nil
}

// SetSaveDirectory sets the directory where battery-backed save files will be stored
func (gb *GameBoyCore) SetSaveDirectory(dir string) {
	gb.batterySaveDir = dir
	log.Printf("[Core] Battery save directory set to: %s", dir)
}
