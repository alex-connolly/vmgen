package vmgen

// Memory interface represent an abstract memory location
type Memory interface {
	Set()
	Get()
	Size()
}
