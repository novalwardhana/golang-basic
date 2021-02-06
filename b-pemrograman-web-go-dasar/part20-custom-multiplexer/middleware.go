package main

import (
	"encoding/json"
	"net/http"
)

type customMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *customMux) registerMiddleware(next func(http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}

func middlewareMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			response := struct {
				Code    uint   `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusMethodNotAllowed,
				Message: "Method not allowed",
			}
			responseJSON, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		}
		next.ServeHTTP(w, r)
	})
}
