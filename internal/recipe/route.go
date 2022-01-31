package recipe

import (
	"net/http"
)

type Route struct {
	Get, Post, Put, Patch, Delete, Options *Handler
}

func (r Route) Methods() map[string]*Handler {
	return map[string]*Handler{
		http.MethodGet:     r.Get,
		http.MethodPost:    r.Post,
		http.MethodPut:     r.Put,
		http.MethodPatch:   r.Patch,
		http.MethodDelete:  r.Delete,
		http.MethodOptions: r.Options,
	}
}
