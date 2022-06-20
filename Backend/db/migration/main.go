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
	
	CREATE TABLE IF NOT EXISTS categori (
		id integer not null primary key AUTOINCREMENT,
		name varchar(100) not null
	);

	INSERT INTO categori (id, name) VALUES
		(965321, "Politik"),
		(965322, "Olahraga"),
		(965323, "Kesehatan"),
		(965324, "Teknologi"),
		(965325, "Bisnis"),
		(965326, "Sosial"),
		(965327, "Biology");
	
	CREATE TABLE IF NOT EXISTS books (
		id integer not null primary key AUTOINCREMENT,
		user_id integer not null,
		categori_id integer not null,
		title varchar(100) not null,
		penulis varchar(100) not null,
		berat varchar(100) not null,
		stock integer not null,
		price float not null,
		description text,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (categori_id) REFERENCES categori(id)
	);

	INSERT INTO books (id, user_id, categori_id, title, penulis, berat, stock, price, description) VALUES
		(071235, 110108, 965321, "How To Win An Argument", "abcdef", "200 Gram", 50, 80.999, 
		" Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
		Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
		Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
		Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
		(071236, 120193, 965321, "Matinya Demokrasi dan Kuasa Teknologi", "Jamie Bartlett", "350 Gram", 5, 75.999,
		" Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
		Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
		Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
		Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.");
	
	CREATE TABLE IF NOT EXISTS shopping_cart (
		id integer not null primary key AUTOINCREMENT,
		book_id integer not null,
		quantity integer not null,
		FOREIGN KEY (book_id) REFERENCES books(id)
	);

	INSERT INTO shopping_cart (id, book_id, quantity) VALUES
		(531237, 071235, 1),
		(531238, 071236, 2);

	CREATE TABLE IF NOT EXISTS payment (
		id integer not null primary key AUTOINCREMENT,
		cart_id integer not null,
		metode varchar(100) not null,
		alamat TEXT not null,
		ongkos_kirim float not null,
		time_payment timestamp not null,
		FOREIGN KEY (cart_id) REFERENCES shopping_cart(id)
	);

	INSERT INTO payment (id, cart_id, metode, alamat, ongkos_kirim, time_payment) VALUES
		(331639, 531237, "COD", "Jl. Raya Kedungkandang No.1", 10000, "2022-06-15 00:00:00"),
		(331640, 531238, "COD", "Jl. Raya Kedungkandang No.2", 14000, "2022-06-16 00:00:00");
	`)

	if err != nil {
		panic(err)
	}
}
