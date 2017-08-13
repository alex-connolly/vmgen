package vmgen

import (
	"axia/guardian/go/util"
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	s := new(Stack)
	data := []byte("hello")
	s.Push(data)
	util.Assert(t, s.Size() == len(data), fmt.Sprintf("wrong length %d, expected %d", s.Size(), len(data)))
}

func TestPop(t *testing.T) {
	data := []byte("me")
	util.Assert(t, len(data) == 2, "wrong data length")
	vm := new(VM)
	vm.SetStack(new(Stack))
	util.Assert(t, vm.Stack.Size() == 0, "wrong starting length")
	vm.Stack.Push(data)
	util.Assert(t, vm.Stack.Size() == 2, "wrong length after push")
	bytes := vm.Stack.Pop(2)
	util.Assert(t, vm.Stack.Size() == 0, "wrong length")
	util.Assert(t, len(bytes) == len(data), "wrong popped value")
}

func TestValidate(t *testing.T) {

}
