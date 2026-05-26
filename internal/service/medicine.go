package service

import (
	"github.com/diomidispt/go-app/internal/model"      // the Medicine struct
	"github.com/diomidispt/go-app/internal/repository" // the repository that runs SQL queries
)

// MedicineService holds a reference to the repository
// all business logic for medicines lives here
type MedicineService struct {
	repo *repository.MedicineRepository // the repository this service will call
}

// NewMedicineService creates a new service with the repository
func NewMedicineService(repo *repository.MedicineRepository) *MedicineService {
	return &MedicineService{repo: repo}
}

// GetAll returns all medicines — no business rules, pass straight through to the repository
func (s *MedicineService) GetAll() ([]model.Medicine, error) {
	return s.repo.GetAll()
}

// GetByID returns one medicine by ID — no business rules, pass straight through
func (s *MedicineService) GetByID(id int) (model.Medicine, error) {
	return s.repo.GetByID(id)
}

// Create adds a new medicine — no business rules for medicines, pass straight through
func (s *MedicineService) Create(m model.Medicine) (model.Medicine, error) {
	return s.repo.Create(m)
}

// Update modifies an existing medicine — no business rules, pass straight through
func (s *MedicineService) Update(m model.Medicine) error {
	return s.repo.Update(m)
}

// Delete removes a medicine by ID — no business rules, pass straight through
func (s *MedicineService) Delete(id int) error {
	return s.repo.Delete(id)
}
