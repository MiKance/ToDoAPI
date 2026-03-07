package todolist

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ToDoListPostgres) DeleteList(ctx context.Context, listID, userID int) error {
	query := fmt.Sprintf(`DELETE FROM %s  WHERE id = $1 AND user_id = $2`, postgres.ListsTableName)

	_, err := s.pool.Exec(ctx, query, listID, userID)
	if err != nil {
		return err
	}

	return nil
}
