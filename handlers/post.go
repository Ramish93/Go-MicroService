package handlers

import (
	"net/http"
	"yt-go-microservice/data"
)


func (p *Products) Create(rw http.ResponseWriter, r *http.Request){
	p.l.Println("handle post products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("[DEBUG] inserting product: %#v", prod)
	data.AddProduct(&prod)
}
