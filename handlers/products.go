package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"yt-go-microservice/data"

	"github.com/gorilla/mux"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("handle post products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to parse id", http.StatusBadRequest)
		return
	}
	p.l.Println("handle Put products", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound{
		http.Error(rw, "product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "unable to update product", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
		//validate product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s",err ), http.StatusBadRequest)
			return
		}
		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}