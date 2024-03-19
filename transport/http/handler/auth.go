package handler

import (
	"context"
	"encoding/base64"
	"filmoteca/pkg/logger"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) CheckUser(w http.ResponseWriter, r *http.Request) bool {
	// Basic Auth checker
	encodedData := r.Header.Get("Authorization")
	parts := strings.SplitN(encodedData, " ", 2)
	if len(parts) != 2 || parts[0] != "Basic" {
		logger.Log.Error("Invalid encoded data")
		return false
	}
	decoded, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		logger.Log.Error("Error decoding data:", err)
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	users, err := h.Services.Users.Get(ctx)
	if err != nil {
		logger.Log.Error("can't get admins from db")
	}
	username := string(decoded[:len(decoded)-1])
	for _, v := range users {
		if username == v.Name {
			return true
		}
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return false
}
