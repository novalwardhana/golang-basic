package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Part 45 Regex")
	fmt.Println("")

	fmt.Println("Part 45.1 Penerapan regexp")
	text := "One Ok Rock"
	var regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	regexFindAllStrA := regex.FindAllString(text, 2)
	regexFindAllStrB := regex.FindAllString(text, -1)
	text = "arsenal barcelona chelsea"
	regexFinsAllStrC := regex.FindAllString(text, -1)
	regexFindAllStrD := regex.FindAllString(text, 0)
	text = "arSEnal barCEloNa CheLseA"
	regexFindAllStrE := regex.FindAllString(text, -1)
	regexFindAllStrF := regex.FindAllString(text, 3)
	fmt.Printf("regexFindAllStrA: %v\n", regexFindAllStrA)
	fmt.Printf("regexFindAllStrB: %v\n", regexFindAllStrB)
	fmt.Printf("regexFindAllStrC: %v\n", regexFinsAllStrC)
	fmt.Printf("regexFindAllStrD: %v\n", regexFindAllStrD)
	fmt.Printf("regexFindAllStrE: %v\n", regexFindAllStrE)
	fmt.Printf("regexFindAllStrF: %v\n", regexFindAllStrF)
	regex, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "One Ok Rock"
	regexFindAllStrG := regex.FindAllString(text, -1)
	fmt.Printf("regexFindAllStrG: %v\n", regexFindAllStrG)
	fmt.Println("")

	fmt.Println("Part 45.2 Match string")
	regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "xoxoxo yoyoyo hohoho"
	regexMatchStrA := regex.MatchString(text)
	text = "XoXoXo YoYoYo HoHoHo"
	regexMatchStrB := regex.MatchString(text)
	text = "RADWIMPS ONE OK ROCK"
	regexMatchStrC := regex.MatchString(text)
	text = "BanTUl KoTA YOgyakaRTA slEMaN"
	regexMatchStrD := regex.MatchString(text)
	fmt.Printf("regexMatchStrA: %t\n", regexMatchStrA)
	fmt.Printf("RegexMatchStrB: %t\n", regexMatchStrB)
	fmt.Printf("regexMatchStrC: %t\n", regexMatchStrC)
	fmt.Printf("regexMatchStrD: %t\n", regexMatchStrD)
	regex, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "xoxoxo yoyoyo hohoho"
	regexMatchStrE := regex.MatchString(text)
	text = "WaRRior venTELa ComPasS"
	regexMatchStrF := regex.MatchString(text)
	fmt.Printf("regexMatchStrE: %t\n", regexMatchStrE)
	fmt.Printf("regexMatchStrF: %t\n", regexMatchStrF)
	fmt.Println("")

	fmt.Println("Part 45.3 Find string")
	regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "sate soto bakso"
	regexFindStrA := regex.FindString(text)
	text = "Sate Soto Bakso"
	regexFindStrB := regex.FindString(text)
	text = "WARRIOR VENTELA PATROBAS"
	regexFindStrC := regex.FindString(text)
	text = "xiAoMi SamSUng OpPO"
	regexFindStrD := regex.FindString(text)
	fmt.Printf("regexFindStrA: %s\n", regexFindStrA)
	fmt.Printf("regexFindStrB: %s\n", regexFindStrB)
	fmt.Printf("regexFindStrC: %s\n", regexFindStrC)
	fmt.Printf("regexFindStrD: %s\n", regexFindStrD)
	regex, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "sate soto bakso"
	regexFindStrE := regex.FindString(text)
	text = "Sate Soto Bakso"
	regexFindStrF := regex.FindString(text)
	text = "WARRIOR VENTELA PATROBAS"
	regexFindStrG := regex.FindString(text)
	text = "xiAOMi SamSUng OpPO"
	regexFindStrH := regex.FindString(text)
	fmt.Printf("regexFindSteE: %s\n", regexFindStrE)
	fmt.Printf("regexFindStrF: %s\n", regexFindStrF)
	fmt.Printf("regexFindStrG: %s\n", regexFindStrG)
	fmt.Printf("regexFindStrH: %s\n", regexFindStrH)
	fmt.Println("")

	fmt.Println("Part 45.4 Find string index")
	regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "xoxoxo yoyoyo hohoho"
	regexFindIndexStrA := regex.FindStringIndex(text)
	text = "XoXoXo YoYoYo HoHoHo"
	regexFindIndexStrB := regex.FindStringIndex(text)
	text = "WARRIOR VENTELA PATROBAS"
	regexFindIndexStrC := regex.FindStringIndex(text)
	text = "WaRRior venTELa ComPasS"
	regexFindIndexStrD := regex.FindStringIndex(text)
	fmt.Printf("regexFindIndexStrA: %v\n", regexFindIndexStrA)
	fmt.Printf("regexFindIndexStrB: %v\n", regexFindIndexStrB)
	fmt.Printf("regexFindIndexStrC: %v\n", regexFindIndexStrC)
	fmt.Printf("regexFindIndexStrD: %v\n", regexFindIndexStrD)
	regex, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "xoxoxo yoyoyo hohoho"
	regexFindIndexStrE := regex.FindStringIndex(text)
	text = "XoXoXo YoYoYo HoHoHo"
	regexFindIndexStrF := regex.FindStringIndex(text)
	text = "WARRIOR VENTELA PATROBAS"
	regexFindIndexStrG := regex.FindStringIndex(text)
	text = "WaRRior venTELa ComPasS"
	regexFindIndexStrH := regex.FindStringIndex(text)
	fmt.Printf("regexFindIndexStrE: %v\n", regexFindIndexStrE)
	fmt.Printf("regexFindIndexStrF: %v\n", regexFindIndexStrF)
	fmt.Printf("regexFindIndexStrF: %v\n", regexFindIndexStrG)
	fmt.Printf("regexFindIndexStrH: %v\n", regexFindIndexStrH)
	fmt.Println("")

	fmt.Println("Part 45.5 Find all string")
	regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "sate soto bakso"
	regexFindAllStrH := regex.FindAllString(text, -1)
	text = "Sate Soto Bakso"
	regexFindAllStrI := regex.FindAllString(text, 2)
	text = "sATe soTO bAKsO"
	regexFindAllStrJ := regex.FindAllString(text, 3)
	text = "SATE SOTO BAKSO"
	regexFindAllStrK := regex.FindAllString(text, -1)
	fmt.Printf("regexFindAllStrH: %v\n", regexFindAllStrH)
	fmt.Printf("regexFindAllStrI: %v\n", regexFindAllStrI)
	fmt.Printf("regexFindAllStrJ: %v\n", regexFindAllStrJ)
	fmt.Printf("regexFindAllStrK: %v\n", regexFindAllStrK)
	regex, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "sate soto bakso"
	regexFindAllStrL := regex.FindAllString(text, -1)
	text = "Sate Soto Bakso"
	regexFindAllStrM := regex.FindAllString(text, 2)
	text = "sATe soTO bAKsO"
	regexFindAllStrN := regex.FindAllString(text, 3)
	text = "SATE SOTO BAKSO"
	regexFindAllSTrO := regex.FindAllString(text, -1)
	fmt.Printf("regexFindAllStrL: %v\n", regexFindAllStrL)
	fmt.Printf("regexFindAllStrM: %v\n", regexFindAllStrM)
	fmt.Printf("regexFindAllStrN: %v\n", regexFindAllStrN)
	fmt.Printf("regexFindAllStrO: %v\n", regexFindAllSTrO)
	fmt.Println("")

	fmt.Println("Part 45.6 Replace all string")
	regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "lenovo acer asus dell"
	regexReplaceAllStrA := regex.ReplaceAllString(text, "laptop")
	text = "Arsenal Barcelona Chelsea"
	regexReplaceAllStrB := regex.ReplaceAllString(text, "Tottenham FC")
	text = "InTEr miLAn juVEntUs"
	regexReplaceAllStrC := regex.ReplaceAllString(text, "Lazio FC")
	text = "WARRIOR PATROBAS VENTELA"
	regexReplaceAllStrD := regex.ReplaceAllString(text, "Nike")
	fmt.Printf("regexReplaceAllStrA: %s\n", regexReplaceAllStrA)
	fmt.Printf("regexReplaceAllStrB: %s\n", regexReplaceAllStrB)
	fmt.Printf("regexReplaceAllStrC: %s\n", regexReplaceAllStrC)
	fmt.Printf("regexReplaceAllStrD: %s\n", regexReplaceAllStrD)
	regex, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "lenovo acer asus dell"
	regexReplaceAllStrE := regex.ReplaceAllString(text, "LAPTOP")
	text = "Arsenal Barcelona Chelsea"
	regexReplaceAllStrF := regex.ReplaceAllString(text, "Tottenhan FC-")
	text = "InTEr miLAn juVEntUs"
	regexReplaceAllStrG := regex.ReplaceAllString(text, "Lazio FC")
	text = "WARRIOR PATROBAS VENTELA"
	regexReplaceAllStrH := regex.ReplaceAllString(text, "Nike")
	fmt.Printf("regexReplaceAllStrE: %s\n", regexReplaceAllStrE)
	fmt.Printf("regexReplaceAllStrF: %s\n", regexReplaceAllStrF)
	fmt.Printf("regexReplaceAllStrG: %s\n", regexReplaceAllStrG)
	fmt.Printf("regexReplaceAllSTrH: %s\n", regexReplaceAllStrH)
	fmt.Println("")

	fmt.Println("Part 45.7 Relace all string func")
	regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "burger pizza kebab"
	regexReplaceAllStrFuncA := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "pizza" {
			return "ayam geprek"
		}
		return each
	})
	regexReplaceAllStrFuncB := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "Pizza" {
			return "ayam geprek"
		}
		return each
	})
	text = "BarCelona arSenal AC MiLan Inter MiLan"
	regexReplaceAllStrFuncC := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "ar" || each == "an" {
			return "Xoxo"
		}
		return each
	})
	text = "RYzen zenBook TIzen"
	regexReplaceAllStrFuncD := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "zen" {
			return "Book"
		}
		return each
	})
	fmt.Printf("regexReplaceAllStrFuncA: %s\n", regexReplaceAllStrFuncA)
	fmt.Printf("regexReplaceAllStrFuncB: %s\n", regexReplaceAllStrFuncB)
	fmt.Printf("regexReplaceAllStrFuncC: %s\n", regexReplaceAllStrFuncC)
	fmt.Printf("regexReplaceAllStrFuncD: %s\n", regexReplaceAllStrFuncD)
	regex, err = regexp.Compile(`[A-Z]+`)
	if err != nil {
		fmt.Printf("..error when compile regexp: %s\n", err.Error())
	}
	text = "BURGER PIZZA KEBAB"
	regexReplaceAllStrFuncE := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "PIZZA" {
			return "Roti Bakar"
		}
		return each
	})
	regexReplaceAllStrFuncF := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "Pizza" {
			return "Roti Panggang"
		}
		return each
	})
	text = "bARcelona ARsenal MANCHESTER United MANCHESTER City"
	regexReplaceAllStrFuncG := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "AR" || each == "MANCHESTER" {
			return "Yoyo"
		}
		return each
	})
	text = "thinkPAD ideaPAD notePAD"
	regexReplaceAllStrFuncH := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "PAD" {
			return "book"
		}
		return each
	})
	fmt.Printf("regexReplaceAllStrFuncE: %s\n", regexReplaceAllStrFuncE)
	fmt.Printf("regexReplaceAllStrFuncF: %s\n", regexReplaceAllStrFuncF)
	fmt.Printf("regexReplaceAllStrFuncG: %s\n", regexReplaceAllStrFuncG)
	fmt.Printf("regexReplaceAllStrFuncH: %s\n", regexReplaceAllStrFuncH)
	fmt.Println("")

	fmt.Println("Part 45.8 String split")
	regex, _ = regexp.Compile(`[a-c]+`)
	text = "ac milan inter milan chelsea"
	regexSplitA := regex.Split(text, -1)
	regex, _ = regexp.Compile(`[s-t]+`)
	text = "soto sate sarden tempe tahu"
	regexSplitB := regex.Split(text, 4)
	regex, _ = regexp.Compile(`[C-E]+`)
	text = "BarCElona Real MaDrid BayErn Munich"
	regexSplitC := regex.Split(text, -1)
	regex, _ = regexp.Compile(`[A-Z]+`)
	text = "NovaLIta12 KuSUma16 WarDHana20"
	regexSplitD := regex.Split(text, -7)
	fmt.Printf("regexSplitA: %#v\n", regexSplitA)
	fmt.Printf("regexSplitB: %#v\n", regexSplitB)
	fmt.Printf("regexSplitC: %#v\n", regexSplitC)
	fmt.Printf("regexSplitD: %#v\n", regexSplitD)
}
