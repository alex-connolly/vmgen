package vmgen

import (
	"log"
)

// ExecuteFile parses opcodes from a file
func (vm *VM) ExecuteFile(path string) []string {
	return vm.ExecuteBytes(GetFileBytes(path))
}

// ExecuteString parses opcodes from a string
func (vm *VM) ExecuteString(data string) []string {
	return vm.ExecuteBytes([]byte(data))
}

// ExecuteBytes parses opcodes from a byte array
func (vm *VM) ExecuteBytes(bytes []byte) []string {
	log.Printf("bytes length: %d", len(bytes))
	vm.Input = new(BasicInput)
	vm.Input.Code().Append(bytes...)
	log.Println("xx")
	vm.assignParameters()
	log.Println("yy")
	for vm.Input.Code().HasNext() {
		vm.executeInstruction(vm.nextInstruction())
	}
	return nil
}

func (vm *VM) executeInstruction(i *instruction) {
	log.Println("a")
	if i.hasExecuteFunction {
		i.execute(vm)
	}
	vm.stats.operations++
	log.Println("b")
	if i.hasFuelFunction {
		vm.stats.fuelConsumption += i.fuel
	} else {
		vm.stats.fuelConsumption += i.fuelFunction(vm)
	}
}

// ExecuteHexFile parses hex opcodes from a file
func (vm *VM) ExecuteHexFile(path string) []string {
	return vm.ExecuteHexBytes(GetFileBytes(path))
}

// ExecuteHexString ...
func (vm *VM) ExecuteHexString(hex string) []string {
	return vm.ExecuteHexBytes([]byte(hex))
}

// ExecuteHexBytes parses opcodes from a byte array
func (vm *VM) ExecuteHexBytes(bytes []byte) []string {
	if len(bytes)%2 != 0 {
		return []string{"Invalid Hex Input"}
	}
	return vm.ExecuteBytes(FromHexBytes(bytes))
}
