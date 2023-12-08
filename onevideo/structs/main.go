package main

import "fmt"

func main() {
	var c1 = Customer{id: 1, name: "Bekocan", age: 25, address: Address{city: "Sakarya"}}
	c1.ToString()
	c1.changeName("Yusuf Berkay Girgin")
	c1.ToString()

	// fmt.Println("customer1: ", c1)
	//
	// c1.age = 26
	// fmt.Println("customer1: ", c1)

}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city string
}

// pythondaki ve typescriptte classlarin icinde tanimladigimiz ve self veya this ile birbirine bagladigimiz metodlari
// golang'da bu sekilde veriyoruz. direkt olarak alistigimiz oop yapisi mevcut degil
func (customer *Customer) ToString() { // bu customer struct'inin bir fonksiyonu
	fmt.Printf("Name: %s, Age: %d", customer.name, customer.age)
}

// structlar icinde farkli bir yerden cagirilan ve degisiklik yapacak olan metodlar eger pointer ile gonderilmez ise
// bir degisiklik yapamaz.
// dolayisiyla pass by reference kullanarak yapmaliyiz
// bu sekilde tanimlama pek tercih edilmez
// func changeName(c *Customer) {
// 	c.name = "Yusuf Berkay Girgin"
// }

// bu kisimda customer'in bir fonksiyonu ama yine degerine kendisi kopyalandigi icin
// degisiklik yapamayacak
// yani yine pointer kullanmamiz lazim
// yukaridaki gibi sadece print kullanacagimiz bir sekilde tanimlama yaptiysak
// o zaman reference vermeye gerek yok direkt olarak value gonderebiliriz
// func (c Customer) changeName(newName string) {
func (c *Customer) changeName(newName string) {
	c.name = newName
}
