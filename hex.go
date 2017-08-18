package vmgen

import "encoding/hex"

// FromHexBytes ...
func FromHexBytes(bytes []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(bytes)))
	_, err := hex.Decode(dst, bytes)
	if err != nil {
		return nil
	}
	return dst
}

// FromHexString ...
func FromHexString(str string) []byte {
	return FromHexBytes([]byte(str))
}
