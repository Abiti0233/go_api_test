package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/Abiti0233/go_api_test/backend/usecase"
)

type TodoHandler struct {
	todoUC usecase.TodoUseCase
}

func NewTodoHandler(todoUC usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{todoUC: todoUC}
}

// 全TODO取得
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoUC.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, todos, http.StatusOK)
}

// 単一TODO取得
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	todo, err := h.todoUC.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if todo == nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}
	respondJSON(w, todo, http.StatusOK)
}

// 新規作成
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.todoUC.CreateTodo(req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, todo, http.StatusCreated)
}

// 更新
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req struct {
		Text string
		Done bool
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.todoUC.UpdateTodo(id, req.Text, req.Done)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, todo, http.StatusOK)
}

// 削除
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
	}
	if err := h.todoUC.DeleteTodo(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
			json.NewEncoder(w).Encode(data)
	}
}

