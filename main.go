package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	http.HandleFunc("/", func(rw http.ResponseWriter, req*http.Request) {
		log.Println("Hello, world!")
		data, err := ioutil.ReadAll(req.Body)

		fmt.Fprintf(rw, "hello %s\n", data)

		if err!=nil {
			http.Error(rw, "opps error", http.StatusBadRequest)
			return
		}
	})
	
	http.HandleFunc("/goodby", func(http.ResponseWriter, *http.Request){
		log.Println("goodby, world!")
	})

	http.ListenAndServe(":9090", nil)
}