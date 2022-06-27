package repository

import "time"

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Token    string `db:"token"`
}

type Profile struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type Book struct {
	ID           int     `db:"id"`
	CategoryID   int     `db:"category_id"`
	CategoryName *string `db:"category_name"`
	BookName     string  `db:"book_name"`
	Penulis      string  `db:"penulis"`
	Penerbit     string  `db:"penerbit"`
	Kondisi      string  `db:"kondisi"`
	Berat        string  `db:"berat"`
	Stock        int     `db:"stock"`
	Harga        int     `db:"harga"`
	Deskripsi    string  `db:"deskripsi"`
	Image        string  `db:"image"`
}

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type OrderCart struct {
	OrderID      int     `db:"order_id"`
	BookID       int     `db:"book_id"`
	BookName     string  `db:"book_name"`
	CategoryID   int     `db:"category_id"`
	CategoryName *string `db:"category_name"`
	Penulis      string  `db:"penulis"`
	Penerbit     string  `db:"penerbit"`
	Quantity     int     `db:"quantity"`
	Harga        int     `db:"harga"`
}

type GetBookRequest struct {
	BookName     string `json:"book_name"`
	Penulis      string `json:"penulis"`
	Penerbit     string `json:"penerbit"`
	CategoryName string `json:"category_name"`
}

type Payment struct {
	ID               int        `db:"id"`
	OrderID          int        `db:"order_id"`
	UserID           int        `db:"user_id"`
	MetodePembayaran string     `db:"metode_pembayaran"`
	Alamat           string     `db:"alamat"`
	OngkosKirim      int        `db:"ongkos_kirim"`
	TotalPayment     int        `db:"total_payment"`
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
