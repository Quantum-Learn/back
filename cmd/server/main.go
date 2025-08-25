package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"MVP_project/internal/apihttp"
	"MVP_project/internal/auth"
	"MVP_project/internal/handlers"
	"MVP_project/internal/storage"
)

func main() {
	// store (пока in‑memory)
	store := storage.NewInMemoryUserStore()

	// бизнес‑хэндлеры
	authHandler := handlers.NewAuthHandler(store)

	// http‑обвязка
	httpAuth := &apihttp.AuthHTTP{H: authHandler, Store: store}

	r := chi.NewRouter()

	// публичные ручки
	r.Post("/auth/register", httpAuth.Register)
	r.Post("/auth/login", httpAuth.Login)

	// защищённые ручки
	r.Group(func(pr chi.Router) {
		pr.Use(auth.JWTMiddleware)
		pr.Get("/users/me", httpAuth.Me)
	})

	log.Println("server :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
