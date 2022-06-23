package api

import (
	"backend/repository"
	"encoding/json"
	"net/http"
)

type BookListErrorResponse struct {
	Error string `json:"error"`
}

type Book struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Penulis      string `json:"penulis"`
	Penerbit     string `json:"penerbit"`
	CategoryName string `json:"category_name"`
	Kondisi      string `json:"kondisi"`
	Berat        string `json:"berat"`
	Stock        int64  `json:"stock"`
	Harga        int64  `json:"harga"`
	Deskripsi    string `json:"deskripsi"`
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
