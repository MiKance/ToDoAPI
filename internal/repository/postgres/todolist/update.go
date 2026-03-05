package todolist

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ToDoListPostgres) UpdateList(ctx context.Context, list models.ToDoList, userID int) error {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE id = $3 AND $4 = 
(SELECT user_id FROM %s WHERE list_id = $3)`, postgres.TodoListsTableName, postgres.UsersListTableName)
	_, err := s.pool.Exec(ctx, query, list.Title, list.Description, list.ID, userID)
	if err != nil {
		return err
	}
	return nil
}
