package vmgen

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestBytecodeAddCommand(t *testing.T) {
	b := new(Bytecode)
	c := Command{
		mnemonic: "ADD",
	}
	b.AddCommand(c)
	goutil.Assert(t, b.Length() == 1, "wrong length")
}

func TestBytecodeAdd(t *testing.T) {
	b := new(Bytecode)
	b.Add("ADD")
	goutil.Assert(t, b.Length() == 1, "wrong length")
}

func TestBytecodeConcat(t *testing.T) {
	b := new(Bytecode)
	b.Add("ADD")
	goutil.Assert(t, b.Length() == 1, "wrong b length")

	o := new(Bytecode)
	o.Add("ADD")
	goutil.Assert(t, o.Length() == 1, "wrong o length")

	b.Concat(o)
	goutil.Assert(t, b.Length() == 2, "wrong total length")
}
