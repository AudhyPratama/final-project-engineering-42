# PesBuk

# Available APIs
* `POST` : `/api/register`
* `POST` : `/api/user/login`
* `POST` : `/api/user/logout`
* `GET`  : `/api/products`
* `GET`  : `/api/book?book_name=<book_name>`
* `GET`  : `/api/book?penulis=<penulis>`
* `GET`  : `/api/book?penerbit=<penerbit>`
* `GET`  : `/api/carts`
* `GET`  : `/api/cart/delete`

# How to run Service
* Database : Masuk ke directory `Backend`, lalu run `go run db\main.go`, untuk running database SQLite

* Main : Masuk ke directory `Backend`, lalu run `go run main.go` untuk running main Service


# API Documentation
* Register (User harus mendaftar akun terlebih dahulu)
* Login (User dapat login dengan akun yang telah terdaftar)

# Register

* ![Capture1](https://user-images.githubusercontent.com/100668235/174640317-fb4d911f-38fa-408d-98ff-d183b99f3ac6.PNG)


+ request Body

* ![Capture2](https://user-images.githubusercontent.com/100668235/174641111-a54cabaf-4df2-412f-9cf4-3011c1c000c4.PNG)


+ response 

* ![Capture3](https://user-images.githubusercontent.com/100668235/174641228-730691c4-b790-48bf-b87d-575a9c969ff2.PNG)


# Login

* ![Capture4](https://user-images.githubusercontent.com/100668235/174641727-f9174f8a-c1d6-4158-9241-8864f58a5ac5.PNG)


+ request Body

* ![Capture5](https://user-images.githubusercontent.com/100668235/174641853-347cf19c-ddc4-4434-99b2-0db4399ba8d0.PNG)


+ response 

* ![Capture6](https://user-images.githubusercontent.com/100668235/174642004-1d404703-d9c3-41e8-9220-668b4884fc25.PNG)


# List Products / List Book

![path](https://user-images.githubusercontent.com/100668235/174638909-2aad9d16-3da8-4218-ae9c-981e91b92cda.PNG)


![response](https://user-images.githubusercontent.com/100668235/174639074-9c4b17ce-4498-44c2-ae60-6a4e0cad4cd3.PNG)
