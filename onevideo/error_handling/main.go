package main

import (
	"fmt"
)

func main() {
	// var x int // primitive/ilkel degerleri go direkt olarak 0 olarak tanimliyor
	// var y float32
	// var pointer1 *int
	//
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(pointer1)
	//
	// if pointer1 == nil {
	// 	fmt.Println("pointer herhangi bir degeri tasimiyor")
	// }

	// fileData, err := os.ReadFile("sample.txt") // eger dosya yoksa direkt olarak error doner
	// // fmt.Println("filedata, err", fileData, err)
	// if err != nil {
	// 	fmt.Println("error while reading file: ", err)
	// } else {
	// 	fmt.Printf("filedata: %c", fileData)
	// }

	// result, err := divide(10, 0)
	// fmt.Println("result: ", result)
	// fmt.Println("error: ", err)

	productService := ProductService{}
	err := productService.Add(Product{id: 1, name: "", price: 3000})
	if err != nil {
		fmt.Println(err)
	}
}

// func divide(x int, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("Payda sifir olamaz")
// 	}
// 	return x / y, nil
// }

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		// return errors.New("urun ismi bos olmaz")
		return ValidationError{text: "Urun ismi bos olamaz", code: "1001"}
	}

	if product.price == 0 {
		return ValidationError{text: "urun fiyati bos olamaz", code: "1002"}
	}
	fmt.Println("urun eklendi")
	return nil
}

type ValidationError struct {
	code string
	text string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Hata :%s, Kod: %s", v.text, v.code)
}
