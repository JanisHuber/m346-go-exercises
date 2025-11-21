package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Blabla1",
		117: "Blabla2",
		346: "Cloud Computing",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[104] = "Blabla1"
	// TODO: replace one
	modules[117] = "Blabla2 - Replaced"
	fmt.Println(modules)
}
