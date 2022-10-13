package main

import (
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello, world!")
	})
	
	http.HandleFunc("/googby", func(http.ResponseWriter, *http.Request){
		log.Println("goodby, world!")
	})

	http.ListenAndServe(":9090", nil)
}