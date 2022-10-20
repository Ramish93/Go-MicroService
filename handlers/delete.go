package handlers

import (
	"net/http"
	"strconv"
	"yt-go-microservice/data"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])
	
	p.l.Println("Handle delete request", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound{
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err!= nil {
		http.Error(rw, "product not found", http.StatusInternalServerError)
	}
}