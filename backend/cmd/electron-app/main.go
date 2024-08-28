package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/rs/cors"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		// Create an instance of the Response struct
		response := Response{
			Message: "Hello, World!",
		}

		// Encode the struct to JSON and write it to the response
		json.NewEncoder(w).Encode(response)
	})

	handler := cors.Default().Handler(mux)

	fmt.Println("Go server is running on port 8080")
	http.ListenAndServe(":8080", handler)
}