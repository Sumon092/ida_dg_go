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
	_, err := s.db.Exec("INSERT INTO users (name, email,phone_no,address) VALUES ($1, $2,$3,$4)", user.Name, user.Email, user.PhoneNo, user.Address)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetAllUsers() ([]*User, error) {
	rows, err := s.db.Query("SELECT id, name, email, phone_no, address FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNo, &user.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUser(userID int) (*User, error) {
	user := &User{}
	err := s.db.QueryRow("SELECT id, name, email, phone_no, address FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNo, &user.Address)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(userID int, updatedUser *User) error {
	_, err := s.db.Exec("UPDATE users SET name=$1, email=$2, phone_no=$3, address=$4 WHERE id=$5",
		updatedUser.Name, updatedUser.Email, updatedUser.PhoneNo, updatedUser.Address, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(userID int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		return err
	}
	return nil
}
