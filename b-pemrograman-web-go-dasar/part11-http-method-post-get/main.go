package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan []byte)
		go func() {
			data := ""
			if r.Method == "POST" {
				data = "Method is POST"
			} else if r.Method == "GET" {
				data = "Method is GET"
			} else {
				data = "Method not allowed"
			}
			result <- []byte(data)
		}()
		w.Write(<-result)
	})

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.WriteTimeout = 10 * time.Second
	server.ReadTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Server run at: %s", address))
	server.ListenAndServe()
}
