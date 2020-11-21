package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 16 Slice")
	fmt.Println("")

	fmt.Println("16.1 Inisiai Slice")
	slc01 := []string{"Thinkpad", "Ideapad", "Legion"}
	for i, v := range slc01 {
		fmt.Printf("Data ke-%v: %v\n", i, v)
	}
	fmt.Println("")

	fmt.Println("16.2 Hubungan slice dengan array & operasi slice")
	slc02 := []string{"Java", "Golang", "PHP", "Javascript", "Clojure", "C", "C++"}
	fmt.Printf("slc02[0:2]: %v\n", slc02[0:2])
	fmt.Printf("slc02[3:7]: %v\n", slc02[3:7])
	fmt.Printf("slc02[0:0]: %v\n", slc02[0:0])
	fmt.Printf("slc02[:]: %v\n", slc02[:])
	fmt.Printf("slc02[3:]: %v\n", slc02[3:])
	fmt.Printf("slc02[:4]: %v\n", slc02[:4])
	fmt.Println("")

	fmt.Println("16.3 slice merupakan tipe data reference")
	slc03 := []string{"Compass", "Ventela", "Piero", "Patrobas", "Warrior"}
	fmt.Printf("slc03: %v\n", slc03)
	slc03A := slc03[0:2]
	fmt.Printf("slc03A: %v\n", slc03A)
	slc03A[0] = "Vans"
	fmt.Printf("slc03A: %v\n", slc03A)
	fmt.Printf("slc03: %v\n", slc03)
	slc03B := slc03[2:]
	fmt.Printf("Slc03B: %v\n", slc03B)
	slc03B[0] = "Adidas"
	fmt.Printf("slc03: %v\n", slc03)
	fmt.Println("")

	fmt.Println("16.4 Fungsi len")
	slc04 := []int{1, 3, 5, 7, 9, 11, 13, 15, 18, 19}
	fmt.Printf("slc04: %v\n", slc04)
	fmt.Printf("Panjang slc04: %v\n", len(slc04))
	fmt.Println("")

	fmt.Println("16.5 Fungsi Cap")
	slc05A := []string{"Thinkpad", "Ideapad", "Legion", "Yoga"}
	fmt.Printf("slc05A: %v\n", slc05A)
	fmt.Printf("slc05A len: %v\n", len(slc05A))
	fmt.Printf("slc05A cap: %v\n", cap(slc05A))
	slc05A01 := slc05A[0:3]
	fmt.Printf("slc05A01: %v\n", slc05A01)
	fmt.Printf("slc05A01 len: %v\n", len(slc05A01))
	fmt.Printf("slc05A01 cap: %v\n", cap(slc05A01))
	slc05A02 := slc05A[1:4]
	fmt.Printf("slc05A02: %v\n", slc05A02)
	fmt.Printf("slc05A02 len: %v\n", len(slc05A02))
	fmt.Printf("slc05A02 cap: %v\n", cap(slc05A02))
	slc05B := []string{"Asus", "Lenovo", "Huawei", "Samsung", "Sony", "HP"}
	fmt.Printf("slc05B: %v\n", slc05B)
	fmt.Printf("slc05B len: %v\n", len(slc05B))
	fmt.Printf("slc05B cap: %v\n", cap(slc05B))
	slc05B01 := slc05B[0:5]
	fmt.Printf("slc05B01: %v\n", slc05B01)
	fmt.Printf("slc05B01 len: %v\n", len(slc05B01))
	fmt.Printf("slc05B01 cap: %v\n", cap(slc05B01))
	slc05B02 := slc05B[1:6]
	fmt.Printf("slc05B02: %v\n", slc05B02)
	fmt.Printf("slc05B02 len: %v\n", len(slc05B02))
	fmt.Printf("slc05B02 cap: %v\n", cap(slc05B02))
	fmt.Println("")

	fmt.Println("16.6 Fungsi append")
	slc06A := []string{"Bantul", "Gunungkidul", "Kota Yogyakarta"}
	fmt.Printf("slc06A: %v\n", slc06A)
	slc06A = append(slc06A, "Kulon Progo")
	fmt.Printf("slc06A: %v\n", slc06A)
	slc06A = append(slc06A, "Sleman")
	fmt.Printf("slc06A: %v\n", slc06A)
	slc06B := []int{}
	fmt.Printf("slc06B: %v\n", slc06B)
	for i := 0; i < 15; i++ {
		slc06B = append(slc06B, i)
	}
	fmt.Printf("slc06B: %v\n", slc06B)
	fmt.Println("")

	fmt.Println("16.7 Copy")
	slc07A := make([]string, 2)
	slc07B := []string{"Asus", "Lenovo", "Huawei", "Samsung", "Sony", "HP"}
	fmt.Printf("slc06A: %v\n", slc07A)
	fmt.Printf("slc06B: %v\n", slc07B)
	copy(slc07A, slc07B)
	fmt.Printf("slc06A: %v\n", slc07A)
	fmt.Printf("slc06B: %v\n", slc07B)
	slc07C := []string{"Toshiba", "IBM"}
	fmt.Printf("slc07C: %v\n", slc07C)
	copy(slc07C, slc07A)
	fmt.Printf("slc07C: %v\n", slc07C)
	fmt.Println("")

	fmt.Println("16.8 Pengaksesan elemen slice dengan 3 indeks")
	slc08 := []string{"Real Madrid", "Manchester United", "Arsenal", "Juventus", "Chelsea"}
	fmt.Printf("slc08: %v\n", slc08)
	fmt.Printf("slc08 len: %v\n", len(slc08))
	fmt.Printf("slc08 cap: %v\n", cap(slc08))
	slc08A := slc08[2:4]
	fmt.Printf("slc08A: %v\n", slc08A)
	fmt.Printf("slc08A len: %v\n", len(slc08A))
	fmt.Printf("slc08A cap: %v\n", cap(slc08A))
	slc08B := slc08[0:2]
	fmt.Printf("slc08B: %v\n", slc08B)
	fmt.Printf("slc08B len: %v\n", len(slc08B))
	fmt.Printf("slc08B cap: %v\n", cap(slc08B))
	slc08C := slc08[0:2:2]
	fmt.Printf("slc08C: %v\n", slc08C)
	fmt.Printf("slc08C len: %v\n", len(slc08C))
	fmt.Printf("slc08C cap: %v\n", cap(slc08C))

}
