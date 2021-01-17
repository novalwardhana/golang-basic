package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type request struct {
	RequestType string
	RequestData interface{}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := make(chan []byte)
		go func() {
			data := []byte("Welcome to my API.....")
			result <- data
		}()
		w.Write(<-result)
	})

	http.HandleFunc("/new-message", routeNewMessage)
	http.HandleFunc("/show-message", routeShowMessage)
	http.HandleFunc("/new-merchant", routeNewMerchant)
	http.HandleFunc("/show-merchant", routeShowMerchant)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Server run at: %s", address))
	server.ListenAndServe()
}

func routeNewMessage(w http.ResponseWriter, r *http.Request) {
	var request request
	type params struct {
		Title string
	}
	var requestData = make(chan params)
	go func() {
		newRequestData := params{
			Title: "Message - Input",
		}
		requestData <- newRequestData
	}()
	request.RequestType = "Message Input"
	request.RequestData = <-requestData
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.New("new-message").ParseFiles("views/new-message.html"))
	err := tmpl.ExecuteTemplate(w, "new-message", request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeShowMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Methot not allowed", http.StatusInternalServerError)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")      // bisa menggunakan r.FormValue atau r.Form.get
	message := r.Form.Get("message") // bisa menggunakan r.FormValue atau r.Form.get
	data := map[string]string{
		"name":    name,
		"message": message,
	}
	tmpl := template.Must(template.New("show-message").ParseFiles("views/show-message.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error generate to template", http.StatusConflict)
	}
}

func routeNewMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var request request
	type params struct {
		Title string
	}
	var requestData = make(chan params)
	go func() {
		newRequestData := params{
			Title: "New Merchant",
		}
		requestData <- newRequestData
	}()
	request.RequestType = "Create new merchant"
	request.RequestData = <-requestData
	tmpl := template.Must(template.New("new-merchant").ParseFiles("views/new-merchant.html"))
	if err := tmpl.Execute(w, request); err != nil {
		http.Error(w, "Error execute content", http.StatusNoContent)
	}
}

func routeShowMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed...", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parse request", http.StatusBadRequest)
		return
	}
	var request request
	type params struct {
		Title string
		Data  map[string]string
	}
	var requestData = make(chan params)
	go func() {
		newRequestData := params{
			Title: "Show Merchant",
			Data: map[string]string{
				"Name":        r.FormValue("name"),
				"Email":       r.FormValue("email"),
				"PhoneNumber": r.Form.Get("phone_number"),
			},
		}
		requestData <- newRequestData
	}()
	request.RequestType = "Show New Merchant"
	request.RequestData = <-requestData
	tmpl := template.Must(template.New("show-merchant").ParseFiles("views/show-merchant.html"))
	if err := tmpl.Execute(w, request); err != nil {
		http.Error(w, "Error execute content", http.StatusNoContent)
	}
}
