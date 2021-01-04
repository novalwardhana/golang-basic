package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.Handle("/file", http.FileServer(http.Dir("assets")))

	http.Handle("/file-css/",
		http.StripPrefix("/file-css/",
			http.FileServer(http.Dir("assets"))))

	http.Handle("/my-lyric/",
		http.StripPrefix("/my-lyric/",
			http.FileServer(http.Dir("lyric"))))

	address := "localhost:9000"
	fmt.Printf("Lister to server: %s\n", address)
	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
