package repository

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres/auth"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres/item"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres/todolist"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, username, password string) (models.User, error)
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

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPostgres(pool),
		ToDoList:      todolist.NewToDoListPostgres(pool),
		ToDoItem:      item.NewItemService(pool),
	}
}
