package api

import (
	"encoding/json"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name    string  `json:"name"`
	Penulis string  `json:"penulis"`
	Price   float64 `json:"price"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	products, err := api.productRepo.FetchProducts()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ProductListErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, product := range products {
		response.Products = append(response.Products, Product{
			Name:    product.Title,
			Penulis: product.Penulis,
			Price:   product.Price,
		})
	}

	encoder.Encode(response)
}
