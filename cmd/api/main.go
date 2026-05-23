package main // every Go file belongs to a package. "main" is special — it means this is the entry point of the program

import (
	"fmt"      // used for printing to the terminal and writing responses
	"net/http" // Go's built-in HTTP server — no extra install needed
	"os"       // used to read environment variables (like PORT)
)

func main() { // Go always starts here — this is the first function that runs
	port := os.Getenv("PORT") // read the PORT environment variable (e.g. export PORT=9090)
	if port == "" {           // if no PORT is set, fall back to 8080
		port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { // register a route: when someone visits /health, run this function
		fmt.Fprintln(w, "ok") // write "ok" back to whoever made the request (w = response writer)
	})

	fmt.Println("Server starting on port " + port + "...") // print to terminal so you know the server started
	http.ListenAndServe(":"+port, nil)                      // start the server on the configured port and keep it running forever
}
