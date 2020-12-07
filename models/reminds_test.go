package models

import (
	"database/sql"
	"log"
	"regexp"
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
		log.Fatalf("Error on DB init: '%s'", err)
	}

	return fakeDb, mock
}

func TestGetUserLastRemind(t *testing.T) {
	fakeDb, mock := NewMock()
	defer fakeDb.Close()
	db = fakeDb

	query := regexp.QuoteMeta(`SELECT remind_id, content FROM reminds WHERE user_id=$1 ORDER BY remind_id DESC LIMIT 1;`)
	rows := sqlmock.NewRows([]string{"remind_id", "content"}).AddRow(r.RemindID, r.Content)

	mock.ExpectQuery(query).WithArgs(r.UserID).WillReturnRows(rows)

	remind, err := GetUserLastRemind(r.UserID)

	if err != nil {
		t.Errorf("GetUserLastRemind returned an error: %s", err)
	} else if remind != *r {
		t.Errorf("GetUserLastRemind didn't return a valid remind")
	}
}

func TestInsertRemind(t *testing.T) {
	fakeDb, mock := NewMock()
	defer fakeDb.Close()
	db = fakeDb

	query := regexp.QuoteMeta("INSERT INTO reminds (content, user_id) VALUES ($1, $2) RETURNING remind_id;")
	rows := sqlmock.NewRows([]string{"remind_id"}).AddRow(0)

	mock.ExpectQuery(query).
		WithArgs(r.Content, r.UserID).
		WillReturnRows(rows)

	insertedRemind, err := InsertRemind(r.UserID, r.Content)

	if err != nil {
		t.Errorf("InsertRemind returned an error: %s", err)
	} else if insertedRemind != *r {
		t.Errorf("InsertRemind didn't return a valid remind")
	}
}
