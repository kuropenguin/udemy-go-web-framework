package framework

import (
	"errors"
	"net/http"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{},
	}
}

type Router struct {
	routingTable map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) error {
	if r.routingTable == nil {
		r.routingTable = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	if r.routingTable[path] != nil {
		return errors.New("exist path")
	}
	r.routingTable[path] = handler
	return nil
}

func (e *Engine) Run() {
	http.ListenAndServe("localhost:8080", e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handler := e.Router.routingTable[r.URL.Path]
		if handler == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		handler(w, r)
		return
	}

}
