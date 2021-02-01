package main

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int32
}

var students = []*student{}

func init() {
	students = append(students, &student{ID: "s001", Name: "Noval", Grade: 3})
	students = append(students, &student{ID: "s002", Name: "Kusuma", Grade: 4})
	students = append(students, &student{ID: "s003", Name: "Wardhana", Grade: 5})
}

func getStudentData(id string) <-chan student {
	result := make(chan student)
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

func getStudentsData() <-chan []*student {
	result := make(chan []*student)
	go func() {
		result <- students
	}()
	return result
}

func outputStudentData(w http.ResponseWriter, data interface{}) {
	responseJSON, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
