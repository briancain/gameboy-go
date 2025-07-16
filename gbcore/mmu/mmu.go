package gbcore

import (
	"log"
)

// Memory map:
// 0000-3FFF: ROM Bank 0
// 4000-7FFF: ROM Bank 1-n
// 8000-9FFF: Video RAM (VRAM)
// A000-BFFF: External RAM
// C000-CFFF: Work RAM (WRAM) Bank 0
// D000-DFFF: Work RAM (WRAM) Bank 1-n
// E000-FDFF: Echo RAM (mirror of C000-DDFF)
// FE00-FE9F: Sprite attribute table (OAM)
// FEA0-FEFF: Not Usable
// FF00-FF7F: I/O Registers
// FF80-FFFE: High RAM (HRAM)
// FFFF     : Interrupt Enable Register (IE)

type MemoryManagedUnit struct {
	// Memory regions
	bios      [0x100]byte // 0x0000-0x00FF (only during boot)
	rom       []byte      // Cartridge ROM
	vram      [0x2000]byte // 0x8000-0x9FFF
	eram      [0x2000]byte // 0xA000-0xBFFF (Cartridge RAM)
	wram      [0x2000]byte // 0xC000-0xDFFF (Work RAM)
	oam       [0x100]byte  // 0xFE00-0xFE9F (Sprite attribute table)
	io        [0x80]byte   // 0xFF00-0xFF7F (I/O registers)
	hram      [0x7F]byte   // 0xFF80-0xFFFE (High RAM)
	ie        byte         // 0xFFFF (Interrupt Enable register)

	// Control flags
	biosActive bool // Whether BIOS is active

	// References to other components
	cartridge Cartridge
}

// Cartridge interface for memory banking
type Cartridge interface {
	ReadByte(addr uint16) byte
	WriteByte(addr uint16, value byte)
}

// Initialize a new MMU
func NewMMU() *MemoryManagedUnit {
	mmu := &MemoryManagedUnit{
		biosActive: true,
	}
	return mmu
}

// Reset the MMU to initial state
func (m *MemoryManagedUnit) Reset() {
	log.Print("Resetting MMU")
	
	// Clear all memory regions
	for i := range m.vram {
		m.vram[i] = 0
	}
	for i := range m.eram {
		m.eram[i] = 0
	}
	for i := range m.wram {
		m.wram[i] = 0
	}
	for i := range m.oam {
		m.oam[i] = 0
	}
	for i := range m.io {
		m.io[i] = 0
	}
	for i := range m.hram {
		m.hram[i] = 0
	}
	m.ie = 0
	
	// Initialize I/O registers to their default values
	m.io[0x05] = 0x00 // TIMA
	m.io[0x06] = 0x00 // TMA
	m.io[0x07] = 0x00 // TAC
	m.io[0x10] = 0x80 // NR10
	m.io[0x11] = 0xBF // NR11
	m.io[0x12] = 0xF3 // NR12
	m.io[0x14] = 0xBF // NR14
	m.io[0x16] = 0x3F // NR21
	m.io[0x17] = 0x00 // NR22
	m.io[0x19] = 0xBF // NR24
	m.io[0x1A] = 0x7F // NR30
	m.io[0x1B] = 0xFF // NR31
	m.io[0x1C] = 0x9F // NR32
	m.io[0x1E] = 0xBF // NR33
	m.io[0x20] = 0xFF // NR41
	m.io[0x21] = 0x00 // NR42
	m.io[0x22] = 0x00 // NR43
	m.io[0x23] = 0xBF // NR44
	m.io[0x24] = 0x77 // NR50
	m.io[0x25] = 0xF3 // NR51
	m.io[0x26] = 0xF1 // NR52
	m.io[0x40] = 0x91 // LCDC
	m.io[0x42] = 0x00 // SCY
	m.io[0x43] = 0x00 // SCX
	m.io[0x45] = 0x00 // LYC
	m.io[0x47] = 0xFC // BGP
	m.io[0x48] = 0xFF // OBP0
	m.io[0x49] = 0xFF // OBP1
	m.io[0x4A] = 0x00 // WY
	m.io[0x4B] = 0x00 // WX
	m.ie = 0x00 // IE
}

// Set the cartridge
func (m *MemoryManagedUnit) SetCartridge(cart Cartridge) {
	m.cartridge = cart
}

// Load BIOS
func (m *MemoryManagedUnit) LoadBIOS(data []byte) error {
	if len(data) > len(m.bios) {
		return nil // Error: BIOS data too large
	}
	
	copy(m.bios[:], data)
	m.biosActive = true
	return nil
}

// Disable BIOS
func (m *MemoryManagedUnit) DisableBIOS() {
	m.biosActive = false
}

// Read a byte from memory
func (m *MemoryManagedUnit) ReadByte(addr uint16) byte {
	switch {
	case addr < 0x100 && m.biosActive:
		// BIOS (if active)
		return m.bios[addr]
	case addr < 0x8000:
		// ROM banks
		return m.cartridge.ReadByte(addr)
	case addr < 0xA000:
		// VRAM
		return m.vram[addr-0x8000]
	case addr < 0xC000:
		// External RAM (in cartridge)
		return m.cartridge.ReadByte(addr)
	case addr < 0xE000:
		// Work RAM
		return m.wram[addr-0xC000]
	case addr < 0xFE00:
		// Echo RAM (mirror of C000-DDFF)
		return m.wram[addr-0xE000]
	case addr < 0xFEA0:
		// OAM
		return m.oam[addr-0xFE00]
	case addr < 0xFF00:
		// Not usable
		return 0xFF
	case addr < 0xFF80:
		// I/O registers
		return m.readIO(addr)
	case addr < 0xFFFF:
		// High RAM
		return m.hram[addr-0xFF80]
	default:
		// Interrupt Enable register
		return m.ie
	}
}

// Write a byte to memory
func (m *MemoryManagedUnit) WriteByte(addr uint16, value byte) {
	switch {
	case addr < 0x8000:
		// ROM banks - handled by cartridge
		m.cartridge.WriteByte(addr, value)
	case addr < 0xA000:
		// VRAM
		m.vram[addr-0x8000] = value
	case addr < 0xC000:
		// External RAM (in cartridge)
		m.cartridge.WriteByte(addr, value)
	case addr < 0xE000:
		// Work RAM
		m.wram[addr-0xC000] = value
	case addr < 0xFE00:
		// Echo RAM (mirror of C000-DDFF)
		m.wram[addr-0xE000] = value
	case addr < 0xFEA0:
		// OAM
		m.oam[addr-0xFE00] = value
	case addr < 0xFF00:
		// Not usable
		// Do nothing
	case addr < 0xFF80:
		// I/O registers
		m.writeIO(addr, value)
	case addr < 0xFFFF:
		// High RAM
		m.hram[addr-0xFF80] = value
	default:
		// Interrupt Enable register
		m.ie = value
	}
}

// Read a 16-bit word
func (m *MemoryManagedUnit) ReadWord(addr uint16) uint16 {
	low := uint16(m.ReadByte(addr))
	high := uint16(m.ReadByte(addr + 1))
	return (high << 8) | low
}

// Write a 16-bit word
func (m *MemoryManagedUnit) WriteWord(addr uint16, value uint16) {
	m.WriteByte(addr, byte(value&0xFF))
	m.WriteByte(addr+1, byte(value>>8))
}

// Special handling for I/O register reads
func (m *MemoryManagedUnit) readIO(addr uint16) byte {
	// Handle special I/O registers
	switch addr {
	case 0xFF00: // Joypad
		// TODO: Implement joypad reading
		return 0xFF
	default:
		return m.io[addr-0xFF00]
	}
}

// Special handling for I/O register writes
func (m *MemoryManagedUnit) writeIO(addr uint16, value byte) {
	// Handle special I/O registers
	switch addr {
	case 0xFF00: // Joypad
		// TODO: Implement joypad writing
		m.io[0] = value
	case 0xFF04: // DIV - Divider register
		// Writing any value resets DIV to 0
		m.io[0x04] = 0
	case 0xFF46: // DMA - OAM DMA transfer
		m.doDMATransfer(value)
	default:
		m.io[addr-0xFF00] = value
	}
}

// Perform DMA transfer from ROM/RAM to OAM
func (m *MemoryManagedUnit) doDMATransfer(value byte) {
	// DMA transfers 160 bytes from XX00-XX9F to FE00-FE9F
	// where XX is the value written to FF46
	baseAddr := uint16(value) << 8
	for i := uint16(0); i < 160; i++ {
		m.oam[i] = m.ReadByte(baseAddr + i)
	}
	m.io[0x46] = value
}
