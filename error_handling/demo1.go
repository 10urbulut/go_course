package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	file, error := os.Open("error_handling/demo1.txt")

	if error != nil {

		if pathError, ok := error.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı: ", pathError.Error())
			return
		}

		fmt.Println("Dosya bulunamadı.", error.Error())
		return
	}
	fmt.Println(file.Name())

}
