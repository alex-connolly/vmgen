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
	Instructions       map[string]instruction
	Parameters         map[string]int
	AssignedParameters map[string][]byte
	PC                 int
	Current            instruction
	stats              *stats
	Stack              *Stack
	Memory             []interface{}
	Environment        Environment
	State              State
	Input              Input
	Bytecode           []byte
	NumOpcodes         int
	opcodes            map[string]string
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
	opcode         []byte
	description    string
	execute        ExecuteFunction
	fuel           int
	fuelFunction   FuelFunction
	disasmFunction DisasmFunction
	count          int
}

func (vm *VM) nextInstruction() instruction {
	return vm.Instructions[string(vm.Input.Code().Next(1))]
}

const prototype = "vmgen.efp"

func (vm *VM) executeInstruction(opcode string) {
	i := vm.Instructions[opcode]
	i.execute(vm)
	vm.stats.operations++
	if i.fuelFunction != nil {
		vm.stats.fuelConsumption += i.fuel
	} else {
		vm.stats.fuelConsumption += i.fuelFunction(vm)
	}
}

func (vm *VM) assignParameters() {
	for k, v := range vm.Parameters {
		log.Println("ASSIGNING")
		vm.AssignedParameters[k] = vm.Input.Code().Next(v)
		log.Printf("Assigned %s to byte array of length %d\n", k, len(vm.AssignedParameters[k]))
	}
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", 0, 0, 1)
}
