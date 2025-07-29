# GameBoy Go

A GameBoy emulator written in Golang.

This project aims to create a fully functional GameBoy emulator that can run commercial games.
It's designed to be educational and help others learn about GameBoy architecture and emulation techniques.

It's heavily based off the GameBoy CPU manual found online for the Z80.

## Current Status

The emulator is under active development with the following components implemented:

- ✅ **CPU emulation** - Z80-like with all instructions and proper timing
- ✅ **Memory management** - MMU with proper memory mapping
- ✅ **Cartridge support** - MBC1-5 with battery saves, RTC, and rumble
- ✅ **Timer system** - Complete with proper interrupt handling
- ✅ **Interrupt handling** - V-Blank, STAT, timer, joypad, and serial interrupts
- ✅ **Input handling** - Joypad with interrupt support
- ✅ **Picture Processing Unit (PPU)** - Complete graphics rendering with:
  - Core PPU timing and modes
  - Background and window rendering with scrolling
  - Sprite rendering (8x8 and 8x16) with hardware-accurate priority
  - Proper STAT and V-Blank interrupt generation
- ✅ **Visual output** - Real-time display with Ebiten graphics engine

## Ready to Implement (PPU)

The following PPU features are ready to be implemented with existing infrastructure:

- 📝 **Graphics Output Integration**: Connect to a graphics library like ebiten for visual display
- 📝 **Color Palette Customization**: Add support for custom color palettes beyond monochrome
- 📝 **Advanced Rendering Features**: Sprite-to-background priority, transparent colors

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

### Setting Up Test ROMs

To properly test the emulator, you'll need to download test ROMs. These are not included in the repository due to size and licensing considerations.

```bash
# From the project root directory
mkdir -p test/data
cd test/data
git clone https://github.com/retrio/gb-test-roms.git
cd ../..

# Test your emulator with Blargg's CPU instruction test
./bin/gameboy-go -rom-file test/data/gb-test-roms/cpu_instrs/cpu_instrs.gb -debug

# Test PPU rendering with the acid2 visual test
./bin/gameboy-go -rom-file test/data/gb-test-roms/acid2.gb -scale 3
```

**Recommended Test ROMs:**
- `cpu_instrs/cpu_instrs.gb` - Comprehensive CPU instruction test
- `instr_timing/instr_timing.gb` - Instruction timing validation
- `acid2.gb` - PPU rendering accuracy test
- `dmg_sound/dmg_sound.gb` - Sound system test (for future implementation)

These test ROMs are homebrew software specifically designed for emulator testing and are freely available.

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
