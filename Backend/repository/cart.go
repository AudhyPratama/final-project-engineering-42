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

	sqlStatement := `
	SELECT
		o.order_id,
		o.quantity,
		b.book_name,
		b.penulis,
		c.name as category_name,
		b.penerbit,
		b.harga
	FROM orders o
	INNER JOIN books b ON o.book_id = b.id
	INNER JOIN categories c ON b.categori_id = c.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []OrderCart
	for rows.Next() {
		var cart OrderCart
		err := rows.Scan(
			&cart.OrderID,
			&cart.Quantity,
			&cart.BookName,
			&cart.Penulis,
			&cart.CategoryName,
			&cart.Penerbit,
			&cart.Harga)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

func (c *CartRepository) FetchCartByBookID(order_id int) (OrderCart, error) {
	var orderCart OrderCart

	sqlStatement := `
	SELECT
		o.order_id,
		o.book_id,
		o.quantity,
		b.book_name,
		b.penulis,
		b.penerbit,
		b.harga
	FROM orders o
	INNER JOIN books b ON o.book_id = b.id
	WHERE o.book_id = ?`

	row := c.db.QueryRow(sqlStatement, order_id)
	err := row.Scan(
		&orderCart.OrderID,
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

func (c *CartRepository) InsertToCart(cart OrderCart) error {

	_, err := c.db.Exec(`INSERT INTO orders (book_id, quantity) VALUES (?, ?)`, cart.BookID, cart.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) UpdateCart(cart OrderCart) error {

	_, err := c.db.Exec(`UPDATE orders SET quantity = quantity + ? WHERE book_id = ?`, cart.Quantity, cart.OrderID)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) ResetCart() error {

	_, err := c.db.Exec(`DELETE FROM orders`)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) TotalPrice() (int, error) {

	var totalPrice int
	err := c.db.QueryRow(`SELECT SUM(b.harga * o.quantity) FROM orders o INNER JOIN books b ON o.book_id = b.id`).Scan(&totalPrice)
	if err != nil {
		return 0, err
	}

	return totalPrice, nil
}

func (c *CartRepository) DeleteCart(order_id int) error {

	_, err := c.db.Exec(`DELETE FROM orders WHERE order_id = ?`, order_id)
	if err != nil {
		return err
	}

	return nil
}
