package vmgen

import "testing"

func TestAddMemory(t *testing.T) {
	CreateVM("example.vm", nil, nil, nil, nil)
}

func TestGetMemory(t *testing.T) {
	CreateVM("example.vm", nil, nil, nil, nil)
}
