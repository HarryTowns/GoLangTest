package main

import (
	"fmt"
	"net/http"
)

// Task: Create a simple Go web server that handles basic routes and displays messages.

func main() {
	// Set up the routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server on port 8080
	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// Home route handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go sample project!")
}

// Hello route handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go developer!")
}

// About route handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a sample project to learn Go with a basic web server.")
}
