package pgxnativedemo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Нативный pgx: свой пул, без database/sql.
func userNames(ctx context.Context, dsn string, limit int) ([]string, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(ctx, `SELECT name FROM users LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, rows.Err()
}
