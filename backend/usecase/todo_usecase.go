package usecase

import (
	"errors"

	"github.com/Abiti0233/go_api_test/backend/domain"
)

type TodoUseCase interface {
	GetTodos() ([]*domain.Todo, error)
	GetTodo(id int) (*domain.Todo, error)
	CreateTodo(text string) (*domain.Todo, error)
	UpdateTodo(id int, text string, done bool) (*domain.Todo, error)
	DeleteTodo(id int) error
}

type todoUseCase struct {
	repo domain.TodoRepository
}

func NewTodoUseCase(repo domain.TodoRepository) TodoUseCase {
	return &todoUseCase{repo: repo}
}

func (uc *todoUseCase) GetTodos() ([]*domain.Todo, error) {
	return uc.repo.GetAll()
}

func (uc *todoUseCase) GetTodo(id int) (*domain.Todo, error) {
	return uc.repo.GetByID(id)
}

func (uc *todoUseCase) CreateTodo(text string) (*domain.Todo, error) {
	if text == "" {
		return nil, errors.New("text cannnot be empty")
	}
	newTodo := &domain.Todo{Text: text, Done: false}
	return uc.repo.Create(newTodo)
}

func (uc *todoUseCase) UpdateTodo(id int, text string, done bool) (*domain.Todo, error) {
	existing, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("todo not found")
	}
	existing.Text = text
	existing.Done = done
	return uc.repo.Update(existing)
}

func (uc *todoUseCase) DeleteTodo(id int) error {
	return uc.repo.Delete(id)
}


