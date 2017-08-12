package vmgen

import "fmt"

// VM ...
type VM struct {
	Name           string
	Author         string
	Receiver       string
	Instructions   map[string]instruction
	ProgramCounter int
	Current        instruction
	stats          stats
	// memory interfaces
	Stack  *Stack
	Memory []interface{}
}

// Instruction for the current FireVM instance
type instruction struct {
	opcode       string
	description  string
	execute      string
	fuel         int64
	fuelFunction string
	count        int
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", 0, 0, 1)
}
