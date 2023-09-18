package channels

import (
	"fmt"
)

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Çift sayı: ", i)

	}
	ciftSayiCn <- toplam

}
func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Tek sayı: ", i)

	}
	tekSayiCn <- toplam
}
