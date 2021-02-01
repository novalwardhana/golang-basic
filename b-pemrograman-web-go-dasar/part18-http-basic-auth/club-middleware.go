package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

func authClubBasic(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "Username or password not found",
		}
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return false
	}
	if username != usernameAuth || password != passwordAuth {
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "Username or password not suitable",
		}
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return false
	}
	return true
}

func authClubAPIKey(w http.ResponseWriter, r *http.Request) bool {
	//case success -> bm92YWw6bm92YWxub3ZhbA==
	//case slice cannot parsed -> bmlkYXVsZmlhaHVzbmE=
	//case username & password not suitable -> bm92YWw6bmlkYQ==

	token := r.Header["Authorization"][0]
	tokenSlc := strings.Split(token, " ")
	if len(tokenSlc) < 2 {
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Error parsing auth data",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return false
	}

	tokenDecode, err := base64.StdEncoding.DecodeString(tokenSlc[1])
	if err != nil {
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Error decode auth data",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return false
	}
	tokenDecodeStr := string(tokenDecode)
	tokenDecodeSlc := strings.Split(tokenDecodeStr, ":")
	if len(tokenDecodeSlc) < 2 {
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Slice cannot be parsed",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
		return false
	}
	username := tokenDecodeSlc[0]
	password := tokenDecodeSlc[1]
	if username != usernameAuth || password != passwordAuth {
		response := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusUnauthorized,
			Message: "Username and password not suitable",
		}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-type", "application/json")
		w.Write(responseJSON)
		return false
	}

	return true
}
