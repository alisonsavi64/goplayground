package main

import (
	"log"
	"net/http"

	"hello-world/handlers"
	"hello-world/utils"
)

func main() {
	utils.Log("Starting server on :8080")

	http.HandleFunc("/", handlers.Hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
