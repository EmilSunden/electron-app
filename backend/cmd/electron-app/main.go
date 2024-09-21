package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EmilSunden/electron-app/backend/internal/database"
	"github.com/EmilSunden/electron-app/backend/pkg/handlers"
	"github.com/EmilSunden/electron-app/backend/pkg/routes"
	"github.com/EmilSunden/electron-app/backend/pkg/services"
)

func main() {
	// Initialize the db connection
	database, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer database.Close()

	// Initialize the AppService
	appService := services.NewAppService(database)

	// Initialize the Handler with the AppService
	h := handlers.NewHandler(appService)

	// Set up routes with the Handler
	handler := routes.SetupRoutes(h)

	fmt.Println("Go server is running on port 8080")
	http.ListenAndServe(":8080", handler)

}
