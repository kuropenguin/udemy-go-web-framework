package framework

import (
	"net/http"
	"strings"
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

func (r *Router) Get(pathname string, handler func(ctx *MyContext)) error {
	pathname = strings.TrimSuffix(pathname, "/")
	existedHandler := r.routingTable.Search(pathname)
	if existedHandler != nil {
		panic("path already existed")
	}
	r.routingTable.Insert(pathname, handler)
	return nil
}

func (e *Engine) Run() {
	http.ListenAndServe("localhost:8080", e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewMyContext(w, r)
	if r.Method == http.MethodGet {
		pathname := r.URL.Path
		handler := e.Router.routingTable.Search(pathname)
		if handler == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		handler(ctx)
		return
	}

}
