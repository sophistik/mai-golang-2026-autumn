package connectdemo

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // регистрирует драйвер "pgx"
)

func open(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn) // соединений ещё НЕТ
	if err == nil {
		err = db.PingContext(ctx) // первая реальная проверка
	}
	return db, err
}
