package infrastructure

import (
	"database/sql"
	"errors"

	"github.com/Abiti0233/go_api_test/backend/domain"
)

type postgresTodoRepository struct {
	db *sql.DB
}

func NewPostgresTodoRepository(db *sql.DB) domain.TodoRepository {
	return &postgresTodoRepository{db: db}
}

func (r *postgresTodoRepository) GetAll() ([]*domain.Todo, error) {
	rows, err := r.db.Query("SELECT id, text, done FROM todos ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*domain.Todo
	for rows.Next() {
		var t domain.Todo
		if err := rows.Scan(&t.ID, &t.Text, &t.Done); err != nil {
			return nil, err
		}
		todos = append(todos, &t)
	}
	if todos == nil {
		todos = []*domain.Todo{}
	}
	return todos, nil
}

func (r *postgresTodoRepository) GetByID(id int) (*domain.Todo, error) {
	row := r.db.QueryRow("SELECT id, text, done FROM todos WHERE id = $1", id)
	var t domain.Todo
	if err := row.Scan(&t.ID, &t.Text, &t.Done); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *postgresTodoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	// RETURNING を使ってINSERT時のIDを受け取る
	err := r.db.QueryRow(`
			INSERT INTO todos (text, done)
			VALUES ($1, $2)
			RETURNING id
	`, todo.Text, todo.Done).Scan(&todo.ID)
	if err != nil {
			return nil, err
	}
	return todo, nil
}

func (r *postgresTodoRepository) Update(todo *domain.Todo) (*domain.Todo, error) {
	result, err := r.db.Exec(`
			UPDATE todos
			SET text = $1, done = $2
			WHERE id = $3
	`, todo.Text, todo.Done, todo.ID)
	if err != nil {
			return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
			return nil, err
	}
	if rowsAffected == 0 {
			return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (r *postgresTodoRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
			return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
			return err
	}
	if rowsAffected == 0 {
			return errors.New("todo not found")
	}
	return nil
}