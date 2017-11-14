package vmgen

// Bytecode is composed of commands
type Bytecode struct {
	commands []Command
}

// Command ...
type Command struct {
	mnemonic   string
	parameters []byte
}

// Length ...
func (b *Bytecode) Length() int {
	return len(b.commands)
}

// AddCommand to the current bytecode
func (b *Bytecode) AddCommand(c Command) {
	if b.commands == nil {
		b.commands = make([]Command, 0)
	}
	b.commands = append(b.commands, c)
}

// Add ...
func (b *Bytecode) Add(mnemonic string, parameters ...byte) {
	c := Command{
		mnemonic:   mnemonic,
		parameters: parameters,
	}
	b.AddCommand(c)
}

// Concat another bytecode struct onto ours
func (b *Bytecode) Concat(other *Bytecode) {
	if other == nil || other.commands == nil {
		return
	}
	b.commands = append(b.commands, other.commands...)
}
