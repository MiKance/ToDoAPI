package todolist

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
)

func (s *ToDoListService) GetLists(ctx context.Context, userID int) (*[]models.ToDoList, error) {
	return s.repo.GetLists(ctx, userID)
}

func (s *ToDoListService) GetListByID(ctx context.Context, listID, userID int) (*models.ToDoList, error) {
	return s.repo.GetListByID(ctx, listID, userID)
}
