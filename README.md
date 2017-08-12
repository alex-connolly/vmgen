# VMGen

An efp-based generator for virtual machines.


```go
name = "Example"
author = "[7][7][7]"
type = "VM"

instruction("ADD"){
    description = "Finds the sum of two numbers."
    pop = 2
    push = 1
    fuel = 100
}

instruction("SUB"){
    description = "Finds the difference between two numbers."
    pop = 2
    push = 1
    fuel = 100
}
```
