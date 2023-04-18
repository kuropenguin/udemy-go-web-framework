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
		targetNode := e.Router.routingTable.Search(pathname)

		if targetNode == nil || targetNode.handler == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		node := targetNode

		paramArr := strings.Split(pathname, "/")
		paramDicts := make(map[string]string)
		for i := len(paramArr) - 1; i >= 0; i-- {
			if isGeneral(node.param) {
				paramDicts[node.param] = paramArr[i]
			}
			node = node.parent
		}

		targetNode.handler(ctx)
		return
	}

}
