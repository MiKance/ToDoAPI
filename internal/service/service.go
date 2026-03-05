package service

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository"
	"github.com/MiKance/ToDoAPI/internal/service/auth"
	"github.com/MiKance/ToDoAPI/internal/service/todolist"
)

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	CreateList(ctx context.Context, list models.ToDoList, userID int) (int, error)
	GetLists(ctx context.Context, userID int) ([]*models.ToDoList, error)
	GetListByID(ctx context.Context, listID, userID int) (*models.ToDoList, error)
	UpdateList(ctx context.Context, list models.ToDoList, userID int) error
	DeleteList(ctx context.Context, listID, userID int) error
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuthService(repo.Authorization),
		ToDoList:      todolist.NewToDoListService(repo.ToDoList),
	}
}
