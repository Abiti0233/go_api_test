package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(todoHandler *TodoHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.User(middleware.Recoverer)

	// エンドポイント
	r.Get("/todos", todoHandler.GetTodos)
	r.Get("/todos/{id}", todoHandler.GetTodo)
	r.Post("/todos", todoHandler.CreateTodo)
	r.Put("/todos/{id}", todoHandler.UpdateTodo)
	r.Delete("/todos/{id}", todoHandler.DeleteTodo)

	return r
}