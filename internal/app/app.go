package app

import (
	"database/sql" // the DB connection type
	"net/http"     // for the static file server

	"github.com/go-chi/chi/v5"                          // the router
	"github.com/diomidispt/go-app/internal/handler"    // HTTP handlers
	"github.com/diomidispt/go-app/internal/repository" // SQL queries
	"github.com/diomidispt/go-app/internal/service"    // business logic
)

// App holds the router — main.go only needs this to start the server
type App struct {
	Router chi.Router
}

// New wires all layers together and returns a ready-to-run App
// every new entity (patients, prescriptions) gets added here — main.go never changes
func New(db *sql.DB) *App {
	r := chi.NewRouter()

	// medicines
	medicineRepo := repository.NewMedicineRepository(db)
	medicineSvc := service.NewMedicineService(medicineRepo)
	medicineHandler := handler.NewMedicineHandler(medicineSvc)
	medicineHandler.RegisterRoutes(r)

	// patients
	patientRepo := repository.NewPatientRepository(db)
	patientSvc := service.NewPatientService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientSvc)
	patientHandler.RegisterRoutes(r)

	// prescriptions — needs medicineRepo to check stock in the service layer
	prescriptionRepo := repository.NewPrescriptionRepository(db)
	prescriptionSvc := service.NewPrescriptionService(prescriptionRepo, medicineRepo)
	prescriptionHandler := handler.NewPrescriptionHandler(prescriptionSvc)
	prescriptionHandler.RegisterRoutes(r)

	// serve frontend static files — must be last so API routes take priority
	r.Handle("/*", http.FileServer(http.Dir("./frontend")))

	return &App{Router: r}
}
