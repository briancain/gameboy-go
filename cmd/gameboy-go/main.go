package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/briancain/gameboy-go/internal/core"
	"github.com/briancain/gameboy-go/internal/display"
	"github.com/briancain/gameboy-go/version"
)

var (
	CartridgePath  string
	Help           bool
	DebugOutput    bool
	Scale          int
	Headless       bool
	BatterySaveDir string
)

func init() {
	flag.BoolVar(&Help, "help", false, "Displays help")
	flag.StringVar(&CartridgePath, "rom-file", "", "A path to a cartridge ROM file")
	flag.BoolVar(&DebugOutput, "debug", false, "Displays debug output")
	flag.IntVar(&Scale, "scale", 2, "Screen scale factor (1-4)")
	flag.BoolVar(&Headless, "headless", false, "Run without display (for testing)")
	// Default to current directory for save files
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}
	flag.StringVar(&BatterySaveDir, "battery-save-dir", currentDir, "Directory to store battery-backed save files from cartridges (e.g., game progress)")
}

func startEmulator() error {
	// Create and initialize the Game Boy core
	gb, err := core.NewGameBoyCore(DebugOutput)
	if err != nil {
		log.Print("Failed to create new core: ", err)
		return err
	}

	// Set the save directory
	gb.SetSaveDirectory(BatterySaveDir)

	if err := gb.Init(CartridgePath); err != nil {
		log.Print("[ERROR] Failed to initialize new core!\n", err)
		return err
	}

	// Check if running in headless mode
	if Headless {
		log.Println("Running in headless mode...")
		// Set up signal handling for graceful shutdown
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// Start the emulator in a goroutine
		errChan := make(chan error)
		go func() {
			errChan <- gb.Run()
		}()

		// Wait for either an error or a signal
		select {
		case err := <-errChan:
			return err
		case sig := <-sigChan:
			log.Printf("Received signal %v, shutting down...", sig)
			gb.Exit()
			return <-errChan
		}
	} else {
		// Create and run the ebiten display
		log.Println("Starting visual display...")
		ebitenDisplay := display.NewEbitenDisplay(gb, gb, Scale, DebugOutput)
		return ebitenDisplay.Run()
	}
}

func main() {
	log.Print("Starting gameboy-go ... ")
	versionInfo := version.Get()
	log.Print("Version: ", versionInfo)

	flag.Parse()
	if Help {
		flag.Usage()
		return
	}

	if CartridgePath == "" {
		log.Println("! ERROR: You must define a cartridge ROM file path with '-rom-file'")
		flag.Usage()
		os.Exit(1)
	}

	if err := startEmulator(); err != nil {
		log.Printf("[ERROR] Emulator exited with error: %v", err)
		os.Exit(1)
	}
}
