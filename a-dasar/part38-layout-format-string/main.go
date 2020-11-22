package main

import (
	"fmt"
)

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "Noval",
	height:      167.2,
	age:         25,
	isGraduated: true,
	hobbies:     []string{"Coding", "Jogging", "Touring"},
}

type balok struct {
	panjang int
	lebar   int
	tinggi  int
}

var data02 = balok{10, 8, 6}

type user struct {
	id       int
	username string
	name     string
	email    string
	role     []string
}

var data03 = user{
	id:       12,
	username: "novalwardhana",
	name:     "Noval",
	email:    "novalita.k.wardhana@gmail.com",
	role:     []string{"Admin", "Superadmin"},
}

var data04 = struct {
	player string
	club   string
}{
	player: "Toni Kross",
	club:   "Real Madrid",
}

func main() {
	fmt.Println("Part 38 Layout Format String")
	fmt.Println("")

	fmt.Println("Part 38.1 Persiapan")
	fmt.Printf("Data: %v\n", data)
	fmt.Println("")

	fmt.Println("Part 38.2 Format biner") // Harus dalam format numeric, iteger, atau float64
	fmt.Printf("Data height: %b\n", data.height)
	fmt.Printf("Data age: %b\n", data.age)
	fmt.Printf("Data 02 panjang: %b\n", data02.panjang)
	fmt.Printf("Data 02 lebar: %b\n", data02.lebar)
	fmt.Printf("Data 02 tinggi: %b\n", data02.tinggi)
	fmt.Println("")

	fmt.Println("Part 38.3 Format unicode")
	fmt.Printf("1400: %c\n", 1400)
	fmt.Printf("1500: %c\n", 1500)
	fmt.Printf("1600: %c\n", 1600)
	fmt.Printf("1800: %c\n", 1800)
	fmt.Println("")

	fmt.Println("38.4 Format decimal") // Harus dalam format integer
	fmt.Printf("Data age: %d\n", data.age)
	fmt.Printf("125: %d\n", 125)
	fmt.Printf("155: %d\n", 155)
	fmt.Println("")

	fmt.Println("Part 38.5 Format scientific") // Harus dalam format float64
	fmt.Printf("Data height: %e\n", data.height)
	fmt.Printf("Data age: %E\n", float64(data.age))
	fmt.Printf("125.5: %e\n", 125.5)
	fmt.Printf("155.5: %E\n", 155.5)
	fmt.Println("")

	fmt.Println("Part 38.6 Format decimal dengan kustom precision") // Harus dalam format float64
	fmt.Printf("Data height: %.2f\n", data.height)
	fmt.Printf("Data age: %.2f\n", float64(data.age))
	fmt.Printf("125.375: %.3f\n", 125.375)
	fmt.Printf("155.2: %.1f\n", 155.2)
	fmt.Printf("Data 02 panjang: %.2f\n", float64(data02.panjang))
	fmt.Printf("Data 02 lebar: %.3f\n", float64(data02.lebar))
	fmt.Printf("Data02 tinggi: %.4f\n", float64(data02.tinggi))
	fmt.Println("")

	fmt.Println("Part 38.7 Format decimal tanpa precision") // Harus dalam format float64
	fmt.Printf("Data height: %g\n", data.height)
	fmt.Printf("Data age: %g\n", float64(data.age))
	fmt.Printf("189.21234: %g\n", 189.21234)
	fmt.Printf("1919.2: %g\n", 1919.2)
	fmt.Println("")

	fmt.Println("Part 38.8 Layout format numerik menjadi oktal string") // Harus dalam format integer
	fmt.Printf("Data height: %o\n", int(data.height))
	fmt.Printf("Data age: %o\n", data.age)
	fmt.Printf("85: %o\n", 85)
	fmt.Printf("93: %o\n", 93)
	fmt.Println("")

	fmt.Println("Part 38.9 Layout format pointer")
	fmt.Printf("Data: %p\n", &data)
	fmt.Printf("Data name: %p\n", &data.name)
	fmt.Printf("Data isGraduated: %p\n", &data.isGraduated)
	fmt.Printf("Data hobbies: %p\n", &data.hobbies)
	fmt.Printf("Data 03: %p\n", &data03)
	fmt.Printf("Data 03 id: %p\n", &data03.id)
	fmt.Printf("Data 03 username: %p\n", &data03.username)
	fmt.Printf("Data 03 email: %p\n", &data03.role)
	fmt.Println("")

	fmt.Println("Part 38.10 Layout format escape")
	fmt.Printf("%q\n", `" name \ height "`)
	fmt.Printf("%q\n", "Hello World")
	fmt.Println("")

	fmt.Println("Part 38.11 Layout format string")
	fmt.Printf("Data name: %s\n", data.name)
	fmt.Printf("RM: %s\n", "Real Madrid")
	fmt.Printf("MU: %s\n", "Manchester United")
	fmt.Printf("Barca: %s\n", "Barcelona")
	fmt.Println("")

	fmt.Println("Part 38.12 Layout format boolean")
	fmt.Printf("Data isGraduated: %t\n", data.isGraduated)
	fmt.Printf("Bayern Munich won champions league: %t\n", true)
	fmt.Printf("Paris Saint Germain won champions league: %t\n", false)
	fmt.Printf("Real Madrid won LaLiga: %t\n", true)
	fmt.Println("")

	fmt.Println("Part 38.13 Layout format tipe variabel")
	fmt.Printf("Data: %T\n", data)
	fmt.Printf("Data name: %T\n", data.name)
	fmt.Printf("Data height: %T\n", data.height)
	fmt.Printf("Data age: %T\n", data.age)
	fmt.Printf("Data isGraduated: %T\n", data.isGraduated)
	fmt.Printf("Data hobbies: %T\n", data.hobbies)
	fmt.Printf("Data 04 player: %T\n", data04.player)
	fmt.Printf("Data 04 club: %T\n", data04.club)
	fmt.Println("")

	fmt.Println("Part 38.14 Layout format interface")
	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Data name: %v\n", data.name)
	fmt.Printf("Data isGraduated: %v\n", data.isGraduated)
	fmt.Printf("Data hobbies: %v\n", data.hobbies)
	fmt.Println("")

	fmt.Println("Part 38.15 Layout format berurutan")
	fmt.Printf("Data: %+v\n", data)
	fmt.Printf("Data 02: %+v\n", data02)
	fmt.Printf("Data 03: %+v\n", data03)
	fmt.Printf("Data 04: %+v\n", data04)
	fmt.Println("")

	fmt.Println("Part 38.16 Layout format sesuai struktur package")
	fmt.Printf("Data: %#v\n", data)
	fmt.Printf("Data 02: %#v\n", data02)
	fmt.Printf("Data 03: %#v\n", data03)
	fmt.Printf("Data 04: %#v\n", data04)
	fmt.Println("")

	fmt.Println("Part 38.17 Layout format hexadecimal") //Harus dalam bentuk integer
	fmt.Printf("Data age: %x\n", data.age)
	fmt.Printf("25: %x\n", 25)
	fmt.Printf("125: %x\n", 125)
	fmt.Printf("155: %x\n", 155)
	fmt.Printf("625: %x\n", 625)
	var17A := data.name
	fmt.Printf("%x %x %x %x\n", var17A[0], var17A[1], var17A[2], var17A[3])
	var17B := data04.player
	fmt.Printf("%x %x %x %x %x\n", var17B[0], var17B[1], var17B[2], var17B[3], var17B[4])
	fmt.Printf("Data 02 panjang: %x\n", data02.panjang)
	fmt.Printf("Data 02 lebar: %x\n", data02.lebar)
	fmt.Printf("Data 02 tinggi: %x\n", data02.tinggi)
	fmt.Println("")

	fmt.Println("Part 37.18 Menulis tanda persen")
	fmt.Printf("%%\n")
}
