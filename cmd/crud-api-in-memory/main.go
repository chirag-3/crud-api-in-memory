package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")
	r := mux.NewRouter()
	r.HandleFunc("/", standardFunc).Methods("POST")
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println("Not able to start server", err)
	}
}

func standardFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a standard path")
}
