package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"part22-simple-configuration/conf"
	"path/filepath"
	"time"
)

type customMux struct {
	http.ServeMux
}

func (c *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Configuration().Log.Verbose {
		log.Println("Request from", r.Host, " acsessing ", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	mux := new(customMux)
	mux.HandleFunc("/", routeIndex)
	mux.HandleFunc("/data", routeGetData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)
	server.ReadTimeout = conf.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = conf.Configuration().Server.WriteTimeout * time.Second

	if conf.Configuration().Log.Verbose {
		log.Printf("Run server at:%d", conf.Configuration().Server.Port)
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func routeIndex(w http.ResponseWriter, r *http.Request) {
	result := make(chan []byte)
	go func() {
		result <- []byte("Welcome to my API")
	}()
	w.Write(<-result)
}

type food struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func routeGetData(w http.ResponseWriter, r *http.Request) {
	var output = make(chan []food)
	go func() {
		var foods []food
		basepath, e := os.Getwd()
		if e != nil {
			return
		}
		dataJSON, e := ioutil.ReadFile(filepath.Join(basepath, "data.json"))
		if e != nil {
			return
		}
		e = json.Unmarshal(dataJSON, &foods)
		if e != nil {
			return
		}
		output <- foods
	}()
	outputJSON, e := json.Marshal(<-output)
	if e != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(outputJSON)
}
