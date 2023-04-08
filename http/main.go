package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	h := &Handler{}
	http.ListenAndServe("localhost:8080", h)
}
