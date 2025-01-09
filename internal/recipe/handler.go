package recipe

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strings"
	"sync/atomic"
)

type Handler struct {
	Sequential bool        `yaml:"sequential"`
	Responses  []*Response `yaml:"responses"`
	index      atomic.Int64
	single     bool
}

func (m *Handler) Handler(fp string) (http.Handler, error) {
	for _, res := range m.Responses {
		if err := res.init(path.Dir(fp)); err != nil {
			return nil, err
		}
	}
	if len(m.Responses) == 1 {
		m.single = true
	} else if m.Sequential {
		// initialize at -1 so the first Inc call sets it to 0
		m.index.Store(-1)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sb strings.Builder
		fmt.Fprintf(&sb, "%7s\t%s\t%s", r.Method, r.URL.Path, r.URL.RawQuery)
		_, _ = io.Copy(&sb, r.Body)
		log.Println(sb.String())
		var index int64
		if m.single {
			index = 0
		} else if m.Sequential {
			index = m.index.Add(1)
			// make sure we are inside the bounds of the responses but don't touch the atomic counter
			index %= int64(len(m.Responses))
		} else {
			index = int64(rand.Intn(len(m.Responses)))
		}

		resp := m.Responses[index]
		if resp.Code != 0 {
			w.WriteHeader(resp.Code)
		}
		if err := resp.write(w, r); err != nil {
			log.Println("error writing data", err)
		}
	}), nil
}
