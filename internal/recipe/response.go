package recipe

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Response struct {
	Code int    `yaml:"code"`
	Data string `yaml:"data"`
	tmpl *template.Template
}

func (resp *Response) init() (err error) {
	resp.tmpl, err = template.New("resp").Parse(resp.Data)
	return
}

func (resp *Response) write(r *http.Request, w http.ResponseWriter) error {
	return resp.tmpl.Execute(w, struct {
		Vars   map[string]string
		Params url.Values
	}{
		Vars:   mux.Vars(r),
		Params: r.URL.Query(),
	})
}
