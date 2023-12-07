package main

import "fmt"

func main() {
	// var sumValue = add(3, 4)
	// sumValue := add(3, 4)
	// fmt.Println("sumValue: ", sumValue)

	// total, diff := calculation(10, 20)
	// fmt.Println(total, diff)

	// var numbers = []int{1, 2, 3, 4, 5, 6}
	// fmt.Println("numbers", sum(numbers))

	fmt.Println("numbers", sum(1, 2, 3, 4, 5, 6))
}

func sum(numbers ...int) int { // istedigim kadar param verebilirim demek verilen tum bireysel int degerlerini bir array haline getiricek ve icinde array olarak islem yapmamiza izin verecek
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// func sum(numbers []int) int {
// 	total := 0
// 	for _, num := range numbers {
// 		total += num
// 	}
// 	return total
// }

func calculation(x int, y int) (int, int) {
	return x + y, x - y
}

// func add(x int, y int) int {
// 	return x + y
// }
