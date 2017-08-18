package vmgen

import "testing"

func TestStats(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	vm.Stats()
}
