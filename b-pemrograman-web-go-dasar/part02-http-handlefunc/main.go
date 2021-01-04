package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	handleIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my API Golang"))
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world bro"))
	})

	address := "localhost:9000"
	fmt.Printf("lister to server: %s\n", address)
	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
