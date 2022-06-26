# PesBuk

# Available APIs
* `POST` : `/api/registrasi`
* `POST` : `/api/login`
* `POST` : `/api/logout`
* `POST` : `api/login/forgot-password`
* `POST` : `api/profile`
* `GET`  : `/api/products`
* `GET`  : `/api/product/book?book_name=<book_name>`
* `GET`  : `/api/product/book?penulis=<penulis>`
* `GET`  : `/api/product/book?penerbit=<penerbit>`
* `GET`  : `/api/carts`
* `POST` : `/api/cart/add`
* `GET`  : `/api/cart/delete-all`
* `POST`  : `/api/cart/delete`
* `POST`  : `/api/admin/product/add`
* `POST`  : `/api/admin/product/update`
* `POST`  : `/api/admin/product/delete`

# How to run Service
* Migration : Masuk ke directory `Backend`, lalu run `go run db\main.go`, untuk running Migration database SQLite

* Main : Masuk ke directory `Backend`, lalu run `go run main.go` untuk running main Service

* Test Case : Tetap berada pada directory `final-project-engineering-42`, lalu run `Backend\repository` with `grader-cli`

# API Documentation

# Register
![signup](https://user-images.githubusercontent.com/100668235/175809219-4cf374cd-065f-42d5-8dba-643c83e2287e.PNG)


# Login
![loginn](https://user-images.githubusercontent.com/100668235/175809168-b5b9693f-8c62-46e2-99a5-17538cac20d9.PNG)


# Logout
![logout](https://user-images.githubusercontent.com/100668235/175809244-9b746261-79d2-47c3-930f-ff23c941f98c.PNG)


# Forgot-Password
![forgot-password](https://user-images.githubusercontent.com/100668235/175809155-4990654a-5288-4e62-80f2-cdadeeb1cd64.PNG)


# Profile
![profile](https://user-images.githubusercontent.com/100668235/175809212-865226cb-872f-40f1-adce-60bce3936f7c.PNG)


# List Products / List Book
![product](https://user-images.githubusercontent.com/100668235/175809200-78ea1574-bd22-45a7-9fb0-33f86c73af1c.PNG)


# List Cart
![carts](https://user-images.githubusercontent.com/100668235/175809102-b4b13ed8-f17f-45e5-a96d-52981282490a.PNG)


# Get Request Book berdasarkan nama Buku
![book_name](https://user-images.githubusercontent.com/100668235/175809072-33ae8155-ba77-430b-8610-615e7a0fa6a7.PNG)


# Get Request Book berdasarkan Penulis buku
![penulis](https://user-images.githubusercontent.com/100668235/175809191-9e4b06c4-d092-4849-b7ee-e13a6044c60b.PNG)


# Get Request Book berdasarkan Penerbit buku
![penerbit](https://user-images.githubusercontent.com/100668235/175809178-3b2ed2a9-72e2-4d47-9aea-717aaf1a6796.PNG)


# Add To Cart
![carts add](https://user-images.githubusercontent.com/100668235/175809087-ea17e8ac-0707-4204-a3e5-26fcadc6e9f4.PNG)


# Delete All Cart
![delete-all](https://user-images.githubusercontent.com/100668235/175809136-e21f7d40-df1b-44b4-a7bf-fd2d7dbb70cc.PNG)


# Delete Book in Cart
![delete](https://user-images.githubusercontent.com/100668235/175809118-29fbee67-770d-47bc-82f3-98de6eb34c67.PNG)


