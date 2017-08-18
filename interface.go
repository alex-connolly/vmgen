package vmgen

import (
	"encoding/hex"
	"ender/efp"
	"fmt"
	"log"
	"strconv"
)

// AddBytecode ...
func (vm *VM) AddBytecode(mnemonic string, params ...byte) {
	if vm.Bytecode == nil {
		vm.Bytecode = make([]byte, 0)
	}
	if i, ok := vm.Instructions[vm.opcodes[mnemonic]]; ok {
		vm.Bytecode = append(vm.Bytecode, i.opcode...)
		vm.Bytecode = append(vm.Bytecode, params...)
		vm.NumOpcodes++
	} else {
		log.Printf("Invalid Instruction %s\n", mnemonic)
	}
}

// CreateVM creates a new FireVM instance
func CreateVM(path string, parameters map[string]int,
	executes map[string]ExecuteFunction, fuels map[string]FuelFunction,
	disasms map[string]DisasmFunction) (*VM, []string) {
	p, errs := efp.PrototypeFile(prototype)
	if errs != nil {
		fmt.Printf("Invalid Prototype File\n")
		fmt.Println(errs)
		return nil, errs
	}
	e, errs := p.ValidateFile(path)
	if errs != nil {
		fmt.Printf("Invalid VM File %s\n", path)
		log.Println(errs)
		return nil, errs
	}

	var vm VM
	// no need to check for nil: would have errored
	vm.Author = e.FirstField("author").Value()
	vm.Name = e.FirstField("name").Value()

	vm.Parameters = parameters
	vm.AssignedParameters = make(map[string][]byte)

	vm.Instructions = make(map[string]instruction)
	vm.opcodes = make(map[string]string)
	for _, e := range e.Elements("instruction") {
		var i instruction
		i.mnemonic = e.Parameter(0).Value()

		src := []byte(e.Parameter(1).Value())

		dst := make([]byte, hex.DecodedLen(len(src)))
		_, err := hex.Decode(dst, src)
		if err != nil {
			log.Fatal(err)
		}
		i.opcode = dst

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
		log.Printf("Mnemonic: %s, Opcode: %s\n", i.mnemonic, string(i.opcode))
		i.execute = executes[i.mnemonic]

		vm.Instructions[string(i.opcode)] = i
		vm.opcodes[i.mnemonic] = string(i.opcode)
	}

	vm.stats = new(stats)
	vm.Stack = new(Stack)
	return &vm, nil
}
