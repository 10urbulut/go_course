package functions

func Demo1(number int, number2 int) (int, int, int, float32) {
	toplam := number + number2
	fark := number - number2
	carpim := number * number2
	bolum := float32(number / number2)

	return toplam, fark, carpim, bolum

}
