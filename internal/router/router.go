package router

import (
	"net/http"
	"strings"
)

type Route struct {
	Value     string
	Handlers  map[string]http.Handler
	subroutes map[string]*Route
	wild      *Route
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

	last := ro.Top
	for _, part := range parts {
		route := last.subroutes[part]
		if route == nil {
			route = last.wild
		}
		if route == nil {
			ro.NotFoundHandler.ServeHTTP(w, r)
			return
		}
		last = route
	}
	if handler := last.Handlers[r.Method]; handler != nil {
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
			if len(part) > 0 && part[0] == ':' {
				// is wildcard
				if last.wild == nil {
					route = newRoute(part)
					last.wild = route
				} else {
					if last.wild.Value == part {
						route = last.wild
					} else {
						// TODO: bubble up the error correctly
						panic("multiple wild routes")
					}
				}
			} else {
				route = newRoute(part)
				last.subroutes[part] = route
			}
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
