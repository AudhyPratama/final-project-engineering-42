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
	ID           int64  `db:"id"`
	CategoryID   int64  `db:"category_id"`
	CategoryName string `db:"category_name"`
	BookName     string `db:"book_name"`
	Penulis      string `db:"penulis"`
	Penerbit     string `db:"penerbit"`
	Kondisi      string `db:"kondisi"`
	Berat        string `db:"berat"`
	Stock        int64  `db:"stock"`
	Harga        int64  `db:"harga"`
	Deskripsi    string `db:"deskripsi"`
}

type Category struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type OrderCart struct {
	OrderID      int64  `db:"order_id"`
	BookID       int64  `db:"book_id"`
	BookName     string `db:"book_name"`
	CategotyID   int64  `db:"category_id"`
	CategoryName string `db:"category_name"`
	Penulis      string `db:"penulis"`
	Penerbit     string `db:"penerbit"`
	Quantity     int64  `db:"quantity"`
	Harga        int64  `db:"harga"`
}

type GetBookRequest struct {
	BookName string `json:"book_name"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}

type GetCartRequest struct {
	OrderID int64 `json:"order_id"`
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

func (r GetCartRequest) IsEmptyRequest() bool {
	if r.OrderID == 0 {
		return true
	}

	return false
}

func (r GetCartRequest) IsValidRequest() bool {
	if r.OrderID == 0 {
		return false
	}
	if r.OrderID != 0 {
		return false
	}

	return true
}
