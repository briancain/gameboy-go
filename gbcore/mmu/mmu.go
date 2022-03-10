package gbcore

import (
	"log"
)

type MemoryManagedUnit interface {
	Reset()
}

type MMU struct {
	MemoryMap []byte
	bios      [256]byte // 0x0000-0x00FF
}

// Read a byte from memory
func (m *MMU) ReadBytes(addr byte) (byte, error) {
	return m.MemoryMap[addr], nil
}

// Read a 16-bit word
func ReadWord(addr byte) (byte, error) {
	return 0, nil
}

func (m *MMU) LoadCart() error {
	return nil
}

func (m *MMU) Reset() {
	log.Print("Resetting MMU")
}
