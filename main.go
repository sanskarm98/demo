package main

import (
	"log"
	"net/http"

	"demo/handlers"
	"demo/router"
	"demo/store"
)

func main() {
	empStore := store.NewEmployeeStore()
	server := handlers.NewServer(empStore)
	r := router.NewRouter(server)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
