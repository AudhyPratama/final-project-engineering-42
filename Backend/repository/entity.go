package repository

type User struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Token    string `db:"token"`
}
