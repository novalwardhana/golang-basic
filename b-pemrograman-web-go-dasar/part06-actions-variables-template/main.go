package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/sirupsen/logrus"
)

type info struct {
	Affliation string `json:"affliation"`
	Address    string `json:"address"`
}

type person struct {
	Name    string   `json:"name"`
	Gender  string   `json:"gender"`
	Hobbies []string `json:"hobbies"`
	Info    info     `json:"info"`
}

type club struct {
	Name         string   `json:"name"`
	Competitions []string `json:"competition"`
	Stadium      string   `json:"stadium"`
	Origin       string   `json:"origin"`
	Sponsors     []string `json:"sponsors"`
	Info         info     `json:"info"`
	DataJSON     string   `json:"data_json"`
	Status       bool     `json:"status"`
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

	http.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan person)
		go func() {
			data := person{
				Name:    "Cristiano Ronaldo",
				Gender:  "Male",
				Hobbies: []string{"Football", "Hiking"},
				Info: info{
					Affliation: "Nike",
					Address:    "Portugal",
				},
			}
			result <- data
		}()
		tmpl := template.Must(template.ParseFiles("views/person.html"))
		err := tmpl.Execute(w, <-result)
		if err != nil {
			logrus.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/club", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan club)
		go func() {
			data := club{
				Name: "Manchester United",
				Competitions: []string{
					"English Primeir League",
					"FA CUP",
					"Carabao Cup",
					"UEFA Champions League",
					"UEFA League",
				},
				Stadium: "Old Trafford",
				Origin:  "England",
				Sponsors: []string{
					"Chevrolert",
					"AON",
					"Adidas",
					"Kohler",
				},
				Info: info{
					Address:    "Manchester",
					Affliation: "Real Madrid",
				},
				Status: true,
			}
			dataJSON, _ := json.Marshal(data)
			data.DataJSON = string(dataJSON)
			result <- data
		}()
		tmpl := template.Must(template.ParseFiles("views/club.html"))
		err := tmpl.ExecuteTemplate(w, "club", <-result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Start server in adress: %s\n", address))
	server.ListenAndServe()
}
