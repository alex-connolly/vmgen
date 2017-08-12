package vmgen

import "testing"

func TestPush(t *testing.T) {
	s := new(DefaultStack)
	s.Push([]byte("hello"))
}

func TestPop(t *testing.T) {

}

func TestValidate(t *testing.T) {

}
