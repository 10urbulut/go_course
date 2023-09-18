package defer_statement

import "fmt"

type product struct {
	unitPrice   int
	productName string
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi: ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")

}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = product{productName: "Mouse", unitPrice: 5000}

	fmt.Println("İşlem başarılı")
	fmt.Println(p.productName)
}
