package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		fmt.Println("Çift")
		return "Çift Sayı"
	}
	if sayi%2 != 0 {
		fmt.Println("Tek")
		return "Tek Sayı"
	}
	return "Belirsiz"
}
func Test() {
	Demo2(9)

}
func DeferFunc() {
	fmt.Println("Bitti")

}
