package psql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mstee1/universal_select/internal/config"
)

type req struct {
	cfg  *config.Config
	pool *pgxpool.Pool
}

func NewReq(cfg *config.Config, pool *pgxpool.Pool) (Requests, error) {

	return &req{
		cfg:  cfg,
		pool: pool,
	}, nil
}
