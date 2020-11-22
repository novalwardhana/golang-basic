package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, data string) {
	defer wg.Done()
	fmt.Println(data)
}

func main() {
	runtime.GOMAXPROCS(2)

	fmt.Println("Part 59 Sync Wait Group")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("Data ke-%d", i)
		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()

	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("Iterasi ke-%d", i)
		wg2.Add(1)
		go doPrint(&wg2, data)
	}
	wg2.Wait()

	var wg3 sync.WaitGroup
	for i := 0; i < 3; i++ {
		var data = fmt.Sprintf("Send data-%d", i)
		wg3.Add(1)
		go doPrint(&wg3, data)
	}
	wg3.Wait()

	var wg4 sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("Insert data-%d", i)
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			fmt.Println(data)
		}()
	}
	wg4.Wait()
}
