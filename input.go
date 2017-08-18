package vmgen

// Input ...
type Input interface {
	Code() *Byterable
	SetCode(*Byterable)
	/*Input() *Byterable
	Address(string) Address
	Bytes(string) []byte
	Hash(string) Hash
	Offset() int
	Next(int) []byte
	Value() *big.Int*/
}

// BasicInput ...
type BasicInput struct {
	code *Byterable
}

func (b *BasicInput) Code() *Byterable {
	return b.code
}

func (b *BasicInput) SetCode(by *Byterable) {
	b.code = by
}

// Byterable = Bytes + Iterable
type Byterable struct {
	Data   []byte
	Offset int
}

// NewByterable ...
func NewByterable(bytes []byte) *Byterable {
	return &Byterable{
		Data:   bytes,
		Offset: 0,
	}
}

// Next ...
func (b *Byterable) Next(size int) []byte {
	b.Offset += size
	return b.Data[b.Offset-size : b.Offset]
}

// HasNext ...
func (b *Byterable) HasNext() bool {
	return b.Offset < len(b.Data)
}
