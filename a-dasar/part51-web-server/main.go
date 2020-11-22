package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintln(w, "Welcome to index page..")
}

func main() {
	fmt.Println("Part 51 Web Server")

	http.HandleFunc("/", func(w http.ResponseWriter, h *http.Request) {
		fmt.Fprintln(w, "Selamat datang..")
	})

	http.HandleFunc("/barcelona", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"club":    "Barcelona",
			"stadium": "Camp Nou",
			"captain": "Lionel Messi",
		}
		var t, err = template.ParseFiles("/home/novalwardhana/novalAgung/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/real-madrid", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"club":    "Real Madrid",
			"stadium": "Santiago Bernabeu",
			"captain": "Sergio Ramos",
		}
		var t, err = template.ParseFiles("/home/novalwardhana/novalAgung/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)

}
