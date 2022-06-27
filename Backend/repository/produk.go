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

	sqlStatement := `
	SELECT
		b.id,
		b.categori_id,
		c.name as category_name,
		b.book_name,
		b.penulis,
		b.penerbit,
		b.kondisi,
		b.berat,
		b.stock,
		b.harga,
		b.deskripsi,
		b.image
	FROM books b
	LEFT JOIN categories c ON b.categori_id = c.id`

	rows, err := b.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		var book Book
		
		err := rows.Scan(
			&book.ID,
			&book.CategoryID,
			&book.CategoryName,
			&book.BookName,
			&book.Penulis,
			&book.Penerbit,
			&book.Kondisi,
			&book.Berat,
			&book.Stock,
			&book.Harga,
			&book.Deskripsi,
			&book.Image)
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

	sqlStatement = `SELECT 
		b.id,
		b.categori_id,
		c.name as category_name,
		b.book_name,
		b.penulis,
		b.penerbit,
		b.kondisi,
		b.berat,
		b.stock,
		b.harga,
		b.deskripsi,
		b.image
	 FROM books b
	 INNER JOIN categories c ON b.categori_id = c.id
	 WHERE book_name LIKE '%` + request.BookName + `%'`

	if isValidRequest := request.IsValidRequest(); !isValidRequest {
		return nil, fmt.Errorf("bad request")
	}

	if request.BookName != "" {
		sqlStatement += ` AND book_name LIKE '%` + request.BookName + `%'`
	}

	if request.Penulis != "" {
		sqlStatement += ` AND penulis LIKE '%` + request.Penulis + `%'`
	}

	if request.Penerbit != "" {
		sqlStatement += ` AND penerbit LIKE '%` + request.Penerbit + `%'`
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
			&book.CategoryID,
			&book.CategoryName,
			&book.BookName,
			&book.Penulis,
			&book.Penerbit,
			&book.Kondisi,
			&book.Berat,
			&book.Stock,
			&book.Harga,
			&book.Deskripsi,
			&book.Image)
		// &book.CategoryName)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepository) InsertBook(categori_id int, book_name string, penulis string, penerbit string, kondisi string, berat string, stock int, harga int, deskripsi string, image string) error {
	sqlStatement := `INSERT INTO books (categori_id, book_name, penulis, penerbit, kondisi, berat, stock, harga, deskripsi, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := b.db.Exec(sqlStatement, categori_id, book_name, penulis, penerbit, kondisi, berat, stock, harga, deskripsi, image)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) UpdateBook(id int, categori_id int, book_name string, penulis string, penerbit string, kondisi string, berat string, stock int, harga int, deskripsi string, image string) error {
	sqlStatement := `UPDATE books SET categori_id = ?, book_name = ?, penulis = ?, penerbit = ?, kondisi = ?, berat = ?, stock = ?, harga = ?, deskripsi = ?, image = ? WHERE id = ?`
	_, err := b.db.Exec(sqlStatement, categori_id, book_name, penulis, penerbit, kondisi, berat, stock, harga, deskripsi, image, id)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) DeleteBook(id int) error {
	sqlStatement := `DELETE FROM books WHERE id = ?`
	_, err := b.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
