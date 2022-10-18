package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"yt-go-microservice/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Hello world, the web server
	l:= log.New(os.Stdout, "product-api", log.LstdFlags,)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	
// server:=
	s:= &http.Server{
		Addr: ":9090",
		Handler: sm,
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
	//c=channel
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