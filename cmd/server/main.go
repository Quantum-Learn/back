package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Пример маршрута
	r.Get("/courses", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of courses"))
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
