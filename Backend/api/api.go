package api

import (
	"backend/repository"
	"fmt"
	"net/http"
)

type API struct {
	userRepo repository.UserRepository
	mux      *http.ServeMux
}

func NewAPI(userRepo repository.UserRepository) API {
	mux := http.NewServeMux()
	api := API{
		userRepo, mux,
	}
	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}
func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
