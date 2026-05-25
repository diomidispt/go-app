package main // every Go file belongs to a package. "main" is special — it means this is the entry point of the program

import (
	"fmt"      // used for printing to the terminal and writing responses
	"net/http" // Go's built-in HTTP server — no extra install needed
	"os"       // used to read environment variables (like PORT)

	"github.com/diomidispt/go-app/internal/config" // our own config package — contains the Connect() function
	"github.com/joho/godotenv"                      // loads .env file into environment variables automatically
)

func main() { // Go always starts here — this is the first function that runs
	godotenv.Load() // load .env file into environment variables — works locally; in Docker/ECS the vars are already injected so this is safely ignored

	port := os.Getenv("PORT") // read the PORT environment variable (e.g. export PORT=9090)
	if port == "" {           // if no PORT is set, fall back to 8080
		port = "8080"
	}

	// step 1 — connect to the database before starting the server
	db, err := config.Connect() // Connect() returns two values: the connection and an error
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1) // if we can't reach the database, there's no point starting the server
	}
	defer db.Close() // when main() exits, close the DB connection cleanly

	fmt.Println("Connected to database successfully")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { // register a route: when someone visits /health, run this function
		fmt.Fprintln(w, "our go application is working") // write "ok" back to whoever made the request (w = response writer)
	})

	fmt.Println("Server starting on port " + port + "...") // print to terminal so you know the server started
	http.ListenAndServe(":"+port, nil)                      // this blocks forever — nothing below this line will ever run
}
