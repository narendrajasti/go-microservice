package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/narendrajasti/go-microservice/data"
)

// Products handler
type Products struct {
	l *log.Logger
}

// NewProduct retuns list of products
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts returns the product list
func (p Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct adds a product to list
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST request")

	product := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("product: %#v", product)
	data.AddProduct(&product)
}

// UpdateProducts updates the product if exists in the db
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to covert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product, recieved id: {}", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// KeyProduct used
type KeyProduct struct{}

// MiddlewareValidationProduct returns an error if json deseralization fails
func (p Products) MiddlewareValidationProduct(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := data.Product{}

		err := product.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[Error] deserializing product")
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		handler.ServeHTTP(rw, r)
	})
}
