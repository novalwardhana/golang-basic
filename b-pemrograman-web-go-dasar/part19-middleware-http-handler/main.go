package main

import (
	"log"
	"net/http"
	"time"
)

const usernameAuth string = "noval"
const passwordAuth string = "novalnoval"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan []byte)
		go func() {
			result <- []byte("Welcome to my API")
		}()
		w.Write(<-result)
	})

	muxStudent := http.DefaultServeMux
	muxStudent.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		if id := r.URL.Query().Get("id"); id != "" {
			outputStudentData(w, <-getStudentData(id))
			return
		}
		outputStudentData(w, <-getAllStudentData())
	})
	var handlerStudent http.Handler = muxStudent
	handlerStudent = middlewareMethodStudent(handlerStudent)
	handlerStudent = middlewareAuthStudent(handlerStudent)

	muxFood := http.DefaultServeMux
	muxFood.HandleFunc("/food", func(w http.ResponseWriter, r *http.Request) {
		if id := r.URL.Query().Get("id"); id != "" {
			var output = make(chan food)
			go func() {
				dataFood := <-getFoodData(id)
				output <- dataFood
			}()
			outputFoodData(w, <-output)
			return
		}
		var output = make(chan []*food)
		go func() {
			dataFood := <-getAllFoodData()
			output <- dataFood
		}()
		outputFoodData(w, <-output)
	})
	var handlerFood http.Handler = muxFood
	handlerFood = middlewareAuthFood(handlerFood)
	handlerFood = middlewareMethodFood(handlerFood)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.Handler = handlerStudent
	//server.Handler = handlerFood
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Run service at: %s", address)
	server.ListenAndServe()
}
