package api

import (
	"backend/repository"
	"fmt"
	"net/http"
)

type API struct {
	userRepo repository.UserRepository
	bookRepo repository.BookRepository
	cartRepo repository.CartRepository
	mux      *http.ServeMux
}

func NewAPI(userRepo repository.UserRepository, bookRepo repository.BookRepository, cartRepo repository.CartRepository) API {
	mux := http.NewServeMux()
	api := API{
		userRepo, bookRepo, cartRepo, mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))
	mux.Handle("/api/signup", api.POST(http.HandlerFunc(api.signup)))

	mux.Handle("/api/products", api.GET(http.HandlerFunc(api.booktList)))
	mux.Handle("/api/book", api.GET(http.HandlerFunc(api.getBook)))
	mux.Handle("/api/carts", api.GET(http.HandlerFunc(api.cartlist)))
	mux.Handle("/api/cart/clear", api.GET(http.HandlerFunc(api.clearCart)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}
func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
