package framework

import (
	"net/http"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{
			routingTable: Constructor(),
		},
	}
}

type Router struct {
	routingTable TreeNode
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) error {
	r.routingTable.Insert(path, handler)
	return nil
}

func (e *Engine) Run() {
	http.ListenAndServe("localhost:8080", e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pathname := r.URL.Path
		handler := e.Router.routingTable.Search(pathname)
		if handler == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		handler(w, r)
		return
	}

}
