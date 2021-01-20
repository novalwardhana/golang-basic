package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type request struct {
	RequestType string
	RequestData interface{}
}

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func routeUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var params = make(chan *request)
	go func() {
		params <- &request{
			RequestType: "UR",
			RequestData: struct {
				Title    string
				UserRole []string
			}{
				Title:    "User Registration",
				UserRole: []string{"Superadmin", "Admin", "Employee", "Guest"},
			},
		}
	}()
	tmpl := template.Must(template.New("user").ParseFiles("views/user.html"))
	if err := tmpl.ExecuteTemplate(w, "user", *<-params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func routeUserSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var payload *user
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&payload)
		wg.Done()
	}()
	wg.Wait()
	response := fmt.Sprintf("Your data: %#v", *payload)
	w.Write([]byte(response))
}
