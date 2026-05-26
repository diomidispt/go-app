package service

import (
	"github.com/diomidispt/go-app/internal/model"
	"github.com/diomidispt/go-app/internal/repository"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) GetAll() ([]model.Patient, error) {
	return s.repo.GetAll()
}

func (s *PatientService) GetByID(id int) (model.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *PatientService) Create(p model.Patient) (model.Patient, error) {
	return s.repo.Create(p)
}

func (s *PatientService) Update(p model.Patient) error {
	return s.repo.Update(p)
}

func (s *PatientService) Delete(id int) error {
	return s.repo.Delete(id)
}
