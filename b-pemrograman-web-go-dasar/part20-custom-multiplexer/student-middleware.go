package main

import (
	"encoding/json"
	"net/http"
)

type customMuxStudent struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *customMuxStudent) registerMiddleware(next func(http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *customMuxStudent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}

func middlewareMethodStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "POST" {
			response := struct {
				Code    uint   `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusMethodNotAllowed,
				Message: "Method not allowed",
			}
			responseJSON, _ := json.Marshal(response)
			w.Header().Set("Content-type", "application/json")
			w.Write(responseJSON)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func middlewareAuthStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, status := r.BasicAuth()
		if !status {
			response := struct {
				Message string `json:"message"`
			}{
				Message: "Status unauthorized",
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
