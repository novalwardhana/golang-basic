package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 15 Array")
	fmt.Println("")

	fmt.Println("Pengisian element array")
	arrayOne := []string{"Bantul", "Kota Yogyakarta", "Sleman", "Gunungkidul", "Kulon Progo"}
	fmt.Printf("arrayOne: %v\n", arrayOne)
	fmt.Println("")

	fmt.Println("Inisiasi nilai awal array")
	arrayTwo := []string{"Bantul", "Kota Yogyakarta", "Sleman", "Gunungkidul", "Kulon Progo"}
	fmt.Printf("arrayTwo: %v\n", arrayTwo)
	fmt.Printf("panjang arrayTwo: %v\n", len(arrayTwo))
	fmt.Println("")

	fmt.Println("Inisiasi dengan gaya vertikal")
	arrayThree := []string{
		"Bantul",
		"Kota Yogyakarta",
		"Sleman",
		"Gunungkidul",
		"Kulon Progo",
	}
	fmt.Printf("arrayThree: %v\n", arrayThree)
	fmt.Println("")

	fmt.Println("Inisiasi nilai awal tanpa jumlah elemen")
	arrayFour := [...]string{"Bantul", "Kota Yogyakarta", "Sleman", "Gunungkidul", "Kulon Progo"}
	fmt.Printf("arrayFour: %v\n", arrayFour)
	fmt.Printf("panjang arrayFour: %v\n", len(arrayFour))
	fmt.Println("")

	fmt.Println("Array multidimensi")
	arrayFiveA := [2][3]string{
		[3]string{
			"One ok Rock",
			"Radwimps",
			"7Billion Dots",
		},
		[3]string{
			"Green Day",
			"Metallica",
			"Queen",
		},
	}
	fmt.Printf("arraySixA: %v\n", arrayFiveA)
	fmt.Printf("panjang arraySixA: %v\n", len(arrayFiveA))
	arrayFiveB := [3][2][3]string{
		[2][3]string{
			[3]string{"Xiaomi", "Redmi", "Mi"},
			[3]string{"BlackShark", "Poco", "Hongmi"},
		},
		[2][3]string{
			[3]string{"Oppo", "Oneplus", "Realme"},
			[3]string{"Vivo", "Aqioo"},
		},
		[2][3]string{
			[3]string{"Huawei", "Honor", "Nova"},
			[3]string{"Samsung A series", "Samsung Note Series", "Samsung S Series"},
		},
	}
	fmt.Printf("arraySixB: %v\n", arrayFiveB)
	fmt.Printf("panjang arraySixB: %v\n", len(arrayFiveB))
	arrayFiveC := [3][2][3]string{
		[2][3]string{
			{"Xiaomi", "Redmi", "Mi"},
			{"BlackShark", "Poco", "Hongmi"},
		},
		[2][3]string{
			{"Oppo", "Oneplus", "Realme"},
			{"Vivo", "Aqioo"},
		},
		[2][3]string{
			{"Huawei", "Honor", "Nova"},
			{"Samsung A series", "Samsung Note Series", "Samsung S Series"},
		},
	}
	fmt.Printf("arraySixB: %v\n", arrayFiveC)
	fmt.Printf("panjang arraySixB: %v\n", len(arrayFiveC))
	arrayFiveD := [3][3][3]int{
		[3][3]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[3][3]int{
			{2, 3, 4},
			{5, 6, 7},
			{8, 9, 0},
		},
		[3][3]int{
			{3, 4, 5},
			{6, 7, 8},
			{9, 0, 1},
		},
	}
	fmt.Printf("arraySixB: %v\n", arrayFiveD)
	fmt.Printf("panjang arraySixB: %v\n", len(arrayFiveD))
	fmt.Println("")

	fmt.Println("For")
	var arrSix = [...]string{"Bantul", "Kota Yogyakarta", "Sleman", "Gunungkidul", "Kulon Progo"}
	for i := 0; i < len(arrSix); i++ {
		fmt.Printf("Iterasi ke-%v dengan nilai %v\n", i, arrSix[i])
	}
	fmt.Println("")

	fmt.Println("For-range")
	var arrSeven = [...]string{"Bantul", "Kota Yogyakarta", "Sleman", "Gunungkidul", "Kulon Progo"}
	for i, v := range arrSeven {
		fmt.Printf("Iterasi ke-%v dengan value: %v\n", i, v)
	}
	fmt.Println("")

	fmt.Println("Penggunaan elemen underscore dalam perulangan")
	var arrEightA = [...]string{"Juventus", "Paris Saint Germain", "Borussia Dortmund", "Arsenal", "Tottenham"}
	for _, v := range arrEightA {
		fmt.Printf("Klub: %v\n", v)
	}
	fmt.Println("")
	var arrEightB = [2][3][4]int{
		[3][4]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		[3][4]int{},
	}
	fmt.Println(arrEightB)
	fmt.Println("")

	fmt.Println("Penggunaan elemen array dengan keyword make")
	var arrNineA = make([]string, 2)
	arrNineA[0] = "Sate"
	arrNineA[1] = "Soto"
	fmt.Printf("arrNineA: %v\n", arrNineA)
}
