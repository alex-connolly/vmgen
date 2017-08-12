package vmgen

// Stack ...
type Stack interface {
	Push([]byte)
	Pop(uint)
	Size() uint
}

// SetStack ...
func (vm *VM) SetStack(s *Stack) {
	vm.Stack = s
}

// DefaultStack implementation
type DefaultStack struct {
	data []byte
}

// Push pushes bytes onto the stack
func (s *DefaultStack) Push(data []byte) {
	s.data = append(s.data, data...)
}

// Pop returns and removes the first @size bytes from the stack
func (s *DefaultStack) Pop(size int) []byte {
	if s.validate(size) {
		values := s.data[len(s.data)-1-size:]
		s.data = s.data[:len(s.data)-1-size]
		return values
	}
	return nil
}

func (s *DefaultStack) validate(size int) bool {
	return len(s.data) > size
}
