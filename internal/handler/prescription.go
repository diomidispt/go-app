package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/diomidispt/go-app/internal/model"
	"github.com/diomidispt/go-app/internal/service"
)

type PrescriptionHandler struct {
	svc *service.PrescriptionService
}

func NewPrescriptionHandler(svc *service.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{svc: svc}
}

func (h *PrescriptionHandler) RegisterRoutes(r chi.Router) {
	r.Get("/prescriptions", h.GetAll)
	r.Get("/prescriptions/{id}", h.GetByID)
	r.Post("/prescriptions", h.Create)
	r.Put("/prescriptions/{id}", h.Update)
	r.Delete("/prescriptions/{id}", h.Delete)
}

func (h *PrescriptionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	prescriptions, err := h.svc.GetAll()
	if err != nil {
		http.Error(w, "failed to get prescriptions", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prescriptions)
}

func (h *PrescriptionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	prescription, err := h.svc.GetByID(id)
	if err != nil {
		http.Error(w, "prescription not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prescription)
}

func (h *PrescriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p model.Prescription
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	created, err := h.svc.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity) // 422 — request was valid but business logic rejected it (e.g. out of stock)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *PrescriptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var p model.Prescription
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	p.ID = id
	if err := h.svc.Update(p); err != nil {
		http.Error(w, "failed to update prescription", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *PrescriptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.svc.Delete(id); err != nil {
		http.Error(w, "failed to delete prescription", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
