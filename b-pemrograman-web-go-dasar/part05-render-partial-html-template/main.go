package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type dataResponse map[string]interface{}

func main() {
	tmpl, err := template.ParseGlob("views/*")
	if err != nil {
		logrus.Error(err.Error())
		panic(err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan string)
		go func() {
			data := "Hello, welcome to my api"
			result <- data
		}()
		w.Write([]byte(<-result))
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var result = make(chan dataResponse)
		go func() {
			var data = dataResponse{"name": "Noval", "category": "Index"}
			result <- data
		}()
		err := tmpl.ExecuteTemplate(w, "index", <-result)
		if err != nil {
			logrus.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var result = make(chan dataResponse)
		go func() {
			var data = dataResponse{"name": "New Year's Eve", "category": "About"}
			result <- data
		}()
		err := tmpl.ExecuteTemplate(w, "about", <-result)
		if err != nil {
			logrus.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/index-two", func(w http.ResponseWriter, r *http.Request) {
		var result = make(chan dataResponse)
		go func() {
			var data = dataResponse{"name": "NovalWardhana93", "category": "Index Two"}
			result <- data
		}()
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(w, "index", <-result)
		if err != nil {
			logrus.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about-two", func(w http.ResponseWriter, r *http.Request) {
		var result = make(chan dataResponse)
		go func() {
			var data = dataResponse{"name": "NvlWrdhn93", "category": "About Two"}
			result <- data
		}()
		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(w, "about", <-result)
		if err != nil {
			logrus.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	address := "localhost:9000"
	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Run port: %s\n", address))
	server.ListenAndServe()
}
