package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type food struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var foods []*food

func init() {
	foodInit := []string{"Mie Ayam", "Bakso", "Sate Ayam", "Sate Sapi"}
	for i, data := range foodInit {
		foods = append(foods, &food{ID: i + 1, Code: fmt.Sprintf("%05d", i+1), Name: data, Price: 10000 + ((i + 1) * 500)})
	}
}

func foodGetData(w http.ResponseWriter, r *http.Request) {
	if id := r.FormValue("id"); id != "" {
		responseData := make(chan food)
		go func() {
			data := <-foodGetDataSingle(id)
			responseData <- data
		}()
		responseDataJSON, _ := json.Marshal(<-responseData)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseDataJSON)
		return
	}
	responseData := make(chan []*food)
	go func() {
		data := <-foodGetDataAll()
		responseData <- data
	}()
	responseDataJSON, _ := json.Marshal(<-responseData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseDataJSON)
}

func foodGetDataSingle(id string) <-chan food {
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		idNumber = 0
	}
	result := make(chan food)
	go func() {
		for _, data := range foods {
			if (*data).ID == idNumber {
				result <- *data
				break
			}
		}
		result <- food{}
	}()
	return result
}

func foodGetDataAll() <-chan []*food {
	result := make(chan []*food)
	go func() {
		result <- foods
	}()
	return result
}

func foodAddData(w http.ResponseWriter, r *http.Request) {
	foodData := make(chan food)
	err := make(chan error)
	go func() {
		data := food{}
		err <- json.NewDecoder(r.Body).Decode(&data)
		foodData <- data
	}()
	if (<-err) != nil {
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Request body not valid",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return
	}
	newData := <-foodData
	foods = append(foods, &newData)
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Successfully add new data",
	}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
