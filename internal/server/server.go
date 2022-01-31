package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jimmykodes/mocksponse/internal/recipe"
)

func New(filename string, port int) (*server, error) {
	rec, err := recipe.New(filename)
	if err != nil {
		return nil, err
	}
	router := mux.NewRouter()
	if rec.Default != nil {
		router.NotFoundHandler, err = rec.Default.Handler(filename)
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
			router.Handle(routePath, handler).Methods(methodString)
		}
	}

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	return &server{
		svr: svr,
	}, nil
}

type server struct {
	svr *http.Server
}

func (s server) Run() error {
	log.Printf("running at %s\n", s.svr.Addr)
	return s.svr.ListenAndServe()
}
