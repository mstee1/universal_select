package main

import (
	"context"
	"fmt"

	"github.com/mstee1/universal_select/internal/config"
	"github.com/mstee1/universal_select/internal/database"
	"github.com/mstee1/universal_select/internal/methods"
	"github.com/mstee1/universal_select/internal/storage/psql"
	"github.com/spf13/viper"
)

func main() {

	v := viper.New()
	cfg, err := config.GetConfig(v)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.Version != "" {
		fmt.Println(cfg.Version)
		return
	}

	ctx := context.Background()

	conn, pool, err := database.NewDb(ctx, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.CloseDb()

	req, err := psql.NewReq(pool)
	if err != nil {
		fmt.Println(err)
		return
	}

	methods.Select1(ctx, cfg, req)

	methods.Select2(ctx, cfg, req)

	methods.Select3(ctx, cfg, req)

	methods.Select4(ctx, cfg, req)

}
