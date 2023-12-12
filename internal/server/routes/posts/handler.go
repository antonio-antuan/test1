// Package users is the HTTP handler for the users.
package posts

import (
	"github.com/go-chi/chi"
)

type handler struct {
	proc   Processor
	logger Logger
}

func NewHandler(logger Logger, proc Processor) *chi.Mux {
	h := &handler{
		proc:   proc,
		logger: logger,
	}
	router := chi.NewRouter()
	router.Get("/", h.getPosts)
	router.Options("/", h.options)
	return router
}
