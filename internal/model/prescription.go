package model

import "time"

// Prescription mirrors the prescriptions table in the database
type Prescription struct {
	ID           int        `json:"id"`
	PatientID    int        `json:"patient_id"`    // foreign key — links to patients table
	MedicineID   int        `json:"medicine_id"`   // foreign key — links to medicines table
	Date         string     `json:"date"`          // YYYY-MM-DD format
	Instructions string     `json:"instructions"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
