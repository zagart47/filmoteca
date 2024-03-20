package handler

import (
	"filmoteca/internal/service"
	"net/http"
)

type Handler struct {
	Services service.Services
}

func NewHandler(s service.Services) Handler {
	return Handler{
		Services: s,
	}
}

func (h *Handler) Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/actors", h.Actors)
	mux.HandleFunc("/actors/new", h.NewActor)
	mux.HandleFunc("/actors/edit", h.EditActor)
	mux.HandleFunc("/actors/delete", h.DeleteActor)
	mux.HandleFunc("/movies", h.Movies)
	mux.HandleFunc("/movies/new", h.CreateMovie)
	mux.HandleFunc("/movies/edit", h.EditMovie)
	mux.HandleFunc("/movies/delete", h.EditMovie)
	return mux
}
