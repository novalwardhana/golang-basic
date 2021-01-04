package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}

type iteration struct {
	sync.Mutex
	val int
}

func (i *iteration) addValue() {
	i.Lock()
	i.val++
	i.Unlock()
}

func add(itr *iteration, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		itr.addValue()
	}
}

func main() {
	fmt.Println("Part 59 Sync Wait Group")

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		str := fmt.Sprintf("Iterasi ke: %d", i)
		go doPrint(str, &wg)
	}
	wg.Wait()

	fmt.Println("-----")

	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		str := fmt.Sprintf("Iterasi ke: %d", i)
		go func() {
			defer wg2.Done()
			fmt.Println(str)
		}()
	}
	wg2.Wait()

	fmt.Println("-----")

	var wg3 sync.WaitGroup
	var itr3 iteration
	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			for i := 0; i < 1000; i++ {
				itr3.addValue()
			}
		}()
	}
	wg3.Wait()
	fmt.Printf("Iteration 3: %d\n", itr3.val)

	fmt.Println("-----")

	var wg4 sync.WaitGroup
	var itr4 iteration
	for i := 0; i < 300; i++ {
		wg4.Add(1)
		go func() {
			for i := 0; i < 300; i++ {
				itr4.Lock()
				itr4.val++
				itr4.Unlock()
			}
			wg4.Done()
		}()
	}
	wg4.Wait()
	fmt.Printf("Iteration 4: %d\n", itr4.val)

	fmt.Println("-----")
	var wg5 sync.WaitGroup
	var itr5 iteration
	for i := 0; i < 1000; i++ {
		wg5.Add(1)
		go add(&itr5, &wg5)
	}
	wg5.Wait()
	fmt.Printf("Iteration 5: %d\n", itr5.val)
}
