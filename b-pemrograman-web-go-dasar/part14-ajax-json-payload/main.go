package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/student", routeStudent)
	http.HandleFunc("/student-submit", routeStudentSubmit)
	http.HandleFunc("/user", routeUser)
	http.HandleFunc("/user-submit", routeUserSubmit)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Start server at: %s\n", address)
	server.ListenAndServe()
}
