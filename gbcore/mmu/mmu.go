package gbcore

import (
	"log"

	gbcore "github.com/briancain/gameboy-go/gbcore/cartridge"
)

type MemoryManagedUnit interface {
	Reset()
}

type MMU struct {
	cart gbcore.Cartridge // 0x0000-0x3FFF
	bios [256]byte        // 0x0000-0x00FF
}

// Read a byte from memory
func ReadBytes(addr byte) error {
	return nil
}

// Read a 16-bit word
func ReadWord(addr byte) error {
	return nil
}

func (m *MMU) LoadCart() error {
	return nil
}

func (m *MMU) Reset() {
	log.Print("Resetting MMU")
}
