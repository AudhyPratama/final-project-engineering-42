# PesBuk

# procedure
* Terdapat dua role, yaitu user dan admin
* user and admin must register an account first
* user and admin can login with a registered account

# Available APIs
* POST : /api/signup
* POST : /api/user/login
* POST : /api/user/logout

# How to run Service
Run the following code in the terminal:
1. Migration : run main.go inside directory Backend\db\migration to Migration database SQLite

go run Backend\db\migration\main.go

2. Main : run main.go inside directory Backend to running main Service

go run Backend\main.go