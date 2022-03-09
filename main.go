package main

import (
	"flag"
	"log"
	"os"

	gbcore "github.com/briancain/gameboy-go/gbcore"
	version "github.com/briancain/gameboy-go/version"
)

var (
	CartridgePath string
	Help          bool
	DebugOutput   bool
)

func init() {
	flag.BoolVar(&Help, "help", false, "Displays help")
	flag.StringVar(&CartridgePath, "rom-file", "", "A path to a cartridge ROM file")
	flag.BoolVar(&DebugOutput, "debug", false, "Displays debug output")
}

func startServer() error {
	loadGBCore()

	// TODO(briancain): spin up a server for display and accept inputs
	return nil
}

func loadGBCore() error {
	gb, err := gbcore.NewGameBoyCore(DebugOutput)
	if err != nil {
		log.Print("Failed to create new gbcore: ", err)
		os.Exit(1)
	}
	if err := gb.Init(CartridgePath); err != nil {
		log.Print("[ERROR] Failed to initialize new gbcore!\n", err)
		os.Exit(1)
	}

	gb.Run()

	return nil
}

func main() {
	log.Print("Starting gameboy-go ... ")
	version := version.Get()
	log.Print("Version loaded: ", version)

	flag.Parse()
	if Help {
		flag.Usage()
	}
	if CartridgePath == "" {
		log.Println("! ERROR: You must define a cartridge ROM file path with '-rom-file'\n")
		flag.Usage()
		os.Exit(1)
	}

	startServer()
}
