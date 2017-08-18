package vmgen

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestExecuteFile(t *testing.T) {
	vm, errs := CreateVM("example.vm", nil, nil, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.ExecuteFile("example.bytes")
}

func TestExecuteHexSingleParameter(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	err := vm.ExecuteHexString("01")
	goutil.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexMultipleParameter(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	err := vm.ExecuteHexString("01AAAA")
	goutil.Assert(t, err == nil, "err should be nil")
}

func TestVersion(t *testing.T) {
	goutil.Assert(t, version() != "", "version shouldn't be empty")
}
