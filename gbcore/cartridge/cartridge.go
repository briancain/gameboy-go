package gbcore

import (
	"bufio"
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

func (c *Cartridge) LoadCartridge() error {
	log.Println("[DEBUG] loading cart from path ", c.filePath)
	// Load file on path and read bytes into memory

	romFile, err := os.Open(c.filePath)
	if err != nil {
		return err
	}
	defer romFile.Close()

	stats, err := romFile.Stat()
	if err != nil {
		return err
	}

	size := stats.Size()
	bytes := make([]byte, size)

	bufReader := bufio.NewReader(romFile)
	_, err = bufReader.Read(bytes)
	if err != nil {
		return err
	}

	c.rom = bytes

	log.Println("[DEBUG] Loaded rom file of size", size, "bytes.")
	// TODO(briancain): a function to read the title and cart type from the byte stream?

	return nil
}

func NewCartridge(cartPath string) (*Cartridge, error) {
	if _, err := os.Stat(cartPath); err != nil {
		return nil, errors.New(
			fmt.Sprintf("The ROM file at %q does not exist on the file system", cartPath))
	}

	// we'll load cart title directly from the ROM data on init
	return &Cartridge{title: "", filePath: cartPath}, nil
}
