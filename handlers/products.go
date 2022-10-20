package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"yt-go-microservice/data"
)

type productsResponse struct {
	Body []data.Product
}

//swagger:parameters, deleteProduct
type ProductIDParameterWrapper struct {
	//id of product to delete from the database
	ID int `json:"id"`
}

type productsNoContent struct {}

type Products struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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