package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name      string `json:"Name"`
	AgeNumber int    `json:"Age"`
}

type balok struct {
	P int `json:"panjang"`
	L int `json:"lebar"`
	T int `json:"tinggi"`
}

func main() {
	fmt.Println("Part 53 JSON Data")
	fmt.Println("")

	fmt.Println("Part 53.1 Decode json ke struct")
	jsonStrA := `{"Name": "Novalita Kusuma wardhana", "Age": 25 }`
	jsonByteA := []byte(jsonStrA)
	var jsonA person
	var err = json.Unmarshal(jsonByteA, &jsonA)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonStrA: %s\n", jsonStrA)
	fmt.Printf("JsonA: %v\n", jsonA)
	jsonStrB := `{"panjang": 10, "lebar": 7, "tinggi": 8}`
	jsonByteB := []byte(jsonStrB)
	var jsonB balok
	err = json.Unmarshal(jsonByteB, &jsonB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonStrB: %s\n", jsonStrB)
	fmt.Printf("jsonB: %v\n", jsonB)
	jsonStrC := `{"Name": "Nida", "Age": 27}`
	jsonByteC := []byte(jsonStrC)
	var jsonC *person
	err = json.Unmarshal(jsonByteC, &jsonC)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonStrC: %s\n", jsonStrC)
	fmt.Printf("jsonC: %v\n", *jsonC)
	jsonStrD := `{"panjang": 12, "lebar": 5, "tinggi": 22}`
	jsonByteD := []byte(jsonStrD)
	var jsonD balok
	err = json.Unmarshal(jsonByteD, &jsonD)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("jsonStrD: %s\n", jsonStrD)
	fmt.Printf("jsonD: %v\n", jsonD)
	fmt.Println("")

	fmt.Println("Part 53.2 Decode json ke interface")
	jsonStrInterfaceA := `{"panjang": 12, "lebar": 8, "tinggi": 5}`
	jsonByteInterfaceA := []byte(jsonStrInterfaceA)
	var jsonInterfaceA map[string]interface{}
	err = json.Unmarshal(jsonByteInterfaceA, &jsonInterfaceA)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonInterfaceA: %v\n", jsonInterfaceA)
	fmt.Printf("jsonInterfaceA panjang: %0.f\n", jsonInterfaceA["panjang"])
	fmt.Printf("jsonInterfaceA lebar: %0.f\n", jsonInterfaceA["lebar"])
	fmt.Printf("jsonInterfaceA tinggi: %0.f\n-----\n", jsonInterfaceA["tinggi"])
	jsonStrInterfaceB := `{"name": "noval", "age": 26}`
	jsonByteInterfaceB := []byte(jsonStrInterfaceB)
	var jsonInterfaceB map[string]interface{}
	err = json.Unmarshal(jsonByteInterfaceB, &jsonInterfaceB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonInterfaceB: %v\n", jsonInterfaceB)
	fmt.Printf("jsonInterfaceB name: %s\n", jsonInterfaceB["name"])
	fmt.Printf("jsonInterfaceB age: %0.f\n-----\n", jsonInterfaceB["age"])
	jsonStrInterfaceC := `{"club": "Arsenal", "stadium": "Emirated"}`
	jsonByteInterfaceC := []byte(jsonStrInterfaceC)
	var jsonInterfaceC map[string]interface{}
	err = json.Unmarshal(jsonByteInterfaceC, &jsonInterfaceC)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonInterfaceC: %v\n", jsonInterfaceC)
	fmt.Printf("jsonInterfaceC club: %v\n", jsonInterfaceC["club"])
	fmt.Printf("jsonInterfaceC stadium: %v\n-----\n", jsonInterfaceC["stadium"])
	jsonStrInterfaceD := `{"brand": "lenovo", "seri": "Ideapad SLim 7", "price": 900}`
	jsonByteInterfaceD := []byte(jsonStrInterfaceD)
	var jsonInterfaceD interface{}
	err = json.Unmarshal(jsonByteInterfaceD, &jsonInterfaceD)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonInterfaceMapD := jsonInterfaceD.(map[string]interface{})
	fmt.Printf("jsonInterfaceMapD: %v\n", jsonInterfaceMapD)
	fmt.Printf("jsonInterfaceMapD brand: %v\n", jsonInterfaceMapD["brand"])
	fmt.Printf("jsonInterfaceMapD seri: %v\n", jsonInterfaceMapD["seri"])
	fmt.Printf("jsonInterfaceMapD price: %0.f\n-----\n", jsonInterfaceMapD["price"])
	jsonStrInterfaceE := `{"player": "Toni Kross", "club": "Real Madrid", "country": "Germany", "Price": 125.12}`
	jsonByteInterfaceE := []byte(jsonStrInterfaceE)
	var jsonInterfaceE interface{}
	err = json.Unmarshal(jsonByteInterfaceE, &jsonInterfaceE)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonInterfaceMapE := jsonInterfaceE.(map[string]interface{})
	fmt.Printf("jsonInterfaceE: %v\n", jsonInterfaceMapE)
	fmt.Printf("jsonInterfaceE player: %s\n", jsonInterfaceMapE["player"])
	fmt.Printf("jsonInterfaceE club: %s\n", jsonInterfaceMapE["club"])
	fmt.Printf("jsonInterfaceE country: %s\n", jsonInterfaceMapE["country"])
	fmt.Printf("jsonInterfaceE price: %g\n-----\n", jsonInterfaceMapE["Price"])
	jsonStrInterfaceF := `{"username": "root@localhost.dev", "role": "root"}`
	jsonByteInterfaceF := []byte(jsonStrInterfaceF)
	var jsonInterfaceF interface{}
	err = json.Unmarshal(jsonByteInterfaceF, &jsonInterfaceF)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonInterfaceMapF := jsonInterfaceF.(map[string]interface{})
	fmt.Printf("jsonInterfaceMapF: %v\n", jsonInterfaceMapF)
	fmt.Printf("jsonInterfaceMapF username: %v\n", jsonInterfaceMapF["username"])
	fmt.Printf("jsonInterfaceMapF role: %v\n", jsonInterfaceMapF["role"])
	fmt.Println("")

	fmt.Println("Part 53.3 Decode array json ke array objek")
	jsonStrArrayA := `[
		{"Name": "Noval", "Age": 26},
		{"Name": "Deni", "Age": 27},
		{"Name": "Novan", "Age": 27},
		{"Name": "Adrie", "Age": 29}
	]`
	jsonByteArrayA := []byte(jsonStrArrayA)
	var jsonArrayA []person
	err = json.Unmarshal(jsonByteArrayA, &jsonArrayA)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonArrayA: %v\n", jsonArrayA)
	for i, v := range jsonArrayA {
		fmt.Printf("jsonArrayA %d: %v\n", i, v)
	}
	fmt.Println("-----")
	jsonStrArrayB := `[
		{"panjang": 10, "lebar": 5, "tinggi": 8},
		{"panjang": 12, "lebar": 8, "tinggi": 13},
		{"panjang": 14, "lebar": 11, "tinggi": 18},
		{"panjang": 16, "lebar": 14, "tinggi": 23},
		{"panjang": 18, "lebar": 17, "tinggi": 30}
	]`
	jsonByteArrayB := []byte(jsonStrArrayB)
	var jsonArrayB []balok
	err = json.Unmarshal(jsonByteArrayB, &jsonArrayB)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("jsonArrayB: %v\n", jsonArrayB)
	for i, v := range jsonArrayB {
		fmt.Printf("jsonArrayB %d: %v\n", i, v)
	}
	fmt.Println("-----")
	type club struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		UclNumber int    `json:"ucl_number"`
	}
	jsonStrArrayC := `[
		{"name": "Manchester United", "country": "England", "ucl_number": 3},
		{"name": "Liverpool", "country": "England", "ucl_number": 6},
		{"name": "Real Madrid", "country": "Spain", "ucl_number": 13},
		{"name": "Barcelona", "country": "Spain", "ucl_number": 5},
		{"name": "Bayern Munich", "country": "Germany", "ucl_number": 6}
	]`
	jsonByteArrayC := []byte(jsonStrArrayC)
	var jsonArrayC []club
	err = json.Unmarshal(jsonByteArrayC, &jsonArrayC)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("jsonArrayC: %v\n", jsonArrayC)
	for i, v := range jsonArrayC {
		fmt.Printf("jsonArrayC %d: %v\n", i, v)
	}
	fmt.Println("")

	fmt.Println("Part 53.4 Encode json")
	jsonObjA := []club{
		{Name: "Manchester United", Country: "England", UclNumber: 3},
		{Name: "Chelsea", Country: "England", UclNumber: 1},
		{Name: "Everton", Country: "England", UclNumber: 1},
	}
	jsonObjByteA, err := json.Marshal(jsonObjA)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonObjStrA := string(jsonObjByteA)
	fmt.Printf("jsonObjByteA: %s\n", jsonObjStrA)
	jsonObjB := person{Name: "Novalita", AgeNumber: 26}
	jsonObjByteB, err := json.Marshal(jsonObjB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonObjStrB := string(jsonObjByteB)
	fmt.Printf("jsonObjStrB: %s\n", jsonObjStrB)
	jsonObjC := []balok{
		{P: 10, L: 2, T: 6},
		{P: 15, L: 9, T: 10},
		{P: 20, L: 16, T: 14},
		{P: 25, L: 23, T: 18},
		{P: 30, L: 25, T: 22},
	}
	jsonObjByteC, err := json.Marshal(jsonObjC)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonObjStrC := string(jsonObjByteC)
	fmt.Printf("jsonObjStrC: %s\n", jsonObjStrC)
}
