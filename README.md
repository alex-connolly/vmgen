# VMGen

An [efp-based](https://www.github.com/end-r/efp) generator for virtual machines, initially developed for use with [FireVM](https://www.github.com/end-r/firevm). 

## Example


In ```example.vm```:

```go
name = "Example"
author = "[7][7][7]"
receiver = "VM"

instruction("ADD"){
    description = "Finds the sum of two numbers."
    pop = 2
    push = 1
    fuel = 100
}

instruction("PUSH"){
    hex = 0
    description = "Pushes a number onto the stack."
    push = 1
    fuel = 30
}
```

We could either use IR bytecode in ```example.fire```:

```go
PUSH 1
PUSH 2
ADD
PUSH 5
ADD
```

Or a fully compiled version:

```go
0x010201010201020102010502
```

Now, our Go program:

```go
package main

import (
    "github.com/end-r/vmgen"
    "math/big"
)

const fuel = 1000

var (
    fuels = map[string]vmgen.FuelFunction{

    }
    executes = map[string]vmgen.ExecuteFunction{
        "ADD": Add,
        "PUSH": Push,
    }
)

func main(){
    vm := vmgen.CreateVM("example.vm", executes, fuels)
    vm.ExecuteFile(fuel, "example.fire")
    vm.Stats()
}

func Add(vm *vmgen.VM){
    a := vm.Stack.Pop()
    b := vm.Stack.Pop()
    c := new(big.Int).Add(a, b)
    vm.Stack.Push(c)
}

func Push(vm *vmgen.VM){
    size := vm.Contract.Next(1)
    value := vm.Contract.Next(size)
    vm.Stack.Push(value)
}
```

## Instruction Syntax

See ```vmgen.efp``` for the prototype file, below is a full example instruction:

```go
instruction("OPCODE", "0x33"){
    fuel = 100
}
```

## Fuel

vmgen was built as a generator for costly virtual machines, where each instruction is given a fixed or variable cost and 'charged' against an initial balance, preventing infinite loops and sidestepping the halting problem. The fuel for an instruction can be provided in one of two ways:

1. By assigning an unsigned integer to the fuel field:

```fuel = 100```

2. By providing a fuel function mapping:

```go
instruction("HI", "0x40")
```

```go
fuels = map[string]vmgen.FuelFunction{
    "HI": getFuel,
}

getFuel(vm *vmgen.VM) int{

}
```

If neither of these are set, the instruction will have its fuel cost set to 0.

## Disassembly

vmgen provides support for generalised disassembly.

```go
DisasmString(data string)
DisasmBytes(bytes byte)
DisasmFile(path string)
```

Generally, instructions will be printed in the following format:

| 0x01 | ADD |

By using the ```disasm = customDisasm``` field in the ```.vm``` file, you can assign a mapped disassembly function to an instruction.

In your go code:

```go
disasms = map[string]vmgen.DisasmFunction{
    "customDisasm": customDisasm,
}

func customDisasm(offset int, bytecode []byte)([]string, int){
    // everything will be in format | o1 | o2 | o3 |
    // return []string{o1, o2, o3}, totalOffsetChange
    // this is done for pretty formatting reasons
}
```

Thus, you can produce output like:

| PUSH | 0x02 | 0xFF00 |
