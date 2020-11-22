package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Part 31 Channel")
	fmt.Println("")
	runtime.GOMAXPROCS(2)

	fmt.Println("Part 31.1 Penerapan channel")
	channel01 := make(chan string)
	var func01 = func(s string) {
		result := fmt.Sprintf("Hello: %v", s)
		channel01 <- result
	}
	go func01("Real Madrid")
	go func01("Barcelona")
	go func01("Atletico Madrid")
	var channel01A = <-channel01
	var channel01B = <-channel01
	var channel01C = <-channel01
	fmt.Printf("channel01A: %v\n", channel01A)
	fmt.Printf("channel01B: %v\n", channel01B)
	fmt.Printf("channel01C: %v\n", channel01C)
	fmt.Println("")

	fmt.Println("Part 31.2 Channel sebagai tipe data parameter")
	channel02 := make(chan string)
	for _, v := range []string{"Juventus", "AC Milan", "InterMilan"} {
		go func(club string) {
			channel02 <- club
		}(v)
	}
	channel02A := <-channel02
	channel02B := <-channel02
	channel02C := <-channel02
	fmt.Printf("channel02A: %v\n", channel02A)
	fmt.Printf("channel02B: %v\n", channel02B)
	fmt.Printf("channel02C: %v\n", channel02C)
	fmt.Println("")

	fmt.Println("Part 31.3 Tambahan")
	channel03 := make(chan balok)
	for i := 0; i < 3; i++ {
		go func(i int) {
			result := balok{
				panjang: (i + 2) * (i * 9),
				lebar:   (i + 3) * (i * 5),
				tinggi:  (i + 8) * (i * 2),
			}
			channel03 <- result
		}(i)
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("Balok %v: %v\n", (i + 1), <-channel03)
	}

}

type balok struct {
	panjang int
	lebar   int
	tinggi  int
}
