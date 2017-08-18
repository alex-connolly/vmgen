package vmgen

import (
	"encoding/hex"
	"log"
	"os"
)

// ExecuteFile parses opcodes from a file
func (vm *VM) ExecuteFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	bytes := make([]byte, fi.Size())
	_, err = f.Read(bytes)
	if err != nil {
		return err
	}
	return vm.ExecuteBytes(bytes)
}

// ExecuteString parses opcodes from a string
func (vm *VM) ExecuteString(data string) error {
	return vm.ExecuteBytes([]byte(data))
}

// ExecuteBytes parses opcodes from a byte array
func (vm *VM) ExecuteBytes(bytes []byte) error {
	vm.Input = new(BasicInput)
	vm.Input.SetCode(NewByterable(bytes))
	log.Println("about to assign")
	vm.assignParameters()
	log.Println("finished assigning")
	for vm.Input.Code().HasNext() {
		log.Println("HI")
		vm.nextInstruction().execute(vm)
	}
	return nil
}

func (vm *VM) ExecuteHexString(hexString string) error {
	src := []byte(hexString)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return err
	}
	return vm.ExecuteBytes(dst)
}
