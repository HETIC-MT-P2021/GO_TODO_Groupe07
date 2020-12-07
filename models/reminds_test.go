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

func TestGetUserLastRemind(t *testing.T) {
	fakeDb, mock := NewMock()
	defer fakeDb.Close()
	db = fakeDb

	query := "SELECT remind_id, content FROM reminds WHERE user_id=$1 ORDER BY remind_id DESC LIMIT 1;"
	rows := sqlmock.NewRows([]string{"remind_id", "content"})
	mock.ExpectQuery(query).WithArgs(r.UserID).WillReturnRows(rows)

	remind, err := GetUserLastRemind(r.UserID)

	if err != nil {
		t.Errorf("GetUserLastRemind returned an error: %s", err)
	} else if &remind != r {
		t.Errorf("GetUserLastRemind didn't return a valid remind")
	}
}

func TestInsertReminds(t *testing.T) {
	fakeDb, mock := NewMock()
	defer fakeDb.Close()
	db = fakeDb

	query := "INSERT INTO reminds (content, user_id) VALUES ($1, $2) RETURNING remind_id;"
	rows := sqlmock.NewRows([]string{"remind_id"})

	mock.ExpectQuery(query).
		WithArgs(r.Content, r.UserID).
		WillReturnRows(rows)

	insertedRemind, err := InsertReminds(r.UserID, r.Content)

	if err != nil {
		t.Errorf("InsertReminds returned an error: %s", err)
	} else if &insertedRemind != r {
		t.Errorf("InsertReminds didn't return a valid remind")
	}
}
