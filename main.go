package main

import (
	"fmt"

	"github.com/mstee1/universal_select/internal/config"
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

	fmt.Println(cfg)

}
