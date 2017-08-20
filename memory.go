package vmgen

// Memory ...
type Memory interface {
	DisplayContents()
	Size() int
	Usage() int
}
