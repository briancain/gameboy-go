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

	// Initialize Timer with reference to MMU
	gb.Timer = timer.NewTimer(gb.Mmu)

	// Set the timer in the MMU
	gb.Mmu.SetTimer(gb.Timer)

	// Initialize Sound
	gb.Sound = sound.NewSound()

	// Initialize hardware controller
	gb.Controller = controller.NewKeyboard()
	gb.Controller.Init()

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
		gb.Controller.Update()

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

// Takes a snapshot of the current state of the system that can be loaded later
func (gb *GameBoyCore) SaveState() error {
	// TODO: Implement save state functionality
	return nil
}

// Set the exit flag to true to stop the emulator
func (gb *GameBoyCore) Exit() {
	// Save battery RAM if available
	if gb.Cartridge != nil && gb.Cartridge.GetMBC() != nil {
		log.Println("[Core] Saving battery RAM...")
		gb.Cartridge.GetMBC().SaveBatteryRAM()
	}

	gb.exit = true
}

// SetSaveDirectory sets the directory where battery-backed save files will be stored
func (gb *GameBoyCore) SetSaveDirectory(dir string) {
	gb.batterySaveDir = dir
	log.Printf("[Core] Battery save directory set to: %s", dir)
}
