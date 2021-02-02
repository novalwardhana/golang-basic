package main

import (
	"encoding/json"
	"net/http"
)

func middlewareAuthFood(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, status := r.BasicAuth()
		if !status {
			response := struct {
				Code    uint   `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusBadRequest,
				Message: "Bad request",
			}
			responseJSON, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		}

		if username != usernameAuth || password != passwordAuth {
			response := struct {
				Code    uint   `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusUnauthorized,
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

func middlewareMethodFood(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
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
