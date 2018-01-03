package vmgen

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestGetFileBytes(t *testing.T) {
	bytes := GetFileBytes("example.bytes")
	goutil.Assert(t, bytes != nil, "bytes should not be nil")
	goutil.Assert(t, len(bytes) == 4, fmt.Sprintf("wrong bytes length %d, expected %d", len(bytes), 4))
}

func TestGetFileBytesErrors(t *testing.T) {
	bytes := GetFileBytes("not_example.bytes")
	goutil.Assert(t, bytes == nil, "bytes should not be nil")
}

func TestNonExistentVM(t *testing.T) {
	vm, errs := CreateVM("notRealVM.vm", nil, nil, nil, nil)
	goutil.Assert(t, vm == nil, "vm should be nil")
	goutil.Assert(t, errs != nil, "errs shouldn't be nil")
}
