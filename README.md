# VMGen

An [efp-based](https://www.github.com/end-r/efp) generator for virtual machines, initially developed for use with [FireVM](https://www.github.com/end-r/firevm)


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

We store out bytecode in ```example.fire```

```go
PUSH 1
PUSH 2
ADD
PUSH 5
ADD
```

Now, our Go program:

```go
package vmgen

const fuel = 1000

func main(){
    vm := ParseFile("example.vm")
    vm.Run(fuel, "example.fire")
    vm.Stats()
}
```
