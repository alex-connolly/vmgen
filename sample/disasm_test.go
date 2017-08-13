package sample

import (
	"fmt"
	"testing"

	"github.com/end-r/vmgen"
)

func TestDisasm(t *testing.T) {
	vm := newVM()
	opts := vmgen.DisasmOption{
		ParameterNames: []string{
			"Size Byte",
		},
	}
	opts.Functions = map[string]vmgen.DisasmFunction{
		"PUSH": func(offset int, bytecode []byte) (string, int) {
			size := int(ops.Parameters["Size Byte"])

			return fmt.Sprintf("| PUSH | %d | %x |", size, bytecode[offset+1:offset+1+size]), size + 1
		},
	}
	vm.DisasmString("01", opts)
}
