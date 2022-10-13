package main

import (
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	http.HandleFunc("/", func(rw http.ResponseWriter, req*http.Request) {
		
	})
	
	http.HandleFunc("/goodby", func(http.ResponseWriter, *http.Request){
		log.Println("goodby, world!")
	})

	http.ListenAndServe(":9090", nil)
}