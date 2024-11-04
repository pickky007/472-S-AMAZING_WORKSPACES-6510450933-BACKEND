package models

type Activity struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Priority    int    `json:"priority" db:"priority"`
	StartDate   string `json:"start_date" db:"start_date"`     // หรือใช้ time.Time
	EndDate     string `json:"end_date" db:"end_date"`         // หรือใช้ time.Time
	SectionID   int    `json:"section_id" db:"section_id"`     // foreign key จาก Section
	WorkspaceID string `json:"workspace_id" db:"workspace_id"` // foreign key จาก Workspace
}
