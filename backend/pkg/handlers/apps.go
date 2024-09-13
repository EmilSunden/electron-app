package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EmilSunden/electron-app/backend/pkg/models"
	"github.com/EmilSunden/electron-app/backend/pkg/services"
)

// Create a simple handler that returns hello world
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

var appService = services.NewAppService("../../mockdb/apps.json")

func ReadApps(w http.ResponseWriter, r *http.Request) {
	apps := appService.GetApps()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apps)
}

func CreateApp(w http.ResponseWriter, r *http.Request) {
	var newApp models.App
	if err := json.NewDecoder(r.Body).Decode(&newApp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	appService.CreateApp(newApp)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newApp)
}
