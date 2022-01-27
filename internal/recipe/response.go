package recipe

import (
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Response struct {
	Code int    `yaml:"code"`
	Data string `yaml:"data"`
	File string `yaml:"file"`
	tmpl *template.Template
}

func (resp *Response) init(fp string) error {
	var err error
	if resp.File != "" {
		var data []byte
		data, err = os.ReadFile(filepath.Join(fp, resp.File))
		if err != nil {
			return err
		}
		resp.tmpl, err = template.New("resp").Parse(string(data))
	} else {
		resp.tmpl, err = template.New("resp").Parse(resp.Data)
	}
	return err
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
