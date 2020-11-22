package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func func02A(s string) (bool, error) {
	if strings.TrimSpace(s) == "" {
		return false, errors.New("String is not filled")
	}
	return true, nil
}

func func02B(s string) (bool, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return true, nil
	}
	return false, errors.New("Cannot convert to int")
}

func func03A(s string) (bool, error) {
	if strings.TrimSpace(s) == "" {
		return false, errors.New("String is not filled")
	}
	return true, nil
}

func func03B(s string) (bool, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return true, nil
	}
	return false, errors.New("Cannot parse to int")
}

func func04Recover() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Can run perfectly")
	}
}

func func04(s string) (bool, error) {
	if strings.TrimSpace(s) == "" {
		return false, errors.New("String is null")
	}
	return true, nil
}

func main() {
	fmt.Println("Part 37 Error Panic Recover")
	fmt.Println("")
	fmt.Println("Part 37.1 Pemanfaatan error")
	var var01A string
	fmt.Println("Type your number...")
	fmt.Scanln(&var01A)
	var var01Err error
	var var01B int
	var01B, var01Err = strconv.Atoi(var01A)
	if var01Err != nil {
		fmt.Println(var01Err)
	} else {
		fmt.Printf("Can convert to number: %d\n", var01B)
	}
	if var01C, var01CErr := strconv.Atoi(var01A); var01CErr != nil {
		fmt.Printf("Can't convert to int: %v\n", var01CErr)
	} else {
		fmt.Printf("Can convert to int: %v\n", var01C)
	}
	fmt.Println("")

	fmt.Println("Part 37.2 Custom error")
	var var02A string
	fmt.Scanln(&var02A)
	if _, err := func02A(var02A); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("String is filled")
	}
	var var02B string
	fmt.Scanln(&var02B)
	if valid, err := func02B(var02B); valid {
		fmt.Printf("%v can convert to int\n", var02B)
	} else {
		fmt.Printf("%v %v\n", var02B, err)
	}
	fmt.Println("")

	fmt.Println("Part 37.3 Penggunaan panic")
	var var03A string
	fmt.Scanln(&var03A)
	if _, err := func03A(var03A); err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("%s is not empty string\n", var03A)
	}
	var var03B string
	fmt.Scanln(&var03B)
	if valid, err := func03B(var03B); valid {
		fmt.Printf("%s can convert to int\n", var03B)
	} else {
		panic(err.Error())
	}
	fmt.Println("")

	fmt.Println("Part 37.4 Pengggunaan recover")
	defer func04Recover()
	var var04 string
	fmt.Scanln(&var04)
	if valid, err := func04(var04); valid {
		fmt.Println("String is filled")
	} else {
		panic(err.Error())
	}
	fmt.Println("")

	fmt.Println("Part 37.5 Penggunaan recover IIFE")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Panic error because", r)
	// 	} else {
	// 		fmt.Println("System can work properly")
	// 	}
	// }()
	// panic("Some error happen")
	var05 := []string{"Real Madrid", "Manchester United", "Paris Saint Germain", "Bayern Munich"}
	fmt.Printf("var05: %v\n", var05)
	for _, club := range var05 {
		func() {
			defer func(club string) {
				if r := recover(); r != nil {
					fmt.Printf("Panic error in: %v %v\n", club, r)
				}
			}(club)
			panic(fmt.Sprintf("Error in %v\n", club))
		}()
	}
	var05B := []string{"Legion", "Pegasus", "Alienware", "ROG", "Predator"}
	fmt.Printf("var05B: %v\n", var05B)
	for _, gadget := range var05B {
		func() {
			defer func(gadget string) {
				if r := recover(); r != nil {
					fmt.Printf("Panic error in: %v %v\n", gadget, r)
				}
			}(gadget)
			panic(fmt.Sprintf("Error occured in: %v\n", gadget))
		}()
	}
}
