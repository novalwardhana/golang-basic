package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func fetchStudents() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/students", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchStudent(id string) (student, error) {
	var err error
	var data student
	var client = &http.Client{}

	var params = url.Values{}
	params.Set("id", id)
	payload := bytes.NewBufferString(params.Encode())

	request, err := http.NewRequest("POST", baseURL+"/student", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

type club struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	OrderNum int    `json:"order_num"`
}

func fetchClubs() ([]club, error) {
	var err error
	var client = &http.Client{}
	var data []club

	request, err := http.NewRequest("POST", baseURL+"/clubs", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchClub(id string) (club, error) {
	var data club
	var client = &http.Client{}
	var err error

	var params = url.Values{}
	params.Set("id", id)
	payload := bytes.NewBufferString(params.Encode())

	request, err := http.NewRequest("POST", baseURL+"/club", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, nil
	}

	return data, nil
}

type notebook struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Tipe  string `json:"tipe"`
	Seri  string `json:"seri"`
	Price int    `json:"price"`
}

func fetchNotebooks() ([]notebook, error) {
	//var data []notebook
	var client = &http.Client{}
	var err error
	var data []notebook

	request, err := http.NewRequest("POST", baseURL+"/notebooks", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchNotebook(id string) (notebook, error) {
	var data notebook
	var err error
	var client = &http.Client{}

	params := url.Values{}
	params.Set("id", id)
	payload := bytes.NewBufferString(params.Encode())

	request, err := http.NewRequest("POST", baseURL+"/notebook", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)

	return data, nil
}

func main() {
	fmt.Println("Part 55 Simple Client HTTP Request")
	var students []student
	var err error
	fmt.Println("Get all student data")
	students, err = fetchStudents()
	if err != nil {
		fmt.Printf("Cannot get students data: %s\n", err.Error())
		return
	}
	for i, each := range students {
		fmt.Printf("Student %d: %#v\n", i, each)
	}
	fmt.Println("-----")

	fmt.Println("Get student data")
	student, err := fetchStudent("2")
	if err != nil {
		fmt.Printf("Cannot get student data: %s\n", err.Error())
		return
	}
	fmt.Printf("Student: %#v\n", student)
	fmt.Println("-----")

	fmt.Println("Get all club data")
	var clubs []club
	clubs, err = fetchClubs()
	if err != nil {
		fmt.Printf("Cannot get clubs data: %s\n", err.Error())
		return
	}
	for i, each := range clubs {
		fmt.Printf("Club %d: %#v\n", i, each)
	}
	fmt.Println("-----")

	fmt.Println("Get club data")
	var club club
	club, err = fetchClub("2")
	if err != nil {
		fmt.Printf("Cannot get club data: %s\n", err.Error())
		return
	}
	fmt.Printf("Club: %#v\n", club)
	fmt.Println("-----")

	fmt.Println("Get all notebook data")
	var notebooks []notebook
	notebooks, err = fetchNotebooks()
	if err != nil {
		fmt.Printf("Cannot get notebook data: %s\n", err.Error())
		return
	}
	for i, each := range notebooks {
		fmt.Printf("Notebook %d: %#v\n", i, each)
	}
	fmt.Println("-----")

	fmt.Println("Get notebook data")
	var notebook notebook
	notebook, err = fetchNotebook("2")
	if err != nil {
		fmt.Printf("Cannot get notebook data: %s\n", err.Error())
		return
	}
	fmt.Printf("Notebook data: %#v\n", notebook)
	fmt.Println("-----")
}
