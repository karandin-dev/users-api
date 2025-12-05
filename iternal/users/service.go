package users

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{storage: s}
}

// CREATE
func (s *Service) Create(name, email string, age int) (User, error) {
	return s.storage.Create(User{
		Name:  name,
		Email: email,
		Age:   age,
	})
}

// GET ALL
func (s *Service) GetAll() ([]User, error) {
	return s.storage.GetAll()
}

// GET BY ID
func (s *Service) GetByID(id int) (User, error) {
	return s.storage.GetByID(id)
}

// UPDATE (partial update)
func (s *Service) Update(id int, input UpdateUserInput) (User, error) {
	user, err := s.storage.GetByID(id)
	if err != nil {
		return User{}, err
	}

	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Age != nil {
		user.Age = *input.Age
	}

	return s.storage.Update(user)
}

// DELETE
func (s *Service) Delete(id int) error {
	return s.storage.Delete(id)
}
