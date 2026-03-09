package student

import (
	"gorm.io/gorm"
)

// StudentMigration handles student model migrations
type StudentMigration struct{}

func (sm *StudentMigration) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Student{})
}

// NewStudentMigration creates a new student migration
func NewStudentMigration() *StudentMigration {
	return &StudentMigration{}
}
