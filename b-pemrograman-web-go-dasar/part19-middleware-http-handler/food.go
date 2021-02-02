package main

import (
	"encoding/json"
	"net/http"
)

type food struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

var foods = []*food{}

func init() {
	foods = append(foods, &food{ID: "f001", Name: "Sate Kambing", Price: 20000})
	foods = append(foods, &food{ID: "f002", Name: "Sate Ayam", Price: 15000})
	foods = append(foods, &food{ID: "f003", Name: "Bakso", Price: 15000})
	foods = append(foods, &food{ID: "foo4", Name: "Mie Ayam", Price: 11000})
	foods = append(foods, &food{ID: "f005", Name: "Burger", Price: 18000})
}

func getFoodData(id string) <-chan food {
	output := make(chan food)
	go func() {
		for _, data := range foods {
			if (*data).ID == id {
				output <- *data
			}
			output <- food{}
		}
	}()
	return output
}

func getAllFoodData() <-chan []*food {
	output := make(chan []*food)
	go func() {
		output <- foods
	}()
	return output
}

func outputFoodData(w http.ResponseWriter, data interface{}) {
	responseJSON, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
