package main

import (
	"fmt"
	"net/http"

	"binus.ac.id/controller"
)

// Main Application
func main() {
	// Configure routes
	http.HandleFunc("/", controller.HomeController)
	http.HandleFunc("/echo", controller.EchoController)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
