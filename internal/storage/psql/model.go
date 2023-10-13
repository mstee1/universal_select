package psql

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type req struct {
	pool *pgxpool.Pool
}

func NewReq(pool *pgxpool.Pool) (Requests, error) {

	return &req{
		pool: pool,
	}, nil
}
