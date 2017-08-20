package vmgen

// SetStack ...
func (vm *VM) SetStack(s *Stack) {
	vm.Stack = s
}

// Stack implementation
type Stack struct {
	items [][]byte
}

// Size ...
func (s *Stack) Size() int {
	return len(s.items)
}

// Push pushes bytes onto the stack
func (s *Stack) Push(items ...[]byte) {
	s.items = append(s.items, items...)
}

// Pop returns and removes the top byte array from the stack
func (s *Stack) Pop() []byte {
	values := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return values
}

func (s *Stack) back(num int) []byte {
	return s.items[s.Size()-num-1]
}

// Swap ...
// TODO: check evm implementation???
// TODO: swapping s.Size()-num with s.Size() - 1
func (s *Stack) Swap(num int) {
	s.items[s.Size()-1-num], s.items[s.Size()-1] = s.items[s.Size()-1], s.items[s.Size()-1-num]
}

// Dup ...
func (s *Stack) Dup(num int) {
	src := s.items[s.Size()-num:]
	dst := make([][]byte, len(src))
	copy(dst, src)
	s.Push(dst...)
}

func (s *Stack) validate(size int) bool {
	return len(s.items) >= size
}
