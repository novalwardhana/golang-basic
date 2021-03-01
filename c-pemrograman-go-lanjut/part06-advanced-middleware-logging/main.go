package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := echo.New()

	r.Use(middlewareOne)
	r.Use(middlewareTwo)
	r.Use(echo.WrapMiddleware(middlewareThree))
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "uri: ${uri} | method: ${method} | status: ${status} \n\n",
	}))
	r.Use(middlewareLog)

	r.HTTPErrorHandler = errorHandling

	r.GET("/", routeIndex)
	r.GET("/foods", routeFoods)

	lock := make(chan error)
	go func(lock chan error) {
		lock <- r.Start(":9000")
	}(lock)

	time.Sleep(1 * time.Millisecond)
	makeLog(nil).Info("Start API service, application run without ssl/tls")

	err := <-lock
	if err != nil {
		makeLog(nil).Errorf("Error start service")
	}
}

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("%s\n", "--- Middleware 1 ---")
		return next(c)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("%s\n", "--- Middleware 2 ---")
		return next(c)
	}
}

func middlewareThree(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s\n", "--- Middleware 3 ---")
		next.ServeHTTP(w, r)
	})
}

func middlewareLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLog(c).Info("Middleware log info")
		return next(c)
	}
}

func makeLog(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"time": time.Now().Format("2006-01-02 15:04:5"),
		})
	}
	return log.WithFields(log.Fields{
		"url":    c.Request().URL,
		"method": c.Request().Method,
		"time":   time.Now().Format("2006-01-02 15:04:05"),
	})
}

func errorHandling(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report.Message = fmt.Sprintf("%s", "An error occured")
	} else {
		report.Message = fmt.Sprintf("%d : %s", report.Code, report.Message)
	}
	makeLog(c).Errorf("Error when hit endpoint")
	c.HTML(report.Code, report.Message.(string))
}

func routeIndex(c echo.Context) error {
	output := make(chan interface{})
	go func() {
		data := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusOK,
			Message: "Welcome to my API...",
		}
		output <- data
	}()
	return c.JSON(http.StatusOK, <-output)
}

func routeFoods(c echo.Context) error {
	output := make(chan interface{})
	go func() {
		data := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    []struct {
				ID    string  `json:"id"`
				Name  string  `json:"name"`
				Price float64 `json:"price"`
			} `json:"data"`
		}{
			Code:    http.StatusOK,
			Message: "Data successfully generated",
			Data: []struct {
				ID    string  `json:"id"`
				Name  string  `json:"name"`
				Price float64 `json:"price"`
			}{
				{ID: "F0001", Name: "Garang Asem", Price: 12000},
				{ID: "F0002", Name: "Sate Ayam", Price: 15000},
				{ID: "F0003", Name: "Pecel Lele", Price: 9000},
			},
		}
		output <- data
	}()
	return c.JSON(http.StatusOK, <-output)
}
