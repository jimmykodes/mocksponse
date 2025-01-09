package server

import (
	"fmt"
	"net/http"

	"github.com/jimmykodes/mocksponse/internal/recipe"
	"github.com/jimmykodes/mocksponse/internal/router"
)

func New(filename string, port int) (*http.Server, error) {
	mux := router.New()
	rec, err := recipe.New(filename)
	if err != nil {
		return nil, err
	}
	if rec.Default != nil {
		mux.NotFoundHandler, err = rec.Default.Handler(filename)
		if err != nil {
			return nil, err
		}
	}
	for routePath, route := range rec.Routes {
		for methodString, method := range route.Methods() {
			if method == nil {
				continue
			}
			handler, err := method.Handler(filename)
			if err != nil {
				return nil, err
			}
			mux.Register(routePath, methodString, handler)
		}
	}
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	return svr, nil
}
