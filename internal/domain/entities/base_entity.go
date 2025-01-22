package entities

import "time"

// BaseEntity is global base entities
type BaseEntity struct {
	CreatedAt time.Time `json:"created_at" example:"2025-01-01 00:00:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-01-01 00:00:00"`
	DeletedAt time.Time `json:"deleted_at" example:"2025-01-01 00:00:00"`
}
