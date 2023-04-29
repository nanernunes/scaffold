package scaffold

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Internal *mux.Router
}

func NewRouter(router *mux.Router) *Router {
	r := router
	if r == nil {
		r = mux.NewRouter()
	}

	return &Router{Internal: r}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Internal.ServeHTTP(w, req)
}

func (r *Router) Resource(route string, controller Controller) {
	uuidRoute := fmt.Sprintf("%s/{uuid}", route)

	r.GET(route, controller.Index)
	r.POST(route, controller.Create)
	r.GET(uuidRoute, controller.Show)
	r.PATCH(uuidRoute, controller.Update)
	r.DELETE(uuidRoute, controller.Delete)
}

func (r *Router) OPTIONS(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodOptions)
}

func (r *Router) CONNECT(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodConnect)
}

func (r *Router) HEAD(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodHead, http.MethodOptions)
}

func (r *Router) TRACE(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodTrace)
}

func (r *Router) GET(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodGet, http.MethodOptions)
}

func (r *Router) POST(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodPost, http.MethodOptions)
}

func (r *Router) PATCH(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodPatch, http.MethodOptions)
}

func (r *Router) PUT(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodPut, http.MethodOptions)
}

func (r *Router) DELETE(route string, callback func(context *Context)) {
	r.Internal.HandleFunc(route, ContextAdapter(callback)).Methods(http.MethodDelete, http.MethodOptions)
}
