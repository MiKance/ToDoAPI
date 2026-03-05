package todolist

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ToDoListPostgres) DeleteList(ctx context.Context, listID, userID int) error {
	query := fmt.Sprintf(`DELETE FROM %s USING %s WHERE %s.id = $1 AND %s.list_id = ANY(
		SELECT list_id FROM %s WHERE user_id = $2)`, postgres.TodoListsTableName, postgres.UsersListTableName,
		postgres.TodoListsTableName, postgres.UsersListTableName, postgres.UsersListTableName)

	_, err := s.pool.Exec(ctx, query, listID, userID)
	if err != nil {
		return err
	}

	return nil
}
