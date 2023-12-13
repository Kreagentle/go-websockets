package main

import (
	"log"
	"net/http"
)

// create some basic routes
func main() {
	routes := routes()

	log.Println("Server starts on port :8080")

	_ = http.ListenAndServe(":8080", routes)
}
