package queryrowdemo

import (
	"context"
	"database/sql"
	"errors"
)

func nameByID(ctx context.Context, db *sql.DB, id int) (string, error) {
	var name string
	err := db.QueryRowContext(ctx,
		`SELECT name FROM users WHERE id = $1`, id).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		return "", errors.New("user not found") // в бою: своя ErrNotFound через %w
	}
	return name, err
}
