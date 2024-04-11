package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	router := http.NewServeMux()

	router.HandleFunc("GET /api/launches", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("List Launches")
	})

	router.HandleFunc("GET /api/launches/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Printf("Get Launch based on id: %v", id)
	})

	router.HandleFunc("POST /api/launches", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create Launch")
	})

	router.HandleFunc("DELETE /api/launches/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Printf("Delete Launch based on id: %v", id)
	})

	http.ListenAndServe(":8080", router)
}
