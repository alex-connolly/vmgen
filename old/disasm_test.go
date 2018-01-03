package vmgen

/*
func TestEmptyDisasm(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	vm.DisasmString("")
}

func TestDisasmFile(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	vm.DisasmFile("example.bytes")
}

func TestDisasmHexSingleParameter(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.DisasmString("01")
}

func TestDisasmHexMultipleParameter(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.DisasmString("01AAAA")
}

func TestDisasmHexSingleParameterSingleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.DisasmString("0101")
}

func TestDisasmHexMultipleParameterSingleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.DisasmString("01AAAA01")
}

func TestDisasmHexSingleParameterMultipleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size": 1,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.DisasmString("010101")
}

func TestDiasmHexMultipleParameterMultipleInstruction(t *testing.T) {
	vm, errs := CreateVM("example.vm",
		map[string]int{
			"size":  1,
			"width": 2,
		}, executes, nil, nil)
	goutil.Assert(t, vm != nil, "vm shouldn't be nil")
	goutil.Assert(t, errs == nil, "errs should be nil")
	vm.DisasmString("01AAAA0101")
}

*/
