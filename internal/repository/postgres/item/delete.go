package item

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ItemPostgres) DeleteItem(ctx context.Context, itemID, userID int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 AND list_id = ANY (
			SELECT id FROM %s WHERE user_id = $2)`, postgres.ItemsTableName, postgres.ListsTableName)

	_, err := s.pool.Exec(ctx, query, itemID, userID)
	if err != nil {
		return err
	}

	return nil
}
