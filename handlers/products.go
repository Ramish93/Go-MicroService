package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	//handle an update
	if r.Method == http.MethodPost{
		p.addProduct(rw,r)
		return
	}
	// get id out of URI
	if r.Method == http.MethodPut {
		//expect Id in URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g:= reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) !=1{
			http.Error(rw, "Invalid uri", http.StatusBadRequest)
			return
		}
		if len(g[0]) !=1{
			http.Error(rw, "Invalid uri", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid uri", http.StatusBadRequest)
			return
		}

		p.l.Println("got id:", id)
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

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("handle post products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmashell json", http.StatusBadRequest)
	}
	// p.l.Printf("prod: %#v",prod)
	data.AddProduct(prod)
}