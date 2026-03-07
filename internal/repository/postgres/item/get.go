package item

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *ItemPostgres) GetItemsByListID(ctx context.Context, listId, userID int) (*[]models.ToDoItem, error) {
	var items []models.ToDoItem

	query := fmt.Sprintf(`SELECT title, description, done FROM %s WHERE list_id = $1 AND list_id = ANY (
							SELECT id FROM %s WHERE user_id = $2)`,
		postgres.ItemsTableName, postgres.ListsTableName)

	rows, err := s.pool.Query(ctx, query, listId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.ToDoItem
		if err := rows.Scan(&item.Title, &item.Description, &item.Done); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func (s *ItemPostgres) GetItemsByID(ctx context.Context, itemID, userID int) (*models.ToDoItem, error) {
	var item models.ToDoItem

	query := fmt.Sprintf(`SELECT title, description, done FROM %s WHERE id = $1 AND list_id = ANY (
							SELECT id FROM %s WHERE user_id = $2)`,
		postgres.ItemsTableName, postgres.ListsTableName)

	row := s.pool.QueryRow(ctx, query, itemID, userID)
	if err := row.Scan(&item.Title, &item.Description, &item.Done); err != nil {
		return nil, err
	}

	return &item, nil
}
