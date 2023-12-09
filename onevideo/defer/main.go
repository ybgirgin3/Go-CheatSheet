package main

import "fmt"

// ertelemek anlamina geliyor
// fonksiyon bittiginde bu defer ile cagirdigimiz statement kullaniliyor.

// programlarda temizleme islemi yapacagimiz zaman
// herhangi error durumunda db baglantisinin kapatilmasi
// dosyanin kapatilmasi
// arraylarin ramden temizlenmesi gibi durumlarin, bir sekilde islem bittikten sonra yapilmasi gerektiginden dolayi
// `defer` bu kisimlarda en cok kullanilan anahtar oluyor.

func main() {
	// // output:
	// // Worldd
	// // Hellooww
	//
	// defer fmt.Println("Hellooww")
	// fmt.Println("Worldd")
	//
	// //  ------------------------------
	//
	// /*
	// 	ilk ertelenen en son gelir
	// 	ilk giren son cikar mantigi var
	// 		main fonksiyonu
	// 		3. defer
	// 		2. defer
	// 		1. defer
	// */
	// defer fmt.Println("1. defer")
	// defer fmt.Println("2. defer")
	// defer fmt.Println("3. defer")
	//
	// fmt.Println("main fonksiyonu")
	//
	// //  ------------------------------
	//
	// // condition  true oldugundan dolay if statement icinde fonk bitti zaten
	// // yani ikinci defer ile karsilasmadi bile dolayisiyla o calismaz. kuyruga dahi alinmaz
	//
	// var condition = true
	// defer fmt.Println("1. Defer")
	// if condition {
	// 	return
	// }
	// defer fmt.Println("2. Defer")
	//
	// //  ------------------------------
	//
	// // ONEMLI: defer degeri en son calisiyor ama defer tanimlandigi anda ram'de tutulan degisken degerleri ne ise onu tutarlar
	// // yani asagidaki kod parcasinin ikinci kisminda x ve y'nin degerleri degistirilmis olmasina ragmen
	// // ilk degerleri ile defer kuyruguna alindigi icin degisen degerlerden etkilenmez
	// x := 10
	// y := 20
	//
	// defer fmt.Println("x: ", x)
	// defer fmt.Println("y: ", y)
	//
	// x = 100
	// y = 100
	// fmt.Println(x)
	// fmt.Println(y)

	// 	----
	var condition2 = true

	defer cleanup() // defer program panic atsa bile calistirilir. diger standard metodlar calistirilmaz

	if condition2 {
		panic("An error occurred") // works like error raising in python
	}

}

func cleanup() {
	fmt.Println("cleanup worked")
}
