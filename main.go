package main

import (
	"github.com/briancain/gameboy-go/cpu"
	"github.com/briancain/gameboy-go/version"
	"log"
)

func main() {
	log.Print("Starting gameboy-go ... ")
	version := version.Get()
	log.Print("Version loaded: ", version)
	gbcpu := cpu.NewCPU()
	display := cpu.DisplayCPUFrame(*gbcpu)
	log.Print("CPU Frame:\n", display)
}
