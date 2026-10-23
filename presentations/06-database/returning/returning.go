package returningdemo

import (
	"context"
	"database/sql"
)

// LastInsertId в Postgres не работает — нужен RETURNING.
func create(ctx context.Context, db *sql.DB, name string) (int, error) {
	var id int
	err := db.QueryRowContext(ctx,
		`INSERT INTO users (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	return id, err
}

func rename(ctx context.Context, db *sql.DB, id int, name string) (int64, error) {
	res, err := db.ExecContext(ctx,
		`UPDATE users SET name = $1 WHERE id = $2`, name, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected() // сколько строк реально изменили
}
