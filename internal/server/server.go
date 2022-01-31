package server

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"sync"

	"github.com/gorilla/mux"

	"github.com/jimmykodes/mocksponse/internal/recipe"
)

func New(filename string, port int) (*server, error) {
	rec, err := recipe.New(filename)
	if err != nil {
		return nil, err
	}
	router := mux.NewRouter()
	var once sync.Once
	for _, route := range rec.Routes {
		handler, err := route.Handler(path.Dir(filename))
		if err != nil {
			return nil, err
		}
		r := router.Handle(route.Path, handler)
		if route.Fallback {
			once.Do(func() {
				router.NotFoundHandler = handler
				router.MethodNotAllowedHandler = handler
			})
		}
		if len(route.Methods) > 0 {
			r.Methods(route.Methods...)
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
