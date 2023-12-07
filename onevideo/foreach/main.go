package main

import "fmt"

func main() {
	// var numbers = []int{1, 2, 3, 4}

	// for index := 0; index < len(numbers); index++ {
	//     fmt.Println(numbers[index])
	// }
	//

	// for _, value := range numbers {
	// 	fmt.Println(value)
	// }

	// var lang = "Golang"
	//
	// for _, char := range lang {
	// 	fmt.Printf("%c\n", char)
	// }

	var names = map[string]int{
		"Yusuf":  10,
		"Berkay": 20,
		"Girgin": 30,
	}

	for key, val := range names {
		fmt.Println(key, val)
	}
}
