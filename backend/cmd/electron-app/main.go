package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EmilSunden/electron-app/backend/internal/database"
	"github.com/EmilSunden/electron-app/backend/pkg/routes"
)

func main() {
	// Initialize the db connection
	database, err := database.InitDB();
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	defer database.Close()

	handler := routes.SetupRoutes();

	fmt.Println("Go server is running on port 8080")
	http.ListenAndServe(":8080", handler)
}