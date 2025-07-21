package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	gbcore "github.com/briancain/gameboy-go/gbcore"
	version "github.com/briancain/gameboy-go/version"
)

var (
	CartridgePath string
	Help          bool
	DebugOutput   bool
	Scale         int
	Headless      bool
)

func init() {
	flag.BoolVar(&Help, "help", false, "Displays help")
	flag.StringVar(&CartridgePath, "rom-file", "", "A path to a cartridge ROM file")
	flag.BoolVar(&DebugOutput, "debug", false, "Displays debug output")
	flag.IntVar(&Scale, "scale", 2, "Screen scale factor (1-4)")
	flag.BoolVar(&Headless, "headless", false, "Run without display (for testing)")
}

func startEmulator() error {
	// Create and initialize the Game Boy core
	gb, err := gbcore.NewGameBoyCore(DebugOutput)
	if err != nil {
		log.Print("Failed to create new gbcore: ", err)
		return err
	}

	if err := gb.Init(CartridgePath); err != nil {
		log.Print("[ERROR] Failed to initialize new gbcore!\n", err)
		return err
	}

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
