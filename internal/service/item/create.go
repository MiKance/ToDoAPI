package item

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func (s *ToDoItemService) CreateItem(ctx context.Context, item models.ToDoItem, userID int) (int, error) {
	return s.repo.CreateItem(ctx, item, userID)
}
