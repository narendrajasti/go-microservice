// Package handlers classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
// 	BasePath: /products
// 	Version: 1.0.0
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/narendrajasti/go-microservice/data"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

// swagger:parameter updateProduct
type productIDParameterWrappergo struct {
	// in: path
	// required: true
	ID int `json:"id"`
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// Products handler
type Products struct {
	l *log.Logger
}

// NewProduct retuns list of products
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// KeyProduct used
type KeyProduct struct{}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}

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

		err = product.Validate()

		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)

		handler.ServeHTTP(rw, r)
	})
}
