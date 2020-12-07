package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var r = &Remind{
	RemindID: 0,
	Content:  "Reminder",
	UserID:   "Somebody",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	fakeDb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error on DB ini: '%s'", err)
	}

	return fakeDb, mock
}

func TestInsertReminds(t *testing.T) {
	fakeDb, mock := NewMock()
	defer fakeDb.Close()

	db = fakeDb
	mock.ExpectExec("INSERT INTO reminds (content, user_id) VALUES ($1, $2) RETURNING remind_id;").
		WithArgs(r.Content, r.UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	insertedRemind, err := InsertReminds(r.UserID, r.Content)

	if err != nil {
		t.Errorf("InsertReminds returned an error: %s", err)
	}
	if &insertedRemind != r {
		t.Errorf("InsertReminds didn't return a valid remind")
	}
}
