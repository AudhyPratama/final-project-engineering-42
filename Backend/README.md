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
* Migration : Masuk ke directory `Backend`, lalu run `go run db\main.go`, untuk running Migration database SQLite

* Main : Masuk ke directory `Backend`, lalu run `go run main.go` untuk running main Service


# API Documentation

# Register

* ![Capture1](https://user-images.githubusercontent.com/100668235/174640317-fb4d911f-38fa-408d-98ff-d183b99f3ac6.PNG)


+ request Body

* ![Capture2](https://user-images.githubusercontent.com/100668235/174641111-a54cabaf-4df2-412f-9cf4-3011c1c000c4.PNG)


+ response 

* ![Capture3](https://user-images.githubusercontent.com/100668235/174641228-730691c4-b790-48bf-b87d-575a9c969ff2.PNG)


# Login

* ![api_register](https://user-images.githubusercontent.com/100668235/175238397-8d050661-6535-43ab-8f76-02a5b4f9d5ab.PNG)



+ request Body

* ![api_register1](https://user-images.githubusercontent.com/100668235/175238480-71a80662-823b-4099-bae6-c04e65bde831.PNG)



+ response 

* ![Capture6](https://user-images.githubusercontent.com/100668235/174642004-1d404703-d9c3-41e8-9220-668b4884fc25.PNG)


# List Products / List Book

![api_products](https://user-images.githubusercontent.com/100668235/175234134-983b8b58-2988-4759-a8b2-4095e89316f8.PNG)


+ response
![api_products1](https://user-images.githubusercontent.com/100668235/175234394-9c4c4572-19a9-48ed-96ab-f8c1f5762bb8.PNG)



# List Cart

![api_carts](https://user-images.githubusercontent.com/100668235/175234522-7af39484-1db3-4ef4-8d0b-8c753c465396.PNG)


+ response

![api_carts1](https://user-images.githubusercontent.com/100668235/175234615-09190c2c-d6b7-4320-a259-fcb9be8669b7.PNG)


# Get Request Book berdasarkan nama Buku

![api_bookname](https://user-images.githubusercontent.com/100668235/175236431-1811db85-294d-47ba-bbd8-0eb2f912bc5c.PNG)


+ response

![api_bookname1](https://user-images.githubusercontent.com/100668235/175236524-bd5d2288-b811-4f8d-a91f-643054c837b2.PNG)


# Get Request Book berdasarkan Penulis buku

![api_bookpenulis](https://user-images.githubusercontent.com/100668235/175236698-327c68f4-9fd8-4484-86dd-23786ce9d3f5.PNG)


+ response

![api_bookpenulis1](https://user-images.githubusercontent.com/100668235/175236793-96bd8c08-6658-423e-82f1-3c51f4300ae6.PNG)


# Get Request Book berdasarkan Penerbit buku

![api_bookpenerbit](https://user-images.githubusercontent.com/100668235/175236931-f6c6de5d-443d-4a9a-9f5e-02bab5f0a701.PNG)


+ response

![api_bookpenerbit1](https://user-images.githubusercontent.com/100668235/175237011-99f03483-22e5-4432-a9b9-89214d258243.PNG)


# Delete All Cart

![api_cartdelete](https://user-images.githubusercontent.com/100668235/175237479-6a638389-ee7c-4c4c-ba01-685c1d44db59.PNG)


+ response

![api_cartdelete1](https://user-images.githubusercontent.com/100668235/175237632-0d9ec1bf-afd8-4246-9765-463e38b9f95b.PNG)
