package handlers

import (
	"net/http"
	"yt-go-microservice/data"
)


func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request){
	//list of products
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal products", http.StatusInternalServerError)
	}
}
