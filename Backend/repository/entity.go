package repository

type User struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Token    string `db:"token"`
}

type Book struct {
	ID          int64   `db:"id"`
	UserID      int64   `db:"user_id"`
	CategoryID  int64   `db:"category_id"`
	Title       string  `db:"title"`
	Penulis     string  `db:"penulis"`
	Berat       int64   `db:"berat"`
	Stock       int64   `db:"stock"`
	Price       float64 `db:"price"`
	Description string  `db:"description"`
}

type DetailBook struct {
	BookID       int64   `db:"book_id"`
	CategoryName string  `db:"category_name"`
	Title        string  `db:"title"`
	Penulis      string  `db:"penulis"`
	Berat        int64   `db:"berat"`
	Stock        int64   `db:"stock"`
	Price        float64 `db:"price"`
	Description  string  `db:"description"`
}

type Categori struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type ItemCart struct {
	ID           int64   `db:"id"`
	UserID       int64   `db:"user_id"`
	BookID       int64   `db:"book_id"`
	Title        string  `db:"title"`
	CategoryName string  `db:"category_name"`
	Penulis      string  `db:"penulis"`
	Quantity     int64   `db:"quantity"`
	Price        float64 `db:"price"`
}
