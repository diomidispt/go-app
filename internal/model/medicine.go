package model

import "time"

// Medicine mirrors the medicines table in the database
// each field maps to a column — same names, same types
type Medicine struct {
	ID        int        `json:"id"`         // maps to the id column
	Name      string     `json:"name"`       // maps to the name column
	Dosage    string     `json:"dosage"`     // maps to the dosage column
	Stock     int        `json:"stock"`      // maps to the stock column
	Price     float64    `json:"price"`      // maps to the price column
	DeletedAt *time.Time `json:"deleted_at"` // null means active — a timestamp means soft deleted, * means it can be nil
}
