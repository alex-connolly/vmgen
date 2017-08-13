package vmgen

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/end-r/efp"
)

// VM ...
type VM struct {
	Name         string
	Author       string
	Receiver     string
	Instructions map[string]instruction
	PC           int
	Current      instruction
	stats        *stats
	Stack        *Stack
	Memory       []interface{}
	Environment  *Environment
	Contract     *Contract
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
	opcode       string
	description  string
	execute      ExecuteFunction
	fuel         int
	fuelFunction FuelFunction
	count        int
}

const prototype = "vmgen.efp"

// CreateVM creates a new FireVM instance
func CreateVM(path string, executes map[string]ExecuteFunction, fuels map[string]FuelFunction) (*VM, []string) {
	p, errs := efp.PrototypeFile(prototype)
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
	vm.Receiver = e.FirstField("receiver").Value()

	vm.Instructions = make(map[string]instruction)
	for _, e := range e.Elements("instruction") {
		var i instruction
		i.description = e.FirstField("description").Value()

		// try to get fuel as an integer
		fuel, err := strconv.ParseInt(e.FirstField("fuel").Value(), 10, 64)

		// if not, it's a fuel function
		if err != nil {
			i.fuel = int(fuel)
		} else {
			i.fuelFunction = fuels[e.FirstField("fuel").Value()]
		}

		i.execute = executes[e.FirstField("execute").Value()]

		opcode := e.Parameter(0).Value()
		i.opcode = opcode

		vm.Instructions[opcode] = i
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

func version() string {
	return fmt.Sprintf("%d.%d.%d", 0, 0, 1)
}
