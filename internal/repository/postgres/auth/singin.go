package auth

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
)

func (s *AuthPostgres) GetUser(ctx context.Context, username, password string) (models.User, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", postgres.UsersTableName)

	var user models.User
	row := s.pool.QueryRow(ctx, query, username, password)
	err := row.Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
