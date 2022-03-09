package gbcore

import (
	"log"

	controller "github.com/briancain/gameboy-go/controller"
	cart "github.com/briancain/gameboy-go/gbcore/cartridge"
	cpu "github.com/briancain/gameboy-go/gbcore/cpu"
	mmu "github.com/briancain/gameboy-go/gbcore/mmu"
	snapshot "github.com/briancain/gameboy-go/gbcore/snapshot"
	sound "github.com/briancain/gameboy-go/gbcore/sound"
)

// Core emulator implementation
type GameBoyCore struct {
	// Core gameboy components
	Cpu   cpu.Z80
	Mmu   mmu.MemoryManagedUnit
	Sound sound.Sound

	// Speed options
	FPS int

	Cartridge *cart.Cartridge

	Controller controller.Controller

	Snapshots []snapshot.Snapshot

	// Private vars?

	// If set to true, will exit on the next frame
	exit  bool
	debug bool
}

func NewGameBoyCore(debug bool) (*GameBoyCore, error) {
	return &GameBoyCore{debug: debug}, nil
}

func (gb *GameBoyCore) Init(cartPath string) error {
	// Initialize core components

	// Initialize and read cartridge file
	crt, err := cart.NewCartridge(cartPath)
	if err != nil {
		return err
	} else {
		gb.Cartridge = crt
	}

	// Load the cartridge rom file from disk
	if err := crt.LoadCartridge(); err != nil {
		return err
	}

	// Initialize hardware controller
	gb.Controller = controller.Keyboard{Name: "Keyboard"}
	gb.Controller.Init()

	return nil
}

// Run runs the main emulator loop by progressing the CPU tick
func (gb *GameBoyCore) Run() error {
	for {
		// debug
		if gb.debug {
			display := gb.Cpu.DisplayCPUFrame()
			log.Print("[DEBUG] CPU Frame:\n", display)
			clockdisplay := gb.Cpu.DisplayClock()
			log.Print("[DEBUG] CPU Clock:\n", clockdisplay)
		}

		// Update CPU Tick Frame
		gb.Update()

		// Process controller input
		gb.Controller.Update()

		if gb.exit {
			log.Println("[Core] Exiting emulator ...")
			// shut down any services here
			return nil
		}
	}
}

// Update is the core game loop function
func (gb *GameBoyCore) Update() error {
	gb.exit = true
	return nil
}

// Takes a snapshot of the current state of the sytem that can be loaded later
func (gb *GameBoyCore) SaveState() error {
	return nil
}
