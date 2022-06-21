package repository

import (
	"database/sql"
	"fmt"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) FetchBooks() ([]*Book, error) {
	var sqlStatement string

	sqlStatement = `SELECT book_name, penulis, penerbit, harga FROM books`

	rows, err := b.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.BookName, &book.Penulis, &book.Penerbit, &book.Harga)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (b *BookRepository) FetchBooksByID(id int64) (Book, error) {
	var book Book

	err := b.db.QueryRow("SELECT book_name, penulis, penerbit, harga FROM books WHERE id = ?", id).Scan(&book.BookName, &book.Penulis, &book.Penerbit, &book.Harga)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *BookRepository) FetchBooksByName(book_name string) (Book, error) {
	var book Book

	err := b.db.QueryRow("SELECT book_name, penulis, penerbit, harga FROM books WHERE book_name = ?", book_name).Scan(&book.BookName, &book.Penulis, &book.Penerbit, &book.Harga)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *BookRepository) FetchDetailBook(request GetBookRequest) ([]Book, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT
		b.id,
		b.book_name,
		b.penulis,
		b.penerbit,
		b.kondisi,
		b.berat,
		b.harga,
		b.stock,
		b.deskripsi,
		c.category_id,
		c.category_name
	FROM books b
	INNER JOIN categories c ON b.category_id = c.category_id`

	if isValidRequest := request.IsValidRequest(); !isValidRequest {
		return nil, fmt.Errorf("Bad Request")
	}

	if request.BookName != "" {
		sqlStatement = fmt.Sprintf("%s AND p.book_name = ?", sqlStatement)
	}

	if request.Penulis != "" {
		sqlStatement = fmt.Sprintf("%s AND p.penulis = ?", sqlStatement)
	}

	if request.Penerbit != "" {
		sqlStatement = fmt.Sprintf("%s AND p.penerbit = ?", sqlStatement)
	}

	rows, err := b.db.Query(sqlStatement, request.BookName, request.Penulis, request.Penerbit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.BookName,
			&book.Penulis,
			&book.Penerbit,
			&book.Kondisi,
			&book.Berat,
			&book.Harga,
			&book.Stock,
			&book.Deskripsi,
			&book.CategoryID,
			&book.CategoryName)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
