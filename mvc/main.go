package main

import (
	"embed"
	"fmt"
	"net/http"

	"binus.ac.id/controller"
)

//go:embed static
var static embed.FS

// Main Application
func main() {
	http.HandleFunc("/", controller.HomeController)
	http.HandleFunc("/echo", controller.EchoController)

	http.Handle("/static/", http.FileServer(http.FS(static)))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
