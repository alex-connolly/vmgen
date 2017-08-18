package vmgen

var (
	executes = map[string]ExecuteFunction{
		"ADD":  add,
		"PUSH": push,
		"TEST": test,
	}
)

func add(vm *VM) {

}

func test(vm *VM) {
	vm.Stack.Push([]byte("a"))
	vm.Stack.Push([]byte("b"))
	vm.Stack.Push([]byte("c"))
}

func push(vm *VM) {
	vm.Stack.Push([]byte("a"))
}
