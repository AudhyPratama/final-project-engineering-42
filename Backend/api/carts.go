package api

import (
	"backend/repository"
	"encoding/json"
	"net/http"
)

type CartErrorResponse struct {
	Error string `json:"error"`
}

type AddToCartResponse struct {
	BookName     string  `json:"book_name"`
	CategoryName *string `json:"category_name"`
	Penulis      string  `json:"penulis"`
}

type AddtoCartRequest struct {
	BookName string `json:"book_name"`
}
type Cart struct {
	OrderID int `json:"order_id"`
}
type CartResponse struct {
	Carts []repository.OrderCart `json:"carts"`
}

func (api *API) insertToCart(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var request AddtoCartRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)

	book, err := api.bookRepo.FetchBooksByName(request.BookName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
		return
	}

	carts, err := api.cartRepo.FetchCartByBookID(book.ID)
	if err == nil && carts.OrderID != 0 {
		err = api.cartRepo.UpdateCart(carts)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(AddToCartResponse{
			BookName:     book.BookName,
			CategoryName: book.CategoryName,
			Penulis:      book.Penulis,
		})
		return
	}

	err = api.cartRepo.InsertToCart(repository.OrderCart{
		BookID:   book.ID,
		Quantity: 1,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddToCartResponse{
		BookName:     book.BookName,
		CategoryName: book.CategoryName,
		Penulis:      book.Penulis,
	})
}

func (api *API) getCart(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	orderCart, err := api.cartRepo.FetchCarts()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	encoder.Encode(CartResponse{Carts: orderCart})
}

func (api *API) deleteAllCart(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)

	err := api.cartRepo.ResetCart()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete All Success"))
}

func (api *API) deleteCart(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)

	var cart Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	encoder := json.NewEncoder(w)
	err = api.cartRepo.DeleteCart(cart.OrderID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(BookListErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Book in Cart Success"))
}
