package user

import (
	"database/sql"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) CreateUser(user *User) error {
	// Example using SQL:
	_, err := s.db.Exec("INSERT INTO users (name, email,phone_no,address) VALUES ($1, $2,$3,$4)", user.Name, user.Email, user.PhoneNo, user.Address)
	if err != nil {
		return err
	}
	return nil
}
