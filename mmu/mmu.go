package mmu

import (
	"github.com/briancain/gameboy-go/cartridge"
	"log"
)

type MemoryManagedUnit interface {
	Reset()
}

type MMU struct {
	cart cartridge.Cartridge
	bios [256]byte // 0x0000-0x00FF
}

// Read a byte from memory
func ReadBytes(addr byte) {

}

// Read a 16-bit word
func ReadWord(addr byte) {

}

func LoadCart() {
}

func Reset(m *MMU) {
	log.Print("Resetting MMU")
}
