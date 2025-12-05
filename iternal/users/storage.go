package users

import "errors"

var ErrNotFound = errors.New("user not found")

type MapStorage struct {
	data map[int]User
	id   int
}

type Storage interface {
	Create(User) (User, error)
	GetAll() ([]User, error)
	GetByID(int) (User, error)
	Update(User) (User, error)
	Delete(int) error
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		data: make(map[int]User),
		id:   1,
	}
}
func (s *MapStorage) Create(u User) User {
	u.ID = s.id
	s.id++
	s.data[u.ID] = u
	return u
}
func (s *MapStorage) GetAll() []User {
	dt := make([]User, 0, len(s.data))
	for _, v := range s.data {
		dt = append(dt, v)
	}
	return dt
}
func (s *MapStorage) GetByID(id int) (User, error) {
	u, ok := s.data[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}
func (s *MapStorage) Update(id int, updated User) (User, error) {
	_, ok := s.data[id]
	if !ok {
		return User{}, ErrNotFound
	}
	s.data[id] = updated
	return updated, nil
}
func (s *MapStorage) Delete(id int) error {
	if _, ok := s.data[id]; !ok {
		return ErrNotFound
	}
	delete(s.data, id)
	return nil
}
