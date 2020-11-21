package main

import (
	"fmt"
)

type hitung interface {
	luasPermukaan() int
	volume() int
	upsize(s int)
}

type bangunRuang interface {
	bangunRuangUpsize
	bangunRuangLuasVolume
}

type bangunRuangUpsize interface {
	upsize(s int)
}

type bangunRuangLuasVolume interface {
	luasPermukaan() int
	volume() int
}

type balok struct {
	panjang int
	lebar   int
	tinggi  int
}

func (b *balok) luasPermukaan() int {
	result := (2 * b.panjang * b.lebar) + (2 * b.lebar * b.tinggi) + (2 * b.panjang * b.tinggi)
	return result
}

func (b *balok) volume() int {
	result := b.panjang * b.lebar * b.tinggi
	return result
}

func (b *balok) upsize(s int) {
	b.panjang = s * b.panjang
	b.lebar = s * b.lebar
	b.tinggi = s * b.tinggi
}

type kubus struct {
	s int
}

func (k kubus) luasPermukaan() int {
	result := 6 * k.s * k.s
	return result
}

func (k kubus) volume() int {
	result := k.s * k.s * k.s
	return result
}

func (k kubus) upsize(s int) {
	k.s = s * k.s
}

func main() {
	fmt.Println("Part 27 Interface")
	fmt.Println("")

	fmt.Println("27.1 Penerapan interface")
	var infc01A hitung
	infc01A = &balok{9, 8, 5}
	fmt.Printf("infc01A: %v\n", infc01A)
	fmt.Printf("infc01A luas permukaan: %v\n", infc01A.luasPermukaan())
	fmt.Printf("infc01A volume: %v\n", infc01A.volume())
	infc01A.upsize(2)
	fmt.Printf("infc01A: %v\n", infc01A)
	fmt.Printf("infc01A luas permukaan: %v\n", infc01A.luasPermukaan())
	fmt.Printf("infc01A volume: %v\n", infc01A.volume())
	var infc01B hitung
	infc01B = kubus{9}
	fmt.Printf("infc01B: %v\n", infc01B)
	fmt.Printf("infc01B luas permukaan: %v\n", infc01B.luasPermukaan())
	fmt.Printf("infc01B volume: %v\n", infc01B.volume())
	infc01B.upsize(4)
	fmt.Printf("infc01B: %v\n", infc01B)
	fmt.Printf("infc01B luas permukaan: %v\n", infc01B.luasPermukaan())
	fmt.Printf("infc01B volume: %v\n", infc01B.volume())
	var infc01C hitung
	infc01C = &balok{12, 5, 8}
	fmt.Printf("infc01C: %v\n", infc01C)
	fmt.Printf("infc01C luas permukaan: %v\n", infc01C.luasPermukaan())
	fmt.Printf("infc01C volume: %v\n", infc01C.volume())
	infc01C.upsize(2)
	fmt.Printf("infc01C: %v\n", infc01C)
	fmt.Printf("infc01C luas permukaan: %v\n", infc01C.luasPermukaan())
	fmt.Printf("infc01C volume: %v\n", infc01C.volume())
	fmt.Println("")

	fmt.Println("Part 27.2 embedded interface")
	var infc02A bangunRuang
	infc02A = &balok{19, 3, 7}
	fmt.Printf("infc02A: %v\n", infc02A)
	fmt.Printf("infc02A luas permukaan: %v\n", infc02A.luasPermukaan())
	fmt.Printf("infc02A volume: %v\n", infc02A.volume())
	infc02A.upsize(3)
	fmt.Printf("infc02A: %v\n", infc02A)
	fmt.Printf("infc02A luas permukaan: %v\n", infc02A.luasPermukaan())
	fmt.Printf("infc02A volume: %v\n", infc02A.volume())
}
