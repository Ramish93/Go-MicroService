package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"yt-go-microservice/data"
	"yt-go-microservice/handlers"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Hello world, the web server

	l:= log.New(os.Stdout, "product-api", log.LstdFlags,)
	v:= data.NewValidation()

	ph := handlers.NewProducts(l, v)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/",ph.Create)
	postRouter.Use(ph.MiddlewareValidateProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//CORS
	ch:= gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	
// server:=
	s:= &http.Server{
		Addr: ":9090",
		Handler: ch(sm),
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