package handlers

import (
	"net/http"
	"strconv"
	"yt-go-microservice/data"

	"github.com/gorilla/mux"
)


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