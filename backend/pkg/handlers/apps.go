package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/EmilSunden/electron-app/backend/pkg/services"
	"github.com/EmilSunden/electron-app/backend/pkg/models"
)

// Create a simple handler that returns hello world
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func AppsHandler(w http.ResponseWriter, r *http.Request) {
	appService := services.NewAppService("../../mockdb/apps.json")

	switch r.Method {
	case http.MethodGet:
		apps := appService.GetApps()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apps)
		
	case http.MethodPost:
		var newApp models.App 
		if err := json.NewDecoder(r.Body).Decode(&newApp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		appService.CreateApp(newApp)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newApp)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}