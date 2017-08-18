package vmgen

import (
	"os"
)

// ExecuteFile parses opcodes from a file
func (vm *VM) ExecuteFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		return []string{err.Error()}
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return []string{err.Error()}
	}
	bytes := make([]byte, fi.Size())
	_, err = f.Read(bytes)
	if err != nil {
		return []string{err.Error()}
	}
	return vm.ExecuteBytes(bytes)
}

// ExecuteString parses opcodes from a string
func (vm *VM) ExecuteString(data string) []string {
	return vm.ExecuteBytes([]byte(data))
}

// ExecuteBytes parses opcodes from a byte array
func (vm *VM) ExecuteBytes(bytes []byte) []string {
	vm.Input = new(BasicInput)
	vm.Input.Code().Append(bytes...)
	vm.assignParameters()
	for vm.Input.Code().HasNext() {
		vm.executeInstruction(vm.nextInstruction())
	}
	return nil
}

func (vm *VM) executeInstruction(i *instruction) {
	i.execute(vm)
	vm.stats.operations++
	if i.fuelFunction != nil {
		vm.stats.fuelConsumption += i.fuel
	} else {
		//vm.stats.fuelConsumption += i.fuelFunction(vm)
	}
}

// ExecuteHexFile parses hex opcodes from a file
func (vm *VM) ExecuteHexFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		return []string{err.Error()}
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return []string{err.Error()}
	}
	bytes := make([]byte, fi.Size())
	_, err = f.Read(bytes)
	if err != nil {
		return []string{err.Error()}
	}
	return vm.ExecuteHexBytes(bytes)
}

// ExecuteHexString ...
func (vm *VM) ExecuteHexString(hex string) []string {
	return vm.ExecuteHexBytes([]byte(hex))
}

// ExecuteHexBytes parses opcodes from a byte array
func (vm *VM) ExecuteHexBytes(bytes []byte) []string {
	return vm.ExecuteBytes(FromHexBytes(bytes))
}
