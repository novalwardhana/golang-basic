package main

import (
	"fmt"
	"strings"
)

func fungsi01(data []string, callback func(each string) bool) []string {
	result := []string{}
	for _, v := range data {
		if validation := callback(v); validation == true {
			result = append(result, v)
		}
	}
	return result
}

type fungsi02AboveAverage func(each float64, average float64) bool

func fungsi02(data []int, funcAboveAverage fungsi02AboveAverage) (float64, int, int, []float64) {
	totalData := 0
	jumlahData := 0
	maxData := 0
	minData := 0
	for i, v := range data {
		if i == 0 {
			maxData = v
			minData = v
		}
		totalData += v
		jumlahData++
	}
	rataRataData := float64(totalData) / float64(jumlahData)
	aboveAverage := []float64{}
	for _, v := range data {
		if validation := funcAboveAverage(float64(v), rataRataData); validation == true {
			aboveAverage = append(aboveAverage, float64(v))
		}

	}
	return rataRataData, maxData, minData, aboveAverage
}

func main() {
	fmt.Println("Part 22 Fungsi sebagai parameter")
	fmt.Println("")

	fmt.Println("22.1 Penerapan fungsi sebagai parameter")
	func01Params := []string{"Machester United", "Real Madrid", "Barcelona", "Liverpool", "Lyon", "Paris Saint Germain"}
	fmt.Printf("func01Params: %v\n", func01Params)
	func01A := fungsi01(func01Params, func(each string) bool {
		return strings.Contains(each, "i")
	})
	fmt.Printf("func01A: %v\n", func01A)
	func01B := fungsi01(func01Params, func(each string) bool {
		if len(each) > 5 {
			return true
		}
		return false
	})
	fmt.Printf("func01B: %v\n", func01B)
	fmt.Println("")

	fmt.Println("22.2 Alias skema closure")
	func02Params := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	fmt.Printf("func02Params: %v\n", func02Params)
	func02RataRata, func02Max, func02Min, func02AboveAverage := fungsi02(func02Params, func(each float64, average float64) bool {
		if each >= average {
			return true
		}
		return false
	})
	fmt.Printf("Rata-rata: %v\n", func02RataRata)
	fmt.Printf("Max: %v\n", func02Max)
	fmt.Printf("Min: %v\n", func02Min)
	fmt.Printf("Diatas Rata-rata: %v\n", func02AboveAverage)

}
