package controller

import "log"

// A keyboard controller
type Keyboard struct {
	Name string
}

func (k Keyboard) Init() {
	log.Println("[DEBUG] Initializing Keyboard controller ...")
}

func (k Keyboard) Update() bool {
	return false
}
