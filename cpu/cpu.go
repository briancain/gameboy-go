package cpu

import ()

type Registers struct {
	// 8-bit registers
	A byte
	B byte
	C byte
	D byte
	E byte
	F byte // Flag register
	// 16-bit registers
	PC byte // Program Counter
	SP byte // Stack Pointer
	// Clock for last instruction
	M byte
	T byte
}

type Clock struct {
	m byte
	t byte
}

func Display() string {
	return "Hello"
}
