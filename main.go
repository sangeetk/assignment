package main

import (
	"log"
	"net/http"

	"assignment/handlers"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
// Test with
//      go test ./...

func main() {

	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/invert", handlers.Invert)
	http.HandleFunc("/flatten", handlers.Flatten)
	http.HandleFunc("/sum", handlers.Sum)
	http.HandleFunc("/multiply", handlers.Multiply)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
