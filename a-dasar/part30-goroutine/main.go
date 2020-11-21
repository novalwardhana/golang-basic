package main

import (
	"fmt"
	"runtime"
)

func func01(iteration int, s string) {
	for i := 0; i < iteration; i++ {
		fmt.Printf("%v %v\n", i, s)
	}
}

func main() {
	fmt.Println("Part 30 Goroutine")
	fmt.Println("")

	fmt.Println("Part 30.1 Penerapan goroutine")
	runtime.GOMAXPROCS(5)
	go func01(5, "Ini Async")
	func01(5, "Sync")
	var input01 string = "test01"
	var input02 string = "test02"
	var input03 string = "test03"
	fmt.Scanln(&input01, &input02, &input03)
}
