package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mstee1/universal_select/internal/config"
)

type connect struct {
	pool *pgxpool.Pool
}

func NewDb(ctx context.Context, cfg *config.Config) (Connect, *pgxpool.Pool, error) {

	pool, err := DbConnect(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}

	return &connect{
		pool: pool,
	}, pool, nil
}
