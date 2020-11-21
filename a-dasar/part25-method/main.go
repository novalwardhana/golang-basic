package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

type balok struct {
	panjang int
	lebar   int
	tinggi  int
}

func main() {
	fmt.Println("Part 25 Method")
	fmt.Println("")

	fmt.Println("25.1 Penerapan method")
	var01A := person{
		name: "Noval",
		age:  26,
	}
	fmt.Printf("var01A: %v\n", var01A)
	var01A.func01A("Haloo")
	var01B := person{"Deni", 27}
	fmt.Printf("var01B: %v\n", var01B)
	var01BNamaPanggilan := var01B.func01B(1)
	fmt.Printf("var01BCall: %v\n", var01BNamaPanggilan)
	fmt.Println("")

	fmt.Println("25.2 Method Pointer")
	var02A := balok{
		panjang: 21,
		lebar:   10,
		tinggi:  8,
	}
	fmt.Printf("var02A: %v\n", var02A)
	var02A.upsize(2)
	fmt.Printf("var02A upsize: %v\n", var02A)
	fmt.Printf("var02A luas permukaan: %v\n", var02A.luasPermukaan())
	fmt.Printf("var02A volume: %v\n", var02A.volume())
	var02B := &balok{
		panjang: 8,
		lebar:   4,
		tinggi:  5,
	}
	fmt.Printf("var02B: %v\n", var02B)
	fmt.Printf("var02B luas permukaan: %v\n", var02B.luasPermukaan())
	fmt.Printf("var02B volume: %v\n", var02B.volume())
	var02B.upsize(3)
	fmt.Printf("var02B upsize: %v\n", var02B)
	fmt.Printf("var02B luas permukaan: %v\n", var02B.luasPermukaan())
	fmt.Printf("var02B volume: %v\n", var02B.volume())
	var var02C *balok
	fmt.Printf("var02C: %v\n", var02C)
	var02C = &var02A
	fmt.Printf("var02C: %v\n", var02C)
	var02C.upsize(5)
	fmt.Printf("var02C: %v\n", var02C)
	fmt.Printf("var02C luas permukaan: %v\n", var02C.luasPermukaan())
	fmt.Printf("var02C volume: %v\n", var02C.volume())
	fmt.Printf("var02A: %v\n", var02A)
	var var02D *balok
	fmt.Printf("var02D: %v\n", var02D)
	var02D = var02B
	fmt.Printf("var02D: %v\n", var02D)
	fmt.Printf("var02D luas permukaan: %v\n", var02D.luasPermukaan())
	fmt.Printf("var02D volume: %v\n", var02D.volume())
	fmt.Println("")

	fmt.Println("25.3 Pointer sebagai parameter fungsi")
	var03A := balok{9, 7, 12}
	fmt.Printf("var03A: %v\n", var03A)
	upsize(&var03A, 2)
	fmt.Printf("var03A upsize: %v\n", var03A)
	fmt.Printf("var03A luas permukaan: %v\n", luasPermukaan(&var03A))
	fmt.Printf("var03A volume: %v\n", volume(var03A))
	var03B := &balok{7, 2, 5}
	fmt.Printf("var03B: %v\n", var03B)
	upsize(var03B, 3)
	fmt.Printf("var03B: %v\n", var03B)
	fmt.Printf("var03B luas permukaan: %v\n", luasPermukaan(var03B))
	fmt.Printf("var03B volume: %v\n", volume(*var03B))
	fmt.Println("")

}

func (s person) func01A(message string) {
	fmt.Println(message, s.name)
}

func (s person) func01B(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (b *balok) upsize(i int) {
	b.panjang = i * b.panjang
	b.lebar = i * b.lebar
	b.tinggi = i * b.tinggi
}

func (b balok) luasPermukaan() int {
	hasil := (2 * b.panjang * b.lebar) + (2 * b.lebar * b.tinggi) + (2 * b.panjang * b.tinggi)
	return hasil
}

func (b balok) volume() int {
	hasil := b.panjang * b.lebar * b.tinggi
	return hasil
}

func upsize(b *balok, i int) {
	b.panjang = b.panjang * i
	b.lebar = b.lebar * i
	b.tinggi = b.tinggi * i
}

func luasPermukaan(b *balok) int {
	hasil := (2 * b.panjang * b.lebar) + (2 * b.lebar * b.tinggi) + (2 * b.panjang * b.tinggi)
	return hasil
}

func volume(b balok) int {
	hasil := b.panjang * b.lebar * b.tinggi
	return hasil
}
