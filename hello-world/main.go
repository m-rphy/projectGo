package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

// Uppercase function makes it public!
// Home is th home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About Page and 2 + 2 is %d", sum))
}

// lower case func name makes the function private!
// addValues add two integer and returns the sum
func addValues(x, y int) int {
	return x + y
}

// The basis for any web application in Go!
func main() {

	// Handler functions for different endpoints
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Beep Boop, starting on port%s", PORT))

	_ = http.ListenAndServe(PORT, nil)
}
