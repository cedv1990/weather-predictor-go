package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func new() Server {
	a := &api{}

	r := mux.NewRouter()

	r.HandleFunc("/generar-prediccion", a.fetchGophers).Methods(http.MethodGet)
	r.HandleFunc("/clima", a.fetchGopher).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func main() {
	server := new()
	log.Fatal(http.ListenAndServe(":1234", server.Router()))
}
