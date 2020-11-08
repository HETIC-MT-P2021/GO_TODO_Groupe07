package models

import (
	"context"
	"errors"
)

// Remind type
type Remind struct {
	RemindID int
	Content  string
	UserID   string
}

// InsertReminds handle request to add a new remind to the db
func InsertReminds(userID string, content string) (Remind, error) {
	insertSQLStatement := `
	INSERT INTO reminds (content, user_id)
	VALUES ($1, $2) RETURNING remind_id;`

	var remind Remind

	row := db.QueryRow(insertSQLStatement, content, userID)
	err := row.Scan(&remind.RemindID)

	if err != nil {
		return remind, err
	}

	remind.Content = content
	remind.UserID = userID

	return remind, nil
}

// GetUserLastRemind handle request to add a new remind to the db
func GetUserLastRemind(userID string) (Remind, error) {
	selectSQL := `SELECT remind_id, content FROM reminds WHERE user_id=$1 ORDER BY remind_id DESC LIMIT 1;`

	var remind Remind

	row := db.QueryRow(selectSQL, userID)
	err := row.Scan(&remind.RemindID, &remind.Content)

	if err != nil {
		return remind, err
	}

	remind.UserID = userID

	return remind, nil
}

// GetUserReminds handle request to add a new remind to the db
func GetUserReminds(ctx context.Context, userID string) ([]Remind, error) {
	selectSQL := `SELECT remind_id, content FROM reminds WHERE user_id=$1 ORDER BY remind_id ASC;`

	var reminds []Remind

	rows, queryErr := db.QueryContext(ctx, selectSQL, userID)

	if queryErr != nil {
		return reminds, queryErr
	}

	for rows.Next() {
		var remind Remind

		if remindErr := rows.Scan(&remind.RemindID, &remind.Content); remindErr != nil {
			return reminds, remindErr
		}

		remind.UserID = userID
		reminds = append(reminds, remind)
	}

	if len(reminds) == 0 {
		return reminds, errors.New("no reminds found for this user")
	}

	return reminds, nil
}

// DeleteRemind handle request to add a new remind to the db
func DeleteRemind(ctx context.Context, remindID string) error {
	deleteSQL := `DELETE FROM reminds WHERE remind_id=$1;`

	_, queryErr := db.QueryContext(ctx, deleteSQL, remindID)

	if queryErr != nil {
		return queryErr
	}

	return nil
}
