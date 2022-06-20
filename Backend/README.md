# PesBuk

# Available APIs
* `POST` : `/api/signup`
* `POST` : `/api/user/login`
* `POST` : `/api/user/logout`

# How to run Service
* Database : Masuk ke directory `Backend`, lalu run `go run db\main.go`, untuk running database SQLite

* Main : Masuk ke directory `Backend`, lalu run `go run main.go` untuk running main Service


# API Documentation
* Register (User harus mendaftar akun terlebih dahulu)
* Login (User dapat login dengan akun yang telah terdaftar)

# Register

`POST : http://localhost:8080/api/signup`

+ request Body

`{`

  `"name": "Dito",`
  `"email": "dito@gmail.com",`
  `"password": "dito7654",`
  `"role": "user"`

`}`

+ response `Status:200 OK`

`{`

  `Account registered successfully`

`}`



# Login

`POST : http://localhost:8080/api/user/login`

+ request Body

`{`

  `"email": "dito@gmail.com",`
  `"password": "dito7654"`

`}`

+ response `Status:200 OK`

`{`

  `"email": "dito@gmail.com",`
  `"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRpdG9AZ21haWwuY29tIiwiUm9sZSI6InVzZXIiLCJleHAiOjE2NTU3MTc5NTF9.xr7gSxZiHzytazAUbQw7XKDaC5Z34kPMPGEr9TAvHtU"`

`}`