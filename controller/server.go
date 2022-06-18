package controller

import (
	"log"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() {
	mux := &Server{http.NewServeMux()}
	mux.routes()
	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux.router,
	}
	log.Printf("Listening on port %v\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
