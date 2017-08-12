# VMGen

An [efp-based](https://www.github.com/end-r/efp) generator for virtual machines, initially developed for use with [FireVM](https://www.github.com/end-r/firevm).


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

import "github.com/end-r/vmgen"

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

func Add(vm *vmgen.VM, params []byte){
    a := vm.Stack.Pop()
    b := vm.Stack.Pop()
    c := a.Add(b)
    vm.Stack.Push(c)
}

func Push(vm *vmgen.VM, params []byte){
    size := vm.Input.Next(1)
    value := vm.Input.Next(size)
    vm.Stack.Push(c)
}
```
