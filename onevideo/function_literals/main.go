package main

import (
	"fmt"
)

func main() {
	// f1()

	func() {
		fmt.Println("f1")
	}()

	func(x int, y int) {
		fmt.Println(x + y)
	}(3, 5) // pythondaki lambda ile ayni sey

	add := func(x int, y int) int {
		return x + y
	} // pythondaki lambda ile ayni sey

	multiply := func(x int, y int) int {
		return x * y
	} // pythondaki lambda ile ayni sey

	// fmt.Println(add(3, 6), reflect.TypeOf(add))

	ret := calculator(3, 5, add)
	fmt.Println(ret)

	ret2 := calculator(3, 5, multiply)
	fmt.Println(ret2)

}

// func f1() {
// 	fmt.Println("f1")
// }

// burda tanimladigimiz gibi fonksiyonlar iclerine int, float gibi degerlerin yani sira
// fonksiyonu da parametre olarak alabilirler
func calculator(x int, y int, f1 func(int, int) int) int {
	return f1(x, y)
}
