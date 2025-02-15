package main

import (
	"log"
	"user-service/app"
	"user-service/middleware"
	"user-service/store"
)

func main() {
	// Connect to PostgreSQL
	db := middleware.CreateConnection()
	defer db.Close()

	// Initialize PostgresStore
	postgresStore := store.NewPostgresStore(db)

	// Start the server
	server := app.NewServer(postgresStore)

	log.Println("Server running on :8080")
	server.Router.Run(":8080")
}
