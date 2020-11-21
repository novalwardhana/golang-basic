package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 21 Fungsi Closure")
	fmt.Println("")

	fmt.Printf("21.1 Closure disimpan sebagai variabel")
	var Fungsi01 = func(data []int) (float64, int, int) {
		totalData := 0
		jumlahData := 0
		maxData := 0
		minData := 0
		for i, v := range data {
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
			totalData += v
			jumlahData++
		}
		rataRata := float64(totalData) / float64(jumlahData)
		return rataRata, maxData, minData
	}
	func01Params := []int{1024, 512, 256, 128, 64, 2, 4, 8, 16, 32}
	fmt.Printf("func01Params: %v\n", func01Params)
	fungsi01RataRata, fungsi01MaxData, fungsi01MinData := Fungsi01(func01Params)
	fmt.Printf("Rata-rata: %v\n", fungsi01RataRata)
	fmt.Printf("Max: %v\n", fungsi01MaxData)
	fmt.Printf("Min: %v\n", fungsi01MinData)
	fmt.Println("")

	fmt.Println("21.2 Closure yang dieksekusi secara langsung")
	func02Params := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	fmt.Printf("func02Params: %v\n", func02Params)
	var Fungsi02A = func(min int) []int {
		result := []int{}
		for _, v := range func02Params {
			if min < v {
				result = append(result, v)
			}
		}
		return result
	}(9)
	fmt.Printf("Filter: %v\n", Fungsi02A)
	var Fungsi02BMax, Fungsi02BMin = func() (int, int) {
		maxData := 0
		minData := 0
		for i, v := range func02Params {
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
		return maxData, minData
	}()
	fmt.Printf("Max: %v\n", Fungsi02BMax)
	fmt.Printf("Min: %v\n", Fungsi02BMin)
}
