package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 23 Pointer")
	fmt.Println("")

	fmt.Println("23.1 Penerapan pointer")
	var var01A int = 4
	var var01B *int
	fmt.Printf("var01A: %v\n", var01A)
	fmt.Printf("var01A Address: %v\n", &var01A)
	fmt.Printf("var01B: %v\n", var01B)
	fmt.Printf("var01B Address: %v\n", &var01B)
	var01B = &var01A
	fmt.Printf("var01B: %v\n", *var01B)
	fmt.Printf("var01B Address: %v\n", var01B)
	fmt.Printf("var01B &Address: %v\n", &var01B)
	fmt.Println("")

	fmt.Println("23.2 Efek perubahan nilai pointer")
	var var02A int = 9
	var var02B *int = &var02A
	fmt.Printf("var02A: %v\n", var02A)
	fmt.Printf("var02A Address: %v\n", &var02A)
	fmt.Printf("var02B: %v\n", *var02B)
	fmt.Printf("var02B Address: %v\n", var02B)
	var02A = 12
	fmt.Printf("var02A: %v\n", var02A)
	fmt.Printf("var02A Address: %v\n", &var02A)
	fmt.Printf("var02B: %v\n", *var02B)
	fmt.Printf("var02B Address: %v\n", var02B)
	*var02B = 25
	fmt.Printf("var02A: %v\n", var02A)
	fmt.Printf("var02A Address: %v\n", &var02A)
	fmt.Printf("var02B: %v\n", *var02B)
	fmt.Printf("var02B Address: %v\n", var02B)
	fmt.Println("")

	fmt.Println("23.3 Parameter pointer")
	var var03A int = 25
	fmt.Printf("var03A: %v\n", var03A)
	fungsi03(&var03A, 22)
	fmt.Printf("var03A: %v\n", var03A)
	var var03B *int
	fmt.Printf("var03B: %v\n", var03B)
	var03B = &var03A
	fmt.Printf("var03B: %v\n", *var03B)
	fungsi03(var03B, 32)
	fmt.Printf("var03A: %v\n", var03A)
	fmt.Printf("var03A Address: %v\n", &var03A)
	fmt.Printf("var03B: %v\n", *var03B)
	fmt.Printf("var03B Address: %v\n", var03B)
	fmt.Printf("var03B &Address: %v\n", &var03B)
}

func fungsi03(varSend *int, varChange int) {
	*varSend = varChange
}
