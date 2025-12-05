package users

import (
	"github.com/jmoiron/sqlx"
)

type SQLStorage struct {
	db *sqlx.DB
}

func NewSQLStorage(db *sqlx.DB) *SQLStorage {
	return &SQLStorage{db: db}
}

func (s *SQLStorage) Create(u User) (User, error) {
	var id int

	query := `
        INSERT INTO users (name, email, age)
        VALUES ($1, $2, $3)
        RETURNING id;
 	`

	err := s.db.QueryRow(query, u.Name, u.Email, u.Age).Scan(&id)
	if err != nil {
		return User{}, err
	}
	u.ID = id
	return u, nil
}
func (s *SQLStorage) GetAll() ([]User, error) {
	var users []User
	query := `
	SELECT id, name, email, age
	FROM users;
	`
	err := s.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (s *SQLStorage) GetByID(id int) (User, error) {
	var user User
	query := `
	SELECT id, name, email, age
	FROM users
	WHERE id = $1;
	`
	if err := s.db.Get(&user, query, id); err != nil {
		return User{}, err
	}
	return user, nil
}
func (s *SQLStorage) Update(u User) (User, error) {
	query := `
	UPDATE users
	SET name = $1, email = $2, age = $3
	WHERE id = $4
	`

	_, err := s.db.Exec(query, u.Name, u.Email, u.Age, u.ID)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *SQLStorage) Delete(id int) error {
	query := `
        DELETE FROM users
        WHERE id = $1;
    `

	_, err := s.db.Exec(query, id)
	return err
}
