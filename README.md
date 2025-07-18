# GameBoy Go

A GameBoy emulator written in Golang.

This project aims to create a fully functional Game Boy emulator that can run commercial games.
It's designed to be educational and help others learn about Game Boy architecture and emulation techniques.

It's heavily based off the GameBoy CPU manual found online for the Z80.

## Features

- CPU emulation (Z80-like)
- Memory management (MMU)
- Cartridge support with MBC1
- Graphics rendering (PPU)
- Basic sound emulation
- Input handling
- Timer system

## Supported Cartridge Types

- ROM Only
- MBC1
- MBC1+RAM
- MBC1+RAM+BATTERY

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

- `-rom-file`: Path to the Game Boy ROM file (required)
- `-debug`: Enable debug output
- `-scale`: Screen scale factor (1-4, default: 2)
- `-headless`: Run without display (for testing)
- `-help`: Display help information

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

## Resources

The implementation is based on various Game Boy documentation sources:

- Game Boy CPU Manual
- Pan Docs
- Game Boy Programming Manual
- Various online resources about Game Boy architecture
