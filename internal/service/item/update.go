package item

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func (s *ToDoItemService) UpdateItem(ctx context.Context, item models.ToDoItem, userID int) error {
	return s.repo.UpdateItem(ctx, item, userID)
}
