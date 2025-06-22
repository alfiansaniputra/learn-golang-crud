package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"belajar-golang-api/config"
	"belajar-golang-api/routes"
)

func main() {
	// Initialize database
	db := config.InitDB()
	defer db.Close()

	// Initialize router
	r := mux.NewRouter()

	// Setup routes
	routes.SetupAuthRoutes(r)

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
