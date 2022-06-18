package handlers

import (
	"fmt"
	"net/http"
)

const (
	content = "Content-type"
	yamlapp = "application/x-yaml"
	jsonapp = "application/json"
	html    = "text/html; charset=utf-8"
	plain   = "text/plain"
)

func HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		greet := "Hello world"
		w.Header().Set(content, html)
		w.Write([]byte(greet))
	}
}

func HandleAdmin() http.HandlerFunc {
	// Here you can prepare a 'private thing' to use in the
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, plain)
		fmt.Fprintf(w, "Hello Admin!")
	}
}
