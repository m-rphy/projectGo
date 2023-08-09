package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

// The basis for any web application in Go!
func main() {

	// Handler functions for different endpoints
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Beep Boop, starting on port%s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
