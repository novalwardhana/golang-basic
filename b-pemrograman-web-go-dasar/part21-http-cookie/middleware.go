package main

import "net/http"

type customMux struct {
	http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (c *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	for _, middleware := range c.middlewares {
		current = middleware(current)
	}
	current.ServeHTTP(w, r)
}
