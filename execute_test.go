package vmgen

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestExecuteHexFile(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		},
		executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.ExecuteHexFile("example.bytes")
}

func TestExecuteFile(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		},
		executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.ExecuteFile("example_binary.bytes")
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

func TestExecuteHexSingleParameterSingleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	err := vm.ExecuteHexString("0101")
	goutil.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexMultipleParameterSingleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	err := vm.ExecuteHexString("01AAAA01")
	goutil.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexSingleParameterMultipleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	err := vm.ExecuteHexString("010101")
	goutil.Assert(t, err == nil, "err should be nil")
}

func TestExecuteHexMultipleParameterMultipleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	err := vm.ExecuteHexString("01AAAA0101")
	goutil.Assert(t, err == nil, "err should be nil")
}
