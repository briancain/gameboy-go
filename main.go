package main

import (
	"github.com/briancain/gameboy-go/cpu"
	"github.com/briancain/gameboy-go/version"
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

func main() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	log.Print("Starting gameboy-go ... ")
	version := version.Get()
	log.Print("Version loaded: ", version)
	display := cpu.Display()
	log.Print(display)
}
