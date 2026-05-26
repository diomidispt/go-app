package service

import (
	"errors" // used to create a custom error message

	"github.com/diomidispt/go-app/internal/model"
	"github.com/diomidispt/go-app/internal/repository"
)

type PrescriptionService struct {
	repo         *repository.PrescriptionRepository
	medicineRepo *repository.MedicineRepository // needed to check stock before creating a prescription
}

func NewPrescriptionService(repo *repository.PrescriptionRepository, medicineRepo *repository.MedicineRepository) *PrescriptionService {
	return &PrescriptionService{repo: repo, medicineRepo: medicineRepo}
}

func (s *PrescriptionService) GetAll() ([]model.Prescription, error) {
	return s.repo.GetAll()
}

func (s *PrescriptionService) GetByID(id int) (model.Prescription, error) {
	return s.repo.GetByID(id)
}

// Create checks that the medicine has stock before creating the prescription
// this is the business logic that lives in the service layer — not in the handler, not in the repository
func (s *PrescriptionService) Create(p model.Prescription) (model.Prescription, error) {
	medicine, err := s.medicineRepo.GetByID(p.MedicineID)
	if err != nil {
		return p, errors.New("medicine not found")
	}

	if medicine.Stock == 0 {
		return p, errors.New("cannot create prescription — medicine is out of stock")
	}

	return s.repo.Create(p)
}

func (s *PrescriptionService) Update(p model.Prescription) error {
	return s.repo.Update(p)
}

func (s *PrescriptionService) Delete(id int) error {
	return s.repo.Delete(id)
}
