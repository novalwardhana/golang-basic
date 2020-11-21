package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 14 Perulangan")
	fmt.Println("")

	fmt.Println("Perulangan dengan for")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iterasi ke-%v\n", i)
	}
	fmt.Println("")

	fmt.Println("Perulangan dengan for-kondisi")
	forTwo := 0
	for forTwo < 5 {
		fmt.Printf("Iterasi ke-%v\n", forTwo)
		forTwo++
	}
	fmt.Println("")

	fmt.Println("Perulangan dengan for-tanpa-kondisi")
	forThree := 0
	for {
		fmt.Printf("Iterasi ke-%v\n", forThree)
		forThree++
		if forThree >= 5 {
			break
		}
	}
	fmt.Println("")

	fmt.Println("Perulangan dengan for-range")
	forFour := []int{0, 1, 2, 3, 4}
	for v, i := range forFour {
		fmt.Printf("Iterasi ke-%v dengan nilai %v\n", i+1, v)
	}
	fmt.Println("")

	fmt.Println("Perulangan dengan for-break-continue")
	forFive := 0
	forFiveLimit := 10
	for {
		if result := forFive % 2; result == 0 {
			forFive++
			continue
		} else if result := forFive % 2; result == 1 {
			fmt.Printf("Iterasi ke-%v\n", forFive)
			forFive++
			if forFive >= forFiveLimit {
				break
			}
		}
	}
	fmt.Println("")

	fmt.Println("Perulangan bersarang")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			k := 10 - i
			if k > j {
				fmt.Printf("%v ", j)
			}
		}
		fmt.Println("")
	}
	fmt.Println("")

	fmt.Println("Perulangan break-outerloop")
outerloop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			k := 10 - i
			if k > j {
				fmt.Printf("%v ", j)
			}
			if i >= 6 {
				break outerloop
			}
		}
		fmt.Println("")
	}
}
