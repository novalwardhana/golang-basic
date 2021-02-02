package main

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Grade int32  `json:"grade"`
}

var students = []*student{}

func init() {
	students = append(students, &student{ID: "w001", Name: "Noval", Grade: 3})
	students = append(students, &student{ID: "w002", Name: "Kusuma", Grade: 5})
	students = append(students, &student{ID: "w003", Name: "Wardhana", Grade: 7})
}

func getStudentData(id string) <-chan student {
	var result = make(chan student)
	go func() {
		for _, data := range students {
			if (*data).ID == id {
				result <- *data
				break
			}
		}
		result <- student{}
	}()
	return result
}

func getAllStudentData() <-chan []*student {
	var result = make(chan []*student)
	go func() {
		result <- students
	}()
	return result
}

func outputStudentData(w http.ResponseWriter, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
