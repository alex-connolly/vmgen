package vmgen

import (
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

func TestVersion(t *testing.T) {
	goutil.Assert(t, version() != "", "version shouldn't be empty")
}
