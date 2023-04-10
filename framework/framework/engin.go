package framework

import (
	"net/http"

	"github.com/kuropenguin/udemy-go-web-framework/framework/controllers"
)

type Engine struct{}

func (e *Engine) Run() {
	http.ListenAndServe("localhost:8080", e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/users" {
			controllers.UsersController(w, r)
		}
		if r.URL.Path == "/listen" {
			controllers.ListenController(w, r)
		}
		if r.URL.Path == "/students" {
			controllers.GetStudent(w, r)
		}
	}
}
