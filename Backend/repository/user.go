package repository

import (
	"database/sql"
	"errors"
	"net/mail"
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

func (u *UserRepository) FetchUsersByEmail(email string) (string, error) {
	var user User

	err := u.db.QueryRow(`SELECT email FROM users WHERE email = ?`, email).Scan(&user.Email)
	if err != nil {
		return "", errors.New("Email not found")
	}
	return user.Email, nil
}

func (u *UserRepository) Signup(name string, email string, password string, role string) error {
	_, err := u.db.Exec(`INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)`, name, email, password, role)
	if err != nil {
		return err
	}
	return nil
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

func (u *UserRepository) Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
