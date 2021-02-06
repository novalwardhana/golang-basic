package main

import (
	"log"
	"net/http"
	"time"
)

const usernameAuth = "Noval"
const passwordAuth = "Noval2021"

func main() {
	mux := new(customMux)
	mux.HandleFunc("/", routingIndex)
	mux.registerMiddleware(middlewareMethod)

	muxStudent := new(customMuxStudent)
	muxStudent.HandleFunc("/student", studentGetData)
	muxStudent.HandleFunc("/student-add", studentAddData)
	muxStudent.registerMiddleware(middlewareMethodStudent)
	muxStudent.registerMiddleware(middlewareAuthStudent)

	muxFood := new(customMuxFood)
	muxFood.HandleFunc("/food", foodGetData)
	muxFood.HandleFunc("/food-add", foodAddData)
	muxFood.registerMiddleware(middlewareMethodFood)
	muxFood.registerMiddleware(middlewareAuthFood)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.Handler = muxFood
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Start service at: %s\n", address)
	server.ListenAndServe()
}

func routingIndex(w http.ResponseWriter, r *http.Request) {
	response := make(chan []byte)
	go func() {
		response <- []byte("Welcome to My API Service")
	}()
	w.Write(<-response)
}
