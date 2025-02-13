package main

import (
	"log"
	"net/http"
	"receipt-processor/routes"
)

func main() {
	r := routes.SetupRoutes()
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
