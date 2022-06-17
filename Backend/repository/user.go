package repository

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var users []User

	rows, err := u.db.Query(`SELECT  email, password FROM users`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Email,
			&user.Password,
		); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) Login(email string, password string) (*string, error) {
	var user User

	err := u.db.QueryRow(`SELECT email FROM users WHERE email = ? AND password = ? `, email, password).Scan(
		&user.Email,
	)

	if err != nil {
		return nil, errors.New("Login Failed")
	}

	return &user.Email, nil
}

func (u *UserRepository) FetchUserRole(email string) (*string, error) {
	var user User

	err := u.db.QueryRow(`SELECT role FROM users WHERE email = ?`, email).Scan(&user.Role)
	if err != nil {
		return nil, errors.New("Role not found")
	}
	return &user.Role, nil
}
