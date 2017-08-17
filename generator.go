package vmgen

import (
	"fmt"
	"math/big"
	"os"
	"reflect"
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
	Contract           *Contract
	Bytecode           []byte
	NumOpcodes         int
	opcodes            map[string]string
}

// Address ...
type Address interface {
}

// Hash ...
type Hash interface {
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

func (vm *VM) getNextInstruction(offset int, bytecode []byte) instruction {
	size := vm.AssignedParameters["Instruction Size"]
	intSize := new(big.Int).SetBytes(size).Int64() // TODO: test
	return vm.Instructions[string(bytecode[offset:int64(offset)+intSize])]
}

const prototype = "vmgen.efp"

func (vm *VM) createParameters(params []string) []reflect.Value {
	var vals []reflect.Value
	vals = append(vals, reflect.ValueOf(vm))
	for _, p := range params {
		vals = append(vals, reflect.ValueOf(p))
	}
	return vals
}

// ExecuteFile parses opcodes from a file
func (vm *VM) ExecuteFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	bytes := make([]byte, fi.Size())
	_, err = f.Read(bytes)
	if err != nil {
		return err
	}
	return nil
}

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
		vm.AssignedParameters[k] = vm.Contract.Code[vm.Contract.Offset : vm.Contract.Offset+v]
		vm.Contract.Offset += v
	}
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", 0, 0, 1)
}
