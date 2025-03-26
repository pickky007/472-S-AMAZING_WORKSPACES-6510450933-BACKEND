package services

import (
	"onez19/config"
	"onez19/models"
)

func CreateMessage(message models.Message) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO message (id, message, date, workspace_id, username) VALUES (?, ?, ?, ?, ?)",
		message.ID, message.Message, message.Date, message.WorkspaceID, message.Username)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}
func DeleteMessage(messageID string) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM message WHERE id = ?", messageID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func GetAllMessagesByWorkspaceID(workspace_id string) ([]models.Message, error) {
	var messages []models.Message

	query := `
		SELECT *
		FROM message
		WHERE workspace_id = ?;
	`

	rows, err := config.DB.Query(query, workspace_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.Message, &message.Date, &message.WorkspaceID, &message.Username); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil

}

func SearchMessagesByText(query string, workspaceID string) ([]models.Message, error) {
	rows, err := config.DB.Query("SELECT * FROM message WHERE workspace_id = ? AND message LIKE ? ORDER BY date DESC", workspaceID, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.Message, &message.Date, &message.WorkspaceID, &message.Username); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func SearchMessagesByRegex(query string, workspaceID string) ([]models.Message, error) {
	rows, err := config.DB.Query("SELECT * FROM message WHERE workspace_id = ? AND message REGEXP ? ORDER BY date DESC", workspaceID, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.Message, &message.Date, &message.WorkspaceID, &message.Username); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
