package vmgen

import (
	"log"
	"os"
)

// DissemblyFunction ...
type DissemblyFunction func(int, []byte) (string, int)

// DissemblyOptions ...
type DissemblyOptions struct {
	Parameters     map[string]byte
	ParameterNames []string
	Functions      map[byte]DissemblyFunction
}

// DisasmBytes ...
func (vm *VM) DisasmBytes(bytecode []byte, options *DissemblyOptions) {
	log.Printf("%s Disassembler", vm.Name)
	for i := 0; i < len(vm.Name)+len(" Disassembler"); i++ {
		log.Printf("=")
	}
	for i, pn := range options.ParameterNames {
		// assign the parameters (used in disasm calcs)
		options.Parameters[pn] = bytecode[i]
		log.Printf("| %s: %x |", pn, bytecode[i])
	}
	for i := len(options.ParameterNames); i < len(bytecode); i++ {
		if f, ok := options.Functions[bytecode[i]]; ok {
			str, offset := f(i, bytecode)
			i += offset
			log.Println(str)
		} else {
			// default is just to print the instruction mnemonic
			log.Printf("| %s |\n", vm.Instructions[bytecode[i]].mnemonic)
		}
	}
}

// DisasmString ...
func (vm *VM) DisasmString(data string, options *DissemblyOptions) {
	vm.DisasmBytes([]byte(data), options)
}

// DisasmFile ...
func (vm *VM) DisasmFile(path string, options *DissemblyOptions) {
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
	vm.DisasmBytes(bytes, options)
}
