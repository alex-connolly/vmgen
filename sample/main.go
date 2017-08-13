package sample

import (
	"log"

	"github.com/end-r/vmgen"
)

const (
	vmPath   = "sample.vm"
	dataPath = "sample.bytes"
)

func newVM() *vmgen.VM {
	ops := map[string]vmgen.ExecuteFunction{
		"ADD":  add,
		"PUSH": push,
	}
	fuel := map[string]vmgen.FuelFunction{}
	vm, err := vmgen.CreateVM(vmPath, ops, fuel)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return vm
}

func add(vm *vmgen.VM) {

}

func push(vm *vmgen.VM) {
}
