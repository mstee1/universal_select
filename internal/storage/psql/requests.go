package psql

import "context"

func (r *req) SelectData(ctx context.Context, query string) ([]interface{}, error) {

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []interface{}

	for rows.Next() {
		line := make([]interface{}, len(rows.FieldDescriptions()))
		for i := range line {
			var val interface{}
			line[i] = &val
		}

		if err := rows.Scan(line...); err != nil {
			return nil, err
		}

		result = append(result, line)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}
