package services

import (
	"onez19/config"
	"onez19/models"
)

func GetAllSectionsByWorkspaceID(workspaceID int) ([]models.Section, error) {
	var sections []models.Section

	// ดึงข้อมูล sections จากฐานข้อมูลตาม workspace_id
	rows, err := config.DB.Query("SELECT id, workspace_id, name FROM section WHERE workspace_id = ?", workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var section models.Section
		if err := rows.Scan(&section.ID, &section.WorkspaceID, &section.Name); err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}
