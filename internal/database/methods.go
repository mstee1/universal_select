package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mstee1/universal_select/internal/config"
)

func (c *connect) CloseDb() {
	c.pool.Close()
}

func DbConnect(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	conf := getConfig(&cfg.Database)
	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func getConfig(cfg *config.Database) *pgxpool.Config {
	parseConfig, _ := pgxpool.ParseConfig("")
	parseConfig.ConnConfig.User = cfg.DbUser
	parseConfig.ConnConfig.Password = cfg.DbPassword
	parseConfig.ConnConfig.Host = cfg.DbHost
	parseConfig.ConnConfig.Port = uint16(cfg.DbPort)
	parseConfig.ConnConfig.Database = cfg.DbName

	parseConfig.MaxConns = 1

	return parseConfig
}
