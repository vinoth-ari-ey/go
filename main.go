package main

import (
	"fmt"
	"go-crud/db"
	userHandler "go-crud/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	//Load .env when run in local
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database
	db.Connect()

	// Set up user routes
	http.HandleFunc("/users/", userHandler.UserHandler)

	// Set up an app route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to My Go Application!")
	})

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
