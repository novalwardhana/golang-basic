package main

import "net/http"

func middlewareMethodStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Method not allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func middlewareAuthStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("username or password must field in authenication"))
			return
		}
		if username != usernameAuth || password != passwordAuth {
			w.Write([]byte("Username or password not suitable"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
