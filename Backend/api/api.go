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

	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/logout", api.POST(http.HandlerFunc(api.logout)))
	mux.Handle("/api/registrasi", api.POST(http.HandlerFunc(api.signup)))
	mux.Handle("/api/login/forgot-password", api.POST(http.HandlerFunc(api.updatePassword)))
	mux.Handle("/api/profile", api.GET(http.HandlerFunc(api.profile)))
	mux.Handle("/api/profile/update", api.POST(http.HandlerFunc(api.updateProfile)))

	mux.Handle("/api/products", api.GET(http.HandlerFunc(api.booktList)))
	mux.Handle("/api/product/book", api.GET(http.HandlerFunc(api.getBook)))
	mux.Handle("/api/cart/add", api.POST(http.HandlerFunc(api.insertToCart)))
	mux.Handle("/api/carts", api.GET(http.HandlerFunc(api.getCart)))
	mux.Handle("/api/cart/delete-all", api.GET(http.HandlerFunc(api.deleteAllCart)))
	mux.Handle("/api/cart/delete", api.POST(http.HandlerFunc(api.deleteCart)))

	mux.Handle("/api/admin/product/add", api.POST(http.HandlerFunc(api.addBook)))
	mux.Handle("/api/admin/product/update", api.POST(http.HandlerFunc(api.updateBook)))
	mux.Handle("/api/admin/product/delete", api.POST(http.HandlerFunc(api.deleteBook)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}
func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
