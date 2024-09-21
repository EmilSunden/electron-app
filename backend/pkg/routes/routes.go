package routes

import (
	"net/http"

	"github.com/EmilSunden/electron-app/backend/pkg/handlers"
	"github.com/rs/cors"
)

// SetupRoutes initializes all the routes and returns the configured http.Handler
func SetupRoutes(h *handlers.Handler) http.Handler {
	mux := http.NewServeMux()

	// Define your routes here

	// setup the apps handler
	mux.HandleFunc("/api/get/apps", h.ReadApps)

	mux.HandleFunc("/api/post/apps", h.CreateApp)

	// Add CORS support (or any other middleware)
	handler := cors.Default().Handler(mux)

	return handler
}
