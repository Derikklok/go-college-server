package student

import "time"

type Student struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
