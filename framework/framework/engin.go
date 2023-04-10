package framework

import (
	"fmt"
	"net/http"

	"github.com/kuropenguin/udemy-go-web-framework/framework/controllers"
)

type Engine struct{}

func (e *Engine) Run() {
	http.ListenAndServe("localhost:8080", e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/students" {
		if r.Method == http.MethodGet {
			controllers.GetStudent(w, r)
		}
	}
	fmt.Fprintf(w, "hello")
}
