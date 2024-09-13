package routes

import (
	"net/http"

	"github.com/EmilSunden/electron-app/backend/pkg/handlers"
	"github.com/rs/cors"
)

// SetupRoutes initializes all the routes and returns the configured http.Handler
func SetupRoutes() http.Handler {
    mux := http.NewServeMux()

    // Define your routes here
    mux.HandleFunc("/api/hello", handlers.HelloHandler)

	// setup the apps handler
	mux.HandleFunc("/api/get/apps", handlers.ReadApps)

    mux.HandleFunc("/api/post/apps", handlers.CreateApp)

    // Add CORS support (or any other middleware)
    handler := cors.Default().Handler(mux)

    return handler
}
