package item

import "github.com/MiKance/ToDoAPI/internal/repository"

type ToDoItemService struct {
	repo repository.ToDoItem
}

func NewToDoItemService(repo repository.ToDoItem) *ToDoItemService {
	return &ToDoItemService{repo: repo}
}
