package main

import "fmt"

// maps dict ile ayni sey
// verileri key value seklinde tutuyor

func main() {
    /* isimler ve onlari icin bir int deger tanimlayacagiz
       pythondaki sekliyle: names: {'bekocan': 1, 'girgin': 2}
    */

    // TANIMLAMA 1
    // var names map[string]int // isim string ona denk gelecek olan deger int
    //
    // names = make(map[string]int, 0)
    //
    // names["Yusuf"] = 1
    // names["Berkay"] = 2
    // names["Girgin"] = 3

    // TANIMLAMA 2
    names := map[string]int{
        "Yusuf": 1,
        "Berkay": 2,
        "Girgin": 3,
    }

    fmt.Println(names) // hepsini yazdir
    fmt.Println(names["Berkay"]) // sadece Berkay'in degerini yazdir
    fmt.Println(names["Bekocan"]) // olmayan bir degeri yazdir bu kisimda, pythondaki gibi hata vermez direkt olarak default 0 degerini doner

    // bir key sil
    delete(names, "Yusuf")

    fmt.Println(names) // hepsini yazdir



}
