package vmgen

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestBytecodeAddCommand(t *testing.T) {
	b := new(Bytecode)
	c := Command{
		mnemonic: "ADD",
	}
	b.AddCommand(&c)
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

	o := Bytecode{}
	o.Add("ADD")
	goutil.Assert(t, o.Length() == 1, "wrong o length")

	b.Concat(o)
	goutil.Assert(t, b.Length() == 2, "wrong total length")
}

func TestBytecodeFinalise(t *testing.T) {
	b := new(Bytecode)
	b.Add("ADD")
	b.AddMarker("PUSH", 10)
	b.Finalise()
	goutil.Assert(t, b.Length() == 2, "wrong total length")
	goutil.Assert(t, b.commands[1].offset == 10, fmt.Sprintf("wrong offset: %d", b.commands[1].offset))
}
