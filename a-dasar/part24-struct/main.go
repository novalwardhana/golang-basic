package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 24 Struct")
	fmt.Println("")

	fmt.Println("24.1 Deklarasi struct")
	type Struct01 struct {
		name string
		age  int
	}
	var01A := Struct01{
		name: "Noval",
		age:  25,
	}
	fmt.Printf("var01A: %v\n", var01A)
	fmt.Printf("var01A name: %v\n", var01A.name)
	fmt.Printf("var01A age: %v\n", var01A.age)
	fmt.Printf("var01A address: %v\n", &var01A)
	fmt.Printf("var01A address name: %v\n", &var01A.name)
	fmt.Printf("var01A address age: %v\n", &var01A.age)
	fmt.Println("")

	fmt.Println("24.2 Penerapan struct")
	type Struct02 struct {
		name  string
		grade int
	}
	var02A := Struct02{}
	var02A.name = "Deni"
	var02A.grade = 2
	fmt.Printf("var02A: %v\n", var02A)
	var02B := &Struct02{
		name:  "Novan",
		grade: 3,
	}
	fmt.Printf("var02B: %v\n", var02B)
	fmt.Printf("var02B name: %v\n", var02B.name)
	fmt.Printf("var02B grade: %v\n", var02B.grade)
	fmt.Println("")

	fmt.Println("24.3 Inisiasi object struct")
	type Struct03 struct {
		name    string
		address string
		age     int
	}
	var03A := Struct03{}
	var03A.name = "Deni"
	var03A.address = "Sleman"
	var03A.age = 27
	fmt.Printf("var03A: %v\n", var03A)
	var03B := Struct03{
		name:    "Novan",
		address: "Gunungkidul",
		age:     28,
	}
	fmt.Printf("var03B: %v\n", var03B)
	var03C := Struct03{"Adrie", "Bengkulu", 29}
	fmt.Printf("var03C: %v\n", var03C)
	fmt.Println("")

	fmt.Println("24.4 Variabel objek pointer")
	type Struct04 struct {
		name        string
		merk        string
		releaseDate int
	}
	var var04A *Struct04
	fmt.Printf("var04A: %v\n", var04A)
	var04B := Struct04{"Redmi Note 8", "Xiaomi", 2019}
	fmt.Printf("var04B: %v\n", var04B)
	fmt.Printf("var04B Address: %v\n", &var04B)
	fmt.Printf("var04B Address name: %v\n", &var04B.name)
	fmt.Printf("var04B Address merk: %v\n", &var04B.merk)
	fmt.Printf("var04B Address releaseDate: %v\n", &var04B.releaseDate)
	var04A = &var04B
	fmt.Printf("var04A: %v\n", var04A)
	fmt.Printf("var04A Address: %v\n", &var04A)
	fmt.Printf("var04A Address name: %v\n", &var04A.name)
	fmt.Printf("var04A Address merk: %v\n", &var04A.merk)
	fmt.Printf("var04A Address releaseDate: %v\n", &var04A.releaseDate)
	fmt.Println("")

	fmt.Println("24.5 Embedded struct")
	type Struct05Person struct {
		name string
		age  int
	}
	type Struct05 struct {
		grade int
		Struct05Person
	}
	var05A := Struct05{
		grade: 2,
	}
	var05A.name = "Adrie"
	var05A.age = 29
	fmt.Printf("var05A: %v\n", var05A)
	var05B := Struct05{}
	var05B.name = "Deni"
	var05B.Struct05Person.age = 27
	var05B.grade = 2
	fmt.Printf("var05B: %v\n", var05B)
	var05C := Struct05{
		grade: 4,
		Struct05Person: Struct05Person{
			name: "Novan",
			age:  28,
		},
	}
	fmt.Printf("var05C: %v\n", var05C)
	fmt.Println("")

	fmt.Println("24.6 Embedded struct dengan nama property yang sama")
	type Struct06Prodi struct {
		name  string
		grade int
	}
	type Struct06 struct {
		name    string
		grade   int
		address string
		Struct06Prodi
	}
	var06A := Struct06{}
	var06A.name = "Novan"
	var06A.grade = 2
	var06A.address = "Gunungkidul"
	var06A.Struct06Prodi.name = "Pendidikan luar sekolah"
	var06A.Struct06Prodi.grade = 3
	fmt.Printf("var06A: %v\n", var06A)
	var06B := Struct06{}
	var06B.name = "Deni"
	var06B.grade = 2
	var06B.address = "Sleman"
	var06B.Struct06Prodi.name = "Teknologi pendidikan"
	var06B.Struct06Prodi.grade = 1
	fmt.Printf("var06B: %v\n", var06B)
	fmt.Println("")

	fmt.Println("24.7 Pengsian nilai sub struct")
	type Struct07Club struct {
		name  string
		state string
		home  string
	}
	type Struct07 struct {
		name string
		age  int
		Struct07Club
	}
	var07AClub := Struct07Club{
		name:  "Manchester United",
		state: "England",
		home:  "Old Trafford",
	}
	var07A := Struct07{
		name:         "David De Gea",
		age:          29,
		Struct07Club: var07AClub,
	}
	fmt.Printf("var07A: %v\n", var07A)
	var07BClub := Struct07Club{"Real Madrid", "Spain", "Santiago Bernabeu"}
	var07B := Struct07{
		name:         "Sergio Ramos",
		age:          30,
		Struct07Club: var07BClub,
	}
	fmt.Printf("var07B: %v\n", var07B)
	fmt.Println("")

	fmt.Println("24.8 Anonymous Struct")
	var08A := struct {
		name    string
		merk    string
		release int
	}{}
	var08A.name = "Lenovo Slim 3i"
	var08A.merk = "Lenovo"
	var08A.release = 2020
	fmt.Printf("var08A: %v\n", var08A)
	var08B := struct {
		panjang int
		lebar   int
		tinggi  int
	}{
		panjang: 10,
		lebar:   8,
		tinggi:  6,
	}
	fmt.Printf("var08B: %v\n", var08B)
	type Struct08 struct {
		name   string
		madeIn string
	}
	var08CStruct08 := Struct08{
		name:   "Asus",
		madeIn: "Taiwan",
	}
	var08C := struct {
		series string
		Struct08
	}{
		series:   "Asus Zenbook UM431",
		Struct08: var08CStruct08,
	}
	fmt.Printf("var08C: %v\n", var08C)
	var08D := struct {
		series string
		Struct08
	}{
		series: "Acer 431",
		Struct08: Struct08{
			name:   "Acer",
			madeIn: "Taiwan",
		},
	}
	fmt.Printf("var08D: %v\n", var08D)
	fmt.Println("")

	fmt.Println("24.9 Kombinasi dan slice dan struct")
	type Struct09 struct {
		panjang int
		lebar   int
		tinggi  int
	}
	var09A := []Struct09{
		Struct09{panjang: 9, lebar: 8, tinggi: 12},
		Struct09{panjang: 12, lebar: 7, tinggi: 8},
		Struct09{panjang: 12, lebar: 5, tinggi: 9},
	}
	fmt.Printf("var09A: %v\n", var09A)
	var09B := []Struct09{}
	for i := 1; i <= 10; i++ {
		var09BDetail := Struct09{
			panjang: 2 * i,
			lebar:   3 * i,
			tinggi:  4 * i,
		}
		var09B = append(var09B, var09BDetail)
	}
	fmt.Printf("var09B: %v\n", var09B)
	fmt.Println("")

	fmt.Println("24.10 Inisiasi anonymous struct")
	var10A := []struct {
		panjang int
		lebar   int
		tinggi  int
	}{
		{9, 7, 10},
		{8, 4, 9},
		{12, 2, 15},
	}
	fmt.Printf("var10A: %v\n", var10A)
	var10B := []struct {
		panjang int
		lebar   int
		tinggi  int
	}{}
	for i := 1; i < 10; i++ {
		var10BDetail := struct {
			panjang int
			lebar   int
			tinggi  int
		}{
			panjang: 3 * i,
			lebar:   4 * i,
			tinggi:  5 * i,
		}
		var10B = append(var10B, var10BDetail)
	}
	fmt.Printf("var10B: %v\n", var10B)
	fmt.Println("")

	fmt.Println("24.11 Deklarasi struct dengan var")
	var var11A struct {
		name    string
		address string
	}
	fmt.Printf("var11A: %v\n", var11A)
	var11A.name = "Noval"
	var11A.address = "Bantul"
	fmt.Printf("var11A: %v\n", var11A)
	type Struct11 struct {
		name    string
		address string
	}
	var var11B struct {
		grade int
		Struct11
	}
	fmt.Printf("var11B: %v\n", var11B)
	var11B.grade = 5
	var11B.Struct11.name = "Noval"
	var11B.Struct11.address = "Bantul"
	fmt.Printf("var11B: %v\n", var11B)
	fmt.Println("")

	fmt.Println("24.12 Nested struct")
	type Struct12A struct {
		person struct {
			name    string
			address string
		}
		grade   int
		hobbies []string
	}
	var12A := Struct12A{
		grade:   3,
		hobbies: []string{"Bike", "Football", "Coding", "Hiking"},
	}
	var12A.person.name = "Noval"
	var12A.person.address = "Bantul"
	fmt.Printf("var12A: %v\n", var12A)
	type Struct12B struct {
		name       string
		categories string
		element    []struct {
			panjang int
			lebar   int
			tinggi  int
		}
	}
	var12B := Struct12B{
		name:       "Balok",
		categories: "Bangun ruang",
	}
	for i := 1; i < 20; i++ {
		var12BElement := struct {
			panjang int
			lebar   int
			tinggi  int
		}{
			panjang: 5 * i,
			lebar:   6 * i,
			tinggi:  7 * i,
		}
		var12B.element = append(var12B.element, var12BElement)
	}
	fmt.Printf("var12B: %v\n", var12B)
	fmt.Println("")

	fmt.Println("24.13 Deklarasi dan inisiai struct secara horizontal")
	type Struct13A struct {
		name    string
		address string
		age     int
		hobbies []string
	}
	var13A := Struct13A{
		name:    "Intrada",
		address: "Sewon Bantul",
		age:     28,
		hobbies: []string{"Mathematics", "Music", "Drawing"},
	}
	fmt.Printf("var13A: %v\n", var13A)
	var13B := struct {
		name    string
		address string
		grade   int
	}{"Deni", "Sleman", 27}
	fmt.Printf("var13B: %v\n", var13B)
	var13C := struct {
		name    string
		age     int
		hobbies []string
	}{"novan", 28, []string{"Music", "Video Editor"}}
	fmt.Printf("var13C: %v\n", var13C)
	fmt.Println("")

	fmt.Println("24.14 Tag property dalam struct")
	type Struct14A struct {
		name    string
		grade   int
		address string
		hobbies []string
	}
	var14A := Struct14A{
		name:    "Novan",
		grade:   12,
		address: "Gunungkidul",
		hobbies: []string{"Music", "Touring", "Video Editor"},
	}
	fmt.Printf("var14A: %v\n", var14A)
	var14B := struct {
		name       string
		categories string
		element    []struct {
			alas   int
			tinggi int
		}
	}{}
	var14B.name = "Segitiga"
	var14B.categories = "Bangun datar"
	for i := 1; i < 25; i++ {
		var14BElement := struct {
			alas   int
			tinggi int
		}{
			alas:   5 * i,
			tinggi: 8 * i,
		}
		var14B.element = append(var14B.element, var14BElement)
	}
	fmt.Printf("var14B: %v\n", var14B)
	fmt.Println("")

	fmt.Println("24.15 Type Alias")
	type Struct15A struct {
		alas   float64
		tinggi float64
	}
	type Struct15AAlias = Struct15A
	var15A := Struct15AAlias{
		alas:   12,
		tinggi: 5,
	}
	fmt.Printf("var15A: %v\n", var15A)
	type Struct15BElement struct {
		panjang int
		lebar   int
		tinggi  int
	}
	type Struct15B struct {
		name       string
		categories string
		element    []Struct15BElement
	}
	type Struct15BAlias = Struct15B
	var15B := Struct15BAlias{
		name:       "Balok",
		categories: "Bangun Ruang",
	}
	for i := 1; i <= 25; i++ {
		var15BElement := Struct15BElement{
			panjang: 7 * i,
			lebar:   5 * i,
			tinggi:  9 * i,
		}
		var15B.element = append(var15B.element, var15BElement)
	}
	fmt.Printf("var15B: %v\n", var15B)
}
