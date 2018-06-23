package mmu

import (
	"log"
)

type MemoryManagedUnit interface {
	Reset()
}

type MMU struct {
	// cart Cartridge
	bios [256]byte // 0x0000-0x00FF
}

// Read a byte from memory
func ReadBytes(addr byte) {

}

// Read a 16-bit word
func ReadWord(addr byte) {

}

func Reset(m *MMU) {
	log.Print("Resetting MMU")
}
