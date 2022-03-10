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

	// the type of cartridge it is
	cartType byte

	// Memory Controller Bank implementation
	MBC MBC

	// TODO move rom and ram information into MBC

	// raw byte stream of ROM data
	rom []byte

	// optional, only if ROM has additional RAM support
	ramSize byte
}

// Reference https://gbdev.io/pandocs/The_Cartridge_Header.html
func (c *Cartridge) LoadCartridge() error {
	log.Println("[DEBUG] Loading cart from path:", c.filePath)
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
	// TODO(briancain): a function to read the title and cart type from the byte stream

	// Cartridge title is always located at 0x134-0x143 and is in all caps
	log.Println("[Cartridge] Game title:", string(c.rom[0x134:0x143]))
	c.title = string(c.rom[0x134:0x143])

	// Cartridge type defines the kind of cartridge we're loading, aka a cameboy camera cart
	// or a cart with a battery, etc. It also defines the kind of memory controller
	// we're working with.
	c.cartType = c.rom[0x147]
	if ct, ok := cartridgeTypeMap[c.cartType]; !ok {
		return errors.New("[ERROR] Unsupported cartridge type loaded from " + c.title)
	} else {
		log.Println("[Cartridge] Cartridge type discovered:", ct)
		// set MBC based on found string
	}

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

/*
* Memory Bank Controller implementations .... lol there's a lot to do here
*
* Realistically most games are either MBC1, MBC3, or MBC5 so lets start with these.
 */

// A generic Memory Bank Controller interface.
type MBC interface {
	ReadRom(uint16) (byte, error)

	ReadRam(uint16) (byte, error)
	WriteRam(uint16, byte) error
}

// https://gbdev.io/pandocs/The_Cartridge_Header.html#0147---cartridge-type
var cartridgeTypeMap = map[byte]string{
	byte(0x00): "ROM ONLY",
	byte(0x01): "MBC1",
	byte(0x02): "MBC1+RAM",
	byte(0x03): "MBC1+RAM+BATTERY",
	byte(0x05): "MBC2",
	byte(0x06): "MBC2+BATTERY",
	byte(0x08): "ROM+RAM",
	byte(0x09): "ROM+RAM+BATTERY",
	byte(0x0B): "MMM01",
	byte(0x0C): "MMM01+RAM",
	byte(0x0D): "MMM01+RAM+BATTERY",
	byte(0x0F): "MBC3+TIMER+BATTERY",
	byte(0x10): "MBC3+TIMER+RAM+BATTERY",
	byte(0x11): "MBC3",
	byte(0x12): "MBC3+RAM",
	byte(0x13): "MBC3+RAM+BATTERY",
	byte(0x15): "MBC4",
	byte(0x16): "MBC4+RAM",
	byte(0x17): "MBC4+RAM+BATTERY",
	byte(0x19): "MBC5",
	byte(0x1A): "MBC5+RAM",
	byte(0x1B): "MBC5+RAM+BATTERY",
	byte(0x1C): "MBC5+RUMBLE",
	byte(0x1D): "MBC5+RUMBLE+RAM",
	byte(0x1E): "MBC5+RUMBLE+RAM+BATTERY",
	byte(0xFC): "POCKET CAMERA",
	byte(0xFD): "BANDAI TAMA5",
	byte(0xFE): "HuC3",
	byte(0xFF): "HuC1+RAM+BATTERY",
}
