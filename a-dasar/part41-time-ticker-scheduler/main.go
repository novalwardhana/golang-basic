package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Part 41 Time Ticker Scheduler")
	fmt.Println("")
	runtime.GOMAXPROCS(2)

	fmt.Println("Part 41.1 Time Sleep")
	fmt.Printf("Start\n...\n")
	time.Sleep(time.Second * 5)
	fmt.Printf("After 5 second\n")
	fmt.Println("")

	fmt.Println("Part 41.2 Scheduler menggunakan time.Sleep")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("...Wake up every 2 seconds...")
	}
	fmt.Println("")

	fmt.Println("Part 41.3 time newTimer")
	newTimer := time.NewTimer(time.Second * 4)
	fmt.Printf("..new timer start one..\n")
	<-newTimer.C
	fmt.Printf("..new timer finish one..\n")
	newTimer = time.NewTimer(time.Second * 10)
	fmt.Printf("..new timer start two..\n")
	<-newTimer.C
	fmt.Printf("..new timer finish two..\n")
	fmt.Println("")

	fmt.Println("Part 41.4 time afterFunc")
	var ch = make(chan bool)
	time.AfterFunc(time.Second*5, func() {
		fmt.Printf("Expired one\n")
		ch <- true
	})
	fmt.Printf("..after func start one..\n")
	<-ch
	fmt.Printf("..after func finish one..\n")
	ch = make(chan bool)
	time.AfterFunc(time.Second*2, func() {
		fmt.Println("Expired two")
		ch <- true
	})
	fmt.Printf("..after func start two..\n")
	<-ch
	fmt.Printf("..after func finish two..\n")
	fmt.Println("")

	fmt.Println("Part 41.5 time after")
	fmt.Printf("..after func start one..\n")
	<-time.After(time.Second * 3)
	fmt.Printf("..after func finish one..\n")
	fmt.Printf("..after func start two..\n")
	<-time.After(time.Second * 2)
	fmt.Printf("..after func finish two..\n")
	fmt.Println("")

	fmt.Println("Part 41.6 Scheduler menggunakan ticker")
	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)
	// go func() {
	// 	time.Sleep(time.Second * 10)
	// 	done <- true
	// }()
	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 		return
	// 	case t := <-ticker.C:
	// 		fmt.Printf("Ticker one-%v\n", t)
	// 	}
	// }
	// doneTicker := make(chan bool)
	// tickerTwo := time.NewTicker(time.Second)
	// go func() {
	// 	time.Sleep(time.Second * 5)
	// 	doneTicker <- true
	// }()
	// for {
	// 	select {
	// 	case <-doneTicker:
	// 		tickerTwo.Stop()
	// 		return
	// 	case t := <-tickerTwo.C:
	// 		fmt.Printf("Ticker two: %v\n", t)
	// 	}
	// }

	// fmt.Println("")

	fmt.Println("Part 41.7 Kombinasi timer & goroutine")
	timeout := 10
	chInput := make(chan bool)
	go timer(chInput, timeout)
	go watcher(chInput, timeout)
	fmt.Println("25 x 5?")
	var input int
	fmt.Scan(&input)
	if input == 125 {
		fmt.Println("Jawaban benar")
	} else {
		fmt.Println("Coba lagi")
	}
}

func timer(ch chan<- bool, timeout int) {
	time.AfterFunc(time.Second*time.Duration(timeout), func() {
		ch <- true
	})
}

func watcher(ch <-chan bool, timeout int) {
	<-ch
	fmt.Printf("Timeout after %d second \n", timeout)
	os.Exit(0)
}
