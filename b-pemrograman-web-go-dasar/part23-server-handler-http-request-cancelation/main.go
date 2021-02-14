package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"part23-server-handler-http-request-cancelation/conf"
	"strings"
	"time"
)

type customMux struct {
	http.ServeMux
}

func (c *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Configuration().Log.Verbose {
		log.Println("Request from: ", r.Host, "access ", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	mux := new(customMux)
	mux.HandleFunc("/", routeIndex)
	mux.HandleFunc("/add-data", routeAddData)
	mux.HandleFunc("/get-data", routeGetData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)
	if conf.Configuration().Log.Verbose {
		log.Printf("Start server at: %d", conf.Configuration().Server.Port)
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

func routeIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println(strings.ToLower(err.Error()))
			} else {
				log.Println("Other error occured: ", err.Error())
			}
		}
	case <-done:
		w.Write([]byte("Welcome to my API...."))
	}
}

type food struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var foods []*food

func routeAddData(w http.ResponseWriter, r *http.Request) {

	done := make(chan bool)
	data := make(chan food)

	go func() {
		dataBody, _ := ioutil.ReadAll(r.Body)
		log.Println(string(dataBody))
		time.Sleep(10 * time.Second)
		done <- true
	}()
	go func() {
		newFood := food{}
		err := json.NewDecoder(r.Body).Decode(&newFood)
		if err != nil {
			data <- food{}
		} else {
			data <- newFood
		}
	}()

	food := <-data
	foods = append(foods, &food)

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println(strings.ToLower(err.Error()))
			} else {
				log.Println("Other error occured: ", err.Error())
			}
		}
	case <-done:
		w.Write([]byte("Data has been processed"))
	}
}

func routeGetData(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	foodsJSON, err := json.Marshal(foods)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println(err.Error())
			} else {
				log.Println("Other error occured: ", err.Error())
			}
		}
	case <-done:
		w.Header().Set("Content-Type", "application/json")
		w.Write(foodsJSON)
	}
}
