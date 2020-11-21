package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 28 Interface kosong")
	fmt.Println("")

	fmt.Println("Part 28.1 Penggunaan interface")
	var infc01A interface{}
	fmt.Printf("infc01A: %v\n", infc01A)
	infc01A = "Novalita"
	fmt.Printf("infc01A: %v\n", infc01A)
	infc01A = []int{2, 4, 8, 16, 32, 64}
	fmt.Printf("infc01A: %v\n", infc01A)
	infc01A = 22.5
	fmt.Printf("infc01A: %v\n", infc01A)
	var infc01B map[string]interface{}
	infc01B = map[string]interface{}{
		"1":     12,
		"dua":   []int{1, 2, 3, 4, 5},
		"three": 22.5,
		"4": map[string]string{
			"Real Madrid": "Santiago Bernabeu",
			"Barcelona":   "Camp Nou",
			"Juventus":    "Allianz Stadium",
		},
	}
	fmt.Printf("infc01B: %v\n", infc01B)
	fmt.Println("")

	fmt.Println("Part 28.2 casting interface")
	var infc02A interface{}
	fmt.Printf("infc02A: %v\n", infc02A)
	infc02A = 5
	fmt.Printf("infc02A: %v\n", infc02A)
	infc02A = infc02A.(int) * 2
	fmt.Printf("infc02A: %v\n", infc02A)
	var infc02B interface{}
	fmt.Printf("infc02B: %v\n", infc02B)
	infc02B = []int{1, 4, 9, 16, 25, 36, 49}
	fmt.Printf("infc02B: %v\n", infc02B)
	infc02B = append(infc02B.([]int), 64)
	infc02B = append(infc02B.([]int), 81)
	fmt.Printf("infc02B: %v\n", infc02B)
	fmt.Println("")

	fmt.Println("Part 28.3 Casting variabel interface kosong ke pointer")
	type Struct03 struct {
		name    string
		address string
	}
	var infc03A interface{}
	fmt.Printf("infc03A: %v\n", infc03A)
	infc03A = &Struct03{name: "Noval", address: "Banguntapan"}
	fmt.Printf("infc03A: %v\n", infc03A)
	fmt.Printf("infc03A name: %v\n", infc03A.(*Struct03).name)
	fmt.Printf("infc03A address: %v\n", infc03A.(*Struct03).address)
	var infc03B interface{}
	infc03B = Struct03{name: "Novan", address: "Berbah"}
	fmt.Printf("infc03B: %v\n", infc03B)
	fmt.Printf("infc03B name: %v\n", infc03B.(Struct03).name)
	fmt.Printf("infc03B address: %v\n", infc03B.(Struct03).address)
	var infc03C interface{}
	infc03C = &Struct03{"Deni", "Berbah"}
	fmt.Printf("infc03C: %v\n", infc03C)
	fmt.Printf("infc03C name: %v\n", infc03C.(*Struct03).name)
	infc03C.(*Struct03).address = "Jogja"
	fmt.Printf("infc03C address: %v\n", infc03C.(*Struct03).address)
	fmt.Println("")

	fmt.Println("Part 28.4 kombinasi map dan interface")
	var infc04A = []map[string]interface{}{
		{"name": "Noval", "age": 26, "hobbies": []string{"Jogging", "Swimming", "Riding"}},
		{"name": "Deni", "age": 27, "address": "Berbah"},
		{"name": "Novan", "age": 28, "skills": []string{"Videgrapher", "Youtuber"}},
	}
	fmt.Printf("infc04A: %v\n", infc04A)
	for i, v := range infc04A {
		fmt.Printf("infc04A %v: %v\n", i, v)
	}
	var infc04B = []interface{}{
		map[string]string{
			"Barcelona":         "Camp Nou",
			"Arsenal":           "Emirates Stadium",
			"Borussia Dortmund": "Signal Iduna Park",
		},
		map[string]interface{}{
			"Real Madrid":         []string{"Benzema", "Bale", "Isco"},
			"Paris Saint Germain": []string{"Neymar", "Mbappe", "Keylor Navas"},
		},
	}
	fmt.Printf("infc04B: %v\n", infc04B)
	for i, v := range infc04B {
		fmt.Printf("infc04B %v: %v\n", i, v)
	}

}
