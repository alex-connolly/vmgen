package vmgen

import (
	"fmt"
	"log"
	"testing"

	"github.com/end-r/goutil"
)

func TestNextInstruction(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	vm.Input = new(BasicInput)
	bytes := []byte("0101")
	vm.Input.Code().Append(FromHexBytes(bytes)...)
	i := vm.nextInstruction()
	goutil.AssertNow(t, i != nil, "next instruction shouldn't be nil")
	goutil.AssertNow(t, i.mnemonic == "ADD", "wrong mnemonic")
}

func TestNextOpcode(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	vm.Input = new(BasicInput)
	bytes := []byte("0101")
	vm.Input.Code().Append(FromHexBytes(bytes)...)
	goutil.Assert(t, vm.Input.Code().Size() == 2, "wrong code size")
	log.Println(vm.Input.Code().Data)
	n := vm.nextOpcode()
	goutil.Assert(t, n == 1, fmt.Sprintf("Next opcode was: %d\n", n))
}

func TestVersion(t *testing.T) {
	goutil.Assert(t, version() != "", "version shouldn't be empty")
}
