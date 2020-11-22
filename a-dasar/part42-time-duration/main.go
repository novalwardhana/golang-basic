package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Part 42 Time Duration")
	fmt.Println("")

	fmt.Println("Part 42.1 Praktek time duration")
	timeNow := time.Now()
	<-time.After(time.Second * 5)
	timeNowSince := time.Since(timeNow)
	fmt.Printf("Time now since seconds: %v\n", timeNowSince.Seconds())
	fmt.Printf("Time now since minutes: %v\n", timeNowSince.Minutes())
	fmt.Printf("Time now since hours: %v\n", timeNowSince.Hours())
	fmt.Println("")

	fmt.Println("Part 42.2 Kalkulasi menggunakan time since")
	timeNow = time.Now()
	ch := make(chan bool)
	time.AfterFunc(time.Second*5, func() {
		fmt.Println("..5 detik kemudian..")
		ch <- true
	})
	<-ch
	timeNowSince = time.Since(timeNow)
	fmt.Printf("Time now sice seconds: %v\n", timeNowSince.Seconds())
	fmt.Printf("Time now since minutes: %v\n", timeNowSince.Minutes())
	fmt.Printf("Time now since hours: %v\n", timeNowSince.Hours())
	fmt.Println("")

	fmt.Println("Part 42.3 Method time duration")
	timeNow = time.Now()
	var timeout = 5
	var chTime = make(chan bool)
	var input int
	fmt.Printf("Masukkan input waktu:\n")
	fmt.Scan(&input)
	go func(input int, ch chan<- bool) {
		time.AfterFunc(time.Second*time.Duration(input), func() {
			ch <- true
		})
	}(timeout, chTime)
	go func(timeout int, ch <-chan bool) {
		if input == 0 {
			<-ch
			fmt.Println("..timeout input after ", timeout, "seconds")
			os.Exit(0)
		}
	}(timeout, chTime)
	timeNowSince = time.Since(timeNow)
	fmt.Printf("Time now since seconds: %v\n", timeNowSince.Seconds())
	fmt.Printf("Time now since minutes: %v\n", timeNowSince.Minutes())
	fmt.Printf("Time now since hours: %v\n", timeNowSince.Hours())
	fmt.Println("")

	fmt.Println("Part 42.4 Kalkulasi durasi antara 2 objek waktu")
	t1 := time.Now()
	fmt.Printf("t1: %s\n", t1.Format("02/01/2006 03:04:05"))
	time.Sleep(time.Second * 2)
	t2 := time.Now()
	fmt.Printf("t2: %s\n", t2.Format("02/01/2006 03:04:05"))
	time.Sleep(time.Second * 5)
	t3 := time.Now()
	fmt.Printf("t3: %s\n", t3.Format("02/01/2006 03:04:05"))
	duration := t2.Sub(t1)
	fmt.Printf("Time between t1 and t2: %f seconds\n", float64(duration.Seconds()))
	duration = t3.Sub(t1)
	fmt.Printf("Time between t1 and t3: %f minutes\n", float64(duration.Minutes()))
	duration = t3.Sub(t2)
	fmt.Printf("Time between t2 and t1: %f hours\n", float64(duration.Hours()))
	fmt.Println("")

	fmt.Println("Part 42.5 Konversi angka ke time duration")
	var timeLength time.Duration = 5
	fmt.Printf("timeLength: %f\n", float64(timeLength))
	timeLengthSecond := timeLength * time.Second
	timeLengthMinute := timeLength * time.Minute
	timeLengthHour := timeLength * time.Hour
	fmt.Printf("timeLengthSecond: %v\n", timeLengthSecond)
	fmt.Printf("timeLengthMinute: %v\n", timeLengthMinute)
	fmt.Printf("timeLengthHour: %v\n", timeLengthHour)
	duration = time.Duration(2) * time.Second
	fmt.Printf("duration second: %v\n", duration)
	duration = time.Duration(5) * time.Minute
	fmt.Printf("duration minute: %v\n", duration)
	duration = 7 * time.Hour
	fmt.Printf("duration hour: %v\n", duration)
}
