package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jmcknight17/band-finder-api/internal/db"
	"github.com/jmcknight17/band-finder-api/internal/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file") //If the .Env file cannot be found log the error
	}

	//Connecting to the mongoDb
	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Initialize router
	r := mux.NewRouter()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(client)

	// Routes
	r.HandleFunc("/api/register", authHandler.Register).Methods("POST")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
