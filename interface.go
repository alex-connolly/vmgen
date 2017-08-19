package vmgen

import (
	"ender/efp"
	"fmt"
	"log"
)

// AddBytecode ...
func (vm *VM) AddBytecode(mnemonic string, params ...byte) error {
	if vm.Input == nil {
		vm.Input = new(BasicInput)
	}
	if i, ok := vm.mnemonics[mnemonic]; ok {
		vm.Input.Code().Append(i.opcode)
		vm.Input.Code().Append(params...)
		vm.NumOpcodes++
		return nil
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

	vm.executes = executes
	vm.fuels = fuels
	vm.disasms = disasms

	vm.Instructions = make(map[byte]*instruction)
	vm.mnemonics = make(map[string]*instruction)

	for _, e := range e.Elements("instruction") {
		errs := vm.AddInstruction(nil, e)
		if errs != nil {
			return nil, errs
		}
	}

	for _, cat := range e.Elements("category") {
		c := new(category)
		c.name = cat.Parameter(0).Value()
		if cat.Fields("description") != nil {
			c.description = cat.FirstField("description").Value()
		}
		for _, e := range cat.Elements("instruction") {
			errs := vm.AddInstruction(c, e)
			if errs != nil {
				return nil, errs
			}
		}
		if vm.categories == nil {
			vm.categories = make(map[string]*category)
		}
		vm.categories[c.name] = c
	}

	vm.stats = new(stats)
	vm.Stack = new(Stack)
	return &vm, nil
}

func stringToOpcode(str string) byte {
	return FromHexString(str)[0]
}

// AddInstruction ...
func (vm *VM) AddInstruction(c *category, e *efp.Element) []string {

	i := new(instruction)
	i.mnemonic = e.Parameter(0).Value()

	i.opcode = stringToOpcode(e.Parameter(1).Value())

	i.description = e.FirstField("description").Value()

	vm.Instructions[i.opcode] = i

	vm.mnemonics[i.mnemonic] = i

	if c != nil {
		if c.instructions == nil {
			c.instructions = make(map[string]*instruction)
		}
		c.instructions[i.mnemonic] = i
	}
	return nil
}
