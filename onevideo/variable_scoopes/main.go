package main

import "fmt"

var x = 10 // bu degisken global scoop'ta dolayisiyla istedigimiz her yerde kullanabiliriz bunu

// degiskenlerin tanimlandigi yerlerin disinda kullanilmasini saglayan scooplar
func main() {

	var condition = true

	if condition {
		// burda tanimladigimiz x degiskenini
		// asagida print edemeyiz buna block scoop denir
		// sadece tanimlandigi scoop icinde kullanilabilir
		var x = 10
		fmt.Println(x)
	}
	fmt.Println(x)
}

func test() {
	// function scoop
	// burda tanimlanmis olan degerler baska hicbir yerde kullanilamaz sadece fonksiyon icinde kullanilir
	var x = 10
	var y = 10
	fmt.Println(x, y)
}
