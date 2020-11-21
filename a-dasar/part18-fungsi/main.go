package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Part 18 Fungsi")
	fmt.Println("")
	rand.Seed(time.Now().Unix())

	fmt.Println("18.1 Penerapan fungsi")
	func01Params := []string{"Ronaldo", "Messi", "Buffon", "Goetze"}
	Fungsi01("Haloo", func01Params)

	fmt.Println("18.2 Fungsi dengan balikan")
	func02A := Fungsi02(1, 8)
	fmt.Printf("func02A: %v\n", func02A)
	func02B := Fungsi02(3, 4)
	fmt.Printf("func02B: %v\n", func02B)
	fmt.Println("")

	fmt.Println("18.3 Penggunaan fungsi rand seed")
	Fungsi03()
	fmt.Println("")

	fmt.Println("18.4 Fungsi dengan parameter map")
	func04Params := map[string]string{
		"Lenovo": "Legion",
		"Asus":   "ROG",
		"Dell":   "Alienware",
		"Acer":   "Predator",
	}
	Fungsi04(func04Params)
	fmt.Println("")

	fmt.Println("18.5 Deklarasi parameter bertipe data sama")
	func05A := Fungsi05(8, 5, 12)
	fmt.Printf("Luas balok: %v\n", func05A)
	func05B := Fungsi05(12, 3, 8)
	fmt.Printf("Luas balok: %v\n", func05B)
	func05C := Fungsi05(10, 9, 5)
	fmt.Printf("luas balok: %v\n", func05C)
	fmt.Println("")

	fmt.Println("18.6 Penggunaan return untuk menghentikan proses dalam fungsi")
	Fungsi06(9, 8)
	Fungsi06(11, 0)
	Fungsi06(8, 5)

}

func Fungsi01(message string, name []string) {
	for _, v := range name {
		fmt.Printf("%v %v\n", message, v)
	}
	allName := strings.Join(name, "-")
	fmt.Println(allName)
	fmt.Println("")
}

func Fungsi02(min int, max int) int {
	result := rand.Int()%(max-min+1) + min
	return result
}

func Fungsi03() {
	fmt.Println(time.Now())
}

func Fungsi04(notebook map[string]string) {
	for i, v := range notebook {
		fmt.Printf("%v: %v\n", i, v)
	}
}

func Fungsi05(p, l, t int) int {
	result := p * l * t
	return result
}

func Fungsi06(a, b int) {
	if b == 0 {
		fmt.Println("Divider cannot run because divider is 0")
		return
	}
	result := a / b
	fmt.Printf("%v / %v = %v\n", a, b, result)
}
