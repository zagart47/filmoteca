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

func (h *Handler) NewActor(w http.ResponseWriter, r *http.Request) { // POST
	switch r.Method {
	case "POST":
		if h.CheckUser(w, r) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			w.Header().Set("Content-Type", "application/json")
			actor := entity.NewActor()
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&actor); err != nil {
				ms := "JSON decoding error: "
				logger.Log.Error(ms, err.Error())
				http.Error(w, ms, http.StatusInternalServerError)
			}
			if err := h.Services.Actors.Create(ctx, actor); err != nil {
				logger.Log.Error("db creating actor error:", err.Error())
				http.Error(w, "actor creating error", http.StatusInternalServerError)
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (h *Handler) Actors(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("actors getting started")
	switch r.Method {
	case "GET":
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		w.Header().Set("Content-Type", "application/json")
		haveId := r.URL.Query().Has("id")
		id := r.URL.Query().Get("id")
		if !haveId {
			actors, err := h.Services.Actors.ReadAll(ctx)
			if err != nil {
				http.Error(w, "all actors db getting error", http.StatusInternalServerError)
				logger.Log.Error("all actors db getting error", err.Error())
			}
			jsonActors, err := json.Marshal(actors)
			fmt.Fprintf(w, string(jsonActors))
			return
		}
		if haveId {
			idN, err := strconv.Atoi(id)
			if idN <= 0 {
				http.Error(w, "actor bad id", http.StatusBadRequest)
				logger.Log.Error("actor bad id")
				return
			}
			actor, err := h.Services.Actors.ReadOne(ctx, id)
			if err != nil {
				http.Error(w, "one actor db getting error", http.StatusInternalServerError)
				logger.Log.Error("one actor db getting error", err.Error())
			}
			jsonActor, err := json.Marshal(actor)
			fmt.Fprintf(w, string(jsonActor))
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (h *Handler) EditActor(w http.ResponseWriter, r *http.Request) {
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
				logger.Log.Error("actor bad id", err.Error())
			}
			if idN == 0 {
				http.Error(w, "actor bad id", http.StatusInternalServerError)
				logger.Log.Error("actor bad id", err.Error())
				return
			}
			if idN > 0 {
				w.Header().Set("Content-Type", "application/json")
				actor := entity.NewActor()
				decoder := json.NewDecoder(r.Body)
				if err := decoder.Decode(&actor); err != nil {
					ms := "JSON decoding error: "
					logger.Log.Error(ms, err.Error())
					http.Error(w, ms, http.StatusInternalServerError)
				}
				actor, err := h.Services.Actors.Update(ctx, id, actor)
				if err != nil {
					http.Error(w, "one actor db getting error", http.StatusInternalServerError)
					logger.Log.Error("one actor db getting error", err.Error())
				}
				jsonActor, err := json.Marshal(actor)
				fmt.Fprintf(w, string(jsonActor))
				return
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		if h.CheckUser(w, r) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			w.Header().Set("Content-Type", "application/json")
			haveId := r.URL.Query().Has("id")
			options := make(map[string]bool)
			name := r.URL.Query().Has("name")
			options["name"] = name
			gender := r.URL.Query().Has("gender")
			options["gender"] = gender
			birthdate := r.URL.Query().Has("birthdate")
			options["birthdate"] = birthdate
			var o []string
			for k, v := range options {
				if v {
					o = append(o, k)
				}
			}
			if !haveId {
				http.Error(w, "actor bad id", http.StatusInternalServerError)
				logger.Log.Error("actor bad id")
				return
			}
			if haveId {
				id := r.URL.Query().Get("id")
				_, err := h.Services.Actors.Delete(ctx, id, o)
				if err != nil {
					http.Error(w, "deleting error", http.StatusInternalServerError)
					logger.Log.Error("deleting error", err.Error())
				}
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
