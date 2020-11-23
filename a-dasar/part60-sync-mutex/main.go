package main

import (
	"fmt"
	"runtime"
	"sync"

	log "github.com/sirupsen/logrus"
)

type counter struct {
	val int
}

func (c *counter) add(x int) {
	c.val++
}

func (c *counter) value() int {
	return c.val
}

type counterMutex struct {
	sync.Mutex
	val int
}

func (c *counterMutex) addMutex(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counterMutex) valueMutex() int {
	return c.val
}

func main() {
	log.Info("Part 60 Sync Mutex")
	fmt.Println("")
	runtime.GOMAXPROCS(2)

	log.Info("Part 60.1 - 60.3 Kondisi race condition")
	var meter counter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				meter.add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Infoln("Hasil meter dengan race condition:", meter.val, " | Seharusnya: 1000000")
	fmt.Println("")
	//go run -race a-dasar/part60-sync-mutex/*.go

	log.Info("Penerapan sync mutex")
	var meterMutex counterMutex
	var wgMutex sync.WaitGroup
	var mtx sync.Mutex
	for i := 0; i < 1000; i++ {
		wgMutex.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				mtx.Lock()
				meterMutex.addMutex(1)
				mtx.Unlock()
			}
			wgMutex.Done()
		}()
	}
	wgMutex.Wait()
	log.Infoln("Hasil meter dengan mutex:", meterMutex.val)
	fmt.Println("")
}
