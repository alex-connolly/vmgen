package vmgen

import (
	"fmt"
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

func TestFromHexStringTwoBytes(t *testing.T) {
	hex := "0101"
	bytes := FromHexString(hex)
	goutil.Assert(t, len(bytes) == 2, "wrong byte length")
	goutil.Assert(t, bytes[0] == 1, "wrong byte 0 value")
	goutil.Assert(t, bytes[1] == 1, "wrong byte 1 value")
}

func TestFromHexBytes(t *testing.T) {
	bytes := []byte("0101")
	goutil.Assert(t, len(bytes) == 4, "wrong byte length")
	bytes = FromHexBytes(bytes)
	goutil.Assert(t, len(bytes) == 2, "wrong byte length")
}

func TestStringFromHexString(t *testing.T) {
	hex := "0101"
	// 0101 --> 5
	str := StringFromHexString(hex)
	goutil.Assert(t, len(str) == 1,
		fmt.Sprintf("wrong string length: %d", len(str)))
	goutil.Assert(t, str == "5", fmt.Sprintf("wrong string value: %s", str))

}
