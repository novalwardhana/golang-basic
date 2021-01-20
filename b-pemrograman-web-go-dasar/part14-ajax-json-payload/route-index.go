package main

import "net/http"

func routeIndex(w http.ResponseWriter, r *http.Request) {
	result := make(chan *[]byte)
	go func() {
		data := []byte("Welcome to my service API.....")
		result <- &data
	}()
	w.Write(*<-result)
}
