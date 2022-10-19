package handlers

import (
	"net/http"
	"yt-go-microservice/data"
)


func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("handle post products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
