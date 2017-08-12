package sample

import (
	"axia/vmgen"
	"log"
)

const (
	vmPath   = "sample.vm"
	dataPath = "sample.bytes"
)

func main() {
	ops := map[string]vmgen.ExecuteFunction{
		"Add":  add,
		"Push": push,
	}
	fuel := map[string]vmgen.FuelFunction{}
	vm, err := vmgen.CreateVM(vmPath, ops, fuel)
	if err != nil {
		log.Fatal(err)
		return
	}

	vm.ExecuteFile(dataPath)
}

func add(vm vmgen.VM, params []byte) {
	a := vm.Stack.Pop(1)
	b := vm.vm.Stack.Pop(1)
	vm.vm.Stack.Push(a + b)
}

func push(vm vmgen.VM, params []byte) {
	vm.vm.Stack.Push(params)
}
