package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Part 46 Encode Decode Base 64")
	fmt.Println("")

	fmt.Println("Part 46.1 Penerapan encodeToString dan decodeToString")
	str := "Novalita Kusuma Wardhana"
	encodeToStrA := base64.StdEncoding.EncodeToString([]byte(str))
	decodeToStrA, _ := base64.StdEncoding.DecodeString(encodeToStrA)
	fmt.Printf("encodeToStrA: %s\n", encodeToStrA)
	fmt.Printf("decodeToStrA Byte: %v\n", decodeToStrA)
	fmt.Printf("decodeToStrA: %v\n", string(decodeToStrA))
	str = "Manchester United"
	encodeToStrB := base64.StdEncoding.EncodeToString([]byte(str))
	decodeToStrB, _ := base64.StdEncoding.DecodeString(encodeToStrB)
	fmt.Printf("encodeToStrB: %s\n", encodeToStrB)
	fmt.Printf("decodeToStrB Byte: %v\n", decodeToStrB)
	fmt.Printf("decodeToStrB: %s\n", string(decodeToStrB))
	str = "One Ok Rock"
	encodeToStrC := base64.StdEncoding.EncodeToString([]byte(str))
	decodeToStrC, _ := base64.StdEncoding.DecodeString(encodeToStrC)
	fmt.Printf("encodeToStrC: %s\n", encodeToStrC)
	fmt.Printf("decodeToStrC Byte: %v\n", decodeToStrC)
	fmt.Printf("decodeToStrC: %s\n", string(decodeToStrC))
	fmt.Println("")

	fmt.Println("Part 46.2 Penerapan fungsi encode & decode")
	str = "Real Madrid"
	var encodedA = make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	base64.StdEncoding.Encode(encodedA, []byte(str))
	fmt.Printf("encodedA Byte: %v\n", encodedA)
	fmt.Printf("encodedA: %v\n", string(encodedA))
	var decodedA = make([]byte, base64.StdEncoding.DecodedLen(len(encodedA)))
	_, errDecodedA := base64.StdEncoding.Decode(decodedA, encodedA)
	if errDecodedA != nil {
		fmt.Printf(".. Error decode..")
	}
	fmt.Printf("decodedA Byte: %v\n", decodedA)
	fmt.Printf("decodedA: %s\n", string(decodedA))
	str = "Piyungan"
	encodedB := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	base64.StdEncoding.Encode(encodedB, []byte(str))
	fmt.Printf("encodedB Byte: %v\n", encodedB)
	fmt.Printf("encodedB: %s\n", string(encodedB))
	decodedB := make([]byte, base64.StdEncoding.DecodedLen(len(encodedB)))
	_, errDecodedB := base64.StdEncoding.Decode(decodedB, []byte(encodedB))
	if errDecodedB != nil {
		fmt.Printf("..Cannot decode..\n")
	}
	fmt.Printf("decodeB Byte: %v\n", decodedB)
	fmt.Printf("decodeB: %s\n", string(decodedB))
	str = "Heha Sky View"
	encodedC := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	base64.StdEncoding.Encode(encodedC, []byte(str))
	fmt.Printf("encodedC Byte: %v\n", encodedC)
	fmt.Printf("encodedC: %s\n", string(encodedC))
	decodedC := make([]byte, base64.StdEncoding.DecodedLen(len(encodedC)))
	_, errDecodedC := base64.StdEncoding.Decode(decodedC, []byte(encodedC))
	if errDecodedC != nil {
		fmt.Printf("..Cannot decode..\n")
	}
	fmt.Printf("decodedC Byte: %v\n", decodedC)
	fmt.Printf("decodedC: %s\n", string(decodedC))
	str = "One Ok Rock"
	encodedD := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	base64.StdEncoding.Encode(encodedD, []byte(str))
	fmt.Printf("encodedD Byte: %v\n", encodedD)
	fmt.Printf("encodedD: %s\n", string(encodedD))
	decodedD := make([]byte, base64.StdEncoding.DecodedLen(len(encodedD)))
	_, errDecodedD := base64.StdEncoding.Decode(decodedD, []byte(encodedD))
	if errDecodedD != nil {
		fmt.Printf("..Cannot decode..\n")
	}
	fmt.Printf("decodedD Byte: %v\n", decodedD)
	fmt.Printf("decodedD: %s\n", string(decodedD))
	fmt.Println("")

	fmt.Println("Part 46.3 Encode & decode data url")
	url := "https://google.com"
	urlEncodedA := base64.URLEncoding.EncodeToString([]byte(url))
	urlDecodedA, _ := base64.URLEncoding.DecodeString(urlEncodedA)
	fmt.Printf("urlEncodedA: %v\n", urlEncodedA)
	fmt.Printf("urlDecodedA Byte: %v\n", urlDecodedA)
	fmt.Printf("urlDecodedA: %s\n", string(urlDecodedA))
	url = "https://www.yahoo.com"
	urlEncodedB := base64.URLEncoding.EncodeToString([]byte(url))
	urlDecodedB, _ := base64.URLEncoding.DecodeString(urlEncodedB)
	fmt.Printf("urlEncodedB: %v\n", urlEncodedB)
	fmt.Printf("urlDecodedB Byte: %v\n", urlDecodedB)
	fmt.Printf("urlDecodedB: %v\n", string(urlDecodedB))
	str = "http://github.co.id"
	urlEncodedC := base64.URLEncoding.EncodeToString([]byte(str))
	urlDecodedC, _ := base64.URLEncoding.DecodeString(urlEncodedC)
	fmt.Printf("urlEncodedC: %v\n", urlEncodedC)
	fmt.Printf("urlDecodedC Byte: %v\n", urlDecodedC)
	fmt.Printf("urlDecodedC: %s\n", string(urlDecodedC))
	str = "http://kaskus.co.id"
	urlEncodedD := base64.URLEncoding.EncodeToString([]byte(str))
	urlDecodedD, _ := base64.URLEncoding.DecodeString(urlEncodedD)
	fmt.Printf("urlEncodedD: %v\n", urlEncodedD)
	fmt.Printf("urlDecodedD Byte: %v\n", urlDecodedD)
	fmt.Printf("urlDecodedD: %s\n", string(urlDecodedD))
}
