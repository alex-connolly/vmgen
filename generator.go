package vmgen

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/end-r/efp"
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
	Environment        *Environment
	Contract           *Contract
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

// CreateVM creates a new FireVM instance
func CreateVM(path string, parameters map[string]int,
	executes map[string]ExecuteFunction, fuels map[string]FuelFunction,
	disasms map[string]DisasmFunction) (*VM, []string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	p, errs := efp.PrototypeFile(exPath + "/" + prototype)
	if errs != nil {
		fmt.Printf("Invalid Prototype File\n")
		fmt.Println(errs)
		return nil, errs
	}
	e, errs := p.ValidateFile(path)
	if errs != nil {
		fmt.Printf("Invalid VM File %s\n", path)
		return nil, errs
	}

	var vm VM
	// no need to check for nil: would have errored
	vm.Author = e.FirstField("author").Value()
	vm.Name = e.FirstField("name").Value()

	vm.Instructions = make(map[string]instruction)
	for _, e := range e.Elements("instruction") {
		var i instruction

		i.mnemonic = e.Parameter(0).Value()
		i.opcode = []byte(e.Parameter(1).Value())

		i.description = e.FirstField("description").Value()

		if f, ok := fuels[i.mnemonic]; ok {
			i.fuelFunction = f
		} else {
			if e.FirstField("fuel") == nil {
				i.fuel = 0
			} else {
				fuel, err := strconv.ParseInt(e.FirstField("fuel").Value(), 10, 64)
				if err != nil {

				} else {
					i.fuel = int(fuel)
				}
			}

		}

		i.execute = executes[i.mnemonic]

		vm.Instructions[string(i.opcode)] = i
	}

	vm.stats = new(stats)
	vm.Stack = new(Stack)
	return &vm, nil
}

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
