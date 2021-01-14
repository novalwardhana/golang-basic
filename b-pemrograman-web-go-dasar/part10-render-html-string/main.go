package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var views = `
	<html>
	<head>
		<title>Views Page</title>
	</head>
	<body>
		Welcome to page views
	</body>
	</html>
	`
var statisticPage = `
	<html>
	<head>
		<title>Statistic Page</title>
		<style type="text/css">
			.table-td-head {padding: 2px 5px; border: 1px solid #333;background: #f2f2f2;}
			.table-td-body {padding: 2px 5px; border: 1px solid #333;text-align: right}
			.left {text-align: left;}
		</style>
	</head>
	<body>
		Welcome to {{title .RequestType}}
		<table>
			<tr>
				<td class="table-td-head">Input Data</td>
			</tr>
			{{range $index, $data := .RequestData}}
			<tr>
				<td class="table-td-body">{{$data}}</td>
			</tr>
			{{end}}
		</table>
		<br><br>
		<table>
			<tr>
				<td colspan="2" class="table-td-head">Result</td>
			</tr>
			<tr>
				<td class="table-td-body left">Mean</td>
				<td class="table-td-body">{{mean .RequestData}}</td>
			</tr>
			<tr>
				<td class="table-td-body left">Min</td>
				<td class="table-td-body">{{min .RequestData}}</td>
			</tr>
			<tr>
				<td class="table-td-body left">Max</td>
				<td class="table-td-body">{{max .RequestData}}</td>
			</tr>
		</table>
	</body>
	</html>
	`

var statisticFunc = template.FuncMap{
	"title": func(s string) template.HTML {
		return template.HTML(s)
	},
	"mean": func(datas []int) float64 {
		result := make(chan float64)
		go func() {
			if len(datas) == 0 {
				result <- float64(0)
			}
			sum := 0
			for _, data := range datas {
				sum += data
			}
			newResult := float64(sum) / float64(len(datas))
			result <- newResult
		}()
		return <-result
	},
	"min": func(datas []int) int {
		result := make(chan int)
		go func() {
			if len(datas) > 0 {
				min := datas[0]
				for _, data := range datas {
					if min > data {
						min = data
					}
				}
				result <- min
			} else {
				result <- 0
			}
		}()
		return <-result
	},
	"max": func(datas []int) int {
		result := make(chan int)
		go func() {
			if len(datas) > 0 {
				max := datas[0]
				for _, data := range datas {
					if max < data {
						max = data
					}
				}
				result <- max
			} else {
				result <- 0
			}
		}()
		return <-result
	},
}

type request struct {
	RequestType string
	RequestData interface{}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/views", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/views", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("views").Parse(views))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/statistic", func(w http.ResponseWriter, r *http.Request) {
		var request request
		var requestData = make(chan []int)
		go func() {
			requestData <- []int{8, 5, 4, 7, 6, 9, 5, 9, 7, 10}
		}()
		request.RequestType = "Statistic Page"
		request.RequestData = <-requestData
		tmpl := template.Must(template.New("statistic-page").Funcs(statisticFunc).Parse(statisticPage))
		err := tmpl.ExecuteTemplate(w, "statistic-page", request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	server := new(http.Server)
	address := "localhost:9000"
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	logrus.Info(fmt.Sprintf("Run server at: %s\n", address))
	server.ListenAndServe()
}
