package vmgen

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestPush(t *testing.T) {
	s := new(Stack)
	data := []byte("hello")
	s.Push(data)
	goutil.Assert(t, s.Size() == 1, fmt.Sprintf("wrong length %d, expected %d", s.Size(), len(data)))
}

func TestPop(t *testing.T) {
	data := []byte("me")
	goutil.Assert(t, len(data) == 2, "wrong data length")
	vm := new(VM)
	vm.SetStack(new(Stack))
	goutil.Assert(t, vm.Stack.Size() == 0, "wrong starting length")
	vm.Stack.Push(data)
	goutil.Assert(t, vm.Stack.Size() == 1, "wrong length after push")
	bytes := vm.Stack.Pop()
	goutil.Assert(t, vm.Stack.Size() == 0, "wrong length")
	goutil.Assert(t, len(bytes) == len(data),
		fmt.Sprintf("wrong popped value %d, expected %d", len(bytes), len(data)))

}

func TestValidate(t *testing.T) {
	data := []byte("me")
	vm := new(VM)
	vm.SetStack(new(Stack))
	goutil.Assert(t, vm.Stack.validate(0), "wrong validate 0")
	goutil.Assert(t, !vm.Stack.validate(1), "wrong validate 1")
	vm.Stack.Push(data)
	goutil.Assert(t, vm.Stack.validate(0), "wrong validate 0")
	goutil.Assert(t, vm.Stack.validate(1), "wrong validate 1")
	goutil.Assert(t, !vm.Stack.validate(2), "wrong validate 2")
	vm.Stack.Pop()
	goutil.Assert(t, vm.Stack.validate(0), "wrong validate 0")
	goutil.Assert(t, !vm.Stack.validate(1), "wrong validate 1")
}
