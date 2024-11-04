package services

import (
	"onez19/models"
	"onez19/config"
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
