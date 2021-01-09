package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type dcSuperhero struct {
	Name   string   `json:"name"`
	Alias  string   `json:"alias"`
	Member []string `json:"member"`
}

type marvelSuperhero struct {
	Name   string   `json:"name"`
	Alias  string   `json:"alias"`
	Lead   string   `json:"lead"`
	Member []string `json:"member"`
}

func (d dcSuperhero) sayHello(str1, str2 string) string {
	// How to call in html template
	// {{.SayHello "Gotham citizen" "You are our hero!"}}
	// {{end}}
	return fmt.Sprintf("%s said: \"%s\"", str1, str2)
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

	http.HandleFunc("/dc-superhero", func(w http.ResponseWriter, r *http.Request) {
		var result = make(chan dcSuperhero)
		var err error
		tmpl := template.Must(template.ParseFiles("views/dc-superhero.html"))
		go func() {
			data := dcSuperhero{
				Name:   "DC Universe",
				Alias:  "DCU",
				Member: []string{"Batman", "Superman", "Aquaman", "Wonder Woman", "Flash", "Cyborg"},
			}
			result <- data
		}()
		err = tmpl.Execute(w, <-result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/marvel-superhero", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan marvelSuperhero)
		go func() {
			data := marvelSuperhero{
				Name:  "Marvel Universe",
				Alias: "MCU",
				Lead:  "Captain America",
				Member: []string{
					"Captain America",
					"Iron Man",
					"Black Widow",
					"Thor",
					"Black Panther",
					"Spiderman",
					"Antman",
				},
			}
			result <- data
		}()
		tmpl := template.Must(template.ParseFiles("views/marvel-superhero.html"))
		err := tmpl.ExecuteTemplate(w, "marvel-superhero", <-result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Start server adress: %s\n", address))
	server.ListenAndServe()
}
