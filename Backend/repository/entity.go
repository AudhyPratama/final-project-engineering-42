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
	ID           int64   `db:"id"`
	CategoryID   int64   `db:"category_id"`
	CategoryName string  `db:"category_name"`
	BookName     string  `db:"book_name"`
	Penulis      string  `db:"penulis"`
	Penerbit     string  `db:"penerbit"`
	Kondisi      string  `db:"kondisi"`
	Berat        string  `db:"berat"`
	Stock        int64   `db:"stock"`
	Harga        float64 `db:"harga"`
	Deskripsi    string  `db:"deskripsi"`
}

type Category struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type OrderCart struct {
	ID           int64   `db:"id"`
	BookID       int64   `db:"book_id"`
	BookName     string  `db:"book_name"`
	CategotyID   int64   `db:"category_id"`
	CategoryName string  `db:"category_name"`
	Penulis      string  `db:"penulis"`
	Penerbit     string  `db:"penerbit"`
	berat        string  `db:"berat"`
	Quantity     int64   `db:"quantity"`
	Harga        float64 `db:"harga"`
}

type GetBookRequest struct {
	BookName string `json:"book_name"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}

func (r GetBookRequest) IsEmptyRequest() bool {
	if r.BookName == "" && r.Penulis == "" && r.Penerbit == "" {
		return true
	}

	return false
}

func (r GetBookRequest) IsValidRequest() bool {
	if r.BookName == "" && r.Penulis == "" && r.Penerbit == "" {
		return false
	}
	if r.BookName != "" && r.Penulis != "" && r.Penerbit != "" {
		return false
	}

	return true
}
