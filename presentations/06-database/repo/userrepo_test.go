package repo

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const getUserSQL = `SELECT id, name FROM users WHERE id = $1`

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(42, "Аня")
	mock.ExpectQuery(regexp.QuoteMeta(getUserSQL)).
		WithArgs(42).
		WillReturnRows(rows)

	r := New(db)
	u, err := r.GetByID(context.Background(), 42)
	if err != nil {
		t.Fatalf("GetByID: %v", err)
	}
	if u.Name != "Аня" {
		t.Errorf("name = %q, want Аня", u.Name)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания выполнены: %v", err)
	}
}

func TestGetByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(getUserSQL)).
		WillReturnError(sql.ErrNoRows)

	r := New(db)
	if _, err := r.GetByID(context.Background(), 1); !errors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}
