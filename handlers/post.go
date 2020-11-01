package handlers

import (
	"net/http"

	"github.com/narendrajasti/go-microservice/data"
)

// AddProduct adds a product to list
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST request")

	product := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("product: %#v", product)
	data.AddProduct(&product)
}
