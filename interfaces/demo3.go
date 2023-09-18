package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)

	fmt.Println(sayi, ok)
}

func Demo3() {
	dogrula("df")
	dogrula(5)
}
