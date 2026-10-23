package queryrowsdemo

import (
	"context"
	"database/sql"
	"fmt"
	"io"
)

type User struct {
	ID   int
	Name string
}

func allUsers(ctx context.Context, db *sql.DB, w io.Writer) error {
	rows, err := db.QueryContext(ctx, `SELECT id, name FROM users ORDER BY id`)
	if err != nil {
		return err
	}
	defer rows.Close() // ОБЯЗАТЕЛЬНО: иначе соединение не вернётся в пул

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return err
		}
		fmt.Fprintf(w, "%d %s\n", u.ID, u.Name)
	}
	return rows.Err() // ошибка итерации проявляется только здесь
}
