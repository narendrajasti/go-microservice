package handlers

import (
	"log"
	"net/http"

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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the product list
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
