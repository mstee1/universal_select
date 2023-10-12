package methods

import (
	"context"
	"fmt"
	"time"

	"github.com/mstee1/universal_select/internal/config"
	"github.com/mstee1/universal_select/internal/storage/psql"
)

func Select1(ctx context.Context, cfg *config.Config, req psql.Requests) {

	data, err := req.SelectData(ctx, cfg.Sql.Select1)
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int32
	for _, line := range data {
		id = line[0].(int32)
		fmt.Println(id)
	}
}

func Select2(ctx context.Context, cfg *config.Config, req psql.Requests) {

	data, err := req.SelectData(ctx, cfg.Sql.Select2)
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int32
	var cm string
	for _, line := range data {
		id = line[0].(int32)
		cm = line[1].(string)
		fmt.Println(id, cm)
	}
}

func Select3(ctx context.Context, cfg *config.Config, req psql.Requests) {

	data, err := req.SelectData(ctx, cfg.Sql.Select3)
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int32
	var cm string
	var dte time.Time
	for _, line := range data {
		id = line[0].(int32)
		cm = line[1].(string)
		dte = line[2].(time.Time)
		fmt.Println(id, cm, dte)
	}
}

func Select4(ctx context.Context, cfg *config.Config, req psql.Requests) {

	type testData struct {
		id  int32
		cm  string
		dte time.Time
	}

	data, err := req.SelectData(ctx, cfg.Sql.Select4)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result []testData
	for _, line := range data {
		result = append(result, testData{
			id:  line[0].(int32),
			cm:  line[1].(string),
			dte: line[2].(time.Time),
		})
	}

	fmt.Println(result)
}
