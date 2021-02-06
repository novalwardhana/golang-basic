package main

import (
	"encoding/json"
	"net/http"
)

type customMuxFood struct {
	http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (c *customMuxFood) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	for _, middleware := range c.middlewares {
		current = middleware(current)
	}
	current.ServeHTTP(w, r)
}

func (c *customMuxFood) registerMiddleware(next func(http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func middlewareMethodFood(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "POST" {
			response := struct {
				Message string `json:"message"`
			}{
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

func middlewareAuthFood(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, status := r.BasicAuth()
		if !status {
			response := struct {
				Message string `json:"message"`
			}{
				Message: "Auth must be filled",
			}
			responseJSON, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		}
		if usernameAuth != username || passwordAuth != password {
			response := struct {
				Message string `json:"message"`
			}{
				Message: "Username or password not suitable",
			}
			responseJSON, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		}
		next.ServeHTTP(w, r)
	})
}
