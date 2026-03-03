package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/MiKance/ToDoAPI/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(ctx context.Context, cfg *config.PostgresConfig) *Storage {
	var s Storage

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	s.pool = pool

	return &s
}
