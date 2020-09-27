package model

import (
	"fmt"
)

// Todo model
type Todo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}

// Todos model
type Todos []Todo

// ToString will stringify the data
func (t *Todo) ToString() string {
	return fmt.Sprintf("Todo<id: %d, description: %s, done: %v, created_at: %v>", t.ID, t.Description, t.Done, t.CreatedAt)
}
