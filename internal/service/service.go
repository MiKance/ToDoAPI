package service

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository"
	"github.com/MiKance/ToDoAPI/internal/service/auth"
	"github.com/MiKance/ToDoAPI/internal/service/item"
	"github.com/MiKance/ToDoAPI/internal/service/todolist"
)

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	CreateList(ctx context.Context, list models.ToDoList, userID int) (int, error)
	GetLists(ctx context.Context, userID int) (*[]models.ToDoList, error)
	GetListByID(ctx context.Context, listID, userID int) (*models.ToDoList, error)
	UpdateList(ctx context.Context, list models.ToDoList, userID int) error
	DeleteList(ctx context.Context, listID, userID int) error
}

type ToDoItem interface {
	CreateItem(ctx context.Context, item models.ToDoItem, userID int) (int, error)
	GetItemsByListID(ctx context.Context, listId, userID int) (*[]models.ToDoItem, error)
	GetItemsByID(ctx context.Context, itemID, userID int) (*models.ToDoItem, error)
	UpdateItem(ctx context.Context, item models.ToDoItem, userID int) error
	DeleteItem(ctx context.Context, itemID, userID int) error
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
		ToDoItem:      item.NewToDoItemService(repo.ToDoItem),
	}
}
