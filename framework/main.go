package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StudentResponse struct {
	Name string `json:"name"`
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)
	if r.URL.Path == "/students" {
		if r.Method == http.MethodGet {
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
			return
		}
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	h := &Handler{}
	http.ListenAndServe("localhost:8080", h)
}
