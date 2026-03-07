package todolist

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ToDoListPostgres) GetLists(ctx context.Context, userID int) (*[]models.ToDoList, error) {
	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE user_id = $1",
		postgres.ListsTableName)

	rows, err := s.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := make([]models.ToDoList, 0)
	for rows.Next() {
		list := models.ToDoList{}
		if err := rows.Scan(&list.ID, &list.Title, &list.Description); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return &lists, nil
}

func (s *ToDoListPostgres) GetListByID(ctx context.Context, listID, userID int) (*models.ToDoList, error) {
	query := fmt.Sprintf(`SELECT title, description FROM %s WHERE id = $1 AND user_id = $2`, postgres.ListsTableName)
	list := &models.ToDoList{}
	row := s.pool.QueryRow(ctx, query, listID, userID)
	if err := row.Scan(&list.Title, &list.Description); err != nil {
		return nil, err
	}
	list.ID = listID
	return list, nil
}
