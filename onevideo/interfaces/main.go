package main

import "fmt"

// PART 1

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

func (b *Book) ShippingCost() int {
	return 5 + b.desi*2
}

func (e *Electronic) ShippingCost() int {
	return 10 + e.desi*3
}

func calculateTotalShippingCost(product []IShippable) int {
	total := 0
	for _, prod := range product {
		total += prod.ShippingCost()
	}
	return total
}

func _main() {
	// book := &Book{desi: 10}
	// fmt.Println(book.ShippingCost())
	// books := []Book{{desi: 10}, {desi: 20}}
	// fmt.Println(calculateTotalShippingCost(books))

	var products = []IShippable{&Book{desi: 10}, &Book{desi: 20}, &Electronic{desi: 30}}
	fmt.Println(calculateTotalShippingCost(products))

	// product = &Book{desi: 10}
	// fmt.Println(product.ShippingCost())
	// product = &Electronic{desi: 10}
	// fmt.Println(product.ShippingCost())
}

// PART 2

// CALISMAYAN

// type XEncoder struct {
// }
// type YEncoder struct {
// }
//
// func (x *XEncoder) Encode(value string) {
// 	fmt.Println("XEncoder ile encode edildi")
// }
// func (y *YEncoder) Encode(value string) {
// 	fmt.Println("YEncoder ile encode edildi")
// }
//
// func (x *XEncoder) Decode(value string) {
// 	fmt.Println("XEncoder ile decode edildi")
// }
// func (y *YEncoder) Decode(value string) {
// 	fmt.Println("YEncoder ile decode edildi")
// }
//
// func main() {
// 	var encoderX *XEncoder
// 	var encoderY *YEncoder
//
// 	encoderX = &XEncoder{}
// 	encoderY = &YEncoder{}
//
// 	encoderX.Encode("123234")
// 	encoderY.Encode("123234")
//
// }

// CALISAN

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}
type XEncoder struct {
}
type YEncoder struct {
}

func (x *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi")
}
func (y *YEncoder) Encode(value string) {
	fmt.Println("YEncoder ile encode edildi")
}

func (x *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi")
}
func (y *YEncoder) Decode(value string) {
	fmt.Println("YEncoder ile decode edildi")
}

func main() {
	// bu degiskenleri direkt olarak struct kullanarak tanimladigimizda
	// struct yapisini miras alan fonksiyonlari kullanabilmek icin kod icin de komple tipleri degistirmemiz gerekecekti
	// fakat interface kullanip bu islem yapildiginda ve degisken tipi olarak interface verildiginde
	// encode ve decode metodlarini direkt olarak cagirabiliriz
	var encoderX IEncoder
	var encoderY IEncoder

	encoderX = &XEncoder{}
	encoderY = &YEncoder{}

	encoderX.Encode("123234")
	encoderY.Encode("123234")
}
