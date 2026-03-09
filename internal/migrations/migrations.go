package migrations

import (
	"log"

	"gorm.io/gorm"
)

// Migrator interface for models that need database migrations
type Migrator interface {
	Migrate(db *gorm.DB) error
}

var registeredMigrations []Migrator

// Register adds a migration to the list
func Register(m Migrator) {
	registeredMigrations = append(registeredMigrations, m)
}

// RunAll executes all registered migrations
func RunAll(db *gorm.DB) error {
	log.Println("Running database migrations...")

	for i, migration := range registeredMigrations {
		log.Printf("Running migration %d/%d...\n", i+1, len(registeredMigrations))
		if err := migration.Migrate(db); err != nil {
			log.Fatalf("Migration failed: %v", err)
			return err
		}
	}

	log.Println("All migrations completed successfully")
	return nil
}
