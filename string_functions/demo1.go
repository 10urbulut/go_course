package string_functions

import (
	"fmt"
	"strings"
)

func Demo1() {
	isim := "Onur"
	fmt.Println(strings.Count(isim, "O"))
	fmt.Println(strings.Contains(isim, "O"))
	fmt.Println(strings.Index(isim, "u"))

	sonuc := strings.Index(isim, "u")

	if sonuc != -1 {
		fmt.Println("Bulundu")
	} else {
		fmt.Println("Yok")
	}

	fmt.Println(strings.ToLower("onur"))
	fmt.Println(strings.ToUpper("onur"))

}
