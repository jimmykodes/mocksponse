package recipe

import (
	"log"
	"math/rand"
	"net/http"
	"sync"

	"go.uber.org/atomic"
)

type Route struct {
	Path       string      `yaml:"path"`
	Sequential bool        `yaml:"sequential"`
	Responses  []*Response `yaml:"responses"`
	index      *atomic.Int64
	once       sync.Once
}

func (route *Route) Handler() (http.Handler, error) {
	// initialize at -1 so the first Inc call sets it to 0
	route.index = atomic.NewInt64(-1)
	for _, res := range route.Responses {
		if err := res.init(); err != nil {
			return nil, err
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("serving route", r.URL.Path)
		var index int64
		if route.Sequential {
			index = route.index.Inc()
		} else {
			index = int64(rand.Intn(len(route.Responses)))
		}

		// make sure we are inside the bounds of the responses
		index = index % int64(len(route.Responses))
		resp := route.Responses[index]
		if resp.Code != 0 {
			w.WriteHeader(resp.Code)
		}
		if err := resp.write(r, w); err != nil {
			log.Println("error writing data", err)
		}
	}), nil
}
