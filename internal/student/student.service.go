package student

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateStudent(req CreateStudentRequest) (*Student, error) {

	student := Student{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	err := s.repo.Create(&student)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Service) GetAllStudents() ([]Student, error) {

	return s.repo.FindAll()
}

func (s *Service) DeleteStudent(id uint) error {

	return s.repo.Delete(id)
}