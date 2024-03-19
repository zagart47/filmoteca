package handler

import (
	"context"
	"filmoteca/internal/entity"
	"filmoteca/pkg/logger"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) Movies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		w.Header().Set("Content-Type", "application/json")
		haveId := r.URL.Query().Has("id")
		id := r.URL.Query().Get("id")
		if !haveId {
			order := r.URL.Query().Get("order")
			field := r.URL.Query().Get("sort")
			search := r.URL.Query().Get("search")
			actor := r.URL.Query().Get("actor")
			title := r.URL.Query().Get("title")
			Opt := entity.Options{
				Field:  field,
				Order:  order,
				Search: search,
				Actor:  actor,
				Title:  title,
			}
			movies, err := h.Services.Movies.ReadAll(ctx, Opt)
			if err != nil {
				http.Error(w, "all movies db getting error", http.StatusInternalServerError)
				logger.Log.Error("all movies db getting error", err.Error())
			}
			jsonMovies, err := json.Marshal(movies)
			fmt.Fprintf(w, string(jsonMovies))
			return
		}
		if haveId {
			idN, err := strconv.Atoi(id)
			if idN <= 0 {
				http.Error(w, "actor bad id", http.StatusBadRequest)
				logger.Log.Error("actor bad id")
				return
			}
			movie, err := h.Services.Movies.ReadOne(ctx, id)
			if err != nil {
				http.Error(w, "one actor db getting error", http.StatusInternalServerError)
				logger.Log.Error("one actor db getting error", err.Error())
			}
			jsonMovie, err := json.Marshal(movie)
			fmt.Fprintf(w, string(jsonMovie))
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if h.CheckUser(w, r) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			w.Header().Set("Content-Type", "application/json")
			id := r.URL.Query().Get("id")
			idN, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "actor bad id", http.StatusMethodNotAllowed)
				logger.Log.Error("actor bad id", err.Error())
			}
			if idN == 0 {
				http.Error(w, "actor bad id", http.StatusInternalServerError)
				logger.Log.Error("actor bad id", err.Error())
				return
			}
			if idN > 0 {
				w.Header().Set("Content-Type", "application/json")
				movie := entity.NewMovie()
				decoder := json.NewDecoder(r.Body)
				if err := decoder.Decode(&movie); err != nil {
					ms := "JSON decoding error: "
					logger.Log.Error(ms, err.Error())
					http.Error(w, ms, http.StatusInternalServerError)
				}
				err := h.Services.Movies.Create(ctx, movie, id)
				if err != nil {
					http.Error(w, "movie db creating error", http.StatusInternalServerError)
					logger.Log.Error("movie db creating error", err.Error())
				}
				actor, err := h.Services.Actors.ReadOne(ctx, id)
				if err != nil {
					http.Error(w, "one actor getting error", http.StatusInternalServerError)
					logger.Log.Error("one actor getting error", err.Error())
				}
				jsonActor, err := json.Marshal(actor)
				if err != nil {
					http.Error(w, "one actor marshalling error", http.StatusInternalServerError)
					logger.Log.Error("one actor marshalling error", err.Error())
				}
				fmt.Fprintf(w, string(jsonActor))
				return
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) EditMovie(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		if h.CheckUser(w, r) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			w.Header().Set("Content-Type", "application/json")
			id := r.URL.Query().Get("id")
			idN, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "bad id", http.StatusMethodNotAllowed)
				logger.Log.Error("movie bad id", err.Error())
			}
			if idN == 0 {
				http.Error(w, "movie bad id", http.StatusInternalServerError)
				logger.Log.Error("movie bad id", err.Error())
				return
			}
			if idN > 0 {
				w.Header().Set("Content-Type", "application/json")
				movie := entity.NewMovie()
				decoder := json.NewDecoder(r.Body)
				if err := decoder.Decode(&movie); err != nil {
					ms := "JSON decoding error: "
					logger.Log.Error(ms, err.Error())
					http.Error(w, ms, http.StatusInternalServerError)
				}
				movie, err := h.Services.Movies.Update(ctx, id, movie)
				if err != nil {
					http.Error(w, "one movie db getting error", http.StatusInternalServerError)
					logger.Log.Error("one movie db getting error", err.Error())
				}
				jsonActor, err := json.Marshal(movie)
				fmt.Fprintf(w, string(jsonActor))
				return
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
