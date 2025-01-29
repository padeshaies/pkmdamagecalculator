package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	// Add Handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Pokemon World!"))
	})
	mux.HandleFunc("/api/calc", handleGetCalc)

	// Start Server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
