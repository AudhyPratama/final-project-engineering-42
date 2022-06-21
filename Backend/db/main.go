package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/store-app.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id integer not null primary key AUTOINCREMENT,
		name varchar(100) not null,
		email varchar(100) not null,
		password varchar(100) not null,
		role varchar(100) not null
	);
	
	INSERT INTO users (id, name, email, password, role) VALUES 
		(110108,"Dito","dito@gmail.com","dito7654","user"), 
		(120193, "Dio", "dio@gmail.com", "dio12", "user");
	
	CREATE TABLE IF NOT EXISTS categories (
		id integer not null primary key AUTOINCREMENT,
		name varchar(100) not null
	);

	INSERT INTO categories (id, name) VALUES
		(965321, "Politik"),
		(965322, "Olahraga"),
		(965323, "Kesehatan"),
		(965324, "Teknologi"),
		(965325, "Bisnis"),
		(965326, "Sosial"),
		(965327, "Biology");
	
	CREATE TABLE IF NOT EXISTS books (
		id integer not null primary key AUTOINCREMENT,
		categori_id integer not null,
		book_name varchar(100) not null,
		penulis varchar(100) not null,
		penerbit varchar(100) not null,
		kondisi varchar(100) not null,
		berat varchar(100) not null,
		stock integer not null,
		harga float not null,
		deskripsi text,
		FOREIGN KEY (categori_id) REFERENCES categories(id)
	);

	INSERT INTO books (id, categori_id, book_name, penulis, penerbit, kondisi, berat, stock, harga, deskripsi) VALUES
		(071235, 965321, "How To Win An Argument", "abcdef", "ghijk", "baru", "200 Gram", 50, 80.999, 
		" Lorem ipsum dolor sit amet, consectetur adipiscing elit"),
		(071236, 965321, "Matinya Demokrasi dan Kuasa Teknologi", "Jamie Bartlett", "lmnop", "baru", "350 Gram", 5, 75.999,
		"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"),
		(071237, 965327, "The Power of Habit", "Charles Duhigg", "qrstu", "baru", "200 Gram", 50, 80.999,
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit"),
		(071238, 965322, "The Ultimate Guide to Soccer", "Jamie Bartlett", "vwxyz", "baru", "350 Gram", 5, 75.999,
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit");
	
	CREATE TABLE IF NOT EXISTS orders (
		id integer not null primary key AUTOINCREMENT,
		book_id integer not null,
		quantity integer not null,
		FOREIGN KEY (book_id) REFERENCES books(id)
	);

	INSERT INTO orders (id, book_id, quantity) VALUES
		(071237, 110108, 1),
		(071238, 120193, 2);

	CREATE TABLE IF NOT EXISTS payment (
		id integer not null primary key AUTOINCREMENT,
		orders_id integer not null,
		metode_pembayaran varchar(100) not null,
		alamat TEXT not null,
		ongkos_kirim float not null,
		waktu_pembayaran timestamp not null,
		FOREIGN KEY (cart_id) REFERENCES shopping_cart(id)
	);

	`)

	if err != nil {
		panic(err)
	}
}
