package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var templateFunc = template.FuncMap{
	"greeting": func() template.HTML {
		result := make(chan string)
		go func() {
			result <- "Welcome to"
		}()
		return template.HTML(<-result)
	},
	"clubName": func(s string) template.HTML {
		result := make(chan string)
		go func() {
			result <- s
		}()
		return template.HTML(<-result)
	},
}

type request struct {
	RequestType string      `json:"request_type"`
	RequestData interface{} `json:"request_data"`
}

type club struct {
	Name    string
	Stadium string
}

type cli struct {
	Length int
	Width  int
	Height int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan []byte)
		go func() {
			data := []byte("Welcome to my API")
			result <- data
		}()
		w.Write(<-result)
	})

	http.HandleFunc("/manchester-united", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("manchester-united-page").Funcs(templateFunc).ParseFiles("views/views.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/real-madrid", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("real-madrid-page").Funcs(templateFunc).ParseFiles("views/views.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/arsenal", func(w http.ResponseWriter, r *http.Request) {
		var request request
		var clubData = make(chan club)
		go func() {
			data := club{
				Name:    "Arsenal",
				Stadium: "Emirates",
			}
			clubData <- data
		}()
		request.RequestType = "Football Club"
		request.RequestData = <-clubData

		tmpl := template.Must(template.New("arsenal-page").Funcs(templateFunc).ParseFiles("views/views.html"))
		err := tmpl.ExecuteTemplate(w, "arsenal-page", request)
		if err != nil {
			logrus.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Start server in address: %s\n", address))
	server.ListenAndServe()
}
