package student

import (
	"github.com/Derikklok/go-college-server/internal/database"
	"github.com/Derikklok/go-college-server/internal/migrations"
)

func init() {
	// Register student migration on package init
	migrations.Register(NewStudentMigration())
}

// Migrate executes all registered migrations (kept for backward compatibility)
func Migrate() {
	migrations.RunAll(database.DB)
}
