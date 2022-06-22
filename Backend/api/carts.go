package api

import (
	"backend/repository"
	"encoding/json"
	"net/http"
)

type CartErrorResponse struct {
	Error string `json:"error"`
}

type AddToCartSuccessResponse struct {
	Name    string  `json:"name"`
	Penulis string  `json:"penulis"`
	Price   float64 `json:"price"`
}

type AddToCartRequest struct {
	Name string `json:"name"`
}

type CartListSuccessResponse struct {
	Carts []repository.OrderCart `json:"carts"`
}

// func (api *API) addToCart(w http.ResponseWriter, req *http.Request) {
// 	api.AllowOrigin(w, req)

// 	var requestBody AddToCartRequest
// 	err := json.NewDecoder(req.Body).Decode(&requestBody)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	encoder := json.NewEncoder(w)

// }

func (api *API) cartlist(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	itemCart, err := api.cartRepo.FetchCarts()
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
	}

	encoder.Encode(CartListSuccessResponse{Carts: itemCart})
}

func (api *API) clearCart(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	err := api.cartRepo.ResetCart()
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reset Cart Success"))
}
