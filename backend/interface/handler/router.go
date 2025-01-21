package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(todoHandler *TodoHandler) *chi.Mux {
	r := chi.NewRouter()

	// CORS設定
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// エンドポイント
	r.Get("/todos", todoHandler.GetTodos)
	r.Get("/todos/{id}", todoHandler.GetTodo)
	r.Post("/todos", todoHandler.CreateTodo)
	r.Put("/todos/{id}", todoHandler.UpdateTodo)
	r.Delete("/todos/{id}", todoHandler.DeleteTodo)

	return r
}
