package main

import (
	"fmt"
	"runtime"
)

func sendMessages(ch chan<- string) {
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Data: %v", i)
		ch <- result
	}
	close(ch)
}

func receiveMessages(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func sendData(ch chan<- string) {
	for i := 0; i < 12; i++ {
		ch <- fmt.Sprintf("Data: %v", i)
	}
	close(ch)
}

func receiveData(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	fmt.Println("Part 34 Range Close")
	fmt.Println("")
	fmt.Println("Part 34.1 Penerapan range close")
	runtime.GOMAXPROCS(2)
	messages := make(chan string)
	go sendMessages(messages)
	receiveMessages(messages)
	fmt.Println("")

	messages02 := make(chan string)
	go sendData(messages02)
	receiveMessages(messages02)
	fmt.Println("")

}
