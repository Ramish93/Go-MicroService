package handlers

import (
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

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request){
	//list of products
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("handle post products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmashell json", http.StatusBadRequest)
	}
	// p.l.Printf("prod: %#v",prod)
	data.AddProduct(prod)
}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to parse id", http.StatusBadRequest)
		return
	}
	p.l.Println("handle Put products", id)

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmashell json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound{
		http.Error(rw, "product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "unable to update product", http.StatusInternalServerError)
		return
	}
}