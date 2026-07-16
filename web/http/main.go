package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)


type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"` // omitempty hides the field if it is blank
	Data    any `json:"data,omitempty"`    // interface{} (or any) allows any data structure
}

// User represents a sample data payload.
type User struct {
	ID         int    `json:"id"`
	Name  	   string `json:"name"`
	Age        int 	 `json:"age"`
	Occupation string  `json:"occupation"`
	Email      string `json:"email"`
}


func userHandler(w http.ResponseWriter, r *http.Request) {	// 1. Set the response header to JSON

	w.Header().Set("Content-Type", "application/json")

	mockUser := User{
		ID:    4,
		Name:  "Keith Thomson",
		Age:   38,
		Occupation: "Cloud Software and Product Engineer",
		Email: "ceo@whalelogic.com",
	}

	apiResponse := Response{
		Success: true,
		Message: "Mucho Successo",
		Data:    mockUser,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse)

}

func main() {
	userHandler := http.HandlerFunc(userHandler)
	http.Handle("/", userHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
