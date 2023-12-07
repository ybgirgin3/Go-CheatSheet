package main

import "fmt"

func main() {
	// index := 1

    // simple for loop
	for index := 1; index <= 10; index++ {
		fmt.Println("current value of index", index)
	}

    // iterate over a array
	index := 0

	var names = [3]string{"Yusuf", "Berkay", "Girgin"}

	for index < len(names) {
		fmt.Println(names[index])
		index++
	}

    // break continue
    for index := 0; index <= 10; index++ {
        fmt.Println(index)
        if index == 3 {
            fmt.Println("loop is dead")
            break
        }
    }

    for index := 0; index <= 10; index++ {
        if index > 3 {
            fmt.Println("loop is started in ", index)
            continue
        }
    }
}
