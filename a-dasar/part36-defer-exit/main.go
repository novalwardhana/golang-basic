package main

import (
	"fmt"
	"os"
)

func func01(s string) {
	defer fmt.Printf("###..Great..###\n")
	defer fmt.Printf("Thanks for choose team\n")
	if s == "Real Madrid" {
		fmt.Printf("La liga Champions 2020: %s\n", s)
		return
	}
	fmt.Printf("You choose team: %s\n", s)
}

func func02(s int, str string) {
	defer fmt.Println("#*#*#*")
	for i := 0; i < s; i++ {
		defer fmt.Printf("Data %s: %d\n", str, i)
	}
}

func main() {
	fmt.Println("Part 36 Defer Exit")
	fmt.Println("")

	fmt.Println("Part 36.1 Penerapan defer")
	func01("Real Madrid")
	func01("Arsenal")
	func01("Real Madrid")
	fmt.Println("")

	fmt.Println("Part 36.2 Kombinasi defer dan IIFE")
	var02A := 3
	fmt.Printf("Data A: 1\n")
	if var02A == 3 {
		defer fmt.Println("")
		func() {
			defer fmt.Printf("Data A: 2\n")
		}()
		defer fmt.Printf("Data A: %d\n", var02A)
	}
	var02B := 10
	if var02B == 10 {
		for i := 0; i < 10; i++ {
			defer fmt.Printf("Data B: %d\n", i)
		}
	}
	var02C := 10
	func02(var02C, "C")
	var02D := 5
	func02(var02D, "D")
	var02E := 7
	func02(var02E, "E")

	fmt.Println("Part 36.3 OS exit")
	defer fmt.Println("Os Exit")
	os.Exit(1)
	fmt.Println("Hello World")
}
