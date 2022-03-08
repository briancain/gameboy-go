package gbcore

import "github.com/briancain/gameboy-go/controller/controller"

// Core emulator implementation
type GameBoyCore struct {
	// Core gameboy components
	Cpu   Z80
	Mmu   MemoryManagedUnit
	Sound Sound

	Cartridge Cartridge

	Controller controller.Controller
}

func (gb *GameBoyCore) Init() error {
}
