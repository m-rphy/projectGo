package main

import (
	"fmt"
	"net/http"
)

// The basis for any web application in Go!
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
