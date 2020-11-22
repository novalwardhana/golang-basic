package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendMessages(ch chan<- string) {
	for i := 0; true; i++ {
		ch <- fmt.Sprintf("Data: %v", i)
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func receiveMessages(ch <-chan string) {
loop:
	for {
		select {
		case message := <-ch:
			fmt.Println(message)
		case <-time.After(time.Second * 5):
			fmt.Println("Sleep timeout")
			break loop
		}
	}
}

func sendMessages02(ch chan<- string) {
	for i := 0; true; i++ {
		ch <- fmt.Sprintf("Data: %d", i)
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func receiveMessages02(ch <-chan string) {
loop:
	for {
		select {
		case message := <-ch:
			fmt.Sprintln(message)
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout")
			break loop
		}
	}
}

func sendMessages03(ch chan<- string) {
	for i := 0; true; i++ {
		ch <- fmt.Sprintf("Data: %d", i)
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func receiveMessages03(ch <-chan string) {
loop:
	for {
		select {
		case message := <-ch:
			fmt.Println(message)
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout 3")
			break loop
		}
	}
}

func main() {
	fmt.Println("Part 35 Channel Timeout")
	fmt.Println("")
	fmt.Println("Part 35.1 Penerapan channel timeout")
	runtime.GOMAXPROCS(3)
	messages := make(chan string)
	go sendMessages(messages)
	receiveMessages(messages)
	fmt.Println("")

	messages02 := make(chan string)
	go sendMessages02(messages02)
	receiveMessages02(messages02)
	fmt.Println("")

	message03 := make(chan string)
	go sendMessages03(message03)
	receiveMessages03(message03)
	fmt.Println("")
}
