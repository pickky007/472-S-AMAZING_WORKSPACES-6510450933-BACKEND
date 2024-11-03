package models

type Section struct {
	ID          int    `json:"id" db:"id"`
	WorkspaceID int    `json:"workspace_id" db:"workspace_id"` // foreign key จาก Workspace
	Name        string `json:"name" db:"name"`
}
