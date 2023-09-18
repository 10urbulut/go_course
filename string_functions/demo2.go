package string_functions

import (
	"fmt"
	"strings"
)

func Demo2() {
	isim := "Onur"

	fmt.Println(strings.HasPrefix(isim, "On"))
	fmt.Println(strings.HasSuffix(isim, "ur"))
	fmt.Println(strings.Index(isim, "ur"))

	harfler := []string{"o", "n", "u", "r"}

	fmt.Println(strings.Join(harfler, " "))

	fmt.Println(strings.Replace(isim, "", "+", 2))

	fmt.Println(strings.Split(isim, ""))

	fmt.Println(strings.Repeat(isim, 4))

}
