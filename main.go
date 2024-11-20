package main

import (
	"net/http"

	"assignment/handlers"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {

	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/invert", handlers.Invert)
	http.HandleFunc("/flatten", handlers.Flatten)
	http.HandleFunc("/sum", handlers.Sum)
	http.HandleFunc("/multiply", handlers.Multiply)

	http.ListenAndServe(":8080", nil)
}
