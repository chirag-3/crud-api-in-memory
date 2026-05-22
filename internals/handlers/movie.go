package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chirag-3/crud-api-in-memory/internals/models"
	"github.com/chirag-3/crud-api-in-memory/internals/store"
	"github.com/gorilla/mux"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var m models.Movie
	json.NewDecoder(r.Body).Decode(&m)
	store.AddMovie(m)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success")
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	returnBody := store.ReturnMovie(id)
	json.NewEncoder(w).Encode(returnBody)
	w.WriteHeader(http.StatusOK)
}
