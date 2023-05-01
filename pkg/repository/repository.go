package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	AuthorizationRepository
	TodoListRepository
	TodoItemRepository
}

type AuthorizationRepository interface {
}

type TodoListRepository interface {
}

type TodoItemRepository interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
