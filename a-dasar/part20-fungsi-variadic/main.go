package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 20 Fungsi Variadic")
	fmt.Println("")

	fmt.Println("20.1 Penerapan fungsi variadic")
	func01 := fungsi01(9, 3, 12, 15, 21, 4, 5, 9, 22, 25, 21, 29, 30)
	fmt.Printf("Data: %v\n", []int{9, 3, 12, 15, 21, 4, 5, 9, 22, 25, 21, 29, 30})
	fmt.Printf("Rata-rata data: %v\n", func01)
	fmt.Println("")

	fmt.Println("20.2 Pengisian data fungsi variadic dengan slice")
	func02Params := []int{1, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Printf("Data: %v\n", func02Params)
	func02 := fungsi02(func02Params...)
	fmt.Printf("Rata-rata data: %v\n", func02)
	fmt.Println("")

	fmt.Println("20.3 Perbedaan fungsi variadic dan fungsi biasa")
	func03Params := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	fmt.Printf("func03Params: %v\n", func03Params)
	fungsi03("Bilangan pangkat 2", func03Params...)
}

func fungsi01(data ...int) float64 {
	totalData := 0
	jumlahData := 0
	for _, v := range data {
		totalData += v
		jumlahData++
	}
	rataRata := float64(totalData) / float64(jumlahData)
	return rataRata
}

func fungsi02(data ...int) float64 {
	totalData := 0
	jumlahData := 0
	for _, v := range data {
		totalData += v
		jumlahData++
	}
	rataRata := float64(totalData) / float64(jumlahData)
	return rataRata
}

func fungsi03(message string, data ...int) {
	fmt.Printf("%v\n", message)
	totalData := 0
	jumlahData := 0
	maxData := 0
	minData := 0
	for i, v := range data {
		totalData += v
		jumlahData++
		if i == 0 {
			maxData = v
			minData = v
		}
		if maxData < v {
			maxData = v
		}
		if minData > v {
			minData = v
		}
	}
	rataRata := float64(totalData) / float64(jumlahData)
	fmt.Printf("Rata-rata: %v\n", rataRata)
	fmt.Printf("Minimal: %v\n", minData)
	fmt.Printf("Maksimal: %v\n", maxData)
}
