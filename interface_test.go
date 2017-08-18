package vmgen

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestAddBytecode(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	goutil.AssertNow(t, vm != nil, "vm is nil")
	err := vm.AddBytecode("PUSH", byte(1), byte(1))
	goutil.AssertNow(t, err == nil, "PUSH error")
	goutil.Assert(t, vm.Input.Code().Size() == 3,
		fmt.Sprintf("wrong bytecode length %d not %d", vm.Input.Code().Size(), 3))
	goutil.Assert(t, vm.NumOpcodes == 1, "wrong opcodes length")
	err = vm.AddBytecode("TEST", byte(1))
	goutil.AssertNow(t, err == nil, "TEST error")
	goutil.Assert(t, vm.Input.Code().Size() == 5,
		fmt.Sprintf("wrong bytecode length %d not %d", vm.Input.Code().Size(), 5))
	goutil.Assert(t, vm.NumOpcodes == 2, "wrong opcodes length")
	err = vm.AddBytecode("ADD")
	goutil.AssertNow(t, err == nil, "ADD error")
	goutil.Assert(t, vm.Input.Code().Size() == 6,
		fmt.Sprintf("wrong bytecode length %d not %d", vm.Input.Code().Size(), 6))
	goutil.Assert(t, vm.NumOpcodes == 3, "wrong opcodes length")
}
