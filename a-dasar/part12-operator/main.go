package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 12 Operator")
	fmt.Println("")

	fmt.Println("Operator Aritmatika")
	var artmOne = 12
	var artmTwo = 18
	artmThree := 3
	artmFour := 7
	fmt.Printf("Hasil %v + %v = %v\n", artmOne, artmTwo, artmOne+artmTwo)
	fmt.Printf("Hasil %v + %v = %v\n", artmTwo, artmOne, artmTwo-artmOne)
	fmt.Printf("Hasil %v mod %v= %v\n", artmTwo, artmThree, artmTwo%artmThree)
	fmt.Printf("Hasil %v mod %v= %v\n", artmTwo, artmFour, artmTwo%artmFour)
	fmt.Printf("Hasil %v x %v = %v\n", artmOne, artmFour, artmOne*artmFour)
	fmt.Printf("Hasil %v / %v = %v\n", artmThree, artmFour, float64(artmThree)/float64(artmFour))
	fmt.Println("")

	fmt.Println("Operator Perbandingan")
	equalOne := 12
	equalTwo := 29.5
	fmt.Printf("%v > %v : %v\n", equalOne, equalTwo, (float64(equalOne) > equalTwo))
	fmt.Printf("%v < %v : %v\n", equalOne, equalTwo, (float64(equalOne) < equalTwo))
	fmt.Println("")

	fmt.Println("Operator Logika")
	logicOne := true
	logicTwo := false
	logicAnd := logicOne && logicTwo
	logicOr := logicOne || logicTwo
	logicNegation := !logicTwo
	fmt.Printf("%v && %v : %v\n", logicOne, logicTwo, logicAnd)
	fmt.Printf("%v || %v : %v\n", logicOne, logicTwo, logicOr)
	fmt.Printf("!%v : %v\n", logicTwo, logicNegation)
}
