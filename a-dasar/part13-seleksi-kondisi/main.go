package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 13 Seleksi Kondisi")
	fmt.Println("")

	fmt.Println("Seleksi kondisi dengan if else-if else")
	condOne := 12
	if condOne > 20 {
		fmt.Printf("%v > 20\n", condOne)
	} else if condOne > 15 {
		fmt.Printf("%v > 15\n", condOne)
	} else if condOne > 10 {
		fmt.Printf("%v > 10\n", condOne)
	} else {
		fmt.Printf("%v < 10\n", condOne)
	}
	fmt.Println("")

	fmt.Println("Variabel temporary if else-if else")
	condTwo := 21
	if condTwoResult := condTwo % 2; condTwoResult != 0 {
		fmt.Printf("%v adalah bilangan ganjil\n", condTwo)
	} else if condTwoResult := condTwo % 2; condTwoResult == 0 {
		fmt.Printf("%v adalah bilangan genap\n", condTwo)
	} else {
		fmt.Printf("%v tak terdefinisi\n", condTwo)
	}
	fmt.Println("")

	fmt.Println("Switch-case")
	condThree := 8
	switch condThree {
	case 1:
		fmt.Printf("%v is one\n", condThree)
	case 7:
		fmt.Printf("%v is seven\n", condThree)
	case 8:
		fmt.Printf("%v is eight\n", condThree)
	case 9:
		fmt.Printf("%v is nine\n", condThree)
	case 10:
		fmt.Printf("%v is ten\n", condThree)
	default:
		fmt.Printf("%v is not found\n", condThree)
	}
	fmt.Println("")

	fmt.Println("Switch-case untuk banyak kondisi")
	condFour := 9
	switch condFour {
	case 1, 2, 3:
		{
			fmt.Printf("%v is low\n", condFour)
		}
	case 4, 5, 6:
		fmt.Printf("%v is middle\n", condFour)
	case 7, 8, 9:
		{
			fmt.Printf("%v is heigh\n", condFour)
		}
	default:
		fmt.Printf("%v is not definited\n", condFour)
	}
	fmt.Println("")

	fmt.Println("Switch-case dengan kurung kurawal")
	condFive := 97
	switch condFive {
	case 81, 82, 83, 84, 85:
		{
			fmt.Printf("%v is good score\n", condFive)
		}
	case 86, 87, 88, 89, 90, 91, 92, 93, 94, 95:
		{
			fmt.Printf("%v is excellent\n", condFive)
		}
	case 96, 97, 98, 99, 100:
		{
			fmt.Printf("%v is perfect\n", condFive)
		}
	default:
		{
			fmt.Printf("%v Not Bad\n", condFive)
		}
	}
	fmt.Println("")

	fmt.Println("Switch-case dengan gaya if-else")
	condSix := 85
	switch {
	case condSix >= 90:
		{
			fmt.Printf("%v score is perfect\n", condSix)
		}
	case condSix >= 75:
		{
			fmt.Printf("%v score is excellent\n", condSix)
		}
	case condSix >= 65:
		{
			fmt.Printf("%v score is not bad\n", condSix)
		}
	default:
		{
			fmt.Printf("%v score is bad\n", condSix)
		}
	}
	fmt.Println("")

	fmt.Println("Switch-case dengan fallthrough")
	condSeven := 85
	switch {
	case condSeven >= 90:
		{
			fmt.Printf("%v score is perfect\n", condSeven)
		}
	case condSeven >= 85:
		{
			fmt.Printf("%v score is almost-perfect\n", condSeven)
		}
		fallthrough
	case condSeven >= 75:
		{
			fmt.Printf("%v score is excellent\n", condSeven)
		}
	case condSeven >= 65:
		{
			fmt.Printf("%v score is not bad\n", condSeven)
		}
	default:
		{
			fmt.Printf("%v score is bad\n", condSeven)
		}
	}
	fmt.Println("")

	fmt.Println("Seleksi kondisi bersarang")
	condEight := 86
	if condEightResult := condEight % 2; condEightResult == 1 {
		switch {
		case condEight >= 90:
			{
				fmt.Printf("%v score is perfect\n", condEight)
			}
		case condEight >= 85:
			{
				fmt.Printf("%v score is almost-perfect\n", condEight)
			}
			fallthrough
		case condEight >= 75:
			{
				fmt.Printf("%v score is excellent\n", condEight)
			}
		case condEight >= 65:
			{
				fmt.Printf("%v score is not bad\n", condEight)
			}
		default:
			{
				fmt.Printf("%v score is bad\n", condEight)
			}
		}
	} else if condEightResult := condEight % 2; condEightResult == 0 {
		if condEight >= 90 {
			fmt.Printf("%v score is perfect\n", condEight)
		} else if condEight >= 85 {
			fmt.Printf("%v score is almost-perfect\n", condEight)
		} else if condEight >= 75 {
			fmt.Printf("%v score is excellent\n", condEight)
		} else if condEight >= 65 {
			fmt.Printf("%v score is not bad\n", condEight)
		} else {
			fmt.Printf("%v score is bad\n", condEight)
		}
	} else {
		fmt.Printf("%v bukan bilangan bulat", condEight)
	}
}
