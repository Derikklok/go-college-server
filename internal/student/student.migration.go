package student

import (
	"log"

	"github.com/Derikklok/go-college-server/internal/database"
)

func Migrate() {
	err := database.DB.AutoMigrate(&Student{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}
