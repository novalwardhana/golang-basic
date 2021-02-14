package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type mapData map[string]interface{}

func main() {
	r := echo.New()
	r.GET("/", func(ctx echo.Context) error {
		data := "Welcome to my api"
		return ctx.String(http.StatusOK, data)
	})

	/* Response data */
	r.GET("/response-string", func(ctx echo.Context) error {
		data := "Welcome to my API - String"
		return ctx.String(http.StatusOK, data)
	})
	r.GET("/response-html", func(ctx echo.Context) error {
		data := "Welcome to my API - HTML"
		return ctx.HTML(http.StatusOK, data)
	})
	r.GET("/response-redirect", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})
	r.GET("/response-json", func(ctx echo.Context) error {
		data := mapData{"message": "Welcome to my API JSON", "code": http.StatusOK}
		return ctx.JSON(http.StatusOK, data)
	})
	/* End response data */

	/* Parsing data */
	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		age := ctx.QueryParam("age")
		log.Println("access: http://localhost:9000?", name, "&age=", age)
		data := mapData{"name": name, "age": age}
		return ctx.JSON(http.StatusOK, data)
	})
	r.GET("/page2/:name/:age", func(ctx echo.Context) error {
		name := ctx.Param("name")
		age := ctx.Param("age")
		log.Println("access: http://localhost:9000/", name, "/", age)
		data := mapData{"name": name, "age": age}
		return ctx.JSON(http.StatusOK, data)
	})
	r.GET("/page3/:name/:age/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		age := ctx.Param("age")
		message := ctx.Param("*")
		log.Println("access: http://localhost:9000/", name, "/", age, "/", message)
		data := mapData{"name": name, "age": age, "message": message}
		return ctx.JSON(http.StatusOK, data)
	})
	r.GET("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		age := ctx.FormValue("age")
		data := mapData{"name": name, "age": age}
		return ctx.JSON(http.StatusOK, data)
	})
	r.GET("/page5/:name/:age/*", func(ctx echo.Context) error {
		type dataForm struct {
			Name    string `json:"name"`
			Age     string `json:"age"`
			Message string `json:"message"`
		}

		var body dataForm
		if err := ctx.Bind(&body); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		data := struct {
			DataQuery struct {
				Name    string `json:"name"`
				Age     string `json:"age"`
				Message string `json:"message"`
			} `json:"data_query"`
			DataPath struct {
				Name    string `json:"name"`
				Age     string `json:"age"`
				Message string `json:"message"`
			} `json:"data_path"`
			DataForm struct {
				Name    string `json:"name"`
				Age     string `json:"age"`
				Message string `json:"message"`
			} `json:"data_form"`
		}{}
		data.DataQuery.Name = ctx.QueryParam("name")
		data.DataQuery.Age = ctx.QueryParam("age")
		data.DataQuery.Message = ctx.QueryParam("message")
		data.DataPath.Name = ctx.Param("name")
		data.DataPath.Age = ctx.Param("age")
		data.DataPath.Message = ctx.Param("*")
		data.DataForm.Name = body.Name
		data.DataForm.Age = body.Age
		data.DataForm.Message = body.Message
		return ctx.JSON(http.StatusOK, data)
	})
	/* End parsing data */

	/* Echo wrap handler */
	r.GET("/index", echo.WrapHandler(http.HandlerFunc(routeIndex)))
	r.GET("/home", echo.WrapHandler(routeHome))
	r.GET("/about", routeAbout)
	r.GET("/data/:name/:age/*", routeData)
	/* End Echo warp handler*/

	r.Start(":9000")
}

var routeIndex = func(w http.ResponseWriter, r *http.Request) {
	output := make(chan []byte)
	go func() {
		output <- []byte("Welcome to index page")
	}()
	w.Write(<-output)
}

var routeHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		output := make(chan []byte)
		go func() {
			output <- []byte("Welcome to home page")
		}()
		w.Write(<-output)
	},
)

var routeAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			output := make(chan []byte)
			go func() {
				output <- []byte("Welcome to about page")
			}()
			w.Write(<-output)
		},
	),
)

var routeData = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			data := struct {
				DataQuery struct {
					Name    string `json:"name"`
					Age     string `json:"age"`
					Message string `json:"message"`
				} `json:"data_query"`
				DataPath struct {
					Name    string `json:"name"`
					Age     string `json:"age"`
					Message string `json:"message"`
				} `json:"data_path"`
				DataForm struct {
					Name    string `json:"name"`
					Age     string `json:"age"`
					Message string `json:"message"`
				} `json:"data_form"`
			}{}
			urlPath := r.URL.Path
			urlPathSlice := strings.Split(urlPath, "/")

			data.DataQuery.Name = r.FormValue("name")
			data.DataQuery.Age = r.FormValue("age")
			data.DataQuery.Message = r.FormValue("message")
			data.DataPath.Name = urlPathSlice[2]
			data.DataPath.Age = urlPathSlice[3]
			data.DataPath.Message = urlPathSlice[4]
			data.DataForm.Name = r.MultipartForm.Value["name"][0]
			data.DataForm.Age = r.MultipartForm.Value["age"][0]
			data.DataForm.Message = r.MultipartForm.Value["message"][0]
			dataJSON, err := json.Marshal(data)
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(dataJSON)
		},
	),
)
