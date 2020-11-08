package models

import (
	"context"
	"errors"
)

// Action type
type Action struct {
	ActionID int
	Content  string
	UserID   string
}

// InsertActions handle request to add a new action to the db
func InsertActions(userID string, content string) (Action, error) {
	insertSQLStatement := `
	INSERT INTO actions (content, user_id)
	VALUES ($1, $2) RETURNING action_id;`

	var action Action

	row := db.QueryRow(insertSQLStatement, content, userID)
	err := row.Scan(&action.ActionID)

	if err != nil {
		return action, err
	}

	action.Content = content
	action.UserID = userID

	return action, nil
}

// GetUserLastAction handle request to add a new action to the db
func GetUserLastAction(userID string) (Action, error) {
	selectSQL := `SELECT action_id, content FROM actions WHERE user_id=$1 ORDER BY action_id DESC LIMIT 1;`

	var action Action

	row := db.QueryRow(selectSQL, userID)
	err := row.Scan(&action.ActionID, &action.Content)

	if err != nil {
		return action, err
	}

	action.UserID = userID

	return action, nil
}

// GetUserActions handle request to add a new action to the db
func GetUserActions(ctx context.Context, userID string) ([]Action, error) {
	selectSQL := `SELECT action_id, content FROM actions WHERE user_id=$1 ORDER BY action_id ASC;`

	var actions []Action

	rows, queryErr := db.QueryContext(ctx, selectSQL, userID)

	if queryErr != nil {
		return actions, queryErr
	}

	for rows.Next() {
		var action Action

		if actionErr := rows.Scan(&action.ActionID, &action.Content); actionErr != nil {
			return actions, actionErr
		}

		action.UserID = userID
		actions = append(actions, action)
	}

	if len(actions) == 0 {
		return actions, errors.New("no actions found for this user")
	}

	return actions, nil
}

// DeleteAction handle request to add a new action to the db
func DeleteAction(ctx context.Context, actionID string) error {
	deleteSQL := `DELETE FROM actions WHERE action_id=$1;`

	_, queryErr := db.QueryContext(ctx, deleteSQL, actionID)

	if queryErr != nil {
		return queryErr
	}

	return nil
}
