package recipe

import (
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/jimmykodes/mocksponse/internal/router"
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

func (resp *Response) write(w http.ResponseWriter, r *http.Request) error {
	return resp.tmpl.Execute(w, struct {
		Vars   map[string]string
		Params url.Values
	}{
		Vars:   router.Vars(r),
		Params: r.URL.Query(),
	})
}
