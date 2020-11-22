package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Part 32 Buffered channel")
	fmt.Println("")

	fmt.Println("Part 32.1 Penerapan buffer")
	runtime.GOMAXPROCS(2)
	messages01 := make(chan int, 2)
	go func() {
		for {
			i := <-messages01
			fmt.Printf("Receive data: %v\n", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Printf("Send data: %v\n", i)
		messages01 <- i
	}
	fmt.Println("")

	messages02 := make(chan int, 5)
	go func() {
		for {
			i := <-messages02
			fmt.Println("Receive message: ", i)
		}
	}()
	for i := 0; i < 20; i++ {
		fmt.Println("Send message: ", i)
		messages02 <- i
	}
}
