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

// Pop returns and removes the first items from the stack
func (s *Stack) Pop(size int) [][]byte {
	if s.validate(size) {
		values := s.data[len(s.data)-size:]
		s.data = s.data[:len(s.data)-size]
		return values
	}
	return nil
}

func (s *Stack) validate(size int) bool {
	return len(s.data) >= size
}
