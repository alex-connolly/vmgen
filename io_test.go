package vmgen

import (
	"fmt"
	"log"
	"testing"

	"github.com/end-r/goutil"
)

func TestGetFileBytes(t *testing.T) {
	bytes := GetFileBytes("example.bytes")
	goutil.Assert(t, bytes != nil, "bytes should not be nil")
	goutil.Assert(t, len(bytes) == 4, fmt.Sprintf("wrong bytes length %d, expected %d", len(bytes), 4))
	log.Println(bytes)
}

func TestNonexistentVM(t *testing.T) {
	vm, errs := CreateVM("notRealVM.vm", nil, nil, nil, nil)
	goutil.Assert(t, vm == nil, "vm should be nil")
	goutil.Assert(t, errs != nil, "errs shouldn't be nil")
}
