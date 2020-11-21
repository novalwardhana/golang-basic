package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 17 Map")
	fmt.Println("")

	fmt.Println("17.1 Penggunaan Map")
	var map01A map[string]int
	map01A = map[string]int{}
	map01A["januari"] = 1
	map01A["februari"] = 2
	map01A["maret"] = 3
	map01A["april"] = 4
	fmt.Printf("map01A: %v\n", map01A)
	map01B := map[string]int{}
	map01B["mei"] = 5
	map01B["juni"] = 6
	map01B["juli"] = 7
	fmt.Printf("map01B: %v\n", map01B)
	fmt.Println("")

	fmt.Println("17.2 Inisiai nilai Map")
	map02A := map[string]int{"agustus": 8, "september": 9, "oktober": 10, "november": 11, "desember": 12}
	fmt.Printf("map02A: %v\n", map02A)
	map02B := map[string]string{
		"Manchester United": "Old Ttafford",
		"Arsenal":           "Emirates Stadium",
		"Chelsea":           "Stanford Bridge",
	}
	fmt.Printf("map02B: %v\n", map02B)
	map02C := map[int]int{
		1: 1,
		2: 4,
		3: 9,
		4: 16,
		5: 25,
		6: 36,
	}
	fmt.Printf("map02C: %v\n", map02C)
	fmt.Println("")

	fmt.Println("17.3 Iterasi map dengan menggunakan for-range")
	map03 := map[string]int{
		"Manchester United": 20,
		"Liverpool":         19,
		"Arsenal":           13,
		"Chelsea":           6,
		"Manchester City":   6,
	}
	for i, v := range map03 {
		fmt.Printf("%v : %v\n", i, v)
	}
	fmt.Println("")

	fmt.Println("17.4 Menghapus item di dalam map")
	map04 := map[string]int{
		"januari":  1,
		"februari": 2,
		"maret":    3,
		"april":    4,
		"mei":      5,
	}
	fmt.Printf("map04: %v\n", map04)
	delete(map04, "april")
	fmt.Printf("map04: %v\n", map04)
	delete(map04, "miii")
	fmt.Printf("map04: %v\n", map04)
	delete(map04, "mei")
	fmt.Printf("map04: %v\n", map04)
	fmt.Println("")

	fmt.Println("17.5 Deteksi keberadaan item dengan key tertentu")
	map05 := map[string]int{
		"januari":  1,
		"februari": 2,
		"maret":    3,
		"april":    4,
	}
	fmt.Printf("map05: %v\n", map05)
	if value, status := map05["februari"]; status == true {
		fmt.Printf("map05['februari']: %v\n", value)
	} else {
		fmt.Printf("map05['februari]: %v\n", status)
	}
	if value, status := map05["mei"]; status == true {
		fmt.Printf("map05['mei]: %v\n", value)
	} else {
		fmt.Printf("map05['mei']: %v\n", status)
	}
	fmt.Println("")

	fmt.Println("17.6 Kombinasi slice dan map")
	map06A := []map[int]int{
		{1: 1, 2: 4, 3: 9, 4: 16, 5: 25},
		map[int]int{
			1: 1,
			2: 8,
			3: 27,
			4: 64,
			5: 125,
		},
	}
	fmt.Printf("map06A: %v\n", map06A)
	map06B := []map[string]string{
		map[string]string{
			"Manchester United": "Old Trafford",
			"Liverpool":         "Anfield",
			"Arsenal":           "Emirates Stadium",
		},
		map[string]string{
			"Real Madrid": "Santiago Bernabeu",
			"Barcelona":   "Camp Nou",
		},
	}
	fmt.Printf("map06B: %v\n", map06B)
	for i, v := range map06B {
		fmt.Printf("%v : %v\n", i, v)
		for j, w := range v {
			fmt.Printf("%v: %v\n", j, w)
		}
	}
}
