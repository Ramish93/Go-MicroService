package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct{

}


func (h*Hello) ServerHTTP(rw http.ResponseWriter, req *http.Request){
	log.Println("Hello, world!")
	data, err := ioutil.ReadAll(req.Body)

	fmt.Fprintf(rw, "hello %s\n", data)

	if err!=nil {
		http.Error(rw, "opps error", http.StatusBadRequest)
		return
	}
}