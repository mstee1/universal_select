package psql

import "context"

type Requests interface {
	SelectData(ctx context.Context, query string) ([][]interface{}, error)
}
