package main

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

var students []*student

func init() {
	students = append(students, &student{ID: "s001", Name: "Noval", Grade: 1})
	students = append(students, &student{ID: "s002", Name: "Kusuma", Grade: 2})
	students = append(students, &student{ID: "s003", Name: "Wardhana", Grade: 3})
}

func studentGetData(w http.ResponseWriter, r *http.Request) {
	if id := r.FormValue("id"); id != "" {
		resultData := make(chan student)
		go func() {
			data := <-studentGetDataSingle(id)
			resultData <- data
		}()
		resultDataJSON, _ := json.Marshal(<-resultData)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resultDataJSON)
		return
	}
	resultData := make(chan []*student)
	go func() {
		data := <-studentGetDataAll()
		resultData <- data
	}()
	resultDataJSON, _ := json.Marshal(<-resultData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultDataJSON)
}

func studentGetDataSingle(id string) <-chan student {
	result := make(chan student)
	go func() {
		for _, student := range students {
			if id == (*student).ID {
				result <- *student
				break
			}
		}
		result <- student{}
	}()
	return result
}

func studentGetDataAll() <-chan []*student {
	result := make(chan []*student)
	go func() {
		result <- students
	}()
	return result
}

func studentAddData(w http.ResponseWriter, r *http.Request) {
	var data student
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Status request not valid", http.StatusBadRequest)
		return
	}
	students = append(students, &data)
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Success add data",
	}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
