package gbcore

import (
	"log"

	controller "github.com/briancain/gameboy-go/controller"
	cart "github.com/briancain/gameboy-go/gbcore/cartridge"
	cpu "github.com/briancain/gameboy-go/gbcore/cpu"
	mmu "github.com/briancain/gameboy-go/gbcore/mmu"
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

	// Private vars?

	// If set to true, will exit on the next frame
	exit bool
}

func NewGameBoyCore() (*GameBoyCore, error) {
	return &GameBoyCore{}, nil
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

	// Initialize hardware controller
	//gb.Controller.Init()

	return nil
}

// Run runs the main emulator loop by progressing the CPU tick
func (gb *GameBoyCore) Run() error {
	for {
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
	return nil
}
