package service

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository"
)

type Authorization interface {
	Create(ctx context.Context, user models.User) (int, error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
