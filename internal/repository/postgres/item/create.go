package item

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ItemPostgres) CreateItem(ctx context.Context, item models.ToDoItem, userID int) (int, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = $1 AND id = $2",
		postgres.ListsTableName)
	row := s.pool.QueryRow(ctx, query, userID, item.ListID)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (list_id, title, description) VALUES ($1, $2, $3) RETURNING id",
		postgres.ItemsTableName)
	row = s.pool.QueryRow(ctx, query, item.ListID, item.Title, item.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
