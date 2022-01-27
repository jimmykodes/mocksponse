package recipe

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Recipe struct {
	Routes []*Route `yaml:"routes"`
}

func New(filename string) (*Recipe, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := &Recipe{}
	if err := yaml.NewDecoder(f).Decode(r); err != nil {
		return nil, err
	}
	return r, nil
}
