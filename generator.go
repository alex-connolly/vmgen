package vmgen

import (
	"fmt"
	"log"
)

// VM ...
type VM struct {
	Name               string
	Author             string
	Receiver           string
	Instructions       map[string]*instruction
	Parameters         map[string]int
	AssignedParameters map[string][]byte
	PC                 int
	Current            *instruction
	stats              *stats
	Stack              *Stack
	Memory             []interface{}
	Environment        Environment
	State              State
	Input              Input
	NumOpcodes         int
	mnemonics          map[string]string
}

// Environment ...
type Environment map[string][]byte

// FuelFunction ...
type FuelFunction func(*VM) int

// ExecuteFunction ...
type ExecuteFunction func(*VM)

// Instruction for the current FireVM instance
type instruction struct {
	mnemonic       string
	opcode         string
	description    string
	execute        ExecuteFunction
	fuel           int
	fuelFunction   FuelFunction
	disasmFunction DisasmFunction
	count          int
}

func (vm *VM) nextInstruction() *instruction {
	idx := string(vm.Input.Code().Next(1))
	log.Println(idx)
	return vm.Instructions[idx]
}

const prototype = "vmgen.efp"

func (vm *VM) assignParameters() {
	for k, v := range vm.Parameters {
		vm.AssignedParameters[k] = vm.Input.Code().Next(v)
	}
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", 0, 0, 1)
}
