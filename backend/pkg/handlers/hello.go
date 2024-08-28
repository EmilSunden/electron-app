package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Create an instance of the Response struct
	response := Response{
		Message: "Hello, World!",
	}

	// Encode the struct to JSON and write it to the response
	json.NewEncoder(w).Encode(response)
}