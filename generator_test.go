package vmgen

import "testing"

func TestAddPushVM(t *testing.T) {
	vm := CreateVM("tests/example.vm")
	vm.ExecuteFile("tests/example.bytes")
	vm.Stats()
}
