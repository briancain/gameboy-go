package gbcore

// Core emulator implementation
type GameBoyCore struct {
	// Core gameboy components
	Cpu Z80
	Mmu MemoryManagedUnit

	Cartridge Cartridge
}

func (gb *GameBoyCore) Init() error {
}
