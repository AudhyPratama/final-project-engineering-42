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
		return "", errors.New("user not found")
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
		return nil, errors.New("user not found")
	}
	return &user.Role, nil
}

func (u *UserRepository) Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (u *UserRepository) UpdateUser(email string, name string, password string) (Profile, error) {
	var user Profile

	_, err := u.db.Exec(`UPDATE users SET name = ?, password = ? WHERE email = ?`, name, password, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) UpdatePassword(email string, password string) error {
	_, err := u.db.Exec(`UPDATE users SET password = ? WHERE email = ?`, password, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetProfile(email string) (Profile, error) {
	var user Profile

	err := u.db.QueryRow(`SELECT name, email, password, role FROM users WHERE email = ?`, email).Scan(
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return user, errors.New("profile not found")
	}
	return user, nil
}
