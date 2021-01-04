package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/hello-world", handleHello)

	address := "localhost:9000"
	fmt.Printf("Started at: %s\n", address)
	// err := http.ListenAndServe(address, nil)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }

	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.ReadTimeout = time.Second * 10
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	message := "Welcome to my API"
	w.Write([]byte(message))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello world"
	w.Write([]byte(message))
}
