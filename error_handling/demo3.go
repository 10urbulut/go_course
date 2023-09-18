package error_handling

import (
	"fmt"
)

type borderException struct {
	message   string
	parameter int
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d --- %s", b.parameter, b.message)

}
func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{parameter: tahmin, message: "Sınırların dışında kaldı"}
	}
	return "Bildiniz", nil
}

func Demo3() {
	fmt.Println(TahminEt2(300))
}
