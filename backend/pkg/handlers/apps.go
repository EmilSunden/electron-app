package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EmilSunden/electron-app/backend/pkg/models"
	"github.com/EmilSunden/electron-app/backend/pkg/services"
)

type Handler struct {
	AppService *services.AppService
}

func NewHandler(appService *services.AppService) *Handler {
	return &Handler{
		AppService: appService,
	}
}

func (h *Handler) ReadApps(w http.ResponseWriter, r *http.Request) {
	apps, err := h.AppService.GetApps()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get apps: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apps)
}

func (h *Handler) CreateApp(w http.ResponseWriter, r *http.Request) {
	var newApp models.App
	if err := json.NewDecoder(r.Body).Decode(&newApp); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	err := h.AppService.CreateApp(&newApp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create app: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newApp)
}
