package sample

import (
	"log"

	"github.com/end-r/vmgen"
)

const (
	vmPath   = "sample.vm"
	dataPath = "sample.bytes"
)

func newVM(disasms map[string]vmgen.DisasmFunction) *vmgen.VM {
	ops := map[string]vmgen.ExecuteFunction{
		"ADD":  add,
		"PUSH": push,
	}
	fuel := map[string]vmgen.FuelFunction{}
	parameters := map[string]int{
		"Size Byte": 1,
	}
	vm, err := vmgen.CreateVM(vmPath, parameters, ops, fuel, disasms)
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
