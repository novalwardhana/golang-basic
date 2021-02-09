package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	mux := new(customMux)
	mux.HandleFunc("/", routeIndex)

	mux.HandleFunc("/add-cookie", routeAddCookie)
	mux.HandleFunc("/add-cookie-food", routeAddCookieFood)
	mux.HandleFunc("/add-cookie-cofee", routeAddCookieCofee)

	mux.HandleFunc("/delete-cookie", routeDeleteCookie)
	mux.HandleFunc("/delete-cookie-food", routeDeleteCookieFood)
	mux.HandleFunc("/delete-cookie-cofee", routeDeleteCookieCofee)

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.Handler = mux
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	log.Printf("Run server at: %v", address)
	server.ListenAndServe()
}

func routeIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").ParseFiles("views/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func routeAddCookie(w http.ResponseWriter, r *http.Request) {
	cookieName := "CookieData"
	c := &http.Cookie{}
	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}
	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = "One Ok Rock"
		c.Expires = time.Now().Add(1 * time.Minute)
		http.SetCookie(w, c)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func routeDeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookieName := "CookieData"
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func routeAddCookieFood(w http.ResponseWriter, r *http.Request) {
	type food struct {
		ID    int    `json:"id"`
		Code  string `json:"code"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	var foods []*food
	for i, data := range []string{"Sate Ayam", "Sate Kambing", "Sate Kuda", "Sate Sapi"} {
		foods = append(foods, &food{ID: i + 1, Code: fmt.Sprintf("foo-%05d", i+1), Name: data, Price: (10000 + (i * 1000))})
	}
	result, _ := json.Marshal(foods)
	resultStr := fmt.Sprintf("%s", string(result))
	resultStr = strings.ReplaceAll(resultStr, "\"", "")

	cookieName := "cookieFood"
	c := &http.Cookie{}
	if setCookie, _ := r.Cookie(cookieName); setCookie != nil {
		c = setCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = resultStr
		c.Expires = time.Now().Add(5 * time.Second)
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func routeDeleteCookieFood(w http.ResponseWriter, r *http.Request) {
	cookieName := "foodCookie"
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func routeAddCookieCofee(w http.ResponseWriter, r *http.Request) {
	type cofee struct {
		ID    int    `json:"id"`
		Code  string `json:"code"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	var cofees []*cofee
	for i, data := range []string{"Espresso", "Cappuccino", "Macchiato", "Americano", "Cafe Latte"} {
		cofees = append(cofees, &cofee{ID: i + 1, Code: fmt.Sprintf("cof-%05d", i+1), Name: data, Price: 15000 + 500*i})
	}
	resultJSON, _ := json.Marshal(cofees)
	resultStr := string(resultJSON)
	resultStr = strings.ReplaceAll(resultStr, "\"", "")

	cookieName := "cookieCofee"
	c := &http.Cookie{}
	if setCookie, _ := r.Cookie(cookieName); setCookie != nil {
		c = setCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = resultStr
		c.Expires = time.Now().Add(10 * time.Minute)
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func routeDeleteCookieCofee(w http.ResponseWriter, r *http.Request) {
	cookieName := "cookieCofee"
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
