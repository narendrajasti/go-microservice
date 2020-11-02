package handlers

import (
	"net/http"

	"github.com/narendrajasti/go-microservice/data"
)

// swagger:route POST / products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// AddProduct adds a product to list
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST request")

	product := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("product: %#v", product)
	data.AddProduct(&product)
}
