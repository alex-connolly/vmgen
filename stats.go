package vmgen

import (
	"fmt"
	"sort"
)

type stats struct {
	operations      int
	fuelConsumption int
}

// Stats ...
func (vm *VM) Stats() {
	fmt.Printf("VM %s\n", version())
	fmt.Printf("%s - %s\n", vm.Name, vm.Author)
	fmt.Printf("Operations Executed: %d\n", vm.stats.operations)
	fmt.Printf("Fuel Used: %d\n", vm.stats.fuelConsumption)
	fmt.Printf("Fuel/operation: %f\n", float64(vm.stats.operations)/float64(vm.stats.fuelConsumption))
}

func (vm *VM) sortByFuelConsumption() []*instruction {
	il := make(instructionList, len(vm.Instructions))
	i := 0
	for _, v := range vm.Instructions {
		il[i] = v
		i++
	}
	sort.Sort(sort.Reverse(il))
	return il
}

type instructionList []*instruction

func (l instructionList) Len() int { return len(l) }
func (l instructionList) Less(i, j int) bool {
	return (l[i].fuel * l[i].count) < (l[j].fuel * l[j].count)
}
func (l instructionList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// DetailedStats ...
func (vm *VM) DetailedStats() {
	si := vm.sortByFuelConsumption()
	for i, op := range si {
		fmt.Printf("| %d ", i+1)
		fmt.Printf("| %s ", op.opcode)
		fmt.Printf("| %d", op.count)
		fmt.Printf("| %d", op.fuel*op.count)
		fmt.Printf("|")
	}
}
