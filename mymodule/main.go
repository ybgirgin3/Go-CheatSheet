package main

import (
	"fmt"

	"mymodule/helper"
	"mymodule/helper/rest"
	restutils "mymodule/utils/rest" // ayni isimdeki moduller farki dizinlerden cekmek istersek sistem
	// bize ayni isimde iki tane module cagiramazsin diye uyari verir
	// bunu ortadan kaldirmak icin pythondaki `as` gibi calisacak sekilde
	// module basina module vermek istedigimiz adi veririz.
	// -----
	// bunu yapmamizi saglayan seyde diger paketteki cagirdigimiz fonksiyonun buyuk harf ile baslamasidir.
	// eger kucuk harf ile baslatirsak dosyamizi disariya cikmaz. import edilemez
)

func main() {
	fmt.Println("Hello World")
	helper.Helper1()
	rest.RestHelper()
	rest.privateHelper() // private oldugu icin kullanilamiyor. (kucuk harf ile baslattik)
	restutils.RestUtil()
}
