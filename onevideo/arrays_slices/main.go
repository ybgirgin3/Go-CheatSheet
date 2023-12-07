package main

import "fmt"

func main() {
	// ARRAYS
	// Fixed
	// tanimlanmis olan array'e yeni eleman ekleyemeyiz
	// var olanlar uzerinden degisiklik yapilabilir fakat yeni +1. indexe eleman eklenemez
	// koseli parantez icinde ne verdiysek o sayi kadar eleman ekleyebiliriz

	// var names = [3]string{"Yusuf", "Berkay", "Girgin"}
	// names[1] = "Bekocan"
	// fmt.Println(names)
	// fmt.Println(names[0:2])

	// slicing
	var namesWithSlice = []string{"Yusuf", "Berkay", "Girgin"}
	namesWithSlice = append(namesWithSlice, "Nigga")
	fmt.Println(namesWithSlice)

}
