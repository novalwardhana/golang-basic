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

var funcStatistic = template.FuncMap{
	"addHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
	"mean": func(datas []int) float64 {
		result := make(chan float64)
		go func() {
			if len(datas) <= 0 {
				return
			}
			var sumData float64
			for _, data := range datas {
				sumData += float64(data)
			}
			meanData := sumData / float64(len(datas))
			result <- meanData
		}()
		return <-result
	},
	"min": func(datas ...int) int {
		result := make(chan int)
		go func() {
			min := 0
			for i, data := range datas {
				if i == 0 {
					min = data
				}
				if min > data {
					min = data
				}
			}
			result <- min
		}()
		return <-result
	},
	"max": func(datas []int) int {
		result := make(chan int)
		go func() {
			max := 0
			for i, data := range datas {
				if i == 0 {
					max = data
				}
				if max < data {
					max = data
				}
			}
			result <- max
		}()
		return <-result
	},
}

var funcBlockSpace = template.FuncMap{
	"title": func(s string) template.HTML {
		return template.HTML(s)
	},
	"surfaceArea": func(l, w, h int) int {
		result := make(chan int)
		go func() {
			surface := (2 * l * w) + (2 * w * h) + (2 * l * h)
			result <- surface
		}()
		return <-result
	},
	"volume": func(l, w, h int) int {
		result := make(chan int)
		go func() {
			volume := l * w * h
			result <- volume
		}()
		return <-result
	},
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

	http.HandleFunc("/statistic", func(w http.ResponseWriter, r *http.Request) {
		var request request
		requestData := make(chan []int)
		go func() {
			data := []int{9, 5, 6, 7, 3, 10, 7, 8, 6, 8}
			requestData <- data
		}()
		request.RequestType = "Statistic"
		request.RequestData = <-requestData
		tmpl := template.Must(template.New("statistic.html").Funcs(funcStatistic).ParseFiles("views/statistic.html"))
		err := tmpl.Execute(w, request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/block-space", func(w http.ResponseWriter, r *http.Request) {
		type blockSpace struct {
			Length int
			Width  int
			Height int
		}
		var request request
		var requestData = make(chan blockSpace)
		go func() {
			data := blockSpace{
				Length: 12,
				Width:  6,
				Height: 15,
			}
			requestData <- data
		}()
		tmpl := template.Must(template.New("block-space.html").Funcs(funcBlockSpace).ParseFiles("views/block-space.html"))
		request.RequestType = "Block Space"
		request.RequestData = <-requestData
		err := tmpl.Execute(w, request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	address := "localhost:9000"
	logrus.Info(fmt.Sprintf("Start server address: %s\n", address))
	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.ListenAndServe()
}
