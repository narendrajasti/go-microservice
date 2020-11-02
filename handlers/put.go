package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/narendrajasti/go-microservice/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Returns a HTTP 201 with no content
// responses:
// 	201: noContent

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
