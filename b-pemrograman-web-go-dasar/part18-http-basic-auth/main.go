package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const usernameAuth = "noval"
const passwordAuth = "novalnoval"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan []byte)
		go func() {
			result <- []byte("Welcome to my API service ...")
		}()
		w.Write(<-result)
	})

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		if !allowOnlyGet(w, r) {
			return
		}
		if !authStudent(w, r) {
			return
		}
		if id := r.URL.Query().Get("id"); id != "" {
			outputStudentData(w, <-getStudentData(id))
			return
		}

		outputStudentData(w, <-getStudentsData())
	})

	http.HandleFunc("/club", func(w http.ResponseWriter, r *http.Request) {
		if !allowOnlyGet(w, r) {
			return
		}
		if !authClubAPIKey(w, r) {
			return
		}
		// if !authClubBasic(w, r) {
		// 	return
		// }
		if id := r.URL.Query().Get("id"); id != "" {
			outputClubData(w, <-getClubData(id))
			return
		}
		outputClubData(w, <-getClubsData())
	})

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Run service at %s\n", address)
	server.ListenAndServe()
}

func allowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return false
	}
	return true
}
