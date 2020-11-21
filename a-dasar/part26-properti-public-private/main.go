package main

import (
	"fmt"
	"part26-properti-public-private/balok"
	"part26-properti-public-private/kubus"
	lib "part26-properti-public-private/library"
)

func main() {
	fmt.Println("Part 26 properti public private")
	lib.Message()
	fmt.Println(lib.Introduce("Noval"))
	fmt.Println("")
	balok01 := balok.Balok{
		Panjang: 12,
		Lebar:   10,
		Tinggi:  9,
	}
	fmt.Printf("balok01: %v\n", balok01)
	fmt.Printf("Luas Permukaan: %v\n", balok01.LuasPermukaan())
	fmt.Printf("Volum: %v\n", balok01.Volume())
	balok01.Upsize(3)
	fmt.Printf("balok01: %v\n", balok01)
	fmt.Printf("Luas Permukaan: %v\n", balok01.LuasPermukaan())
	fmt.Printf("Volum: %v\n", balok01.Volume())
	fmt.Println("")

	//SayHello("Alex")
	fmt.Printf("kubus01: %v\n", kubus.Kubus)
}
