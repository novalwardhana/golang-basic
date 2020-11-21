package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Part 19 Fungsi Multiple Return")
	fmt.Println("")

	fmt.Println("19.1 Penerapan fungsi multiple return")
	func01Params := float64(7)
	fmt.Printf("Jari-jari lingkaran: %v\n", func01Params)
	func01Luas, func01Keliling := fungsi01(func01Params)
	fmt.Printf("Luas: %v\n", func01Luas)
	fmt.Printf("Keliling: %v\n", func01Keliling)
	fmt.Println("")

	fmt.Println("19.2 Fungsi dengan predefined return value")
	func02ParamsP := 9
	func02ParamsL := 6
	func02ParamsT := 12
	fmt.Printf("Panjang Balok: %v\n", func02ParamsP)
	fmt.Printf("Lebar Balok: %v\n", func02ParamsL)
	fmt.Printf("Tinggi Balok: %v\n", func02ParamsT)
	func02LuasPermukaan, func02Volum := fungsi02(func02ParamsP, func02ParamsL, func02ParamsT)
	fmt.Printf("Luas Permukaan: %v\n", func02LuasPermukaan)
	fmt.Printf("Volum: %v\n", func02Volum)
}

func fungsi01(d float64) (float64, float64) {
	luas := math.Pi * math.Pow(d/2, 2)
	keliling := math.Pi * d
	return luas, keliling
}

func fungsi02(p, l, t int) (luasPermukaan int, volum int) {
	luasPermukaan = (2 * p * l) + (2 * l * t) + (2 * p * t)
	volum = p * l * t
	return
}
