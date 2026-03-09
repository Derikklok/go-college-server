package main

import (
	"flag"
	"log"

	"github.com/Derikklok/go-college-server/internal/config"
	"github.com/Derikklok/go-college-server/internal/database"

	// Import student package to register migrations
	"github.com/Derikklok/go-college-server/internal/migrations"
	_ "github.com/Derikklok/go-college-server/internal/student"
)

func main() {
	flag.Parse()

	log.Println("Migration Tool - College Server")
	log.Println("================================")

	// Load configuration
	config.Load()

	// Connect to database
	database.Connect()
	log.Println("Database connected successfully")

	// Run all registered migrations
	if err := migrations.RunAll(database.DB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("================================")
	log.Println("Migration completed successfully!")
}
