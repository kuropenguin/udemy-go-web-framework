package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func GetStudent(w http.ResponseWriter, r *http.Request) {

	queries := r.URL.Query()
	name := queries.Get("name")
	studentResponse := StudentResponse{
		Name: name,
	}

	responseData, err := json.Marshal(studentResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func ListenController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ListenController")
}

func UsersController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UsersController")
}
