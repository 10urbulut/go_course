package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("Tahmin 1 ile 100 arasında olmalı")
	}
	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı giriniz", nil

	}
	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı giriniz", nil

	}

	return "Bildiniz", nil
}

func Demo2() {
	fmt.Println(TahminEt(501))

}
