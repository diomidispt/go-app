package model

import "time"

// Patient mirrors the patients table in the database
type Patient struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	DateOfBirth string     `json:"date_of_birth"` // stored as string (YYYY-MM-DD) for simplicity
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	DeletedAt   *time.Time `json:"deleted_at"` // null means active, a timestamp means soft deleted
}
