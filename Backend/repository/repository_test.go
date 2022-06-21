package repository_test

import (
	"backend/repository"
	"database/sql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	var db *sql.DB
	var err error
	var userRepo *repository.UserRepository

	BeforeEach(func() {
		db, err = sql.Open("sqlite3", "db/store-app.db")
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

		CREATE TABLE IF NOT EXISTS payment (
			id integer not null primary key AUTOINCREMENT,
			cart_id integer not null,
			metode varchar(100) not null,
			alamat TEXT not null,
			ongkos_kirim float not null,
			time_payment timestamp not null,
			FOREIGN KEY (cart_id) REFERENCES shopping_cart(id)
		);

		`)

		if err != nil {
			panic(err)
		}

		userRepo = repository.NewUserRepository(db)
	})

	AfterEach(func() {

		db, err := sql.Open("sqlite3", "db/store-app.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS categori;
		DROP TABLE IF EXISTS books;
		DROP TABLE IF EXISTS shopping_cart;
		DROP TABLE IF EXISTS payment;`)

		if err != nil {
			panic(err)
		}
	})

	Describe("Select All Users", func() {
		When("get all user list from database", func() {
			It("should return all user list", func() {
				userList, err := userRepo.FetchUsers()
				Expect(err).NotTo(HaveOccurred())

				Expect(userList[0].Email).To(Equal("dito@gmail.com"))
				Expect(userList[0].Password).To(Equal("dito7654"))
				Expect(userList[1].Email).To(Equal("dio@gmail.com"))
				Expect(userList[1].Password).To(Equal("dio12"))
			})
		})
	})

	Describe("Login", func() {
		When("Email and Password are correct", func() {
			It("accepts the login", func() {
				res, err := userRepo.Login("dito@gmail.com", "dito7654")
				Expect(err).NotTo(HaveOccurred())
				Expect(*res).To(Equal("dito@gmail.com"))
			})
		})
		When("Email is correct but password is incorrect", func() {
			It("rejects the login", func() {
				_, err := userRepo.Login("dito@gmail.com", "dito123")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Login Failed"))
			})
		})
		When("both email and password is incorrect", func() {
			It("rejects the login", func() {
				_, err := userRepo.Login("juno@gmail.com", "a23885")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Login Failed"))
			})
		})
	})
})
