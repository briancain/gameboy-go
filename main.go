package main

import (
	//"fmt"
	"log"
	"os"
	//"github.com/briancain/gameboy-go/cpu"
	"github.com/briancain/gameboy-go/version"

	"github.com/hashicorp/logutils"
)

func main() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	log.Print("gameboy-go starting... ")
	version := version.Get()
	log.Print("Version loaded: ", version)
}
