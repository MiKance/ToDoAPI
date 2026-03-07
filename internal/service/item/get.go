package item

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func (s *ToDoItemService) GetItemsByListID(ctx context.Context, listId, userID int) (*[]models.ToDoItem, error) {
	return s.repo.GetItemsByListID(ctx, listId, userID)
}

func (s *ToDoItemService) GetItemsByID(ctx context.Context, itemID, userID int) (*models.ToDoItem, error) {
	return s.repo.GetItemsByID(ctx, itemID, userID)
}
