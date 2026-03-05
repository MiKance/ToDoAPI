package todolist

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository"
)

type ToDoListService struct {
	repo repository.ToDoList
}

func NewToDoListService(repo repository.ToDoList) *ToDoListService {
	return &ToDoListService{repo: repo}
}

func (s *ToDoListService) CreateList(ctx context.Context, list models.ToDoList, userID int) (int, error) {
	return s.repo.CreateList(ctx, list, userID)
}
