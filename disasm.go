package vmgen

import (
	"log"
	"os"
)

// DisasmFunction ...
type DisasmFunction func(*VM, int, []byte) ([]string, int)

// DisasmBytes ...
func (vm *VM) DisasmBytes(bytecode []byte) {
	log.Printf("%s Disassembler", vm.Name)
	for i := 0; i < len(vm.Name)+len(" Disassembler"); i++ {
		log.Printf("=")
	}
	count := 0
	vm.assignParameters()
	for k, v := range vm.AssignedParameters {
		log.Printf("| %s: %x |", k, v)
		count += len(v)
	}
	for i := count; i < len(bytecode); i++ {
		str, offset := vm.nextInstruction().disasmFunction(vm, i, bytecode)
		i += offset
		log.Println(str)
	}
}

func defaultDisasm(vm *VM, offset int, bytecode []byte) ([]string, int) {
	// default is just to return the instruction mnemonic
	return nil, 0
}

// DisasmString ...
func (vm *VM) DisasmString(data string) {
	vm.DisasmBytes([]byte(data))
}

// DisasmFile ...
func (vm *VM) DisasmFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return
	}
	bytes := make([]byte, fi.Size())
	_, err = f.Read(bytes)
	if err != nil {
		return
	}
	vm.DisasmBytes(bytes)
}
