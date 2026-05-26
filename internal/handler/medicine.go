package handler

import (
	"encoding/json" // converts Go structs to JSON and JSON to Go structs
	"net/http"      // gives us the request and response writer types
	"strconv"       // converts strings to other types (e.g. "1" → 1)

	"github.com/go-chi/chi/v5"                          // router — lets us extract values from the URL like /medicines/{id}
	"github.com/diomidispt/go-app/internal/model"      // the Medicine struct
	"github.com/diomidispt/go-app/internal/service"    // the service layer we call for business logic
)

// MedicineHandler holds the service it will call
type MedicineHandler struct {
	svc *service.MedicineService
}

// NewMedicineHandler creates a new handler with the service
func NewMedicineHandler(svc *service.MedicineService) *MedicineHandler {
	return &MedicineHandler{svc: svc}
}

// RegisterRoutes wires all medicine endpoints to the router
func (h *MedicineHandler) RegisterRoutes(r chi.Router) {
	r.Get("/medicines", h.GetAll)
	r.Get("/medicines/{id}", h.GetByID)
	r.Post("/medicines", h.Create)
	r.Put("/medicines/{id}", h.Update)
	r.Delete("/medicines/{id}", h.Delete)
}

// GetAll handles GET /medicines — returns all medicines as JSON
func (h *MedicineHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	medicines, err := h.svc.GetAll()
	if err != nil {
		http.Error(w, "failed to get medicines", http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Content-Type", "application/json") // tell the caller the response is JSON
	json.NewEncoder(w).Encode(medicines)               // convert the slice of medicines to JSON and write it to the response
}

// GetByID handles GET /medicines/{id} — returns one medicine by ID
func (h *MedicineHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id")) // extract "id" from the URL and convert it from string to int
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest) // 400 — the ID they gave us is not a number
		return
	}

	medicine, err := h.svc.GetByID(id)
	if err != nil {
		http.Error(w, "medicine not found", http.StatusNotFound) // 404
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicine)
}

// Create handles POST /medicines — creates a new medicine
func (h *MedicineHandler) Create(w http.ResponseWriter, r *http.Request) {
	var m model.Medicine
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil { // read the JSON body from the request and fill a Medicine struct
		http.Error(w, "invalid request body", http.StatusBadRequest) // 400 — the JSON they sent is malformed
		return
	}

	created, err := h.svc.Create(m)
	if err != nil {
		http.Error(w, "failed to create medicine", http.StatusInternalServerError) // 500
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)  // 201 — resource was created (not 200)
	json.NewEncoder(w).Encode(created) // send back the created medicine including the generated ID
}

// Update handles PUT /medicines/{id} — updates an existing medicine
func (h *MedicineHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest) // 400
		return
	}

	var m model.Medicine
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest) // 400
		return
	}
	m.ID = id // make sure we update the correct record — ID comes from the URL, not the body

	if err := h.svc.Update(m); err != nil {
		http.Error(w, "failed to update medicine", http.StatusInternalServerError) // 500
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 — success, nothing to return
}

// Delete handles DELETE /medicines/{id} — removes a medicine
func (h *MedicineHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest) // 400
		return
	}

	if err := h.svc.Delete(id); err != nil {
		http.Error(w, "failed to delete medicine", http.StatusInternalServerError) // 500
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 — success, nothing to return
}
