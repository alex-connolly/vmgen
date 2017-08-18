package vmgen

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestFromHexString(t *testing.T) {
	hex := "01"
	bytes := FromHexString(hex)
	goutil.Assert(t, len(bytes) == 1, "wrong byte length")
	goutil.Assert(t, bytes[0] == 1, "wrong byte value")
}

func TestFromHexStringError(t *testing.T) {
	hex := "0AA"
	bytes := FromHexString(hex)
	goutil.Assert(t, bytes == nil, "bytes should be nil")
}
