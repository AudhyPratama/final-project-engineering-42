package api

import (
	"backend/repository"
	"encoding/json"
	"net/http"
)

type BookListErrorResponse struct {
	Error string `json:"error"`
}
type AddBook struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	BookName   string `json:"book_name"`
	Penulis    string `json:"penulis"`
	Penerbit   string `json:"penerbit"`
	Kondisi    string `json:"kondisi"`
	Berat      string `json:"berat"`
	Stock      int    `json:"stock"`
	Harga      int    `json:"harga"`
	Deskripsi  string `json:"deskripsi"`
	Image      string `json:"image"`
}
type Book struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Penulis      string  `json:"penulis"`
	Penerbit     string  `json:"penerbit"`
	CategoryName *string `json:"category_name"`
	Kondisi      string  `json:"kondisi"`
	Berat        string  `json:"berat"`
	Stock        int     `json:"stock"`
	Harga        int     `json:"harga"`
	Deskripsi    string  `json:"deskripsi"`
	Image        string  `json:"image"`
}

type BookListSuccessResponse struct {
	Books []Book `json:"products"`
}

type DetailBookResponse struct {
	Book []repository.Book `json:"book"`
}

func (api *API) booktList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := BookListSuccessResponse{}
	response.Books = make([]Book, 0)

	books, err := api.bookRepo.FetchBooks()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(BookListErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, book := range books {
		response.Books = append(response.Books, Book{
			ID:           book.ID,
			CategoryName: book.CategoryName,
			Name:         book.BookName,
			Penulis:      book.Penulis,
			Penerbit:     book.Penerbit,
			Kondisi:      book.Kondisi,
			Berat:        book.Berat,
			Stock:        book.Stock,
			Harga:        book.Harga,
			Deskripsi:    book.Deskripsi,
			Image:        book.Image,
		})
	}

	encoder.Encode(response)
}

func (api *API) getBook(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	bookName := req.URL.Query().Get("book_name")
	penulis := req.URL.Query().Get("penulis")
	penerbit := req.URL.Query().Get("penerbit")

	getBookRequest := repository.GetBookRequest{
		BookName: bookName,
		Penulis:  penulis,
		Penerbit: penerbit,
	}

	encoder := json.NewEncoder(w)

	books, err := api.bookRepo.FetchDetailBook(getBookRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	if len(books) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(BookListErrorResponse{Error: "Book not found"})
		return
	}

	encoder.Encode(DetailBookResponse{Book: books})
	w.WriteHeader(http.StatusOK)
}

func (api *API) addBook(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var books AddBook
	err := json.NewDecoder(req.Body).Decode(&books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	encoder := json.NewEncoder(w)
	err = api.bookRepo.InsertBook(books.CategoryID, books.BookName, books.Penulis, books.Penerbit, books.Kondisi, books.Berat, books.Stock, books.Harga, books.Deskripsi, books.Image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}

func (api *API) updateBook(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var books AddBook
	err := json.NewDecoder(req.Body).Decode(&books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	encoder := json.NewEncoder(w)
	err = api.bookRepo.UpdateBook(books.ID, books.CategoryID, books.BookName, books.Penulis, books.Penerbit, books.Kondisi, books.Berat, books.Stock, books.Harga, books.Deskripsi, books.Image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update Book Success"))
}

func (api *API) deleteBook(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var books AddBook
	err := json.NewDecoder(req.Body).Decode(&books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	encoder := json.NewEncoder(w)
	err = api.bookRepo.DeleteBook(books.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Book Success"))
}
