package item

import "context"

func (s *ToDoItemService) DeleteItem(ctx context.Context, itemID, userID int) error {
	return s.repo.DeleteItem(ctx, itemID, userID)
}
