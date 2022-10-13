package main

import (
	"log"
	"net/http"
	"os"
	"yt-go-microservice/handlers"
)

func main() {
	// Hello world, the web server
	l:= log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)

	mux := http.NewServeMux()
	mux.Handle("/", helloHandler)

	
	http.ListenAndServe(":9090", mux)
}