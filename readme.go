package firevm

import "os"

func (vm *FireVM) GenerateReadme(name string) {
	f, err := os.Create(name)
	if err != nil {
		return
	}
	defer f.Close()
    f.Write([]byte("| " + "Opcode"))
    f.Write([]byte("| " + "Description"))
    f.Write([]byte("| " + "Fuel"))
    f.Write("|\n")
	for _, v := range vm.instructions {
        f.Write([]byte("| " + v.opcode)
        f.Write([]byte("| " + v.description)
        f.Write([]byte("| " + v.fuel))
        f.Write([]byte("|\n"))
	}
    f.Sync()
}
