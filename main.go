package main

import (
	"log"

	cpu "github.com/briancain/gameboy-go/cpu"
	"github.com/briancain/gameboy-go/version"
)

func main() {
	log.Print("Starting gameboy-go ... ")
	version := version.Get()
	log.Print("Version loaded: ", version)

	gbcpu, err := cpu.NewCPU()
	if err != nil {
		log.Print("Failed to initialize CPU: ", err)
	}

	display := cpu.DisplayCPUFrame(*gbcpu)
	log.Print("CPU Frame:\n", display)
	clockdisplay := cpu.DisplayClock(*gbcpu)
	log.Print("CPU Clock:\n", clockdisplay)
}
