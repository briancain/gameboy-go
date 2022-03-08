package gbcore

import (
	"log"
)

type Cartridge struct {
	title    string
	filePath string

	// raw byte stream of ROM data
	rom []byte

	// the type of cartridge it is
	cartType CartridgeType
}

type CartridgeType int

// TODO(briancain): Extract existing defined byte strings for each
// cart type and associate it with internal CartridgeType const
var cartTypeMap = map[byte]string{}

const (
	RomOnly    CartridgeType = iota // 32kb ROM
	MBC1                            // Memory Bank Controller 1
	MBC2                            // Memory Bank Controller 2
	MBC3                            // Memory Bank Controller 3
	MBC5                            // Memory Bank Controller 5
	RumbleCart                      // Uses a MBC5 memory bank controller
	HuC1                            // Infrared LED input/output
)

func (c *Cartridge) LoadCartridge() {
	log.Print("Loading cartridge from file path ", c.filePath)
}

func NewCartridge(title string, cartPath string) (*Cartridge, error) {
	return &Cartridge{title: title, filePath: cartPath}, nil
}
