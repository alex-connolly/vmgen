package sample

import (
	"math/big"
	"testing"

	"github.com/end-r/vmgen"
)

func TestDisasm(t *testing.T) {
	vm := newVM(map[string]vmgen.DisasmFunction{
		"PUSH": func(vm *vmgen.VM, offset int, bytecode []byte) ([]string, int) {
			size := vm.AssignedParameters["Size Byte"]
			intSize := new(big.Int).SetBytes(size).Int64()
			return []string{string(size), string(bytecode[offset+1 : int64(offset+1)+intSize])}, int(intSize) + 1
		},
	})
	vm.DisasmString("010101")
}
