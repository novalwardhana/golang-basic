package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 44 Fungsi String")
	fmt.Println("")

	fmt.Println("Part 44.1 String contains")
	strContainA := strings.Contains("One Ok Rock", "Ok")
	strContainB := strings.Contains("One Ok Rock", "OK")
	strContainC := strings.Contains("Yogyakarta", "Jogja")
	strContainD := strings.Contains("Novalita Kusuma Wardhana", "ma Wardh")
	fmt.Printf("strContainA: %t\n", strContainA)
	fmt.Printf("strContainB: %t\n", strContainB)
	fmt.Printf("strContainC: %t\n", strContainC)
	fmt.Printf("strContainD: %t\n", strContainD)
	fmt.Println("")

	fmt.Println("Part 44.2 String Has Prefix")
	strHasPrefixA := strings.HasPrefix("Novalita", "Nov")
	strHasPrefixB := strings.HasPrefix("Novalita", "NOV")
	strHasPrefixC := strings.HasPrefix("Barcelona", "celo")
	strHasPrefixD := strings.HasPrefix("Real Madrid", "Real Ma")
	fmt.Printf("strHasPrefixA: %t\n", strHasPrefixA)
	fmt.Printf("strHasPrefixB: %t\n", strHasPrefixB)
	fmt.Printf("strHasPrefixC: %t\n", strHasPrefixC)
	fmt.Printf("strHasPrefixD: %t\n", strHasPrefixD)
	fmt.Println("")

	fmt.Println("Part 44.3 String Has Suffix")
	strHasSuffixA := strings.HasSuffix("Lenovo", "ovo")
	strHasSuffixB := strings.HasSuffix("AlienWare", "lien")
	strHasSuffixC := strings.HasSuffix("Honda", "a")
	strHasSuffixD := strings.HasSuffix("Yokohama", "Yoko")
	fmt.Printf("strHasSuffixA: %t\n", strHasSuffixA)
	fmt.Printf("strHasSuffixB: %t\n", strHasSuffixB)
	fmt.Printf("strHasSuffixC: %t\n", strHasSuffixC)
	fmt.Printf("strHasSuffixD: %t\n", strHasSuffixD)
	fmt.Println("")

	fmt.Println("Part 44.4 String Count")
	strCountA := strings.Count("One Ok Rock", "O")
	strCountB := strings.Count("Barcelona", "R")
	strCountC := strings.Count("Lenovo Legion", "Le")
	strCountD := strings.Count("Yogyakarta, Surakarta, Surabaya, Purwakarta, Bandung, Jakarta", "kar")
	fmt.Printf("strCountA: %d\n", strCountA)
	fmt.Printf("strCountB: %d\n", strCountB)
	fmt.Printf("strCountC: %d\n", strCountC)
	fmt.Printf("strCoundD: %d\n", strCountD)
	fmt.Println("")

	fmt.Println("Part 44.5 String index")
	strIndexA := strings.Index("Yogyakarta, Surakarta, Surabaya, Purawkarta", "Su")
	strIndexB := strings.Index("Chelsea", "A")
	strIndexC := strings.Index("Chelsea", "a")
	strIndexD := strings.Index("Yamanaka Sai", "S")
	fmt.Printf("strIndexA: %d\n", strIndexA)
	fmt.Printf("StrIndexB: %d\n", strIndexB)
	fmt.Printf("strIndexC: %d\n", strIndexC)
	fmt.Printf("strIndexD: %d\n", strIndexD)
	fmt.Println("")

	fmt.Println("Part 44.6 String Replace")
	text := "Redmi Redmi Redmi"
	find := "d"
	replaceWith := "s"
	strReplaceA := strings.Replace(text, find, replaceWith, 1)
	strReplaceB := strings.Replace(text, find, replaceWith, 2)
	strReplaceC := strings.Replace(text, find, replaceWith, -1)
	fmt.Printf("strReplaceA: %s\n", strReplaceA)
	fmt.Printf("strReplaceB: %s\n", strReplaceB)
	fmt.Printf("strReplaceC: %s\n", strReplaceC)
	text = "Inter-Milan * Inter-Milan * Inter-Milan * Inter-Milan * Inter-Milan"
	find = "Inter"
	replaceWith = "AC"
	strReplaceD := strings.Replace(text, find, replaceWith, 2)
	strReplaceE := strings.Replace(text, find, replaceWith, 1)
	strReplaceF := strings.Replace(text, find, replaceWith, -1)
	fmt.Printf("strReplaceD: %s\n", strReplaceD)
	fmt.Printf("strReplaceE: %s\n", strReplaceE)
	fmt.Printf("strReplaceF: %s\n", strReplaceF)
	fmt.Println("")

	fmt.Println("Part 44.7 String Repeat")
	strRepeatA := strings.Repeat("Xo", 3)
	strRepaatB := strings.Repeat("Real Madrid ", 5)
	strRepeatC := strings.Repeat("y", 10)
	strRepeatD := strings.Repeat("Halo ", 7)
	fmt.Printf("strRepeatA: %s\n", strRepeatA)
	fmt.Printf("strRepeatB: %s\n", strRepaatB)
	fmt.Printf("strRepeatC: %s\n", strRepeatC)
	fmt.Printf("strRepeatD: %s\n", strRepeatD)
	fmt.Println("")

	fmt.Println("Part 44.8 String Split")
	strSplitA := strings.Split("abcdefg", "")
	strSplitB := strings.Split("Novalita Kusuma Wardhana", "")
	strSplitC := strings.Split("Novalita Kusuma Wardhana", " ")
	strSplitD := strings.Split("R-e-a-l- -M-a-d-r-i-d", "-")
	strSplitE := strings.Split("125098", "")
	fmt.Printf("strSplitA: %v, len: %d\n", strSplitA, len(strSplitA))
	for i, v := range strSplitA {
		fmt.Printf("strSplitA[%d]: %s\n", i, v)
	}
	fmt.Printf("strSplitB: %v, len: %d\n", strSplitB, len(strSplitB))
	for i, v := range strSplitB {
		fmt.Printf("strSplitB[%d]: %s\n", i, v)
	}
	fmt.Printf("strSplitC: %v, len: %d\n", strSplitC, len(strSplitC))
	for i, v := range strSplitC {
		fmt.Printf("strSplitC[%d]: %s\n", i, v)
	}
	fmt.Printf("strSplitD: %v, len: %d\n", strSplitD, len(strSplitD))
	for i, v := range strSplitD {
		fmt.Printf("strSplitD[%d]: %s\n", i, v)
	}
	fmt.Printf("strSplitE: %v, len: %d\n", strSplitE, len(strSplitE))
	for i, v := range strSplitE {
		if num, err := strconv.Atoi(v); err == nil {
			fmt.Printf("strSplitE[%d]: %d\n", i, num)
		} else {
			fmt.Printf("strSplitE[%d] cannot convert to numeric\n", i)
		}
	}
	fmt.Println("")

	fmt.Println("Part 44.9 String Join")
	textJoin := []string{"Novalita", "Kusuma", "Wardhana"}
	strJoinA := strings.Join(textJoin, " ")
	fmt.Printf("strJoinA: %s\n", strJoinA)
	textJoin = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	strJoinB := strings.Join(textJoin, "")
	fmt.Printf("strJoinB: %s\n", strJoinB)
	textJoin = []string{"Real Madrid", "Manchester United", "Juventus", "Borussia Dortmund"}
	strJoinC := strings.Join(textJoin, " - ")
	fmt.Printf("strJoinC: %s\n", strJoinC)
	textJoin = []string{"Kota Yogyakarta", "Bantul", "Gunungkidul", "Kulon Progo", "Sleman"}
	strJoinD := strings.Join(textJoin, " // ")
	fmt.Printf("strJoinD: %s\n", strJoinD)
	fmt.Println("")

	fmt.Println("Part 44.10 String toLower")
	strToLowerA := strings.ToLower("VenTELa")
	strTolowerB := strings.ToLower("ARSENAL")
	strToLowerC := strings.ToLower("NovALWardHAna93")
	strToLowerD := strings.ToLower("90MinutES")
	fmt.Printf("strToLowerA: %s\n", strToLowerA)
	fmt.Printf("strToLowerB: %s\n", strTolowerB)
	fmt.Printf("strToLowerC: %s\n", strToLowerC)
	fmt.Printf("strToLowerD: %s\n", strToLowerD)
	fmt.Println("")

	fmt.Println("Part 44.11 String toUpper")
	strToUpperA := strings.ToUpper("aDIdaS")
	strToUpperB := strings.ToUpper("everton")
	strToUpperC := strings.ToUpper("noVAl20WARDhanA93")
	strToUpperD := strings.ToUpper("16perSOnALitIEs")
	fmt.Printf("strToUpperA: %s\n", strToUpperA)
	fmt.Printf("strToUpperB: %s\n", strToUpperB)
	fmt.Printf("strToUpperC: %s\n", strToUpperC)
	fmt.Printf("strToUpperD: %s\n", strToUpperD)
	fmt.Println("")
}
