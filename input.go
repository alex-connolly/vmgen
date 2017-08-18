package vmgen

// Input ...
type Input interface {
	Code() *Byterable
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

// Code ...
func (b *BasicInput) Code() *Byterable {
	if b.code == nil {
		b.code = new(Byterable)
	}
	return b.code
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

func (b *Byterable) Size() int {
	return len(b.Data)
}

func (b *Byterable) Append(bytes ...byte) {
	if b.Data == nil {
		b.Data = make([]byte, 0)
	}
	b.Data = append(b.Data, bytes...)
}
