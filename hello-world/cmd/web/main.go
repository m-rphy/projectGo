package main

import (
	"fmt"
	"net/http"

	"github.com/m-rphy/hello-world-go/pkg/handlers"
)

const PORT = ":8080"

// The basis for any web application in Go!
func main() {

	// Handler functions for different endpoints
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Beep Boop, starting on port%s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
