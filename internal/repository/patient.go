package repository

import (
	"database/sql"

	"github.com/diomidispt/go-app/internal/model"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetAll() ([]model.Patient, error) {
	rows, err := r.db.Query("SELECT id, name, date_of_birth, phone, email FROM patients WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []model.Patient
	for rows.Next() {
		var p model.Patient
		rows.Scan(&p.ID, &p.Name, &p.DateOfBirth, &p.Phone, &p.Email)
		patients = append(patients, p)
	}
	return patients, nil
}

func (r *PatientRepository) GetByID(id int) (model.Patient, error) {
	var p model.Patient
	row := r.db.QueryRow("SELECT id, name, date_of_birth, phone, email FROM patients WHERE id = $1 AND deleted_at IS NULL", id)
	err := row.Scan(&p.ID, &p.Name, &p.DateOfBirth, &p.Phone, &p.Email)
	return p, err
}

func (r *PatientRepository) Create(p model.Patient) (model.Patient, error) {
	err := r.db.QueryRow(
		"INSERT INTO patients (name, date_of_birth, phone, email) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.DateOfBirth, p.Phone, p.Email,
	).Scan(&p.ID)
	return p, err
}

func (r *PatientRepository) Update(p model.Patient) error {
	_, err := r.db.Exec(
		"UPDATE patients SET name=$1, date_of_birth=$2, phone=$3, email=$4 WHERE id=$5 AND deleted_at IS NULL",
		p.Name, p.DateOfBirth, p.Phone, p.Email, p.ID,
	)
	return err
}

func (r *PatientRepository) Delete(id int) error {
	_, err := r.db.Exec("UPDATE patients SET deleted_at = NOW() WHERE id = $1", id)
	return err
}
