package student

import "github.com/Derikklok/go-college-server/internal/database"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(student *Student) error {
	return database.DB.Create(student).Error
}

func (r *Repository) FindAll() ([]Student, error) {

	var students []Student

	err := database.DB.Find(&students).Error

	return students, err
}

func (r *Repository) FindByID(id uint) (*Student, error) {

	var student Student

	err := database.DB.First(&student, id).Error

	return &student, err
}

func (r *Repository) Delete(id uint) error {

	return database.DB.Delete(&Student{}, id).Error
}
