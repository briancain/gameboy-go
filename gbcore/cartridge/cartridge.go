package gbcore

import (
	"errors"
	"fmt"
	"log"
	"os"
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

func NewCartridge(cartPath string) (*Cartridge, error) {
	log.Println("[DEBUG] loading cart from path ", cartPath)

	if _, err := os.Stat(cartPath); err != nil {
		return nil, errors.New(
			fmt.Sprintf("The ROM file at %q does not exist on the file system", cartPath))
	}
	// we'll load cart title directly from the ROM data on init
	return &Cartridge{title: "", filePath: cartPath}, nil
}
