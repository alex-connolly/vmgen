package vmgen

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
	vm.Input = new(BasicInput)
	vm.Input.Code().Append(bytes...)
	vm.assignParameters()
	for vm.Input.Code().HasNext() {
		vm.executeInstruction(vm.nextInstruction())
	}
	return nil
}

func (vm *VM) executeInstruction(i *instruction) {
	if i != nil {
		if vm.executes != nil {
			if e, ok := vm.executes[i.mnemonic]; ok {
				e(vm)
			}
		}
		vm.stats.operations++
		if vm.fuels != nil {
			if f, ok := vm.fuels[i.mnemonic]; ok {
				f(vm)
			} else {
				vm.stats.fuelConsumption += i.fuel
			}
		} else {
			vm.stats.fuelConsumption += i.fuel
		}
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
