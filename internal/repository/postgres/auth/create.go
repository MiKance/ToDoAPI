package auth

import (
	"context"
	"fmt"

	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthPostgres struct {
	pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{pool: pool}
}

func (a *AuthPostgres) CreateUser(ctx context.Context, user models.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id;",
		postgres.UsersTableName)

	var id int
	err := a.pool.QueryRow(ctx, query, user.Name, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
