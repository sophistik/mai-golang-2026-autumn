package txdemo

import (
	"context"
	"database/sql"
	"fmt"
)

func transfer(ctx context.Context, db *sql.DB, from, to, amount int) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // идиома: после Commit — no-op, не помешает

	if _, err := tx.ExecContext(ctx,
		`UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, from); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx,
		`UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, to); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}
