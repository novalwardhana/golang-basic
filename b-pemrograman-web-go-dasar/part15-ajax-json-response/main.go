package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := make(chan *[]byte)
		go func() {
			data := []byte("Welcome to my API... :-)")
			response <- &data
		}()
		w.Write(*<-response)
	})

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusInternalServerError)
			return
		}

		data := []struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}{
			{"Adrie", 18, "Male"},
			{"Deni", 16, "Male"},
			{"Novan", 17, "Male"},
			{"Nida", 16, "Female"},
		}

		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

	})

	http.HandleFunc("/foods", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusInternalServerError)
			return
		}
		result := make(chan []byte)
		go func() {
			data := []struct {
				Name  string `json:"name"`
				Price int    `json:"price"`
			}{
				{"Sate Klatak", 20000},
				{"Sate Ayam", 15000},
				{"Bubur Ayam", 12000},
				{"Ayam Goreng", 12000},
				{"Pecel Lele", 8000},
				{"Garang Asem", 15000},
			}
			dataJSON, _ := json.Marshal(data)
			result <- dataJSON
		}()
		w.Header().Set("Content-Type", "application/json")
		w.Write(<-result)
	})

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Run server at: %s", address)
	server.ListenAndServe()
}
