package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

var ourStudent = []student{
	student{1, "Noval", 3},
	student{2, "Adrie", 3},
	student{3, "Deni", 2},
	student{4, "Novan", 4},
}

type club struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	OrderNum int    `json:"order_num"`
}

var ourClub = []club{
	club{1, "MANU", "Manchester United", 3},
	club{2, "RMA", "Real Madrid", 13},
	club{3, "FCB", "Barcelona", 5},
	club{4, "Juv", "Juventus", 2},
	club{5, "ACM", "AC Milan", 7},
}

type notebook struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Tipe  string `json:"tipe"`
	Seri  string `json:"seri"`
	Price int    `json:"price"`
}

var ourNotebook = []notebook{
	notebook{1, "Lenovo", "Ideapad", "Slim 3", 8300000},
	notebook{2, "Lenovo", "Thinkpad", "P1", 18300000},
	notebook{3, "Asus", "Zenbook", "UM425", 9300000},
	notebook{4, "Asus", "Zenbook", "UM431", 9700000},
}

func main() {
	fmt.Println("Part 54 Web Service API Server")
	fmt.Println("")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to web service")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		data, err := json.Marshal(ourStudent)

		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		if r.Method == "POST" {
			w.Write(data)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	})
	http.HandleFunc("/student", getStudent)

	http.HandleFunc("/clubs", getClubs)
	http.HandleFunc("/club", getClub)

	http.HandleFunc("/notebooks", getNotebooks)
	http.HandleFunc("/notebook", getNotebook)

	http.ListenAndServe(":8080", nil)
}

func getStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	paramID := r.FormValue("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var data []byte
	for _, each := range ourStudent {
		if each.ID == id {
			data, err = json.Marshal(each)
			if err != nil {
				http.Error(w, "", http.StatusBadRequest)
			}
			w.Write(data)
			return
		}
	}

	http.Error(w, "Cannot find data", http.StatusNotFound)
}

func getClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var data []byte
		var err error
		data, err = json.Marshal(ourClub)
		if err != nil {
			http.Error(w, "Error parsing json", http.StatusBadRequest)
			return
		}
		w.Write(data)
		return
	}

	http.Error(w, "Status method not allowed", http.StatusMethodNotAllowed)
}

func getClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var err error
		var id int
		var data []byte
		paramID := r.FormValue("id")
		id, err = strconv.Atoi(paramID)
		if err != nil {
			http.Error(w, "Status bad request", http.StatusBadRequest)
			return
		}
		for _, each := range ourClub {
			if each.ID == id {
				data, err = json.Marshal(each)
				fmt.Println("Haloo", each)
				if err != nil {
					http.Error(w, "Status no content", http.StatusNoContent)
					return
				}
				w.Write(data)
				return
			}
		}
		http.Error(w, "Status not found", http.StatusNotFound)
		return
	}

	http.Error(w, "Status method not allowed", http.StatusMethodNotAllowed)
}

func getNotebooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		data, err := json.Marshal(ourNotebook)
		if err != nil {
			http.Error(w, "Status no content", http.StatusNoContent)
			return
		}
		w.Write(data)
		return
	}

	http.Error(w, "Status method not allowed", http.StatusMethodNotAllowed)
}

func getNotebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		paramsID := r.FormValue("id")
		id, err := strconv.Atoi(paramsID)
		if err != nil {
			http.Error(w, "Status bad request", http.StatusBadRequest)
			return
		}
		for _, each := range ourNotebook {
			if each.ID == id {
				data, err := json.Marshal(each)
				if err != nil {
					http.Error(w, "Data not found", http.StatusNotFound)
					return
				}
				w.Write(data)
				return
			}
		}
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}

	http.Error(w, "Status method not allowed", http.StatusMethodNotAllowed)
}
