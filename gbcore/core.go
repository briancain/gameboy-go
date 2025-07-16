package gbcore

import (
	"log"
	"time"

	controller "github.com/briancain/gameboy-go/controller"
	cart "github.com/briancain/gameboy-go/gbcore/cartridge"
	cpu "github.com/briancain/gameboy-go/gbcore/cpu"
	mmu "github.com/briancain/gameboy-go/gbcore/mmu"
	ppu "github.com/briancain/gameboy-go/gbcore/ppu"
	snapshot "github.com/briancain/gameboy-go/gbcore/snapshot"
	sound "github.com/briancain/gameboy-go/gbcore/sound"
)

// Core emulator implementation
type GameBoyCore struct {
	// Core gameboy components
	Cpu       *cpu.Z80
	Mmu       *mmu.MemoryManagedUnit
	Ppu       *ppu.PPU
	Sound     *sound.Sound
	Cartridge *cart.Cartridge

	// Speed options
	FPS int

	Controller controller.Controller

	Snapshots []snapshot.Snapshot

	// Private vars
	exit  bool
	debug bool
	
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
	crt, err := cart.NewCartridge(cartPath)
	if err != nil {
		return err
	}
	gb.Cartridge = crt

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
		
		// Update timers
		// TODO: Implement timer
		
		cyclesThisFrame += cycles
		
		// Debug output
		if gb.debug {
			log.Printf("[DEBUG] Executed instruction, cycles: %d", cycles)
		}
	}
	
	// Debug output for frame
	if gb.debug {
		display := gb.Cpu.DisplayCPUFrame()
		log.Print("[DEBUG] CPU Frame:\n", display)
		clockdisplay := gb.Cpu.DisplayClock()
		log.Print("[DEBUG] CPU Clock:\n", clockdisplay)
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
	gb.exit = true
}
