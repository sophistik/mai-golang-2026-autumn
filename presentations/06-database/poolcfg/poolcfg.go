package poolcfgdemo

import (
	"database/sql"
	"time"
)

func tune(db *sql.DB) {
	db.SetMaxOpenConns(25)                 // потолок открытых соединений
	db.SetMaxIdleConns(25)                 // сколько держать простаивающих
	db.SetConnMaxLifetime(5 * time.Minute) // пересоздавать по возрасту
	db.SetConnMaxIdleTime(2 * time.Minute) // закрывать долго простаивавшие
}
