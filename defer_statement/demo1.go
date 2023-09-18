package defer_statement

import "fmt"

func A() {
	fmt.Println("A çalıştı")

}
func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("B çalıştı")

}
func C() {

	fmt.Println("C çalıştı")

}
func D() {

	fmt.Println("D çalıştı")

}
