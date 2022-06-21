package repository

import (
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProducts() ([]*Book, error) {
	var sqlStatement string

	sqlStatement = `SELECT title, penulis, price FROM books`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Title, &book.Penulis, &book.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (p *ProductRepository) FetchBookByID(id int64) (*DetailBook, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT
		b.id,
		b.title,
		c.name as category_name,
		b.penulis,
		b.berat,
		b.stock,
		b.price,
		b.description
	FROM books b
	LEFT JOIN category c ON b.category_id = c.id
	WHERE b.id = ?`

	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*DetailBook
	for rows.Next() {
		var book DetailBook
		err := rows.Scan(
			&book.BookID,
			&book.Title,
			&book.CategoryName,
			&book.Penulis,
			&book.Berat,
			&book.Stock,
			&book.Price,
			&book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books[0], nil
}

func (p *ProductRepository) FetchBookByName(name string) (*DetailBook, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT
		b.id,
		c.name as category_name,
		b.title,
		b.penulis,
		b.berat,
		b.stock,
		b.price,
		b.description
	FROM books b
	LEFT JOIN category c ON b.category_id = c.id
	WHERE b.title = ?`

	rows, err := p.db.Query(sqlStatement, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*DetailBook
	for rows.Next() {
		var book DetailBook
		err := rows.Scan(
			&book.BookID,
			&book.CategoryName,
			&book.Title,
			&book.Penulis,
			&book.Berat,
			&book.Stock,
			&book.Price,
			&book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books[0], nil
}

func (p *ProductRepository) FetchBookByPenulis(penulis string) (*DetailBook, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT
		b.id,
		c.name as category_name,
		b.title,
		b.penulis,
		b.berat,
		b.stock,
		b.price,
		b.description
	FROM books b
	LEFT JOIN category c ON b.category_id = c.id
	WHERE b.penulis = ?`

	rows, err := p.db.Query(sqlStatement, penulis)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*DetailBook
	for rows.Next() {
		var book DetailBook
		err := rows.Scan(
			&book.BookID,
			&book.CategoryName,
			&book.Title,
			&book.Penulis,
			&book.Berat,
			&book.Stock,
			&book.Price,
			&book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books[0], nil
}
