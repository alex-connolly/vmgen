package vmgen

// State ...
type State interface {
	getBytes() map[string][]byte
}
