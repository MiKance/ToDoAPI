package todolist

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func (s *ToDoListService) UpdateList(ctx context.Context, list models.ToDoList, userID int) error {
	return s.repo.UpdateList(ctx, list, userID)
}
