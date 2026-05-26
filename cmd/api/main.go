package main

import (
	"encoding/json"
	"log/slog"   // structured logging — built into Go 1.21+, no extra package needed
	"net/http"
	"os"

	"github.com/diomidispt/go-app/internal/app"
	"github.com/diomidispt/go-app/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := config.Connect()
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("connected to database")

	if err := config.RunMigrations(); err != nil {
		slog.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}
	slog.Info("migrations applied")

	a := app.New(db)

	// health check — hits the DB so the load balancer knows if the app is truly healthy
	a.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable) // 503 — tells ALB/ECS this instance is unhealthy
			json.NewEncoder(w).Encode(map[string]string{"status": "unhealthy", "error": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})

	slog.Info("server starting", "port", port)
	if err := http.ListenAndServe(":"+port, a.Router); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
