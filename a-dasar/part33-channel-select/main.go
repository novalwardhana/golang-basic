package main

import (
	"fmt"
	"runtime"
)

func getMax(numbers []int, ch chan int) {
	var result int
	for i, v := range numbers {
		if i == 0 {
			result = v
		}
		if result < v {
			result = v
		}
	}
	ch <- result
}

func getMean(numbers []int, ch chan float64) {
	totalData := len(numbers)
	sumData := 0
	for _, v := range numbers {
		sumData += v
	}
	result := float64(sumData) / float64(totalData)
	ch <- result
}

func getMin(numbers []int, ch chan int) {
	var result int
	for i, v := range numbers {
		if i == 0 {
			result = v
		}
		if result > v {
			result = v
		}
	}
	ch <- result
}

type balok struct {
	panjang int
	lebar   int
	tinggi  int
}

func (b *balok) luasPermukaan(ch chan int) {
	result := (2 * b.panjang * b.lebar) + (2 * b.lebar * b.tinggi) + (2 * b.panjang * b.tinggi)
	ch <- result
}

func (b *balok) volume(ch chan int) {
	result := b.panjang * b.lebar * b.tinggi
	ch <- result
}

func (b *balok) upsize(ch chan balok, s int) {
	b.panjang = s * b.panjang
	b.lebar = s * b.lebar
	b.tinggi = s * b.tinggi
	ch <- *b
}

type bangunRuang interface {
	luasPermukaan(ch chan int)
	volume(ch chan int)
	upsize(ch chan balok, s int)
}

func main() {
	fmt.Println("Part 33 Channel Select")
	runtime.GOMAXPROCS(2)
	fmt.Println("")
	fmt.Println("Penerapan keyword select")
	arr01 := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	fmt.Printf("arr01: %v\n", arr01)
	arr01Max := make(chan int)
	arr01Mean := make(chan float64)
	go getMax(arr01, arr01Max)
	go getMean(arr01, arr01Mean)
	for i := 0; i < 2; i++ {
		select {
		case max := <-arr01Max:
			fmt.Printf("Nilai max: %v\n", max)
		case mean := <-arr01Mean:
			fmt.Printf("Nilai mean: %v\n", mean)
		}
	}
	fmt.Println("")

	arr02 := []int{5, 8, 6, 7, 9, 5, 5, 8, 9, 9}
	fmt.Printf("arr02: %v\n", arr02)
	arr02Max := make(chan int)
	arr02Mean := make(chan float64)
	arr02Min := make(chan int)
	go getMax(arr02, arr02Max)
	go getMin(arr02, arr02Min)
	go getMean(arr02, arr02Mean)
	for i := 0; i < 3; i++ {
		select {
		case max := <-arr02Max:
			fmt.Printf("Nilai max: %v\n", max)
		case min := <-arr02Min:
			fmt.Printf("Nilai min: %v\n", min)
		case mean := <-arr02Mean:
			fmt.Printf("Nilai mean: %v\n", mean)
		}
	}
	fmt.Println("")

	arr03 := balok{10, 5, 8}
	arr03LuasPermukaan := make(chan int)
	arr03Volume := make(chan int)
	fmt.Printf("arr03 balok: %v\n", arr03)
	go arr03.luasPermukaan(arr03LuasPermukaan)
	go arr03.volume(arr03Volume)
	for i := 0; i < 2; i++ {
		select {
		case luas := <-arr03LuasPermukaan:
			fmt.Printf("Luas permukaan: %v\n", luas)
		case volume := <-arr03Volume:
			fmt.Printf("Volume: %v\n", volume)
		}
	}
	fmt.Println("")

	var arr04 bangunRuang
	arr04 = &balok{12, 5, 7}
	fmt.Printf("arr04 balok: %v\n", arr04)
	arr04LuasPermukaan := make(chan int)
	arr04Volume := make(chan int)
	arr04Upsize := make(chan balok)
	go arr04.luasPermukaan(arr04LuasPermukaan)
	go arr04.volume(arr04Volume)
	go arr04.upsize(arr04Upsize, 2)
	for i := 0; i < 3; i++ {
		select {
		case luas := <-arr04LuasPermukaan:
			fmt.Printf("Luas permukaan: %v\n", luas)
		case volume := <-arr04Volume:
			fmt.Printf("Volume: %v\n", volume)
		case upsize := <-arr04Upsize:
			fmt.Printf("Upsize: %v\n", upsize)
		}
	}
}
