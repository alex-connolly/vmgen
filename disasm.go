package vmgen

import (
	"log"
	"os"
)

// DisasmBytes ...
func (vm *VM) DisasmBytes(bytecode []byte) {
	log.Printf("%s Disassembler", vm.Name)
	for i := 0; i < len(vm.Name)+len(" Disassembler"); i++ {
		log.Printf("=")
	}
	log.Printf("Size Byte: %x", bytecode[0])
	for i := 1; i < len(bytecode); i++ {

	}
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
