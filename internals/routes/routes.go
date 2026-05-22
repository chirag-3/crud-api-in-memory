package routes

import (
	"github.com/chirag-3/crud-api-in-memory/internals/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/movie", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", handlers.GetMovie).Methods("GET")
}
