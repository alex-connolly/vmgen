package vmgen

import (
	"ender/efp"
	"fmt"
	"log"
	"strconv"
)

// AddBytecode ...
func (vm *VM) AddBytecode(mnemonic string, params ...byte) error {
	if vm.Input == nil {
		vm.Input = new(BasicInput)
	}
	opcode, ok := vm.mnemonics[mnemonic]
	if ok {
		if i, ok := vm.Instructions[opcode]; ok {
			if i.opcode != "" {
				vm.Input.Code().Append([]byte(i.opcode)...)
				vm.Input.Code().Append(params...)
				vm.NumOpcodes++
			}
			return nil
		}
	}
	return fmt.Errorf("Invalid Instruction %s\n", mnemonic)
}

// CreateVM creates a new FireVM instance
func CreateVM(path string, parameters map[string]int,
	executes map[string]ExecuteFunction, fuels map[string]FuelFunction,
	disasms map[string]DisasmFunction) (*VM, []string) {
	p, errs := efp.PrototypeFile(prototype)
	if errs != nil {
		return nil, errs
	}
	e, errs := p.ValidateFile(path)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}

	var vm VM
	// no need to check for nil: would have errored
	vm.Author = e.FirstField("author").Value()
	vm.Name = e.FirstField("name").Value()

	vm.Parameters = parameters
	vm.AssignedParameters = make(map[string][]byte)
	vm.stats = new(stats)

	vm.Instructions = make(map[string]*instruction)
	vm.mnemonics = make(map[string]string)
	for _, e := range e.Elements("instruction") {
		var i instruction
		i.mnemonic = e.Parameter(0).Value()

		s := e.Parameter(1).Value()

		opcode := FromHexString(s)

		if len(opcode) == 1 {
			i.opcode = string(opcode[:])
		} else {
			log.Println(opcode)
			i.opcode = strconv.Itoa(int(opcode[0]))
		}

		fmt.Printf("opcode: %s\n", i.opcode)

		i.description = e.FirstField("description").Value()

		if f, ok := fuels[i.mnemonic]; ok {
			i.fuelFunction = f
		} else {
			if e.FirstField("fuel") == nil {
				i.fuel = 0
			} else {
				fuel, err := strconv.ParseInt(e.FirstField("fuel").Value(), 10, 64)
				if err != nil {
					return nil, []string{err.Error()}
				} else {
					i.fuel = int(fuel)
				}
			}
		}

		if disasms != nil {
			if d, ok := disasms[i.mnemonic]; ok {
				i.disasmFunction = d
			} else {
				i.disasmFunction = defaultDisasm
			}
		}

		i.execute = executes[i.mnemonic]

		vm.Instructions[string(i.opcode)] = &i
		vm.mnemonics[i.mnemonic] = string(i.opcode)
	}

	vm.stats = new(stats)
	vm.Stack = new(Stack)
	return &vm, nil
}
