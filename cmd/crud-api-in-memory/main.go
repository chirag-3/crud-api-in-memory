package main

import (
	"fmt"
	"net/http"

	"github.com/chirag-3/crud-api-in-memory/internals/routes"
	"github.com/chirag-3/crud-api-in-memory/internals/store"
	"github.com/gorilla/mux"
)

var r *mux.Router

func main() {
	fmt.Println("Hello world")
	r = mux.NewRouter()

	store.InitializeStore()
	routes.RegisterRoutes(r)

	if err := http.ListenAndServe(":8090", r); err != nil {
		fmt.Println("Server not started", err)
	}

}
