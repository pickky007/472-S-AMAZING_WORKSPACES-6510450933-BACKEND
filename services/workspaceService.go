package services

import (
	"onez19/config"
	"onez19/models"

	"github.com/google/uuid"
)

func GetWorkspacesByUsername(username string) ([]models.Workspace, error) {
	var workspaces []models.Workspace

	query := `
        SELECT w.id, w.name, w.description, w.owner
        FROM workspace AS w
        INNER JOIN user_workspace AS uw ON w.id = uw.workspace_id
        WHERE uw.username = ?
    `

	rows, err := config.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var workspace models.Workspace
		if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Description, &workspace.Owner); err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}

	// ตรวจสอบข้อผิดพลาดจาก rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workspaces, nil
}

func CreateWorkspace(workspace models.Workspace) error {

	// วนซ้ำจนกว่า uuid จะไม่ซ้ำ :D
	for {
		workspace.ID = generateShortUUID()
		exists, err := workspaceExists(workspace.ID)
		if err != nil {
			return err
		}
		if !exists {
			break
		}
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO workspace (id, name, description, owner) VALUES (?, ?, ?, ?)",
		workspace.ID, workspace.Name, workspace.Description, workspace.Owner)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO user_workspace (username, workspace_id) VALUES (?, ?)", workspace.Owner, workspace.ID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func JoinWorkspace(username, workspaceID string) error {
	_, err := config.DB.Exec("INSERT INTO user_workspace (username, workspace_id) VALUES (?, ?)", username, workspaceID)
	return err
}

func generateShortUUID() string {
	return uuid.New().String()[:6]
}

func workspaceExists(id string) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM workspace WHERE id = ?)", id).Scan(&exists)
	return exists, err
}
