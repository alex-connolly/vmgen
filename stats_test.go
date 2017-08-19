package vmgen

import "testing"

func TestStats(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	vm.Stats()
}

func TestDetailedStats(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	vm.DetailedStats()
}
