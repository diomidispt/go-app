package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/diomidispt/go-app/internal/model"
	"github.com/diomidispt/go-app/internal/service"
)

type PatientHandler struct {
	svc *service.PatientService
}

func NewPatientHandler(svc *service.PatientService) *PatientHandler {
	return &PatientHandler{svc: svc}
}

func (h *PatientHandler) RegisterRoutes(r chi.Router) {
	r.Get("/patients", h.GetAll)
	r.Get("/patients/{id}", h.GetByID)
	r.Post("/patients", h.Create)
	r.Put("/patients/{id}", h.Update)
	r.Delete("/patients/{id}", h.Delete)
}

func (h *PatientHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	patients, err := h.svc.GetAll()
	if err != nil {
		http.Error(w, "failed to get patients", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

func (h *PatientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	patient, err := h.svc.GetByID(id)
	if err != nil {
		http.Error(w, "patient not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

func (h *PatientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p model.Patient
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	created, err := h.svc.Create(p)
	if err != nil {
		http.Error(w, "failed to create patient", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *PatientHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var p model.Patient
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	p.ID = id
	if err := h.svc.Update(p); err != nil {
		http.Error(w, "failed to update patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *PatientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.svc.Delete(id); err != nil {
		http.Error(w, "failed to delete patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
