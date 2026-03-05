package repository

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres/auth"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres/todolist"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, username, password string) (models.User, error)
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

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPostgres(pool),
		ToDoList:      todolist.NewToDoListPostgres(pool),
	}
}
