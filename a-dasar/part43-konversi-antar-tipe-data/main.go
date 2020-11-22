package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Part 43 Konversi Antar Tipe Data")
	fmt.Println("")

	fmt.Println("Part 43.1 Konversi menggunakan strconv")
	str := "125"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("Convert str to numeric success, num: %d\n", num)
	} else {
		fmt.Printf("..Cannot convert str to numeric..\n")
	}
	str = "a12b"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("Convert str to numeric success, num: %d\n", num)
	} else {
		fmt.Printf("..Cannot convert str to numeric..\n")
	}
	num := 89
	str = strconv.Itoa(num)
	fmt.Printf("Convert num to string success, str: %s\n", str)
	num = 92
	str = strconv.Itoa(num)
	fmt.Printf("Convert num to string, str: %s\n", str)
	str = "125"
	if num, err := strconv.ParseInt(str, 10, 8); err == nil {
		fmt.Printf("Convert str to numeric success, num: %d\n", num)
	} else {
		fmt.Printf("..Cannot convert str to numeric..\n")
	}
	str = "10001110101"
	if num, err := strconv.ParseInt(str, 2, 64); err == nil {
		fmt.Printf("Convert str to numeric success, num(2): %d\n", num)
	} else {
		fmt.Printf("..Cannot conver str to numeric..\n")
	}
	num = 24
	str = strconv.FormatInt(int64(num), 8)
	fmt.Printf("24 to okta: %s\n", str)
	str = strconv.FormatInt(int64(num), 2)
	fmt.Printf("24 to biner: %s\n", str)
	str = strconv.FormatInt(int64(num), 6)
	fmt.Printf("24 to hexa: %s\n", str)
	num = 68
	str = strconv.FormatInt(int64(num), 8)
	fmt.Printf("%d to okta: %s\n", num, str)
	str = strconv.FormatInt(int64(num), 2)
	fmt.Printf("%d to biner: %s\n", num, str)
	str = strconv.FormatInt(int64(num), 6)
	fmt.Printf("%d to hexa: %s\n", num, str)
	str = "95.21"
	if num, err := strconv.ParseFloat(str, 32); err == nil {
		fmt.Printf("Convert str to numeric success, num: %f\n", num)
	} else {
		fmt.Printf("..Cannot convert str to numeric: %s..\n", err.Error())
	}
	str = "21.229"
	if num, err := strconv.ParseFloat(str, 64); err == nil {
		fmt.Printf("Convert str to numeric success, num: %f\n", num)
	} else {
		fmt.Printf("..Cannot convert str to numeric: %s\n", err.Error())
	}
	numFloat := float64(95.125)
	str = strconv.FormatFloat(numFloat, 'f', 5, 32)
	fmt.Printf("numFloat: %s\n", str)
	numFloat = float64(92.123456789)
	str = strconv.FormatFloat(numFloat, 'f', 2, 64)
	fmt.Printf("numFloat: %s\n", str)
	str = "true"
	if numBool, err := strconv.ParseBool(str); err == nil {
		fmt.Printf("numBool: %t\n", numBool)
	} else {
		fmt.Printf("..Cannot convert str to boolean: %s\n", err.Error())
	}
	str = "trueee"
	if numBool, err := strconv.ParseBool(str); err == nil {
		fmt.Printf("numBool: %t\n", numBool)
	} else {
		fmt.Printf("..Cannot convert str to boolean: %s\n", err.Error())
	}
	str = "false"
	if numBool, err := strconv.ParseBool(str); err == nil {
		fmt.Printf("numBool: %t\n", numBool)
	} else {
		fmt.Printf("..Cannot convert str to boolean: %s\n", err.Error())
	}
	numBool := true
	str = strconv.FormatBool(numBool)
	fmt.Printf("numBool string: %s\n", str)
	numBool = false
	str = strconv.FormatBool(numBool)
	fmt.Printf("numBool string: %s\n", str)
	numBool = true
	str = strconv.FormatBool(numBool)
	fmt.Printf("numBool string: %s\n", str)
	fmt.Println("")

	fmt.Println("Part 43.2 Konversi data dengan casting")
	intToFloatA := float32(25)
	intToFloatB := float64(95)
	intToFLoatC := float64(45.789)
	fmt.Printf("intToFloatA: %.2f\n", intToFloatA)
	fmt.Printf("intToFloatB: %.5f\n", intToFloatB)
	fmt.Printf("intToFLoat64: %.3f\n", intToFLoatC)
	floatToIntA := int(21.0000)
	floatToIntB := int32(325.0000000000)
	floatToIntC := int64(625.0000)
	fmt.Printf("floatToIntA: %d\n", floatToIntA)
	fmt.Printf("floatToIntB: %d\n", floatToIntB)
	fmt.Printf("floatToIntC: %d\n", floatToIntC)
	fmt.Println("")

	fmt.Println("Part 43.3 Casting string to byte")
	str = "Novalita K W"
	strToByteA := []byte(str)
	fmt.Printf("strToByte: %v\n", strToByteA)
	for i := range strToByteA {
		fmt.Printf("strToByteA[%d]: %v\n", i, strToByteA[i])
	}
	str = "Real Madrid"
	strToByteB := []byte(str)
	fmt.Printf("strToByteB: %v\n", strToByteB)
	for i := range strToByteB {
		fmt.Printf("strToByteB[%d]: %v\n", i, strToByteB[i])
	}
	str = "Jventus"
	strToByteC := []byte(str)
	fmt.Printf("strToByteC: %v\n", strToByteC)
	str = "Barcelona"
	strToByteD := []byte(str)
	fmt.Printf("strToByteD: %v\n", strToByteD)
	strToByteE := []byte("Legion")
	fmt.Printf("strToByteE: %v\n", strToByteE)
	var strToByteF int64 = int64('N')
	fmt.Printf("strToByteF: %d\n", strToByteF)
	strToByteG := int64('D')
	fmt.Printf("strToByteG: %d\n", strToByteG)
	strToByteH := int64('R')
	fmt.Printf("strToByteH: %d\n", strToByteH)
	byteToStrA := string([]byte{123, 124, 125, 126, 127, 128, 129})
	fmt.Printf("byteToStrA: %s\n", byteToStrA)
	byt := []byte{78, 105, 100, 97}
	byteToStrB := string(byt)
	fmt.Printf("byteToStrB: %s\n", byteToStrB)
	byteToStrC := string([]byte{79, 110, 101, 32, 79, 107, 32, 82, 111, 99, 107})
	fmt.Printf("byteToStrC: %s\n", byteToStrC)
	byteToSrcD := string(85)
	fmt.Printf("byteToSrcD: %s\n", byteToSrcD)
	byteToSrcE := string(90)
	fmt.Printf("byteToStrE: %s\n", byteToSrcE)
	byteToSrcF := string(75)
	fmt.Printf("byteToSrcF: %s\n", byteToSrcF)
	fmt.Println("")

	fmt.Println("Part 43.4 Tipe assertions pada interface kosong")
	var dataInterfaceA = map[string]interface{}{
		"satu":  "One Ok Rock",
		"dua":   68,
		"tiga":  22.5,
		"empat": []int{2, 4, 8, 16, 32, 64},
		"lima":  true,
		"enam": [][]int{
			[]int{1, 3, 5, 7, 8},
			[]int{2, 4, 6, 8, 10},
		},
	}
	fmt.Printf("dataInterfaceA: %v\n", dataInterfaceA)
	fmt.Printf("dataInterfaceA satu: %s\n", dataInterfaceA["satu"].(string))
	fmt.Printf("dataInterfaceA dua: %d\n", dataInterfaceA["dua"].(int))
	fmt.Printf("dataInterfaceA tiga: %g\n", dataInterfaceA["tiga"].(float64))
	fmt.Printf("dataInterfaceA empat: %v\n", dataInterfaceA["empat"].([]int))
	fmt.Printf("dataInterfaceA lima: %T\n", dataInterfaceA["lima"].(bool))
	fmt.Printf("dataInterfaceA enam: %v\n", dataInterfaceA["enam"].([][]int))
	var dataInterfaceB = map[string]interface{}{
		"one": [][]string{
			[]string{"Real Madrid", "Barcelona", "Atletico Madrid"},
			[]string{"Arsenal", "Manchester United", "Chelsea", "Everton", "Manchester City", "Liverpool"},
			[]string{"Borussia Dortmund", "Bayern Munich"},
			[]string{"Paris Saint Germain", "AS Monaco"},
		},
		"two": [][]int{
			[]int{2, 4, 6, 8, 10},
			[]int{3, 6, 9, 12, 15},
			[]int{4, 8, 12, 16, 20, 24},
		},
		"three": "Novalita Kusuma wardhana",
		"four":  22.5,
		"five":  75,
		"six":   true,
	}
	fmt.Printf("dataInterfaceB: %v\n", dataInterfaceB)
	for i, val := range dataInterfaceB {
		switch val.(type) {
		case [][]string:
			{
				fmt.Printf("dataInterfaceB[%s] type is [][]string: %v\n", i, val)
			}
		case [][]int:
			{
				fmt.Printf("dataInterfaceB[%s] type is [][]int: %v\n", i, val)
			}
		case float64:
			{
				fmt.Printf("dataInterfaceB[%s] type is float64: %g\n", i, val)
			}
		case bool:
			{
				fmt.Printf("dataInterfaceB[%s] type is boolean: %T\n", i, val)
			}
		case int:
			{
				fmt.Printf("datInterfaceB[%s] type is int: %d\n", i, val)
			}
		default:
			{
				fmt.Printf("dataInterfaceB[%s] type is unlisted: %v\n", i, val)
			}
		}
	}
}
