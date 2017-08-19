package vmgen

import "os"

// GetFileBytes ...
func GetFileBytes(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return nil
	}
	bytes := make([]byte, fi.Size())
	_, err = f.Read(bytes)
	if err != nil {
		return nil
	}
	// if it ends in a newline, remove it
	if bytes[len(bytes)-1] == 10 {
		bytes = bytes[:len(bytes)-1]
	}
	return bytes
}
