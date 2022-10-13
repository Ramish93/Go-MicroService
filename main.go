package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
// server:=
	s:= &http.Server{
		Addr: ":9090",
		Handler: mux,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	//timeout Context
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	// http.ListenAndServe(":9090", mux)
}