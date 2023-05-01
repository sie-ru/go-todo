package service

import "todo/pkg/repository"

type Service struct {
	AuthorizationService
	TodoListService
	TodoItemService
}

type AuthorizationService interface {
}

type TodoListService interface {
}

type TodoItemService interface {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}

}
