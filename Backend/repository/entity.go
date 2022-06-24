package repository

import "time"

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
	CategoryID   int64  `db:"category_id"`
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

type Payment struct {
	ID               int64      `db:"id"`
	OrderID          int64      `db:"order_id"`
	UserID           int64      `db:"user_id"`
	MetodePembayaran string     `db:"metode_pembayaran"`
	Alamat           string     `db:"alamat"`
	OngkosKirim      int64      `db:"ongkos_kirim"`
	TotalPayment     int64      `db:"total_payment"`
	TimePayment      *time.Time `db:"time_payment"`
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
