package handlers

import (
	"encoding/json"
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

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//list of products
	lp:= data.GetProducts()
	d, err:=json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshal products", http.StatusInternalServerError)
	}

	rw.Write(d)

}