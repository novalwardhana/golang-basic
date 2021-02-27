package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type renderer struct {
	template *template.Template
	location string
	debug    bool
}

func (r *renderer) ReloadTemplates() {
	r.template = template.Must(template.ParseGlob(r.location))
}

func (r *renderer) Render(w io.Writer, location string, data interface{}, c echo.Context) error {
	if r.debug {
		r.ReloadTemplates()
	}

	return r.template.ExecuteTemplate(w, location, data)
}

func newRenderer(location string, debug bool) *renderer {
	tpl := new(renderer)
	tpl.location = location
	tpl.debug = debug
	tpl.ReloadTemplates()
	return tpl
}

type dataMap map[string]interface{}

func main() {
	r := echo.New()
	r.Renderer = newRenderer("*.html", true)

	r.Any("/", routeIndex)
	r.Any("/user", routeUser)
	r.Any("/food", routeFood)

	r.Start(":9000")
}

func routeIndex(c echo.Context) error {
	output := make(chan dataMap)
	template := "index.html"
	go func() {
		data := dataMap{
			"Title":   "Index Page",
			"Code":    http.StatusOK,
			"Message": "Welcome to my API",
		}
		output <- data
	}()
	return c.Render(http.StatusOK, template, <-output)
}

func routeUser(c echo.Context) error {
	output := make(chan dataMap)
	template := "user.html"
	go func() {
		data := dataMap{
			"Code":    http.StatusOK,
			"Message": "Data generated",
			"Title":   "User Page",
			"Data": []struct {
				Email string `json:"username"`
				Name  string `json:"name"`
			}{
				{Email: "noval@ainosi.co.id", Name: "Noval"},
				{Email: "novalita.k.wardhana@gmail.com", Name: "Novalwardhana"},
				{Email: "noval@ainosi.com", Name: "Wardhana"},
			},
		}
		output <- data
	}()
	return c.Render(http.StatusOK, template, <-output)
}

func routeFood(c echo.Context) error {
	output := make(chan dataMap)
	template := "food.html"
	go func() {
		data := dataMap{
			"Title":   "Food Page",
			"Code":    http.StatusOK,
			"Message": "Data Generated",
			"Data": []struct {
				Code  string  `json:"code"`
				Name  string  `json:"name"`
				Price float64 `json:"price"`
			}{
				{Code: "f001", Name: "Sate Ayam", Price: 15000},
				{Code: "f002", Name: "Ayam Goreng", Price: 12000},
				{Code: "f004", Name: "Sate Kambing", Price: 25000},
			},
		}
		output <- data
	}()
	return c.Render(http.StatusOK, template, <-output)
}
