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

func (p *ProductRepository) GetAll() ([]*Book, error) {
	var sqlStatement string

	sqlStatement = `SELECT * FROM book ORDER BY id DESC`

	rows, err := p.db.Query(sqlStatement)
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

	return books, nil
}

func (p *ProductRepository) CreateProductBook(book *Book) error {
	var sqlStatement string

	sqlStatement = `INSERT INTO book (category_id, title, penulis, berat, stock, price, description) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := p.db.Exec(sqlStatement, book.CategoryID, book.Title, book.Penulis, book.Berat, book.Stock, book.Price, book.Description)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateProductBook(book *Book) error {
	var sqlStatement string

	sqlStatement = `UPDATE book SET category_id = ?, title = ?, penulis = ?, berat = ?, stock = ?, price = ?, description = ? WHERE id = ?`
	_, err := p.db.Exec(sqlStatement, book.CategoryID, book.Title, book.Penulis, book.Berat, book.Stock, book.Price, book.Description, book.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) DeleteProductBook(id int64) error {
	var sqlStatement string

	sqlStatement = `DELETE FROM book WHERE id = ?`
	_, err := p.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) FetchBookByID(id int64) (*DetailBook, error) {
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
	FROM book b
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
	FROM book b
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
