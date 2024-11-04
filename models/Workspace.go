package models

type Workspace struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Owner       string `json:"owner" db:"owner"`
}
