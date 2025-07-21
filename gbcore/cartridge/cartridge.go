package gbcore

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// Cartridge types
const (
	CART_ROM_ONLY            = 0x00
	CART_MBC1                = 0x01
	CART_MBC1_RAM            = 0x02
	CART_MBC1_RAM_BAT        = 0x03
	CART_MBC2                = 0x05
	CART_MBC2_BAT            = 0x06
	CART_ROM_RAM             = 0x08
	CART_ROM_RAM_BAT         = 0x09
	CART_MMM01               = 0x0B
	CART_MMM01_RAM           = 0x0C
	CART_MMM01_RAM_BAT       = 0x0D
	CART_MBC3_TIMER_BAT      = 0x0F
	CART_MBC3_TIMER_RAM_BAT  = 0x10
	CART_MBC3                = 0x11
	CART_MBC3_RAM            = 0x12
	CART_MBC3_RAM_BAT        = 0x13
	CART_MBC5                = 0x19
	CART_MBC5_RAM            = 0x1A
	CART_MBC5_RAM_BAT        = 0x1B
	CART_MBC5_RUMBLE         = 0x1C
	CART_MBC5_RUMBLE_RAM     = 0x1D
	CART_MBC5_RUMBLE_RAM_BAT = 0x1E
)

// RAM sizes
const (
	RAM_NONE  = 0x00
	RAM_2KB   = 0x01
	RAM_8KB   = 0x02
	RAM_32KB  = 0x03
	RAM_128KB = 0x04
	RAM_64KB  = 0x05
)

type Cartridge struct {
	title    string
	filePath string

	// the type of cartridge it is
	cartType byte

	// Memory Controller Bank implementation
	mbc MBC

	// raw byte stream of ROM data
	rom []byte

	// Directory to store battery-backed save files
	batterySaveDir string

	// RAM size code
	ramSize byte

	// ROM size code
	romSize byte
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

	// Cartridge title is always located at 0x134-0x143 and is in all caps
	c.title = string(c.rom[0x134:0x143])
	log.Println("[Cartridge] Game title:", c.title)

	// Cartridge type defines the kind of cartridge we're loading
	c.cartType = c.rom[0x147]
	if ct, ok := cartridgeTypeMap[c.cartType]; !ok {
		return errors.New("[ERROR] Unsupported cartridge type loaded from " + c.title)
	} else {
		log.Println("[Cartridge] Cartridge type:", ct)
	}

	// ROM size
	c.romSize = c.rom[0x148]
	romSizeBytes := getROMSize(c.romSize)
	log.Println("[Cartridge] ROM size:", romSizeBytes/1024, "KB")

	// RAM size
	c.ramSize = c.rom[0x149]
	ramSizeBytes := getRAMSize(c.ramSize)
	log.Println("[Cartridge] RAM size:", ramSizeBytes/1024, "KB")

	// Initialize the appropriate MBC
	if err := c.initMBC(); err != nil {
		return err
	}

	return nil
}

// Initialize the appropriate Memory Bank Controller
func (c *Cartridge) initMBC() error {
	ramSizeBytes := getRAMSize(c.ramSize)

	switch c.cartType {
	case CART_ROM_ONLY:
		c.mbc = &ROMOnly{rom: c.rom, ram: make([]byte, ramSizeBytes)}

	case CART_MBC1, CART_MBC1_RAM, CART_MBC1_RAM_BAT:
		c.mbc = NewMBC1(c.rom, int(ramSizeBytes), c.cartType, c.title, c.batterySaveDir)

	case CART_MBC2, CART_MBC2_BAT:
		c.mbc = NewMBC2(c.rom, c.cartType, c.title, c.batterySaveDir)

	case CART_MBC3, CART_MBC3_RAM, CART_MBC3_RAM_BAT, CART_MBC3_TIMER_BAT, CART_MBC3_TIMER_RAM_BAT:
		c.mbc = NewMBC3(c.rom, int(ramSizeBytes), c.cartType, c.title, c.batterySaveDir)

	case CART_MBC5, CART_MBC5_RAM, CART_MBC5_RAM_BAT, CART_MBC5_RUMBLE, CART_MBC5_RUMBLE_RAM, CART_MBC5_RUMBLE_RAM_BAT:
		c.mbc = NewMBC5(c.rom, int(ramSizeBytes), c.cartType, c.title, c.batterySaveDir)

	// Add more MBC types as needed

	default:
		return fmt.Errorf("unsupported cartridge type: %02X", c.cartType)
	}

	return nil
}

// Get ROM size in bytes
func getROMSize(romSize byte) uint32 {
	switch romSize {
	case 0x00:
		return 32 * 1024 // 32KB (2 banks)
	case 0x01:
		return 64 * 1024 // 64KB (4 banks)
	case 0x02:
		return 128 * 1024 // 128KB (8 banks)
	case 0x03:
		return 256 * 1024 // 256KB (16 banks)
	case 0x04:
		return 512 * 1024 // 512KB (32 banks)
	case 0x05:
		return 1024 * 1024 // 1MB (64 banks)
	case 0x06:
		return 2048 * 1024 // 2MB (128 banks)
	case 0x07:
		return 4096 * 1024 // 4MB (256 banks)
	case 0x08:
		return 8192 * 1024 // 8MB (512 banks)
	default:
		return 32 * 1024 // Default to 32KB
	}
}

// Get RAM size in bytes
func getRAMSize(ramSize byte) uint32 {
	switch ramSize {
	case RAM_NONE:
		return 0
	case RAM_2KB:
		return 2 * 1024
	case RAM_8KB:
		return 8 * 1024
	case RAM_32KB:
		return 32 * 1024
	case RAM_128KB:
		return 128 * 1024
	case RAM_64KB:
		return 64 * 1024
	default:
		return 0
	}
}

func NewCartridge(cartPath string) (*Cartridge, error) {
	if _, err := os.Stat(cartPath); err != nil {
		return nil, errors.New(
			fmt.Sprintf("The ROM file at %q does not exist on the file system", cartPath))
	}

	// we'll load cart title directly from the ROM data on init
	return &Cartridge{title: "", filePath: cartPath, batterySaveDir: "."}, nil
}

// SetSaveDirectory sets the directory where battery-backed save files will be stored
func (c *Cartridge) SetSaveDirectory(dir string) {
	c.batterySaveDir = dir
	log.Printf("[Cartridge] Battery save directory set to: %s", dir)
}

// Read a byte from the cartridge
func (c *Cartridge) ReadByte(addr uint16) byte {
	return c.mbc.ReadByte(addr)
}

// Write a byte to the cartridge
func (c *Cartridge) WriteByte(addr uint16, value byte) {
	c.mbc.WriteByte(addr, value)
}

// GetMBC returns the Memory Bank Controller for this cartridge
func (c *Cartridge) GetMBC() MBC {
	return c.mbc
}

// A generic Memory Bank Controller interface.
type MBC interface {
	ReadByte(addr uint16) byte
	WriteByte(addr uint16, value byte)
	SaveBatteryRAM()  // Save battery-backed RAM to file (if supported)
	IsRumbling() bool // Returns true if the cartridge has rumble and it's active
}

// ROM Only (no MBC) implementation
type ROMOnly struct {
	rom []byte
	ram []byte
}

func (r *ROMOnly) ReadByte(addr uint16) byte {
	switch {
	case addr < 0x8000:
		// ROM
		if int(addr) < len(r.rom) {
			return r.rom[addr]
		}
		return 0xFF

	case addr >= 0xA000 && addr < 0xC000:
		// RAM (if present)
		if len(r.ram) > 0 {
			ramAddr := addr - 0xA000
			if int(ramAddr) < len(r.ram) {
				return r.ram[ramAddr]
			}
		}
		return 0xFF

	default:
		return 0xFF
	}
}

func (r *ROMOnly) WriteByte(addr uint16, value byte) {
	// Only RAM is writable
	if addr >= 0xA000 && addr < 0xC000 && len(r.ram) > 0 {
		ramAddr := addr - 0xA000
		if int(ramAddr) < len(r.ram) {
			r.ram[ramAddr] = value
		}
	}
}

// SaveBatteryRAM does nothing for ROM-only cartridges
func (r *ROMOnly) SaveBatteryRAM() {
	// ROM-only cartridges don't have battery-backed RAM
}

// IsRumbling always returns false for ROM-only cartridges
func (r *ROMOnly) IsRumbling() bool {
	return false
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
