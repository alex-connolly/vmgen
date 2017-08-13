package vmgen

import "math/big"

// Contract ...
type Contract struct {
	// CallerAddress is the result of the caller which initialised this
	// contract. However when the "call method" is delegated this value
	// needs to be initialised to that of the caller's caller.
	Caller Address
	caller Address
	self   Address

	//jumpdests destinations

	Code     []byte
	CodeHash Hash
	CodeAddr *Address
	Input    []byte

	Gas   uint64
	value *big.Int

	Args []byte

	DelegateCall bool
}

// Next @size bytes
func (c *Contract) Next(size int) []byte {
	return nil
}
