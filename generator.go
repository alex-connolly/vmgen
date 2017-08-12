package vmgen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/end-r/efp"
)

// VM ...
type VM struct {
	Name           string
	Author         string
	Receiver       string
	Instructions   map[string]instruction
	ProgramCounter int
	Current        instruction
	stats          *stats
	Stack          *Stack
	Memory         []interface{}
}

// Instruction for the current FireVM instance
type instruction struct {
	opcode       string
	description  string
	execute      string
	fuel         int
	fuelFunction string
	count        int
}

const prototype = "vmgen.efp"

// CreateVM creates a new FireVM instance
func CreateVM(path string) *VM {
	p, errs := efp.PrototypeFile(prototype)
	if errs != nil {
		fmt.Printf("Invalid prototype file\n")
		fmt.Println(errs)
		return nil
	}
	e, errs := p.ValidateFile(path)
	if errs != nil {
		fmt.Printf("Invalid VM file %s\n", path)
		return nil
	}
	var vm VM
	// no need to check for nil: would have errored
	vm.Author = e.FirstField("author").Value()
	vm.Name = e.FirstField("name").Value()
	vm.Receiver = e.FirstField("receiver").Value()

	vm.Instructions = make(map[string]instruction)
	for _, e := range e.Elements("instruction") {
		var i instruction
		i.description = e.FirstField("description").Value()

		// try to get fuel as an integer
		fuel, err := strconv.ParseInt(e.FirstField("fuel").Value(), 10, 64)

		// if not, it's a fuel function
		if err != nil {
			i.fuel = int(fuel)
		} else {
			i.fuelFunction = e.FirstField("fuel").Value()
		}

		i.execute = e.FirstField("execute").Value()

		opcode := e.Parameter(0).Value()
		i.opcode = opcode

		vm.Instructions[opcode] = i
	}

	vm.stats = new(stats)

	vm.Stack = new(Stack)
	return &vm
}

func (vm *VM) createParameters(params []string) []reflect.Value {
	var vals []reflect.Value
	vals = append(vals, reflect.ValueOf(vm))
	for _, p := range params {
		vals = append(vals, reflect.ValueOf(p))
	}
	return vals
}

func splitInstruction(text string) (string, []string) {
	all := strings.Split(text, " ")
	return all[0], all[1:]
}

// ExecuteFile parses opcodes from a file
func (vm *VM) ExecuteFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		opcode, params := splitInstruction(scanner.Text())
		vm.executeInstruction(opcode, params)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (vm *VM) executeInstruction(opcode string, params []string) {
	//i := vm.instructions[opcode]

}

func version() string {
	return fmt.Sprintf("%d.%d.%d", 0, 0, 1)
}
