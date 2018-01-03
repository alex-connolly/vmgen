package vmgen

/*
// DisasmFunction ...
type DisasmFunction func() []string

// DisasmBytes ...
func DisasmBytes(bytecode []byte) {
	vm.Input = new(BasicInput)
	vm.Input.Code().Append(bytecode...)
	fmt.Printf("\n%s Disassembler\n", vm.Name)
	for i := 0; i < len(vm.Name)+len(" Disassembler"); i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
	vm.assignParameters()
	for k, v := range vm.AssignedParameters {
		fmt.Printf("| %s | %x |\n", k, v)
	}
	for i := 0; i < len(vm.Name)+len(" Disassembler"); i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
	for vm.Input.Code().HasNext() {
		i := vm.nextInstruction()
		if i != nil {
			var strs []string
			if d, ok := vm.disasms[i.Mnemonic]; ok {
				strs = d(vm)
			} else {
				strs = defaultDisasm(vm)
			}
			for _, s := range strs {
				fmt.Printf("| %s", s)
			}
			fmt.Printf("|\n")
		} else {
			fmt.Printf("i is nil\n")
		}
	}
}

func defaultDisasm() []string {
	// default is just to return the instruction mnemonic
	return []string{vm.Instructions[vm.nextOpcode()].Mnemonic}
}

// DisasmString ...
func DisasmString(data string) {
	vm.DisasmBytes([]byte(data))
}

// DisasmFile ...
func DisasmFile(path string) {
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
*/
