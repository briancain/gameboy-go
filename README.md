# GameBoy Go

A GameBoy emulator written in Golang.

This project aims to create a fully functional GameBoy emulator that can run commercial games.
It's designed to be educational and help others learn about GameBoy architecture and emulation techniques.

It's heavily based off the GameBoy CPU manual found online for the Z80.

## Current Status

The emulator is under active development with the following components implemented:

- ✅ CPU emulation (Z80-like) with all instructions and proper timing
- ✅ Basic memory management (MMU)
- ✅ Cartridge loading
- ✅ Timer system
- ⚠️ Partial interrupt handling
- ✅ Memory Bank Controllers (MBC1, MBC2, MBC3, MBC5)
  - ✅ Battery-backed save support
  - ✅ Real-Time Clock (RTC) for MBC3
  - ✅ Rumble support for MBC5
- ✅ Input handling with joypad interrupts
- ✅ Picture Processing Unit (PPU)
  - ✅ Core PPU timing and modes (OAM, VRAM, HBLANK, VBLANK)
  - ✅ Background rendering with scrolling (SCX, SCY)
  - ✅ Window rendering with proper edge case handling
  - ✅ Sprite rendering (8x8 and 8x16 modes)
  - ✅ PPU register write handlers with real-time updates
  - ✅ Palette support (BGP, OBP0, OBP1)
  - ✅ LCDC control (LCD on/off, layer enables)
  - ✅ STAT register with interrupt flags
  - ✅ LY/LYC coincidence detection
  - ✅ Proper STAT and V-Blank interrupt generation
  - ✅ Hardware-accurate sprite priority handling

## Ready to Implement (PPU)

The following PPU features are ready to be implemented with existing infrastructure:

- 📝 **Sprite Priority Handling**: Improve sprite-to-sprite priority when X coordinates are the same (15-20 min)
- 📝 **Graphics Output Integration**: Connect to a graphics library like ebiten for visual display (30-45 min)
- 📝 **Color Palette Customization**: Add support for custom color palettes beyond monochrome (15-20 min)
- 📝 **PPU Interrupt Generation**: Generate STAT and VBLANK interrupts properly (20-30 min)
- 📝 **Advanced Rendering Features**: Sprite-to-background priority, transparent colors (30-45 min)

## Planned Features

- 📝 Sound emulation (4 channels)
- 📝 Serial I/O
- 📝 Save states
- 📝 Debugging tools

## Supported Cartridge Types

Currently implemented:
- ROM Only
- MBC1, MBC1+RAM, MBC1+RAM+BATTERY
- MBC2, MBC2+BATTERY
- MBC3, MBC3+RAM, MBC3+RAM+BATTERY, MBC3+TIMER+BATTERY, MBC3+TIMER+RAM+BATTERY
- MBC5, MBC5+RAM, MBC5+RAM+BATTERY, MBC5+RUMBLE, MBC5+RUMBLE+RAM, MBC5+RUMBLE+RAM+BATTERY

## How to Build

```
make build
```

This will create the executable in the `bin/` directory.

## How to Run

```
./bin/gameboy-go -rom-file PATH_TO_ROM
```

### Command Line Options

- `-battery-save-dir` Directory to store battery-backed save files from cartridges (e.g., game progress)
- `-debug`: Enable debug output
- `-headless`: Run without display (for testing)
- `-help`: Display help information
- `-rom-file`: Path to the GameBoy ROM file (required)
- `-scale`: Screen scale factor (1-4, default: 2)

## Controls

- Arrow keys: D-pad
- Z: A button
- X: B button
- Enter: Start button
- Space: Select button

## Project Structure

- `cmd/gameboy-go/`: Main application entry point
- `internal/`: Core emulator components (private)
  - `cartridge/`: Cartridge and MBC implementations
  - `controller/`: Input handling
  - `core/`: Core emulator functionality
  - `cpu/`: CPU implementation
  - `mmu/`: Memory management unit
  - `ppu/`: Picture processing unit (graphics)
  - `snapshot/`: Save state functionality
  - `sound/`: Sound system
  - `timer/`: Timer implementation
- `docs/`: Documentation

## Development

### Running Tests

```
make test
```

## Contributing

Contributions are welcome! The emulator is still in development, and there are many features that need to be implemented. Check the "In Progress" and "Planned Features" sections for areas that need work.

## Resources

The implementation is based on various GameBoy documentation sources:

- GameBoy CPU Manual
- Pan Docs
- GameBoy Programming Manual
- Various online resources about GameBoy architecture
