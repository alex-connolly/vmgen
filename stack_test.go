package vmgen

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestStackPush(t *testing.T) {
	s := new(Stack)
	data := []byte("hello")
	s.Push(data)
	goutil.Assert(t, s.Size() == 1, fmt.Sprintf("wrong length %d, expected %d", s.Size(), len(data)))
}

func TestStackPop(t *testing.T) {
	data := []byte("me")
	goutil.Assert(t, len(data) == 2, "wrong data length")
	s := new(Stack)
	goutil.Assert(t, s.Size() == 0, "wrong starting length")
	s.Push(data)
	goutil.Assert(t, s.Size() == 1, "wrong length after push")
	bytes := s.Pop()
	goutil.Assert(t, s.Size() == 0, "wrong length")
	goutil.Assert(t, len(bytes) == len(data),
		fmt.Sprintf("wrong popped value %d, expected %d", len(bytes), len(data)))

}

func TestStackValidate(t *testing.T) {
	data := []byte("me")
	s := new(Stack)
	goutil.Assert(t, s.validate(0), "wrong validate 0")
	goutil.Assert(t, !s.validate(1), "wrong validate 1")
	s.Push(data)
	goutil.Assert(t, s.validate(0), "wrong validate 0")
	goutil.Assert(t, s.validate(1), "wrong validate 1")
	goutil.Assert(t, !s.validate(2), "wrong validate 2")
	s.Pop()
	goutil.Assert(t, s.validate(0), "wrong validate 0")
	goutil.Assert(t, !s.validate(1), "wrong validate 1")
}

func TestStackBack(t *testing.T) {
	d1 := []byte("me444")
	d2 := []byte("al44")
	d3 := []byte("xx4")
	s := new(Stack)
	s.Push(d1)
	s.Push(d2)
	s.Push(d3)
	goutil.Assert(t, s.Size() == 3, "wrong stack size")
	goutil.Assert(t, len(s.back(2)) == 5, "wrong back value")
}

func TestStackSwap(t *testing.T) {
	d1 := []byte("me444")
	d2 := []byte("al44")
	d3 := []byte("xx4")
	s := new(Stack)
	s.Push(d1)
	s.Push(d2)
	s.Push(d3)
	goutil.Assert(t, s.Size() == 3, "wrong stack size")
	goutil.Assert(t, len(s.back(2)) == 5, "wrong bottom value pre-swap")
	goutil.Assert(t, len(s.back(0)) == 3, "wrong top value pre-swap")
	s.Swap(2)
	goutil.Assert(t, s.Size() == 3, "wrong stack size")
	goutil.Assert(t, len(s.back(2)) == 3, "wrong bottom value post-swap")
	goutil.Assert(t, len(s.back(0)) == 5, "wrong top value post-swap")
	s.Swap(1)
	goutil.Assert(t, s.Size() == 3, "wrong stack size")
	goutil.Assert(t, len(s.back(1)) == 5, "wrong bottom value post-swap 2")
	goutil.Assert(t, len(s.back(0)) == 4, "wrong top value post-swap 2")
}

func TestStackDup(t *testing.T) {
	d1 := []byte("me444")
	s := new(Stack)
	s.Push(d1)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	s.Dup(1)
	goutil.Assert(t, s.Size() == 2, "wrong stack size")
	goutil.Assert(t, len(s.Pop()) == 5, "wrong duplicated data")
	s.Dup(1)
	s.Dup(2)
	goutil.Assert(t, s.Size() == 4, "wrong stack size 1 2")
	goutil.Assert(t, len(s.Pop()) == 5, "wrong duplicated data")
}
