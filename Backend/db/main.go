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
	
	INSERT or REPLACE INTO users (id, name, email, password, role) VALUES 
		(110108,"Dito","dito@gmail.com","dito7654","user"), 
		(120193, "Dio", "dio@gmail.com", "dio12", "user");
	
	CREATE TABLE IF NOT EXISTS categories (
		id integer not null primary key AUTOINCREMENT,
		name varchar(100) not null
	);

	INSERT or REPLACE INTO categories (id, name) VALUES
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
		harga integer not null,
		deskripsi text,
		image varchar(100) not null,
		FOREIGN KEY (categori_id) REFERENCES categories(id)
	);

	INSERT INTO books (id, categori_id, book_name, penulis, penerbit, kondisi, berat, stock, harga, deskripsi, image) VALUES
		(071235, 965321, "How To Win An Argument", "Marcus Tullius Cicero", "Kepustakaan Populer Gramedia", "baru", "525 Gram", 50, 100000, 
		"Marcus Tullius Cicero, sejak kecil terlatih dalam seluk-beluk retorika. Cicero unggul bukan hanya sebagai 
		pembicara publik yang efektif, melainkan juga sebagai teoretikus seni persuasi lisan yang menghasilkan 
		risalah-risalah bertema retorika.", 
		"https://cf.shopee.co.id/file/561e88051cd120db6ed574ec9b822c75"),
		(071236, 965323, "Covid 19: Seluk Beluk Corona Virus", "Prof.Dr.dr.Anies, M.Kes, PKK", "Ar-ruzz Media", "baru", "350 Gram", 13, 40000,
		"Buku ini mengajak para pembaca sekalian untuk mengetahui segala seluk-beluk tentang COVID-19. Selain pemaparan 
		sejarah virus corona yang ditemukan pada 1960-an, Anda akan diajak untuk mengetahui cara  sederhana agar terhindar 
		dari infeksi COVID-19.", 
		"https://cf.shopee.co.id/file/358a4e2b0b46ddc13aeec03ab566bc87"),
		(071237, 965326, "Filsafat Sosial", "Hans Fink", "Pustaka Pelajar", "baru", "200 Gram", 78, 30500,
		"Buku filsafat sosial ini menyajikan ulasan tentang beberapa sistem filsafat sosial yang amat berpengaruh, yang dijadikan 
		sebagai titik-tolak bagi orientasi diskusi modern. Paparan atas berbagai tahap dalam sejarah filsafat sosial dikemukakan 
		dalam kaitannya dengan pembahasan tentang lingkungan sosial yang berubah.",
		"https://cf.shopee.co.id/file/e40fdd4fab9f5460b08a8159eaded544"),
		(071238, 965325, "Jatuh Bangun Jadi Pengusaha", "Ervina Pitasari", "Checklist", "baru", "350 Gram", 15, 55500,
		"Buku ini menorehkan sejumlah kisah hidup para pengusaha yang luar biasa dalam memulai dan mengembangkan usahanya. 
		Dari kisah hidup inilah, kita dapat memetik banyak pelajaran berharga, salah satunya yaitu tidak ada kata “menyerah” 
		jika ingin sukses.",
		"https://cf.shopee.co.id/file/9a97c6bf91f6ab605cefc0e88a92921d");
	
	CREATE TABLE IF NOT EXISTS orders (
		order_id integer not null primary key AUTOINCREMENT,
		book_id integer not null,
		quantity integer not null,
		FOREIGN KEY (book_id) REFERENCES books(id)
	);

	INSERT INTO orders (order_id, book_id, quantity) VALUES
		(11, 071235, 1),
		(12, 071236, 2),
		(13, 071237, 3),
		(14, 071238, 4);


	CREATE TABLE IF NOT EXISTS payment (
		id integer not null primary key AUTOINCREMENT,
		orders_id integer not null,
		metode_pembayaran varchar(100) not null,
		alamat TEXT not null,
		ongkos_kirim float not null,
		waktu_pembayaran timestamp not null,
		FOREIGN KEY (orders_id) REFERENCES orders(order_id)
	);

	`)

	if err != nil {
		panic(err)
	}
}
