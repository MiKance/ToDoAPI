package item

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ItemPostgres) UpdateItem(ctx context.Context, item models.ToDoItem, userID int) error {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2, done = $3 WHERE id = $4 AND list_id IN (
        SELECT id FROM %s WHERE user_id = $5
    );`,
		postgres.ItemsTableName, postgres.ListsTableName)

	_, err := s.pool.Exec(ctx, query, item.Title, item.Description, item.Done, item.ID, userID)
	if err != nil {
		return err
	}
	return nil
}
