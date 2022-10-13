package main

import (
	"log"
	"net/http"
	"os"
	"yt-go-microservice/handlers"
)

func main() {
	// Hello world, the web server
	l:= log.New(os.Stdout, "product-api", log.LstdFlags,)

	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	mux := http.NewServeMux()
	mux.Handle("/", helloHandler)
	mux.Handle("/goodbye", goodbyeHandler)

	
	http.ListenAndServe(":9090", mux)
}