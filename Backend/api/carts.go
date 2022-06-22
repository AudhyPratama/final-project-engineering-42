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
	BookName string `json:"book_name"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
	Harga    int64  `json:"harga"`
}

type AddtoCartRequest struct {
	BookName string `json:"book_name"`
}

type CartResponse struct {
	Carts []repository.OrderCart `json:"carts"`
}

type AddToCartListResponse struct {
	Cart []AddToCartResponse `json:"cart"`
}

func (api *API) addToCart(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)

	var request AddtoCartRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)

	book, err := api.bookRepo.FetchBooksByName(request.BookName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(CartErrorResponse{Error: "Book not found"})
		return
	}

	cart, err := api.cartRepo.FetchCartByID(book.ID)
	if err != nil && cart.OrderID != 0 {
		err := api.cartRepo.UpdateCart(cart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(CartErrorResponse{Error: "Failed to update cart"})
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(AddToCartResponse{
			BookName: book.BookName,
			Penulis:  book.Penulis,
			Penerbit: book.Penerbit,
			Harga:    book.Harga,
		})
		return
	}

	err = api.cartRepo.InsertToCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(CartErrorResponse{Error: "Failed to insert to cart"})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddToCartResponse{
		BookName: book.BookName,
		Penulis:  book.Penulis,
		Penerbit: book.Penerbit,
		Harga:    book.Harga,
	})
}

func (api *API) fetchCart(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)

	encoder := json.NewEncoder(w)

	carts, err := api.cartRepo.FetchCarts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(CartErrorResponse{Error: "Failed to fetch cart"})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(CartResponse{Carts: carts})

}

func (api *API) deleteAllCart(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)

	err := api.cartRepo.ResetCart()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}
