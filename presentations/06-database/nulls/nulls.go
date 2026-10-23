package nullsdemo

import "database/sql"

type Session struct {
	ID   int
	Bio  *string      // NULL → nil
	Seen sql.NullTime // .Valid + время
}

func scan(row *sql.Row) (Session, error) {
	var s Session
	err := row.Scan(&s.ID, &s.Bio, &s.Seen)
	return s, err
}
