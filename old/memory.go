package vmgen

// Memory interface
type Memory interface {
	DisplayContents()
}

// AddMemory ...
func (vm *VM) AddMemory(key string, m Memory) {
	if vm.Memory == nil {
		vm.Memory = make(map[string]Memory)
	}
	vm.Memory[key] = m
}

// GetMemory ...
func (vm *VM) GetMemory(key string) Memory {
	return vm.Memory[key]
}
