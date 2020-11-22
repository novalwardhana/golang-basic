package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Part 47 Hash Sha1")
	fmt.Println("")

	fmt.Println("Part 47.1 Penerapan hash sha 1")
	text := "Novalita Kusuma Wardhana"
	var shaA = sha1.New()
	shaA.Write([]byte(text))
	textHashA := shaA.Sum(nil)
	fmt.Printf("textHashA Byte: %v\n", textHashA)
	fmt.Printf("textHashA: %x\n", textHashA)
	text = "One Ok Rock"
	var shaB = sha1.New()
	shaB.Write([]byte(text))
	textHashB := shaB.Sum(nil)
	fmt.Printf("textHashB Byte: %v\n", textHashB)
	fmt.Printf("textHashB: %x\n", textHashB)
	text = "Warrior"
	var shaC = sha1.New()
	shaC.Write([]byte(text))
	textHashC := shaC.Sum(nil)
	fmt.Printf("textHashC Byte: %v\n", textHashC)
	fmt.Printf("textHashC: %x\n", textHashC)
	text = "Lenovo Ideapad Slim 3"
	var shaD = sha1.New()
	shaD.Write([]byte(text))
	textHashD := shaD.Sum(nil)
	fmt.Printf("textHashD Byte: %v\n", textHashD)
	fmt.Printf("textHashD: %x\n", textHashD)
	fmt.Println("")

	fmt.Println("Part 47.2 Methode salting pada hash sha1")
	hashedA, saltA := saltHash("Novalita Kusuma Wardhana")
	fmt.Printf("hashedA: %s\n", hashedA)
	hashedB, saltB := saltHash("One Ok Rock")
	fmt.Printf("hashedB: %s\n", hashedB)
	hashedC, saltC := saltHash("Warrior")
	fmt.Printf("hashedC: %s\n", hashedC)
	hashedD, saltD := saltHash("Lenovo Ideapad Slim 3")
	fmt.Printf("hashedD: %s\n", hashedD)
	_, _, _, _ = saltA, saltB, saltC, saltD
}

func saltHash(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	combineText := fmt.Sprintf("salt:%s, text:%s", salt, text)
	var sha = sha1.New()
	sha.Write([]byte(combineText))
	textHash := sha.Sum(nil)
	textHashStr := fmt.Sprintf("%x", textHash)
	return textHashStr, salt
}
