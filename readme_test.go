package vmgen

import "testing"

func TestGenerateReadme(t *testing.T) {
	vm, _ := CreateVM("example.vm", nil, nil, nil, nil)
	vm.GenerateReadMe("SPEC.md")
}
