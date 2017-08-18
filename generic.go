package vmgen

// Address ...
type Address interface {
	Bytes() []byte
}

// Hash ...
type Hash interface {
	Bytes() []byte
}
