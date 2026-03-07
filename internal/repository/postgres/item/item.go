package item

import "github.com/jackc/pgx/v5/pgxpool"

type ItemPostgres struct {
	pool *pgxpool.Pool
}

func NewItemService(pool *pgxpool.Pool) *ItemPostgres {
	return &ItemPostgres{pool: pool}
}
