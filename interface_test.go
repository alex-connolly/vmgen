package vmgen

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestAddBytecode(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	goutil.AssertNow(t, vm != nil, "vm is nil")
	vm.AddBytecode("PUSH", byte(1), byte(1))
	goutil.Assert(t, len(vm.Bytecode) == 3,
		fmt.Sprintf("wrong bytecode length %d not %d", len(vm.Bytecode), 3))
	goutil.Assert(t, vm.NumOpcodes == 1, "wrong opcodes length")
	vm.AddBytecode("TEST", byte(1))
	goutil.Assert(t, len(vm.Bytecode) == 5,
		fmt.Sprintf("wrong bytecode length %d not %d", len(vm.Bytecode), 5))
	goutil.Assert(t, vm.NumOpcodes == 2, "wrong opcodes length")
	vm.AddBytecode("ADD")
	goutil.Assert(t, len(vm.Bytecode) == 6,
		fmt.Sprintf("wrong bytecode length %d not %d", len(vm.Bytecode), 6))
	goutil.Assert(t, vm.NumOpcodes == 3, "wrong opcodes length")
}
