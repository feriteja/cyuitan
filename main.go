package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize todos slice
	initTodos()

	// Define routes
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todoHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
