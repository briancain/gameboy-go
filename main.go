package main

import (
	"flag"
	"log"
	"os"

	cpu "github.com/briancain/gameboy-go/cpu"
	version "github.com/briancain/gameboy-go/version"
)

var (
	CartridgePath string
	Help          bool
)

func init() {
	flag.BoolVar(&Help, "help", false, "Displays help")
	flag.StringVar(&CartridgePath, "rom-file", "", "A path to a cartridge ROM file")
}

func startServer() error {
	gbcpu, err := cpu.NewCPU()
	if err != nil {
		log.Print("Failed to initialize CPU: ", err)
		os.Exit(1)
	}

	display := cpu.DisplayCPUFrame(*gbcpu)
	log.Print("CPU Frame:\n", display)
	clockdisplay := cpu.DisplayClock(*gbcpu)
	log.Print("CPU Clock:\n", clockdisplay)

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
}
