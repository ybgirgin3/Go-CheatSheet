package main

import "fmt"

func main() {
	// var a int
	//
	// var p *int
	//
	// p = &a
	//
	// a = 10
	// fmt.Println(a, p, *p)

	// islemi pointer gondererek yaparsak fonksiyondan bir deger return etmemiz gerek yok
	// direkt olarak ramdeki degerini degistirdigimizden kalici bir degisiklik oluyor.
	// aksi halde degeri iki defa tanimlamis oluyoruz uzerinden degisiklik yapmak istedigimizde
	// var a int = 10
	// fmt.Println(a)
	// add12pointer(&a)
	// fmt.Println(a)

	// Bu array tanimlamasinda ise reference kullanmasak bile pass by value seklindeki tanimlamada da
	// biz arrayin uzerinde degisiklik yapabiliriz.
	// bunun sebebi normal integer degerlerinin ilkel; array, slice gibi metodlari
	// reference tipler diye gecer ve bunlar
	// uzerinde direkt olarak degisiklik yapilabilir
	var numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)

}

func changeValue(numbers []int) {
	numbers[0] = 1000
}

func add12(x int) int { // pass by value
	x = x + 12
	return x
}

func add12pointer(x *int) { // pass by reference
	*x = *x + 12
}
