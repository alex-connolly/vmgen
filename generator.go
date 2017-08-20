package vmgen

import (
	"fmt"
)

// VM ...
type VM struct {
	Name               string
	Author             string
	categories         map[string]*category
	Instructions       map[byte]*instruction
	mnemonics          map[string]*instruction
	Parameters         map[string]int
	AssignedParameters map[string][]byte
	PC                 int
	Current            *instruction
	stats              *stats
	Stack              *Stack
	Memory             map[string]Memory
	Environment        Environment
	State              State
	Input              Input
	NumOpcodes         int

	executes map[string]ExecuteFunction
	fuels    map[string]FuelFunction
	disasms  map[string]DisasmFunction
}

// Environment ...
type Environment map[string][]byte

// FuelFunction ...
type FuelFunction func(*VM) int

// ExecuteFunction ...
type ExecuteFunction func(*VM)

// Instruction for the current FireVM instance
type instruction struct {
	mnemonic    string
	opcode      byte
	description string
	fuel        int
	count       int
}

type category struct {
	name         string
	description  string
	instructions map[string]*instruction
}

func (vm *VM) nextInstruction() *instruction {
	return vm.Instructions[vm.nextOpcode()]
}

func (vm *VM) nextOpcode() byte {
	return vm.Input.Code().Next(1)[0]
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
