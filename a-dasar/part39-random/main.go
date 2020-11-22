package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Part 39 Random")
	fmt.Println("")

	fmt.Println("Part 39.1 Pengenalan random")
	fmt.Printf("contoh: %d\n", rand.Int())
	fmt.Println("")

	fmt.Println("Part 39.2 Package math/rand")
	rand.Seed(10)
	for i := 1; i <= 3; i++ {
		fmt.Printf("Random ke-%d: %d\n", i, rand.Int())
	}
	fmt.Println("")

	fmt.Println("Part 39.2 Random unique")
	timeNowUniqueInt := time.Now().UTC().UnixNano()
	rand.Seed(timeNowUniqueInt)
	fmt.Printf("Time now: %d\n", timeNowUniqueInt)
	index := 1
	for {
		if index >= 4 {
			break
		}
		fmt.Printf("Random unique ke-%d: %d\n", index, rand.Int())
		index++
	}
	fmt.Println("")

	fmt.Println("Part 39.4 Random tipe data numerik lainnya")
	timeNowUniqueFLoat32 := time.Now().UTC().UnixNano()
	fmt.Printf("Time now: %d\n", timeNowUniqueFLoat32)
	index = 1
	for {
		fmt.Printf("Random unique float32 ke-%d: %.3f\n", index, rand.Float32())
		if index >= 4 {
			break
		}
		index++
	}
	timeNowUniqueUint32 := time.Now().UTC().UnixNano()
	fmt.Printf("Time Now: %d\n", timeNowUniqueUint32)
	index = 1
	for index <= 4 {
		fmt.Printf("Random unique uint32 ke-%d: %v\n", index, rand.Uint32())
		index++
	}
	fmt.Println("")

	fmt.Println("Part 39.5 Random index tertentu")
	fmt.Printf("Random unique int: %d\n", rand.Intn(2))
	fmt.Printf("Random unique int: %d\n", rand.Intn(3))
	fmt.Println("")

	fmt.Println("Part 39.6 Random tipe data string")
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fmt.Printf("letters: %v\n", string(letters))
	var funcRandomString = func(length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}(10)
	fmt.Printf("Random string-1: %s\n", funcRandomString)
	letters = []rune("AaBbCcDdEeFfGgHhIiJj1234567890")
	fmt.Printf("letters: %s\n", string(letters))
	var funcRandomStringTwo = func(length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}
	fmt.Printf("Random string-2: %s\n", funcRandomStringTwo(12))
	letters = []rune("NoVAliTAKusuMAwarDHAna20")
	fmt.Printf("letters: %s\n", string(letters))
	var funcRandomStringThree = func(length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}
	fmt.Printf("Random string-3: %s\n", funcRandomStringThree(8))
}
