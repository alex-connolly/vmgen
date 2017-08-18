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

// StringFromHexBytes ...
func StringFromHexBytes(bytes []byte) string {
	bytes = FromHexBytes(bytes)
	return string(bytes)
}

// StringFromHexString ...
func StringFromHexString(str string) string {
	return StringFromHexBytes([]byte(str))
}
