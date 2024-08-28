package routes

import (
	"electron-app/backend/pkg/handlers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hello", handlers.HelloHandler)

	return mux
}