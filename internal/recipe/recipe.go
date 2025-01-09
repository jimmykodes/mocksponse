package recipe

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Recipe struct {
	Routes  map[string]*Route `yaml:"routes"`
	Default *Handler          `yaml:"default"`
}

func New(filename string) (*Recipe, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var r Recipe
	if err := yaml.NewDecoder(f).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}
