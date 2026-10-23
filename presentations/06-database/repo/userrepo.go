package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type User struct {
	ID   int
	Name string
}

// Интерфейс объявляет потребитель (сервис), реализует репозиторий — лекция 2.
type UserRepository interface {
	GetByID(ctx context.Context, id int) (User, error)
	Create(ctx context.Context, name string) (int, error)
}

type Postgres struct{ db *sql.DB }

func New(db *sql.DB) *Postgres { return &Postgres{db: db} }

func (r *Postgres) GetByID(ctx context.Context, id int) (User, error) {
	var u User
	err := r.db.QueryRowContext(ctx,
		`SELECT id, name FROM users WHERE id = $1`, id).Scan(&u.ID, &u.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, fmt.Errorf("user %d: %w", id, ErrNotFound)
	}
	return u, err
}

func (r *Postgres) Create(ctx context.Context, name string) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO users (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	return id, err
}
