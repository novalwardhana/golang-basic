package main

import (
	"encoding/json"
	"net/http"
)

func authStudent(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "Username and password not found",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return false
	}

	if username != usernameAuth || password != passwordAuth {
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusNonAuthoritativeInfo,
			Message: "Username or password not suitable",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return false
	}
	return true
}
