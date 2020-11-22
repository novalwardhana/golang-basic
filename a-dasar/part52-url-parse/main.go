package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Part 52 URL Parse")
	urlA := "https://www.novalwardhana.com/profile/?name=noval&age=25"
	fmt.Printf("urlA: %s\n", urlA)
	urlAParse, err := url.Parse(urlA)
	if err != nil {
		fmt.Printf("--An error occured: %s\n", err.Error())
		return
	}
	fmt.Printf("urlA Scheme: %s\n", urlAParse.Scheme)
	fmt.Printf("urlA : %s\n", urlAParse.Host)
	fmt.Printf("urlA Path: %s\n", urlAParse.Path)
	fmt.Printf("urlA Query name: %#v\n", urlAParse.Query()["name"][0])
	fmt.Printf("urlA Query age: %#v\n", urlAParse.Query()["age"][0])
	fmt.Println("")

	urlB := "http://localhost:8080/transaction/report/?merchant=starcross&brand=strcrs&merchant=youthgonewild"
	fmt.Printf("urlB: %s\n", urlB)
	urlBParse, err := url.Parse(urlB)
	if err != nil {
		fmt.Printf("--An error occured: %s\n", err.Error())
		return
	}
	fmt.Printf("UrlB Schema: %s\n", urlBParse.Scheme)
	fmt.Printf("UrlB Host: %s\n", urlBParse.Host)
	fmt.Printf("urlB Path: %s\n", urlBParse.Path)
	fmt.Printf("UrlB Query merchant[0]: %#v\n", urlBParse.Query()["merchant"][0])
	fmt.Printf("UrlB Query merchant[1]: %#v\n", urlBParse.Query()["merchant"][1])
	fmt.Printf("UrlB Query brand: %#v\n", urlBParse.Query()["brand"][0])

}
