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

## In Progress

- 🔄 Picture Processing Unit (PPU)
- 🔄 Input handling

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

- `gbcore/`: Core emulator components
  - `cpu/`: CPU implementation
  - `mmu/`: Memory management unit
  - `ppu/`: Picture processing unit (graphics)
  - `cartridge/`: Cartridge and MBC implementations
  - `sound/`: Sound system
  - `timer/`: Timer implementation
- `controller/`: Input handling
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
