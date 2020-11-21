package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 10 Tipe Data")
	fmt.Println("")

	//Tipe data numerik non desimal
	fmt.Println("Contoh tipe data numerik non decimal")
	var numOne uint8 = 255
	var numTwo uint16 = 8000
	var numThree int = 12345678
	var numFour byte
	var numFive bool = true
	fmt.Printf("numOne: %v\n", numOne)
	fmt.Printf("numTwo: %v\n", numTwo)
	fmt.Printf("numThree: %v\n", numThree)
	fmt.Printf("numFour: %v\n", numFour)
	fmt.Printf("numFive: %v\n", numFive)
	fmt.Println("")

	fmt.Println("Contoh tipe data numerik decimal")
	var numDecOne float32 = 123.456
	var numDecTwo float64 = 190000.231
	fmt.Printf("numDecOne: %v\n", numDecOne)
	fmt.Printf("numDecTwo: %v\n", numDecTwo)
	fmt.Println("")

	fmt.Println("Contoh tipe data boolen")
	var boolOne bool = true
	var boolTwo bool
	var boolThree = false
	boolFour := true
	fmt.Printf("boolOne: %v\n", boolOne)
	fmt.Printf("boolTwo: %v\n", boolTwo)
	fmt.Printf("boolThree: %v\n", boolThree)
	fmt.Printf("boolFour: %v\n", boolFour)
	fmt.Println("")

	fmt.Println("Contoh tipe data string")
	var strOne string = "One"
	strTwo := "Ok"
	var strThree *string
	strThreeDereference := "Rock"
	strThree = &strThreeDereference
	fmt.Printf("strOne: %v\n", strOne)
	fmt.Printf("strTwo: %v\n", strTwo)
	fmt.Printf("strThree: %v\n", *strThree)
}
