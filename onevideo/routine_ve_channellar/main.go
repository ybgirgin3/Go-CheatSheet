package main

import (
	"fmt"
	"sync"
)

/*

 */

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2) // 2 tane go routine bekle demek

	go func() {
		defer wg.Done()
		fmt.Println("f1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
	}()

	wg.Wait() // yukaridaki add kismindaki bekleme kismini 0'a donmesini bekliyor

	fmt.Println("end of the main")
}

// func main() {
// 	go f1()
// 	go f2()
// 	fmt.Println("end of the main")
// 	// time.Sleep(3 * time.Second)
// }
//
// func f1() {
// 	fmt.Println("f1")
// }
//
// func f2() {
// 	fmt.Println("f2")
// }
