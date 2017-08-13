package sample

import "testing"

func TestDisasm(t *testing.T) {
	vm := newVM()
	vm.DisasmString("01")
}
