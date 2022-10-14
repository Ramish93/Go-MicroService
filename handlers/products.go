package handlers

import (
	"log"
	"net/http"
	"yt-go-microservice/data"
)

type Products struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		p.getProducts(rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request){
	//list of products
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal products", http.StatusInternalServerError)
	}
}
