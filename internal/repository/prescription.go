package repository

import (
	"database/sql"

	"github.com/diomidispt/go-app/internal/model"
)

type PrescriptionRepository struct {
	db *sql.DB
}

func NewPrescriptionRepository(db *sql.DB) *PrescriptionRepository {
	return &PrescriptionRepository{db: db}
}

func (r *PrescriptionRepository) GetAll() ([]model.Prescription, error) {
	rows, err := r.db.Query("SELECT id, patient_id, medicine_id, date, instructions FROM prescriptions WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prescriptions []model.Prescription
	for rows.Next() {
		var p model.Prescription
		rows.Scan(&p.ID, &p.PatientID, &p.MedicineID, &p.Date, &p.Instructions)
		prescriptions = append(prescriptions, p)
	}
	return prescriptions, nil
}

func (r *PrescriptionRepository) GetByID(id int) (model.Prescription, error) {
	var p model.Prescription
	row := r.db.QueryRow("SELECT id, patient_id, medicine_id, date, instructions FROM prescriptions WHERE id = $1 AND deleted_at IS NULL", id)
	err := row.Scan(&p.ID, &p.PatientID, &p.MedicineID, &p.Date, &p.Instructions)
	return p, err
}

func (r *PrescriptionRepository) Create(p model.Prescription) (model.Prescription, error) {
	err := r.db.QueryRow(
		"INSERT INTO prescriptions (patient_id, medicine_id, date, instructions) VALUES ($1, $2, $3, $4) RETURNING id",
		p.PatientID, p.MedicineID, p.Date, p.Instructions,
	).Scan(&p.ID)
	return p, err
}

func (r *PrescriptionRepository) Update(p model.Prescription) error {
	_, err := r.db.Exec(
		"UPDATE prescriptions SET patient_id=$1, medicine_id=$2, date=$3, instructions=$4 WHERE id=$5 AND deleted_at IS NULL",
		p.PatientID, p.MedicineID, p.Date, p.Instructions, p.ID,
	)
	return err
}

func (r *PrescriptionRepository) Delete(id int) error {
	_, err := r.db.Exec("UPDATE prescriptions SET deleted_at = NOW() WHERE id = $1", id)
	return err
}
