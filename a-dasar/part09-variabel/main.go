package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 09 Variabel")

	var firstName = "Novalita"
	middleName := "Kusuma"
	var lastName string
	lastName = "Wardhana"
	fmt.Printf("Nama saya: %v %v %v\n", firstName, middleName, lastName)

	var notebookOne, notebookTwo, notebookThree string
	notebookOne = "Dell"
	notebookTwo = "Fujitsu"
	notebookThree = "Lenovo"
	fmt.Printf("Laptop saya: %v %v %v\n", notebookOne, notebookTwo, notebookThree)

	menuOne, menuTwo, menuThree := "Sate", "Mie Ayam", "Salad"
	fmt.Printf("Makanan favorit saya: %v %v %v\n", menuOne, menuTwo, menuThree)

	var musicOne, musicTwo, musicThree string
	musicOne, musicTwo, musicThree = "Andra & The Backbone", "One Ok Rock", "Radwimps"
	fmt.Printf("Musik favorit: %v %v %v\n", musicOne, musicTwo, musicThree)

	var smartphone, _ = "Xiaomi", "Samsung"
	fmt.Printf("Merk smartphone saya: %v\n", smartphone)
	fmt.Println("")

	var pointerOne = new(string)
	fmt.Printf("Nilai pointer pointerOne: %v\n", pointerOne)
	fmt.Printf("Nilai variabel pointerOne: %v\n", *pointerOne)
	*pointerOne = "Haloo"
	fmt.Printf("Nilai pointer pointerOne saat ini: %v\n", pointerOne)
	fmt.Printf("Nilai variabel pointerOne saat ini: %v\n", *pointerOne)
}
