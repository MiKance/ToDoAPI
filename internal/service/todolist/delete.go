package todolist

import (
	"context"
)

func (s *ToDoListService) DeleteList(ctx context.Context, listID, userID int) error {
	return s.repo.DeleteList(ctx, listID, userID)
}
