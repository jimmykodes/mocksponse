package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jimmykodes/mock-sponse/internal/recipe"
)

func New(filename string, port int) (*server, error) {
	rec, err := recipe.New(filename)
	if err != nil {
		return nil, err
	}
	router := mux.NewRouter()

	for _, route := range rec.Routes {
		handler, err := route.Handler()
		if err != nil {
			return nil, err
		}
		router.Handle(route.Path, handler)
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
	log.Println("running")
	return s.svr.ListenAndServe()
}
