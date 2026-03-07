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
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES ($1, $2, $3) RETURNING id;",
		postgres.ListsTableName)
	var id int
	err := s.pool.QueryRow(ctx, query, userID, list.Title, list.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
