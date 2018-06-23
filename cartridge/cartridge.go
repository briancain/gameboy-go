package cartridge

import (
	"log"
)

type Cartridge struct {
	title string
}

func OpenFile(filePath string) {
	log.Print("Loading cartridge from file path ", filePath)
}

func NewCartridge(cartPath string) (*Cartridge, error) {
	return nil, nil
}
