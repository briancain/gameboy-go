package cartridge

import (
	"log"
)

type Cartridge struct {
	title    string
	filePath string
}

func (c *Cartridge) OpenFile() {
	log.Print("Loading cartridge from file path ", c.filePath)
}

func NewCartridge(cartPath string) (*Cartridge, error) {
	return &Cartridge{title: "", filePath: cartPath}, nil
}
