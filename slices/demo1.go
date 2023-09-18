package slices

import "fmt"

func Demo1() {
	names := make([]string, 3)

	names[0] = "Onur"
	names[1] = "Buket"
	names[2] = "Elif"
	fmt.Println(names)
	names = append(names, "Fatma")
	namesTwo := make([]string, len(names))
	copy(namesTwo, names)
	fmt.Println(names)
	fmt.Println(namesTwo[:3])

}
