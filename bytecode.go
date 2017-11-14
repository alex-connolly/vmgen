package vmgen

import (
	"fmt"
)

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
func (b *Bytecode) Concat(other Bytecode) {
	if other.commands == nil {
		return
	}
	b.commands = append(b.commands, other.commands...)
}

// Compare to another bytecode
func (b *Bytecode) Compare(other Bytecode) bool {
	if b.commands == nil && other.commands == nil {
		return true
	}
	if other.commands == nil || b.commands == nil {
		return false
	}
	if len(b.commands) != len(other.commands) {
		return false
	}
	for i, c := range b.commands {
		o := other.commands[i]
		if c.mnemonic != o.mnemonic {
			return false
		}
		for j, p := range c.parameters {
			if p != o.parameters[j] {
				return false
			}
		}
	}
	return true
}

// Format bytecode for output
func (b *Bytecode) Format() string {
	if b.commands == nil {
		return "No commands"
	}
	s := fmt.Sprintf("%d commands\n", len(b.commands))
	for i, c := range b.commands {
		// have to do this manually
		s += fmt.Sprintf("%d | %s %v\n", i+1, c.mnemonic, c.parameters)
	}
	return s
}
