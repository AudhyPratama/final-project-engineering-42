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

func (c *CartRepository) FetchCarts() ([]*ItemCart, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT
		i.id,
		u.id as user_id,
		b.id as book_id,
		b.title,
		c.name as category_name,
		b.penulis,
		i.quantity,
		b.price
	FROM item_cart i
	INNER JOIN user u ON i.user_id = u.id
	INNER JOIN book b ON i.book_id = b.id
	INNER JOIN category c ON b.category_id = c.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []*ItemCart
	for rows.Next() {
		var cart ItemCart
		err := rows.Scan(
			&cart.ID,
			&cart.UserID,
			&cart.BookID,
			&cart.Title,
			&cart.CategoryName,
			&cart.Penulis,
			&cart.Quantity,
			&cart.Price)
		if err != nil {
			return nil, err
		}
		carts = append(carts, &cart)
	}

	return carts, nil
}

func (c *CartRepository) FetchCartByID(id int64) (*ItemCart, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT
		i.id,
		u.id as user_id,
		b.id as book_id,
		b.title,
		c.name as category_name,
		b.penulis,
		i.quantity,
		b.price
	FROM item_cart i
	INNER JOIN user u ON i.user_id = u.id
	INNER JOIN book b ON i.book_id = b.id
	INNER JOIN category c ON b.category_id = c.id
	WHERE i.id = ?`

	rows, err := c.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cart *ItemCart
	for rows.Next() {
		cart = &ItemCart{}
		err := rows.Scan(
			&cart.ID,
			&cart.UserID,
			&cart.BookID,
			&cart.Title,
			&cart.CategoryName,
			&cart.Penulis,
			&cart.Quantity,
			&cart.Price)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}

func (c *CartRepository) InserCart(cart *ItemCart) error {
	var sqlStatement string

	sqlStatement = `INSERT INTO item_cart (user_id, book_id, quantity) VALUES (?, ?, ?)`

	_, err := c.db.Exec(sqlStatement, cart.UserID, cart.BookID, cart.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) UpdateCart(cart *ItemCart) error {
	var sqlStatement string

	sqlStatement = `UPDATE item_cart SET quantity = ? WHERE id = ?`

	_, err := c.db.Exec(sqlStatement, cart.Quantity, cart.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) DeleteCart(id int64) error {
	var sqlStatement string

	sqlStatement = `DELETE FROM item_cart`

	_, err := c.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) TotalPrice() (int, error) {
	var sqlStatement string

	sqlStatement = `
	SELECT SUM(i.quantity * b.price) as total_price
	FROM item_cart i
	INNER JOIN book b ON i.book_id = b.id`

	var totalPrice int
	err := c.db.QueryRow(sqlStatement).Scan(&totalPrice)
	if err != nil {
		return 0, err
	}

	return totalPrice, nil
}
