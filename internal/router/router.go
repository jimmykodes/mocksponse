package router

import (
	"net/http"
	"strings"
)

type Route struct {
	Value     string
	Handlers  map[string]http.Handler
	subroutes map[string]*Route
}

type Router struct {
	Top             *Route
	NotFoundHandler http.Handler
}

func newRoute(value string) *Route {
	return &Route{
		Value:     value,
		Handlers:  make(map[string]http.Handler),
		subroutes: make(map[string]*Route),
	}
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	route := ro.Top
	for _, part := range parts {
		route = route.subroutes[part]
		if route == nil {
			ro.NotFoundHandler.ServeHTTP(w, r)
			return
		}
	}
	if handler := route.Handlers[r.Method]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		ro.NotFoundHandler.ServeHTTP(w, r)
	}
}

func (ro *Router) Register(pattern string, method string, handler http.Handler) {
	parts := strings.Split(pattern, "/")
	last := ro.Top
	for _, part := range parts {
		route, ok := last.subroutes[part]
		if !ok {
			route = newRoute(part)
			last.subroutes[part] = route
		}
		last = route
	}
	last.Handlers[strings.ToUpper(method)] = handler
}

func New() *Router {
	return &Router{
		NotFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("404: not found"))
		}),
		Top: newRoute("/"),
	}
}
