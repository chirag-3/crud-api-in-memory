package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chirag-3/crud-api-in-memory/internals/models"
	"github.com/chirag-3/crud-api-in-memory/internals/store"
	"github.com/gorilla/mux"
)

type response struct {
	Status  int `json:"status"`
	Message any `json:"data"`
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var m models.Movie

	r.Body = http.MaxBytesReader(w, r.Body, 1048576) //This will return a error response as soon as response body size becomes more than 1 MB

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	if err := dec.Decode(&m); err != nil {
		sendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//normally should be returning error and we should cacth it here and return error response, but here we are suing in memeory store,  So that is not needed
	store.AddMovie(m)
	sendResponse(w, http.StatusCreated, "Movie Created Successfully")
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	returnBody := store.ReturnMovie(id)

	if returnBody.Title == "" {
		sendResponse(w, http.StatusNotFound, "No such movie in our collection")
		return
	}

	sendResponse(w, http.StatusFound, returnBody)
}

func sendResponse(w http.ResponseWriter, statusCode int, message any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response{Status: statusCode, Message: message})
}
