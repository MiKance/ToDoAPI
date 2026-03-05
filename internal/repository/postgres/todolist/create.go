package todolist

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ToDoListPostgres struct {
	pool *pgxpool.Pool
}

func NewToDoListPostgres(pool *pgxpool.Pool) *ToDoListPostgres {
	return &ToDoListPostgres{pool: pool}
}

func (s *ToDoListPostgres) CreateList(ctx context.Context, list models.ToDoList, userID int) (int, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id;", postgres.TodoListsTableName)
	var id int
	err = tx.QueryRow(ctx, query, list.Title, list.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id;", postgres.UsersListTableName)
	err = tx.QueryRow(ctx, query, userID, id).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, err
	}

	return id, nil
}
