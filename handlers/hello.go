package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h*Hello) ServerHTTP(rw http.ResponseWriter, req *http.Request){
	h.l.Println("Hello, world!")
	data, err := ioutil.ReadAll(req.Body)

	if err!=nil {
		http.Error(rw, "opps error", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "hello %s\n", data)
}