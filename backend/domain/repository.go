package domain

type TodoRepository interface {
	GetAll() ([]*Todo, error)
	GetByID(id int) (*Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id int) error
}