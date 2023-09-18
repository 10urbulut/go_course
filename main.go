package main

import (
	"golesson/defer_statement"
)

func main() {

	// ciftSayiToplamCn := make(chan int)
	// tekSayiToplamCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiToplamCn)
	// go channels.TekSayilar(tekSayiToplamCn)
	// ciftSayi, tekSayi := <-ciftSayiToplamCn, <-tekSayiToplamCn
	// println("Çarpım: ", (ciftSayi * tekSayi))

	defer_statement.Demo3()

}
