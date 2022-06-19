package api

import (
	"encoding/json"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string  `json:"name"`
	Category int64   `json:"category"`
	Price    float64 `json:"price"`
	Stock    int64   `json:"stock"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	products, err := api.productRepo.GetAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ProductListErrorResponse{err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, product := range products {
		response.Products = append(response.Products, Product{
			Name:     product.Title,
			Category: product.CategoryID,
			Price:    product.Price,
			Stock:    product.Stock,
		})
	}

	encoder.Encode(response)
}
