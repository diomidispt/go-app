package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/diomidispt/go-app/internal/app"    // owns all wiring — repos, services, handlers, routes
	"github.com/diomidispt/go-app/internal/config" // opens the DB connection
	"github.com/joho/godotenv"                      // loads .env file into environment variables
)

func main() {
	godotenv.Load() // load .env — in Docker/ECS vars are already injected so this is safely ignored

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := config.Connect()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	fmt.Println("Connected to database successfully")

	if err := config.RunMigrations(); err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Migrations applied successfully")

	a := app.New(db) // all wiring happens inside here — main.go never needs to change again

	a.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "our go application is working")
	})

	fmt.Println("Server starting on port " + port + "...")
	http.ListenAndServe(":"+port, a.Router)
}
