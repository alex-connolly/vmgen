package vmgen

import (
	"encoding/binary"
	"fmt"
)

// Bytecode is composed of commands
type Bytecode struct {
	commands []*Command
}

// Command ...
type Command struct {
	mnemonic   string
	parameters []byte
	isMarker   bool
	offset     int
}

type CostFunc func(interface{}) int

type Instruction struct {
	Opcode uint
	Cost   CostFunc
}

type InstructionMap map[string]Instruction

func (im InstructionMap) AddAll(m InstructionMap) {
	if im == nil {
		im = make(InstructionMap)
	}
	for k, v := range m {
		im[k] = v
	}
}

// Length ...
func (b *Bytecode) Length() int {
	return len(b.commands)
}

// AddCommand to the current bytecode
func (b *Bytecode) AddCommand(c *Command) {
	if b.commands == nil {
		b.commands = make([]*Command, 0)
	}
	b.commands = append(b.commands, c)
}

// Add ...
func (b *Bytecode) Add(mnemonic string, parameters ...byte) {
	c := Command{
		mnemonic:   mnemonic,
		parameters: parameters,
		isMarker:   false,
	}
	b.AddCommand(&c)
}

// AddMarker ...
func (b *Bytecode) AddMarker(mnemonic string, offset int) {
	c := Command{
		mnemonic: mnemonic,
		offset:   offset,
		isMarker: true,
	}
	b.AddCommand(&c)
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

// Finalise ...
func (b *Bytecode) Finalise() {
	for i, c := range b.commands {
		if c.isMarker {
			c.parameters = make([]byte, 8)
			binary.LittleEndian.PutUint64(c.parameters, uint64(c.offset+i))
		}
	}
}

// CompareMnemonics ...
func (b *Bytecode) CompareMnemonics(test []string) bool {
	if len(test) != len(b.commands) {
		return false
	}
	for i, t := range test {
		if t != b.commands[i].mnemonic {
			return false
		}
	}
	return true
}

type Generator interface {
	Instructions() InstructionMap
}

func (b *Bytecode) Generate(g Generator) []byte {
	bytes := make([]byte, 0)
	is := g.Instructions()
	for _, c := range b.commands {
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(is[c.mnemonic].Opcode))
		bytes = append(bytes, bs...)
		bytes = append(bytes, c.parameters...)
	}
	return bytes
}
