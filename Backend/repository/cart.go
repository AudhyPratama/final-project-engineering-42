package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) FetchCarts() ([]OrderCart, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT
		o.id,
		o.book_id,
		o.quantity,
		b.book_name,
		b.penulis,
		b.penerbit,
		b.price
	FROM orders o
	INNER JOIN book b ON o.book_id = b.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []OrderCart
	for rows.Next() {
		var cart OrderCart
		err := rows.Scan(
			&cart.ID,
			&cart.BookID,
			&cart.Quantity,
			&cart.BookName,
			&cart.Penulis,
			&cart.Penerbit,
			&cart.Harga)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

func (c *CartRepository) FetchCartByID(BookID int64) (OrderCart, error) {
	var orderCart OrderCart
	var sqlStatement string

	sqlStatement = `
	SELECT
		o.id,
		o.book_id,
		o.quantity,
		b.book_name,
		b.penulis,
		b.penerbit,
		b.price
	FROM orders o
	INNER JOIN book b ON o.book_id = b.id
	WHERE o.id = ?`

	row := c.db.QueryRow(sqlStatement, BookID)
	err := row.Scan(
		&orderCart.ID,
		&orderCart.BookID,
		&orderCart.Quantity,
		&orderCart.BookName,
		&orderCart.Penulis,
		&orderCart.Penerbit,
		&orderCart.Harga)
	if err != nil {
		return orderCart, err
	}

	return orderCart, nil

}

func (c *CartRepository) InserCart(cart OrderCart) error {
	var sqlStatement string

	sqlStatement = `INSERT INTO orders (book_id, quantity) VALUES (?, ?)`

	_, err := c.db.Exec(sqlStatement, cart.BookID, cart.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) UpdateCart(cart OrderCart) error {
	var sqlStatement string

	sqlStatement = `UPDATE orders SET quantity = quantity + ? WHERE book_id = ?`

	_, err := c.db.Exec(sqlStatement, cart.Quantity, cart.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) ResetCart() error {
	var sqlStatement string

	sqlStatement = `DELETE FROM orders`

	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) TotalPrice() (int, error) {
	var sqlStatement string

	sqlStatement = `SELECT SUM(b.harga * o.quantity) FROM orders o INNER JOIN book b ON o.book_id = b.id`

	var totalPrice int
	err := c.db.QueryRow(sqlStatement).Scan(&totalPrice)
	if err != nil {
		return 0, err
	}

	return totalPrice, nil
}
