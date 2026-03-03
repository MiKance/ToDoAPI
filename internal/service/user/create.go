package user

import (
	"context"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository"
)

type AuthService struct {
	repo *repository.Authorization
}

func NewAuthService(repo *repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(ctx context.Context, user models.User) (int, error) {
	return 0, nil
}
