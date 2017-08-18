package vmgen

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestNonexistentVM(t *testing.T) {
	vm, errs := CreateVM("notRealVM.vm", nil, nil, nil, nil)
	goutil.Assert(t, vm == nil, "vm should be nil")
	goutil.Assert(t, errs != nil, "errs shouldn't be nil")
}
