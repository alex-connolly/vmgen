package vmgen

// SetStack ...
func (vm *VM) SetStack(s *Stack) {
	vm.Stack = s
}

// Stack implementation
type Stack struct {
	data [][]byte
}

// Size ...
func (s *Stack) Size() int {
	return len(s.data)
}

// Push pushes bytes onto the stack
func (s *Stack) Push(data []byte) {
	s.data = append(s.data, data)
}

// Pop returns and removes the top byte array from the stack
func (s *Stack) Pop() []byte {
	values := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return values
}

func (s *Stack) validate(size int) bool {
	return len(s.data) >= size
}
